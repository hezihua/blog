package impl

import (
	"context"
	"hezihua/apps/blog"
	"hezihua/apps/tag"
)

func (i *impl) CreateTag(ctx context.Context, req *tag.CreateTagRequest) (*tag.Tag, error) {
  
	ins, err := tag.NewTag(req)
	if err != nil {
		return nil, err
	}
	
	i.blog.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.BlogId))

	

	err = i.db.WithContext(ctx).Create(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *impl) QueryTag(ctx context.Context, req *tag.QueryTagRequest) (*tag.TagSet, error) {
	query := i.db.Table("tags").WithContext(ctx)

	if len(req.BlogIds) > 0 {
		
		query = query.Where("blog_id IN (?)", req.BlogIds)
	}

	set := tag.NewTagSet()



	err := query.
		Order("created_at desc").
		Find(&set.Items).
		Error

	if err != nil {
		return nil, err
	}

	set.Total = int64(len(set.Items))

	if err != nil {
		return nil, err
	}

	return set, nil
}

func (i *impl) RemoveTag(ctx context.Context, req *tag.RemoveTagRequest) error {

	return nil
}