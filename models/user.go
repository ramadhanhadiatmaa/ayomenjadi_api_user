package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Prov     string `json:"prov"`
	Kab      string `json:"kab"`
	Form     string `json:"form"`
	Res      int    `json:"res"`
	To1      int    `json:"to1"`
	To2      int    `json:"to2"`
	To3      int    `json:"to3"`
	To4      int    `json:"to4"`
	To5      int    `json:"to5"`
	To6      int    `json:"to6"`
	To7      int    `json:"to7"`
	To8      int    `json:"to8"`
	To9      int    `json:"to9"`
	To10     int    `json:"to10"`
}
