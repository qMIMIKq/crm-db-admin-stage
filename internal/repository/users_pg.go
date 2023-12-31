package repository

import (
	"crm_admin/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UsersPG struct {
	db *sqlx.DB
}

func (u *UsersPG) EditUser(user domain.UserInfo) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	var args []interface{}
	args = append(args, user.Name, user.Login)

	setter := fmt.Sprintf("SET user_name = $1, login = $2")
	if len(user.Password) > 0 {
		setter += ", password = $3 WHERE users.user_id = $4;"
		args = append(args, user.Password)
	} else {
		setter += " WHERE users.user_id = $3;"
	}

	userQuery := fmt.Sprintf(`UPDATE users %s`, setter)
	args = append(args, user.ID)

	_, err = tx.Exec(userQuery, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	rightsQuery := fmt.Sprintf(`
			UPDATE users_rights
				 SET group_id = $1, plot_id = $2
		   WHERE user_id = $3;
	`)

	_, err = tx.Exec(rightsQuery, user.GroupID, user.PlotID, user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (u UsersPG) CreateUser(user domain.UserInfo) (int, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return 0, err
	}

	userQuery := fmt.Sprintf(`
			INSERT INTO users (user_name, login, password)
			VALUES ($1, $2, $3)								 
   RETURNING user_id;
	`)

	var id int
	err = tx.QueryRow(userQuery, user.Name, user.Login, user.Password).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	rightsQuery := fmt.Sprintf(`
			INSERT INTO users_rights (user_id, group_id, plot_id)
			VALUES ($1, $2, $3)
	`)

	_, err = tx.Exec(rightsQuery, id, user.GroupID, user.PlotID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}

func (u *UsersPG) GetUserByID(userId int) (domain.UserInfo, error) {
	query := fmt.Sprintf(`
		SELECT u.user_id, u.login, u.user_name, g.group_name, p.plot_name, g.group_id, p.plot_id
      FROM users_rights ur
           JOIN users u on u.user_id = ur.user_id
           JOIN groups g on g.group_id = ur.group_id
           JOIN plots p on p.plot_id = ur.plot_id
     WHERE u.user_id = $1
  `)

	var user domain.UserInfo
	err := u.db.Get(&user, query, userId)

	return user, err
}

func (u *UsersPG) GetUsers() (domain.Users, error) {
	query := fmt.Sprintf(`
		SELECT u.user_id, u.user_name, g.group_name, p.plot_name
      FROM users_rights ur
           JOIN users u on u.user_id = ur.user_id
           JOIN groups g on g.group_id = ur.group_id
           JOIN plots p on p.plot_id = ur.plot_id
     ORDER BY u.user_name DESC;
	`)

	var users domain.Users
	err := u.db.Select(&users, query)

	return users, err
}

func NewUsersPG(db *sqlx.DB) *UsersPG {
	return &UsersPG{db: db}
}
