package models

type PersonalInfo struct {
	Id            string `json:"id" bson:"id"`
	FirstName     string `json:"first_name" bson:"first_name"`
	LastName      string `json:"last_name" bson:"last_name"`
	Gender        string `json:"gender" bson:"gender"`
	MaritalStatus string `json:"marital_status" bson:"marital_status"`
	Email         string `json:"email" bson:"email"`
	PhoneNumber   string `json:"phone_number" bson:"phone_number"`
}

type Address struct {
	Name    string `json:"name" bson:"name"`
	ZipCode string `json:"zip_code" bson:"zip_code"`
	City    string `json:"city" bson:"city"`
}

type FieldInfo struct {
	Id                      string   `json:"id" bson:"id"`
	PlayerId                string   `json:"player_id" bson:"player_id"`
	YearOfExperience        string   `json:"year_of_experience" bson:"year_of_experience"`
	NumberOfGoalsScored     int      `json:"number_of_goals_scored" bson:"number_of_goals_scored"`
	ClubsHePreviouslyPlayed []string `json:"clubs_he_previously_played" bson:"clubs_he_previously_played"`
	JerseyNumber            int      `json:"jersey_number" bson:"jersey_number"`
	YearJoined              string   `json:"year_joined" bson:"year_joined"`
	PositionOnTheField      string   `json:"position_on_the_field" bson:"position_on_the_field"`
	PlayerClubStatus        string   `json:"player_club_status" bson:"player_club_status"`
}
