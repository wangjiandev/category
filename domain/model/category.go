package model

type Category struct {
	ID                  int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	CategoryName        string `json:"category_name" gorm:"unique_index;not_null;"`
	CategoryLevel       uint32 `json:"category_level"`
	CategoryParent      int64  `json:"category_parent"`
	CategoryImage       string `json:"category_image"`
	CategoryDescription string `json:"category_description"`
}
