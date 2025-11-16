package models

type ReviewerStats struct {
	UserID     string `gorm:"column:user_id"`
	Username   string `gorm:"column:username"`
	AmountOfPR int    `gorm:"column:amount_of_pr"`
}
