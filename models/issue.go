// models/issue.go

package models

import (
            "github.com/jinzhu/gorm"
            "encoding/json"
)

type Issue struct {
    ID uint `json:"id" gorm:"primary_key"`
    Issue json.RawMessage
    ContactInfo json.RawMessage
    WebformFormat json.RawMessage
}

// Return table name
func (i *Issue) TableName() string {
	return "active_issues"
}

// Get all issues in database
func GetAllIssues(db *gorm.DB, issues *[]Issue) (err error) {
	if err = db.Find(issues).Error; err != nil {
		return err
	}
	return nil
}

// Get an issue
func GetAnIssue(db *gorm.DB, issue *Issue, id string) (err error) {
	if err := db.Where("id = ?", id).First(issue).Error; err != nil {
		return err
	}
	return nil
}
