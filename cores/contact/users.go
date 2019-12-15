package contact

import (
	"context"
	"log"
	"sync"

	"github.com/vnotes/workweixin_app/cores/dbs"

	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(ctx context.Context, msg *wxContactMsg) error
	UpdateUser(ctx context.Context, msg *wxContactMsg) error
	DeleteUser(ctx context.Context, msg *wxContactMsg) error
}

var (
	u    User
	once sync.Once
)

func Cli() User {
	once.Do(func() {
		if u == nil {
			u = &UserClient{db: dbs.DB}
		}
	})
	return u
}

type UserClient struct {
	db *sqlx.DB
}

func (u *UserClient) CreateUser(ctx context.Context, msg *wxContactMsg) error {
	_sql := `INSERT INTO users(user_id, user_name, gender, state, email, mobile, create_time)
				VALUES(:user_id, :user_name, :gender, :state, :email, :mobile, :create_time);`
	_, err := u.db.NamedExecContext(ctx, _sql, msg)
	if err != nil {
		log.Printf("create user error %#v", err)
		return err
	}
	return nil

}

func (u *UserClient) UpdateUser(ctx context.Context, msg *wxContactMsg) error {
	_sql := `
		UPDATE
			users
		SET
			user_name = ?,
			gender = ?,
			state = ?,
			email = ?,
			mobile = ?,
			create_time = ?
		WHERE
			user_id = ?;`
	_, err := u.db.ExecContext(ctx, _sql, msg.Name, msg.Gender, msg.Status, msg.Email, msg.Mobile, msg.CreateTime, msg.UserID)
	if err != nil {
		log.Printf("update user(%s-%s) error %#v", msg.UserID, msg.Name, err)
		return err
	}
	return nil
}

func (u *UserClient) DeleteUser(ctx context.Context, msg *wxContactMsg) error {
	_sql := `DELETE FROM users WHERE user_id = ?;`
	_, err := u.db.ExecContext(ctx, _sql, msg.UserID)
	if err != nil {
		log.Printf("delete user(%s-%s) error %#v", msg.UserID, msg.Name, err)
		return err
	}
	return nil
}

const (
	CreateUser = "create_user"
	UpdateUser = "update_user"
	DeleteUser = "delete_user"
)
