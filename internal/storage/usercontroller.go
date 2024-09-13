package storage

import (
	"context"
	"github.com/badprinter/govinchik/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserController реализует всю логику работы с User.
// Для этого используется вспомогательная модель (структура) из пакета "models.User"
type UserController struct{}

func (u *UserController) GetUserByID(id float64) *models.User {
	pool, err := pgxpool.New(context.Background(), *cfg.UrlConnectionStr)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}
	defer pool.Close()

	user := &models.User{}
	err = pool.QueryRow(context.Background(), "SELECT id, name, telegram_id, is_deleted FROM users WHERE id=$1", id).Scan(&user.Id, &user.Name, &user.TelegramId, &user.IsDeleted)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}

	return user
}

func (u *UserController) GetUserByTelegramID(id float64) *models.User {
	pool, err := pgxpool.New(context.Background(), *cfg.UrlConnectionStr)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}
	defer pool.Close()

	user := &models.User{}
	err = pool.QueryRow(context.TODO(), "SELECT id, name, telegram_id, is_deleted FROM users WHERE telegram_id=$1", id).Scan(&user.Id, &user.Name, &user.TelegramId, &user.IsDeleted)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}

	return user
}

func (u *UserController) CreateUser(user *models.User) *models.User {
	pool, err := pgxpool.New(context.Background(), *cfg.UrlConnectionStr)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}
	defer pool.Close()

	// Вносим основную информацию
	_, err = pool.Exec(context.Background(), "INSERT INTO users (name, telegram_id, is_deleted) VALUES ($1, $2, $3)", user.Name, user.TelegramId, user.IsDeleted)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}

	// Забираем из базы данных сгенерированный id
	err = pool.QueryRow(context.Background(), "SELECT id FROM users WHERE telegram_id=$1", user.TelegramId).Scan(&user.Id)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}

	return user
}

func (u *UserController) UpdateUser(user *models.User) *models.User {
	pool, err := pgxpool.New(context.Background(), *cfg.UrlConnectionStr)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), "UPDATE users SET name=$1, telegram_id=$2, is_deleted=$3 where id=$4", user.Name, user.TelegramId, user.IsDeleted, user.Id)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}

	mylog.INFO("Successfully updated user.")
	return user
}

func (u *UserController) DeleteUser(user *models.User) *models.User {
	pool, err := pgxpool.New(context.Background(), *cfg.UrlConnectionStr)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}
	defer pool.Close()
	_, err = pool.Exec(context.Background(), "UPDATE users SET is_deleted=true WHERE id=$1", user.Id)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}
	user.IsDeleted = true
	return user
}

func (u *UserController) RestoreUser(user *models.User) *models.User {
	pool, err := pgxpool.New(context.Background(), *cfg.UrlConnectionStr)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}
	defer pool.Close()
	_, err = pool.Exec(context.Background(), "UPDATE users SET is_deleted=false WHERE id=$1", user.Id)
	if err != nil {
		mylog.ERROR(err.Error())
		return nil
	}
	user.IsDeleted = false
	return user
}
