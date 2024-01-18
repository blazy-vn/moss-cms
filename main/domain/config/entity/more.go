package entity

type More struct {
	ArticleViewsPool      int    `json:"article_views_pool"`      // 更新文章浏览量池容量
	FastOffsetMinPage     int    `json:"fast_offset_min_page"`    // 使用快速查询分页时最小页数
	ViewAllCategoryLimit  int    `json:"view_all_category_limit"` // 前台视图调用所有分类时限制的最大数量
	UniqueTitle           bool   `json:"unique_title"`            // 添加或更新文章时。标题必须唯一
	DefaultDateTimeFormat string `json:"default_datetime_format"` // Date time format for record rendering
	DefaultTimeFormat     string `json:"default_time_format"`     // Date time format for record rendering
}

func NewMore() *More {
	return &More{
		ArticleViewsPool:      100,
		FastOffsetMinPage:     100,
		ViewAllCategoryLimit:  200,
		DefaultDateTimeFormat: "2006-01-02 15:04:05",
		DefaultTimeFormat:     "15:04:05",
	}
}

func (*More) ConfigID() string {
	return "more"
}
