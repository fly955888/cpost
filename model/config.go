package model

type Config struct {
	Id   int `json:"id" gorm:"id"`
	Menu int `json:"menu" gorm:"menu"`
	Val  int `json:"val" gorm:"val"`
}
