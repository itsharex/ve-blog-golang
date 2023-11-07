package response

import "time"

type ArticleDetails struct {
	ID             int       `json:"id"`               // 文章ID
	ArticleCover   string    `json:"article_cover"`    // 文章缩略图
	ArticleTitle   string    `json:"article_title"`    // 标题
	ArticleContent string    `json:"article_content"`  // 内容
	LikeCount      int       `json:"like_count"`       // 点赞量
	ViewsCount     int       `json:"views_count"`      // 浏览量
	Type           int       `json:"type"`             // 文章类型
	OriginalURL    string    `json:"original_url"`     // 原文链接
	CreatedAt      time.Time `json:"created_at"`       // 发表时间
	UpdatedAt      time.Time `json:"updated_at"`       // 更新时间
	CategoryID     int       `json:"category_id"`      // 文章分类ID
	CategoryName   string    `json:"category_name"`    // 文章分类名
	ArticleTagList []*TagDTO `json:"article_tag_list"` // 文章标签列表
}

type ArticleConditionDTO struct {
	ArticleDTOList []*ArticleDetails `json:"article_dto_list"` // 文章列表
	ConditionName  string            `json:"condition_name"`   // 条件名
}

// ArticlePaginationDTO represents an article
type ArticlePaginationDTO struct {
	ID                   int                  `json:"id"`                     // 文章ID
	ArticleCover         string               `json:"article_cover"`          // 文章缩略图
	ArticleTitle         string               `json:"article_title"`          // 标题
	ArticleContent       string               `json:"article_content"`        // 内容
	LikeCount            int                  `json:"like_count"`             // 点赞量
	ViewsCount           int                  `json:"views_count"`            // 浏览量
	Type                 int                  `json:"type"`                   // 文章类型
	OriginalURL          string               `json:"original_url"`           // 原文链接
	CreatedAt            time.Time            `json:"created_at"`             // 发表时间
	UpdatedAt            time.Time            `json:"updated_at"`             // 更新时间
	CategoryID           int                  `json:"category_id"`            // 文章分类ID
	CategoryName         string               `json:"category_name"`          // 文章分类名
	ArticleTagList       []*TagDTO            `json:"article_tag_list"`       // 文章标签列表
	LastArticle          *ArticlePreviewDTO   `json:"last_article"`           // 上一篇文章
	NextArticle          *ArticlePreviewDTO   `json:"next_article"`           // 下一篇文章
	RecommendArticleList []*ArticlePreviewDTO `json:"recommend_article_list"` // 推荐文章列表
	NewestArticleList    []*ArticlePreviewDTO `json:"newest_article_list"`    // 最新文章列表
}

// TagDTO 标签
type TagDTO struct {
	ID      int    `json:"id"`       // 标签ID
	TagName string `json:"tag_name"` // 标签名
}

// CategoryDTO 分类
type CategoryDTO struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"` // 分类名
	ArticleCount int64  `json:"article_count"`
}

// ArticleArchivesDTO 文章预览
type ArticlePreviewDTO struct {
	ID           int       `json:"id"`            // 文章ID
	ArticleCover string    `json:"article_cover"` // 文章缩略图
	ArticleTitle string    `json:"article_title"` // 标题
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
}
