package models

import "strconv"

type Admin struct {
	ID
	Account  string `json:"account"`
	Password string `json:"password"`
	Timestamps
}

func (admin Admin) GetUid() string {
	return strconv.Itoa(int(admin.ID.ID))
}

func (admin Admin) GetAccount() string {
	return admin.Account
}

// TableName 设置User的表名为`admin_user`
func (admin Admin) TableName() string {
	return "admin_user"
}