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

func (u *UserClient) GetDBMap(msg *wxContactMsg, value *map[string]interface{}) {
	if msg.Name != nil {
		(*value)["user_name"] = *msg.Name
	}
	if msg.Mobile != nil {
		(*value)["mobile"] = *msg.Mobile
	}
	if msg.Email != nil {
		(*value)["email"] = *msg.Email
	}
	if msg.Status != nil {
		(*value)["state"] = *msg.Status
	}
	if msg.Gender != nil {
		(*value)["gender"] = *msg.Gender
	}
}

func (u *UserClient) CreateUser(ctx context.Context, msg *wxContactMsg) error {
	var value = map[string]interface{}{
		"user_id":     msg.UserID,
		"user_name":   "",
		"mobile":      "",
		"email":       "",
		"state":       4,
		"gender":      1,
		"create_time": msg.CreateTime,
	}
	u.GetDBMap(msg, &value)
	_sql := `INSERT INTO users(user_id, user_name, gender, state, email, mobile, create_time)
				VALUES(:user_id, :user_name, :gender, :state, :email, :mobile, :create_time);`
	_, err := u.db.NamedExecContext(ctx, _sql, value)
	if err != nil {
		log.Printf("create user error %#v", err)
		return err
	}
	return nil

}

func (u *UserClient) UpdateUser(ctx context.Context, msg *wxContactMsg) error {
	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Printf("begin tx error %v", err)
		return err
	}
	var (
		querySQL = "select user_id, user_name, gender, state, email, mobile, create_time from users where user_id = ?;"

		value = make(map[string]interface{})
	)
	if err := tx.QueryRowxContext(ctx, querySQL, msg.UserID).MapScan(value); err != nil {
		_ = tx.Rollback()
		log.Printf("query user id(%s) error %v", msg.UserID, err)
		return err
	}
	u.GetDBMap(msg, &value)
	if msg.NewUserID != nil {
		value["new_user_id"] = *msg.NewUserID
	} else {
		value["new_user_id"] = msg.UserID
	}

	_sql := `
		UPDATE
			users
		SET
		    user_id = :new_user_id,
			user_name = :user_name,
			gender = :gender,
			state = :state,
			email = :email,
			mobile = :mobile,
			create_time = :create_time
		WHERE
			user_id = :user_id;`
	_, err = tx.NamedExecContext(ctx, _sql, value)
	if err != nil {
		_ = tx.Rollback()
		log.Printf("update user(%s-%v) error %#v", msg.UserID, msg.Name, err)
		return err
	}
	_ = tx.Commit()
	return nil
}

func (u *UserClient) DeleteUser(ctx context.Context, msg *wxContactMsg) error {
	_sql := `DELETE FROM users WHERE user_id = ?;`
	_, err := u.db.ExecContext(ctx, _sql, msg.UserID)
	if err != nil {
		log.Printf("delete user(%s-%v) error %#v", msg.UserID, msg.Name, err)
		return err
	}
	return nil
}

const (
	CreateUser = "create_user"
	UpdateUser = "update_user"
	DeleteUser = "delete_user"
)
