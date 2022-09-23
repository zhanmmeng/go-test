package request

type ArticleCreate struct {
	Title      string `form:"title" json:"title" binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"`
	CategoryId int    `form:"categoryId" json:"categoryId" binding:"required"`
}

func (articleCreate *ArticleCreate) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Title.required":      "文章标题不能为空",
		"Content.required":    "文章内容不能为空",
		"CategoryId.required": "分类id不能为空",
	}
}

type ArticleDel struct {
	Id string `form:"id" json:"id" binding:"required"`
}

func (a *ArticleDel) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required": "文章id不能为空",
	}
}

type ArticleModify struct {
	Id         int `form:"id" json:"id" binding:"required"`
	Title      string `form:"title" json:"title" binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"`
	CategoryId int    `form:"categoryId" json:"categoryId" binding:"required"`
}

func (a *ArticleModify) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required":         "文章id不能为空",
		"Title.required":      "文章标题不能为空",
		"Content.required":    "文章内容不能为空",
		"CategoryId.required": "分类id不能为空",
	}
}

type ArticleList struct {
	Id         string `form:"id" json:"id" binding:"required"`
	Title      string `form:"title" json:"title" binding:"required"`
	CategoryId int `form:"categoryId" json:"categoryId" binding:"required"`
	Page       int `form:"page" json:"page" binding:"required"`
	PageSize   int `form:"pageSize" json:"pageSize" binding:"required"`
}

func (a *ArticleList) GetMessage() ValidatorMessages {
	return ValidatorMessages{

	}
}
