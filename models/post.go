package models

import (
	"fmt"
	u "CRUD-GO/utils"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (post *Post) Validate() (map[string]interface{}, bool) {

	if post.Title == "" {
		return u.Message(false, "Title name should be on the payload"), false
	}

	if post.Desc == "" {
		return u.Message(false, "Description should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (post *Post) Create() map[string]interface{} {

	if resp, ok := post.Validate(); !ok {
		return resp
	}

	GetDB().Create(post)

	resp := u.Message(true, "success")
	resp["post"] = post
	return resp
}

func GetPost(id uint) *Post {

	post := &Post{}
	err := GetDB().Table("posts").Where("id = ?", id).First(post).Error
	if err != nil {
		return nil
	}
	return post
}

func GetPosts() []*Post {

	posts := make([]*Post, 0)
	err := GetDB().Table("posts").Find(&posts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return posts
}

func remove() {

}
