package vktest

type Admin struct {
	Id        int    `json:"-" db:"-"`
	Adminname string `json:"adminname" db:"username" binding:"required"`
	Password  string `json:"password" db:"password" binding:"required"`
}
