package models

type Account struct {
	UserId   *int64
	UserName string
	BioInfo  string
	City     string
	Photo    []byte
}
