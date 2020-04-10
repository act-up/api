// models/issue.go

package models

import (
            //"fmt"
            "github.com/jinzhu/gorm"
            //"database/sql/driver"
            "encoding/json"
            //"log"
            //"errors"
)

type Issue struct {
    ID uint `json:"id" gorm:"primary_key"`
    Issue json.RawMessage
    ContactInfo json.RawMessage
    WebformFormat json.RawMessage
}

/*type Issue struct {
    ID int
    IssueDescription IssueDescription
    ContactInfo ContactInfo
    WebformFormat WebformFormat
}

type IssueDescription struct {
    IssueName   string  `json:"issuename,omitempty"`
    IssueDescription string `json:"description,omitempty"`
    Salutation string `json:"salutation,omitempty"`
    Valediction string `json:"valediction,omitempty"`
    FormLetter string `json:"formletter,omitempty"`
    TwitterDMForm string `json:"twitterdmform,omitempty"`
}

type ContactInfo struct {
    Email string `json:"email,omitempty"`
}

type WebformFormat struct {
    Url string `json:"url,omitempty"`
}*/

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


//func GetAllIssues(db *sql.DB) (issues []Issue, err error) {

    /*var count int       // Number of rows in database

    // Count number of rows in database
    row := db.QueryRow("SELECT COUNT(*) FROM active_issues")
    row.Scan(&count)

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Get all rows in database
    rows, err := db.Query(`SELECT i.id, i.issue FROM active_issues i LIMIT $1;`, count)

    if err != nil {
		log.Fatal(err)
	}
    defer rows.Close()

    var active_issues []Issue

    for rows.Next() {
        err := rows.Scan(&id, &issue)

        if err != nil {
            log.Fatal(err)
        }

        active_issues = append(active_issues, Issue{ID: id, IssueDescription: issue})

    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }*/

//    return issues, err

//}


// Get an issue by ID
func GetAnIssue(db *gorm.DB, issue *Issue, id string) (err error) {
	if err := db.Where("id = ?", id).First(issue).Error; err != nil {
		return err
	}
	return nil
}
/*func GetAnIssue(db *sql.DB, issue *Issue, id int) (err error) {
    if id == 1 {
        fmt.Printf("id: ", id)
    }

    err = db.QueryRow("SELECT * FROM active_issues WHERE id=1", id).Scan(&issue.ID, &issue.IssueDescription)

    if err != nil {
        fmt.Printf("Error: %s\n", err)
	    log.Fatal(err)
    }
    defer db.Close()

    return err

}*/


// Return JSON-encoded representation of Issue struct
/*func (i Issue) Value() (val driver.Value, err error) {
    return json.Marshal(i)
}


// Decode JSON-encoded value into struct fields
func (i *Issue) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(b, &i)
}*/
