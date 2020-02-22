package model

type Category struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	No   int    `db:"no"`
}
