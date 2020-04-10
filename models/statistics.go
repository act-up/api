// models/statistics.go

package models

import (
            "github.com/jinzhu/gorm"
)

type Statistics struct {
    ID uint `json:"issue_id" gorm:"primary_key"`
    NumForms int `json:"num_forms"`
    NumEmails int `json:"num_emails"`
    NumTwitterDMs int `json:"num_twitter_dms"`
    Resolution json.RawMessage
}

// Return table name
func (i *Statistics) TableName() string {
	return "issue_statistics"
}

// Get all statistics of an issue by ID
func GetIssueStatistics(db *gorm.DB, stats *Statistics, id string) (err error) {
	if err := db.Where("id = ?", id).First(stats).Error; err != nil {
		return err
	}
	return nil
}


// Update statistics for an issue
func UpdateStatistics(db *gorm.DB, stats *Statistics, id string) (err error) {
	db.Save(stats)
	return nil
}
