package domain

type UserId int64

type User struct {
	Id             UserId `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	PasswordDigest string `json:"-"`
	CreatedAt      string `json:"-"`
	UpdatedAt      string `json:"-"`
}
