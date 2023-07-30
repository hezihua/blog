package impl

import (
	"context"
	"hezihua/apps/blog"

	"dario.cat/mergo"
)

func (i *impl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	// 校验参数
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	//构建对象
  ins := blog.NewBlog(req)

	// 保存对象

	if err := i.db.WithContext(ctx).Save(ins).Error; err != nil {
		return nil, err
	}



	
	return ins, nil
}
func (i *impl) QueryBlog(ctx context.Context, req *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	
	query := i.db.Table("blogs").WithContext(ctx)

	if req.Author != "" {
		
		query = query.Where("author = ?", req.Author)
	}

	if req.Status != nil {
		query = query.Where("status = ?", req.Status)
	}

	if req.Keywords != "" {
		query = query.Where("title like ?", "%"+req.Keywords+"%")
	}

	set := blog.NewBlogSet()



	err := query.
		Order("created_at desc").
		Offset(int(req.Offset())).
		Limit(int(req.PageSize)).
		Find(&set.Items).
		Error

	if err != nil {
		return nil, err
	}

	err = query.Count(&set.Total).Error

	if err != nil {
		return nil, err
	}


	return set, nil
}


func (i *impl) DescribeBlog(ctx context.Context, req *blog.DescribeBlogRequest) (*blog.Blog, error) {
	query := i.db.Table("blogs").WithContext(ctx)
	ins := blog.NewBlog(blog.NewCreateBlogRequest())
	err := query.Where("id = ?", req.Id).First(ins).Error
	if err != nil {
	  return nil, err	
	}
	return ins, nil
}
func (i *impl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.Id))
	if err != nil {
		return nil, err
	}
	switch req.UpdateMode {
		case blog.PUT:
			ins.CreateBlogRequest = req.CreateBlogRequest
		case blog.PATCH:
			// if req.Author != "" {
			// 	ins.Author = req.Author
			// }
			// if req.Title != "" {
			// 	ins.Title
			// }
			mergo.MergeWithOverwrite(ins.CreateBlogRequest, req.CreateBlogRequest)
	}
	if err := ins.CreateBlogRequest.Validate(); err != nil {
		return nil, err
	}
	err = i.db.
	Table("blogs").
	WithContext(ctx).
	Where("id = ?", req.Id).
	Updates(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}
func (i *impl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.Id))
	if err != nil {
		return nil, err
	}

	err = i.db.
		Table("blogs").
		WithContext(ctx).
		// Where("id = ?", req.Id).
		Delete(&blog.Blog{Id: int64(req.Id)}).
		Error

	if err != nil {
		return nil, err
	}

	

	return ins, nil
}

