package models

type (
	User struct {
		Id        string `db:"id"`
		Username  string `db:"username"`
		Password  string `db:"password"`
		Email     string `db:"email"`
		CreatedAt int64  `db:"created_at"`
	}
)
