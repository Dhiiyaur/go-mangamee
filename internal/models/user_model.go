package models

type User struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	HashPassword []byte `json:"hashpassword"`
	AuthToken    string `json:"auth_token"`
}

type UserHistory struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CoverImg   string `json:"cover_img"`
	Lang       string `json:"lang"`
	LatestRead string `json:"latest_read"`
}
