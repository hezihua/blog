package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	orm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConfig() *Config {
	return &Config{
		Auth: newAuth(),
		Mysql: NewDefaultMysql(),
	}
}

type Config struct {
	Auth *Auth `toml:"auth"`
	Mysql *Mysql `toml:"mysql"`
}

func newAuth() *Auth {
	return &Auth{
		Username: "admin",
		Password: "123456",
	}
}

type Auth struct {
	Username string `toml:"username" env:"AUTH_USERNAME"`
	Password string `toml:"password" env:"AUTH_PASSWORD"`
}

func NewDefaultMysql () *Mysql{
	return &Mysql{
		Host: "127.0.0.1",
		DB: "vblog",
		Port: 3306,
		Password: "123456",
		Username: "root",

		MaxOpenConn: 10,
		MaxIdleConn: 5,
	}
}


type Mysql struct {
	Host string `toml:"host" env:"MYSQL_HOST"`
	DB string `toml:"db" env:"MYSQL_DB"`
	Port int `toml:"port" env:"MYSQL_PORT"`
	Username string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`

	// 高级参数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`
	
	// 线程里创建有并发安全问题
	lock sync.Mutex
	db *gorm.DB
}

func (a *Auth) Validate(username string, password string) error {
	if !(a.Username == username && a.Password == password) {
		return fmt.Errorf("username or password error")
	}
	return nil
}

func (m *Mysql) GetConnPool() (*sql.DB, error){
	var err error
	// multiStatements 让db 可以执行多个语句 select; insert;
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&multiStatements=true",
		m.Username, m.Password, m.Host, m.Port, m.DB)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}

	// 对连接池进行设置
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleConn != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil

}

func (m *Mysql) ORM() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.db == nil {
		// 初始化DB
		// 1.1 获取sql.DB
		p, err := m.GetConnPool()
		if err != nil {
			panic(err)
		}

		// 1.2 使用pool 初始化orm db对象
		m.db, err = gorm.Open(orm_mysql.New(orm_mysql.Config{
			Conn: p,
		}), &gorm.Config{
			// 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
			PrepareStmt: true,
			// 对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行。
			// 但这会降低性能，如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约 30%+ 性能提升
			SkipDefaultTransaction: true,
			// 要有效地插入大量记录，请将一个 slice 传递给 Create 方法
			// CreateBatchSize: 200,
		})
		if err != nil {
			panic(err)
		}
	}
	return m.db
}