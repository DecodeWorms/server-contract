package models

import "gorm.io/gorm"

type PersonalInfo struct {
	gorm.Model
	Id            string  `json:"id" gorm:"id"`
	FirstName     string  `json:"first_name" gorm:"first_name"`
	LastName      string  `json:"last_name" gorm:"last_name"`
	Gender        string  `json:"gender" gorm:"gender"`
	MaritalStatus string  `json:"marital_status" gorm:"marital_status"`
	Email         string  `json:"email" gorm:"email"`
	PhoneNumber   string  `json:"phone_number" gorm:"phone_number"`
	Address       Address `json:"address" gorm:"address"`
}

type Address struct {
	gorm.Model
	PersonalInfoId string `json:"personal_info_id" gorm:"personal_info_id"`
	Name           string `json:"name" gorm:"name"`
	ZipCode        string `json:"zip_code" gorm:"zip_code"`
	City           string `json:"city" gorm:"city"`
}

type FieldInfo struct {
	gorm.Model
	Id                  string `json:"id" gorm:"id"`
	PersonalInfoId      string `json:"player_ifo_id" gorm:"player_ifo_id"`
	YearOfExperience    string `json:"year_of_experience" gorm:"year_of_experience"`
	NumberOfGoalsScored int    `json:"number_of_goals_scored" gorm:"number_of_goals_scored"`
	JerseyNumber        int    `json:"jersey_number" gorm:"jersey_number"`
	YearJoined          string `json:"year_joined" gorm:"year_joined"`
	PositionOnTheField  string `json:"position_on_the_field" gorm:"position_on_the_field"`
	PlayerClubStatus    string `json:"player_club_status" gorm:"player_club_status"`
}

type ClubsHePreviouslyPlayed struct {
	gorm.Model
	PersonalInfoId string `json:"personal_info_id" gorm:"personal_info_id"`
	ClubName       string `json:"club_name", gorm:"club_name"`
	StartedYear    string `json:"started_year" gorm:"started_year"`
	EndedYear      string `json:"ended_year" gorm:"ended_year"`
}
