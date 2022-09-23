package request

type CategoryCreate struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func (categoryCreate *CategoryCreate) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Name.required": "分类名称不能为空",
	}
}

type CategoryDel struct {
	Id string `form:"id" json:"id" binding:"required"`
}

func (categoryDel *CategoryDel) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Id.required": "分类id不能为空",
	}
}

type CategoryModify struct {
	Id   int    `form:"id" json:"id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}

func (categoryModify *CategoryModify) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"Name.required": "分类名称不能为空",
		"Id.required":   "分类id不能为空",
	}
}
