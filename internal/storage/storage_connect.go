package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *StorageAPI) Ping() error {
	conn, err := pgxpool.New(context.Background(), *s.cfg.UrlConnectionStr)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.Ping(context.Background())
	if err != nil {
		return err
	}

	return nil
}
