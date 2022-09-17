// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewPost struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"user_id"`
}

type NewUser struct {
	Name string `json:"name"`
}

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	User  *User  `json:"user"`
}

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
}
