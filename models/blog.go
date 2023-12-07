package models

type Blog struct {
	Id     uint   `json:"id"`
	Title  string `json:title`
	Desc   string `json:desc`
	Image  string `json:image`
	UserID uint   `json:userid`
	// fix this id should be user id that of the user which is logged in
	User User `json:"user";gorm:"foreignKey:UserID"`
}
