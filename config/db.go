// config/db.go

package config

import (
            "fmt"
            "os"
            "log"
            "database/sql"
            _ "github.com/jinzhu/gorm/dialects/postgres"
            _ "github.com/lib/pq"
)

var db *sql.DB

func SetupDB() *sql.DB {

    var (
            connectionName = mustGetenv("CLOUDSQL_CONNECTION")
            user           = mustGetenv("CLOUDSQL_USER")
            dbName         = os.Getenv("CLOUDSQL_DATABASE")
            password       = os.Getenv("CLOUDSQL_PASSWORD")
            socket         = os.Getenv("CLOUDSQL_SOCKET_PREFIX")
        )


    // /cloudsql is used on App Engine.
    if socket == "" {
        socket = "/cloudsql"
    }

    // connection string format: user=USER password=PASSWORD host=/cloudsql/PROJECT_ID:REGION_ID:INSTANCE_ID/[ dbname=DB_NAME]
    dbURI := fmt.Sprintf("user=%s password=%s host=/cloudsql/%s dbname=%s", user, password, connectionName, dbName)
    conn, err := sql.Open("postgres", dbURI)

    if err != nil {
    	fmt.Printf("Error: %s\n", err)
        panic("Failed to connect to database!")
    }

    return conn

}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
