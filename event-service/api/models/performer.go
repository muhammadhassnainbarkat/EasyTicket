package models

type Performer struct {
	ID   uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name string `json:"name"`
}
