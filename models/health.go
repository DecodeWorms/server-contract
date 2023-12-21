package models

type Health struct {
	Id                    string `json:"id" bson:"id"`
	PlayerId              string `json:"player_id" bson:"player_id"`
	DateInjured           string `json:"date_injured" bson:"date_injured"`
	DatePredictedToBeFine string `json:"date_predicted_to_be_fine" bson:"date_predicted_to_be_fine"`
	InjuryStatus          string `json:"injury_status" bson:"injury_status"`
	HealthStatus          string `json:"health_status" bson:"health_status"`
	InjuryTypes           string `json:"injury_types" bson:"injury_types"`
}
