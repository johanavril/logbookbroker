package model

import (
	"github.com/johanavril/logbookbroker/src/database"
	"github.com/johanavril/logbookbroker/src/util"
)

type User struct {
	Id       int    `db:"id"`
	UserId   string `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Name     string `db:"name"`
}

func (u *User) Upsert() error {
	db := database.Get()

	encryptedPassword, err := util.Encrypt(u.Password)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO users (user_id, username, password, name) 
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id)
		DO UPDATE SET
			username = EXCLUDED.username,
			password = EXCLUDED.password,
			name = EXCLUDED.name`,
		u.UserId, u.Username, encryptedPassword, u.Name)

	return err
}

func (u *User) GetByLineUserId(userId string) error {
	db := database.Get()
	return db.Get(u, "SELECT * FROM users WHERE user_id = $1", userId)
}
