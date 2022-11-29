package openrecipe

// User is a representation of a client who can post recipes to the site
type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
}
