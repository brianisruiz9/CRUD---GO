package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreatePost = func(w http.ResponseWriter, r *http.Request) {

	post := &models.Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := post.Create()
	u.Respond(w, resp)
}

var GetPostsFor = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetPosts()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
