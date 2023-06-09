package models

import "gorm.io/gorm"

type Salary struct {
	gorm.Model
	Id           string `json:"id" bson:"id"`
	PlayerId     string `json:"player_id" bson:"player_id"`
	WeeklySalary int    `json:"weekly_salary" bson:"weekly_salary"`
	Bonus        int    `json:"bonus" bson:"bonus"`
	SalaryStatus string `json:"salary_status" bson:"salary_status"`
	BonusStatus  string `json:"bonus_status" bson:"bonus_status"`
}
