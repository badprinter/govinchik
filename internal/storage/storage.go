package storage

import (
	"github.com/badprinter/govinchik/internal/config"
	"github.com/badprinter/govinchik/internal/mylogger"
)

// storage - пакет, реализует работу с базой данных.
// Основной структурой пакета является StorageManager.
// В этом файле определены основные структуры и функции
// которые могут понадобиться пользователю этого пакет

// Для работы с базой данных используется "pgxpool" из "https://github.com/jackc/pgx/v5"
// С поддержкой горутин.
// Для его установки используйте команду "go get github.com/jackc/pgx/v5/pgxpool"

// Глобальные переменные нужны чтобы к ним был доступ из всего пакета.
var (
	cfg   *config.StorageConfig = config.NewConfig().Storage
	mylog *mylogger.MyLogger    = mylogger.NewMyLogger()
)

// StorageManager основная структура для работы с пакетом из вне.
// В ней определенны вспомогательные структуры которые реализует
// какой-либо аспект или абстракцию с базой данной.
type StorageManager struct {
	userController *UserController
}

func NewStorageManager() *StorageManager {
	return &StorageManager{}

}
