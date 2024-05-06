package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UUID     string `json:"uuid" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Job struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Requirements string `json:"requirements"`
	EmployerID   string `json:"employer_id"`
	Status       string `json:"status"`
}

type Application struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	JobID      uint   `json:"job_id"`
	TalentID   string `json:"talent_id"`
	EmployerID string `json:"employer_id"`
	Status     string `json:"status"`
}

type Claims struct {
	Username string `json:"username"`
	UUID     string `json:"uuid"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
