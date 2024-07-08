package model

import (
	"time"
)

type User struct {
	ID           string `json:"id"` // UUID for the user.
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"-"` // This field is excluded from JSON serialization.
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Designation  string    `json:"designation,omitempty"`
	Bio          string    `json:"bio,omitempty"`
	ProfileImage string    `json:"profile_image,omitempty"`
	Location     string    `json:"location,omitempty"`
	GitHubURL    string    `json:"github_url,omitempty"`   // GitHub URL for the user.
	LinkedInURL  string    `json:"linkedin_url,omitempty"` // LinkedIn URL for the user.
}

type Skill struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type UserSkill struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	SkillID int64 `json:"skill_id"`
}

type Experience struct {
	ID          string     `json:"id"`
	CompanyLogo string    `json:"company_logo"`
	Designation string    `json:"designation"`
	Company     string    `json:"company"`
	FromDate    time.Time `json:"from_date"`
	ToDate      time.Time `json:"to_date"`
	Description string    `json:"description"`
	UserID      string      `json:"user_id"`
}

type Education struct {
	ID             string     `json:"id"`
	UniversityLogo string    `json:"university_logo"`
	UniversityName string    `json:"university_name"`
	Degree         string    `json:"degree"`
	FromDate       time.Time `json:"from_date"`
	ToDate         time.Time `json:"to_date"`
	UserID         string      `json:"user_id"`
}

type Notification struct {
	ID      string
	UserID  string
	Message string
	Type    string
}
