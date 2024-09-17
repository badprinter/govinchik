package storage

import (
	"context"
	"github.com/badprinter/govinchik/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountController struct{ pool *pgxpool.Pool }

func (a *AccountController) GetAccaunt(ctx context.Context, u *models.User) *models.Account {
	var acc *models.Account
	err := a.pool.QueryRow(ctx, "SELECT id, username, bio_info, city, photo FROM account WHERE id = $1", u.Id).Scan(
		&acc.UserId, &acc.UserName, &acc.BioInfo, &acc.City, &acc.Photo)
	if err != nil {
		mylog.WARNING(err.Error())
		return nil
	}
	return acc
}
func (a *AccountController) CreateAccaunt(ctx context.Context, accaunt *models.Account) *models.Account {
	_, err := a.pool.Exec(ctx, "INSERT INTO account(id, username, bio_info, photo) VALUES ($1, $2, $3, $4)",
		accaunt.UserId, accaunt.UserName, accaunt.BioInfo, accaunt.Photo,
	)
	if err != nil {
		mylog.WARNING(err.Error())
		return nil
	}

	return accaunt // ХD
}
func (a *AccountController) UpdateAccaunt(ctx context.Context, accaunt *models.Account) *models.Account {
	_, err := a.pool.Exec(ctx, "UPDATE account SET username=$1, bio_info=$2, city=$3, photo=$4 WHERE id=$5",
		accaunt.UserName, accaunt.BioInfo, accaunt.City, accaunt.Photo, accaunt.UserId,
	)
	if err != nil {
		mylog.WARNING(err.Error())
		return nil
	}

	return accaunt
}
func (a *AccountController) DeleteAccaunt(ctx context.Context, accaunt *models.Account) *models.Account {
	_, err := a.pool.Exec(ctx, "DELETE FROM account WHERE id=$1", accaunt.UserId)
	if err != nil {
		mylog.WARNING(err.Error())
		return nil
	}
	return accaunt
}
