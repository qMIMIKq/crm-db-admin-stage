package repository

import (
	"crm_admin/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPG struct {
	db *sqlx.DB
}

func (a AuthPG) CheckUser(user domain.UserAuth) (string, error) {
	query := fmt.Sprintf(`
		SELECT g.group_name
      FROM users u
           JOIN users_rights ur ON u.user_id = ur.user_id
           JOIN groups g ON ur.group_id = g.group_id
     WHERE u.login = $1
       AND u.password = $2
			 AND g.group_name = 'админ';
  `)

	var group string
	err := a.db.Get(&group, query, user.Login, user.Password)

	return group, err
}

func (a AuthPG) CreateUser(user domain.UserCreate) (int, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthPG(db *sqlx.DB) *AuthPG {
	return &AuthPG{
		db: db,
	}
}
