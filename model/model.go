package model

import "time"

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Designation  string    `json:"designation,omitempty"`
	Bio          string    `json:"bio,omitempty"`
	ProfileImage string    `json:"profile_image,omitempty"`
	Location     string    `json:"location,omitempty"`
}

type Skill struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type UserSkill struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user_id"`
	SkillID int64 `json:"skill_id"`
}

type UserSocialMedia struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	URL    string `json:"url"`
}

type SocialMedia struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Experience struct {
	ID          int64     `json:"id"`
	CompanyLogo string    `json:"company_logo"`
	Designation string    `json:"designation"`
	Company     string    `json:"company"`
	FromDate    time.Time `json:"from_date"`
	ToDate      time.Time `json:"to_date"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
}

type Education struct {
	ID             int64     `json:"id"`
	UniversityLogo string    `json:"university_logo"`
	UniversityName string    `json:"university_name"`
	Degree         string    `json:"degree"`
	FromDate       time.Time `json:"from_date"`
	ToDate         time.Time `json:"to_date"`
	UserID         uint      `json:"user_id"`
}
