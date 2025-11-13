package models

type Team struct {
	ID     string
	Name   string
	TeamID string
	Users  []User
}
