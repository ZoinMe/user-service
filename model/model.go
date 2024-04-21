package model

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primary_key"`
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
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Count uint   `json:"count"`
}

type UserSkill struct {
	ID       uint `json:"id" gorm:"primary_key"`
	UserID   uint `json:"user_id"`
	SkillID  uint `json:"skill_id"`
	User     User `json:"user" gorm:"foreignkey:UserID"`
	Skill    Skill `json:"skill" gorm:"foreignkey:SkillID"`
}

type UserSocialMedia struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	UserID   uint   `json:"user_id"`
	URL      string `json:"url"`
	User     User   `json:"user" gorm:"foreignkey:UserID"`
}

type SocialMedia struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type Experience struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	CompanyLogo string    `json:"company_logo"`
	Designation string    `json:"designation"`
	Company     string    `json:"company"`
	FromDate    time.Time `json:"from_date"`
	ToDate      time.Time `json:"to_date"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"user" gorm:"foreignkey:UserID"`
}

type Education struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	UniversityLogo string    `json:"university_logo"`
	UniversityName string    `json:"university_name"`
	Degree         string    `json:"degree"`
	FromDate       time.Time `json:"from_date"`
	ToDate         time.Time `json:"to_date"`
	UserID         uint      `json:"user_id"`
	User           User      `json:"user" gorm:"foreignkey:UserID"`
}
