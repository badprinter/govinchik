package storage

import (
	"context"
	"github.com/badprinter/govinchik/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type UserManager struct{}

func (u *UserManager) getUserByID(id int64, url *string) *models.User {

	conn, err := pgxpool.New(context.Background(), *url)
	if err != nil {
		return nil
	}
	defer conn.Close()
	user := models.User{}

	err = conn.QueryRow(context.Background(), "SELECT id, name, telegram_id, is_deleted FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.TelegramId, &user.IsDeleted)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &user

}
func (u *UserManager) getUserByTelegramId(id int64, url *string) *models.User {
	conn, err := pgxpool.New(context.Background(), *url)
	if err != nil {
		return nil
	}
	defer conn.Close()
	user := models.User{}

	err = conn.QueryRow(context.Background(), "SELECT id, name, telegram_id, is_deleted FROM users WHERE telegram_id = $1", id).Scan(&user.Id, &user.Name, &user.TelegramId, &user.IsDeleted)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &user
}

func (u *UserManager) updateUser(user *models.User, url *string) *models.User {
	conn, err := pgxpool.New(context.Background(), *url)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	result := models.User{}
	err = conn.QueryRow(context.TODO(), "UPDATE users SET telegram_id=$1, name=$2, is_deleted=$3", user.TelegramId, user.Name, user.IsDeleted).Scan(&result.Id, &result.Name, &result.TelegramId, &result.IsDeleted)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &result
}
func (u *UserManager) createUser(user *models.User, url *string) *models.User {
	conn, err := pgxpool.New(context.Background(), *url)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer conn.Close()

	result := models.User{}
	err = conn.QueryRow(context.TODO(), "INSERT INTO users (name, telegram_id, is_deleted) VALUES ($1, $2, $3)", user.Name, user.TelegramId, user.IsDeleted).Scan(&result.Id, &result.Name, &result.TelegramId)
	if err != nil {
		log.Println(err)
	}
	return &result
}
func (u *UserManager) deleteUser(user *models.User, url *string) bool {
	conn, err := pgxpool.New(context.Background(), *url)
	if err != nil {
		log.Println(err)
		return false
	}
	defer conn.Close()
	_, err = conn.Query(context.TODO(), "UPDATE users SET is_deleted=1 WHERE id=$1", user.Id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
