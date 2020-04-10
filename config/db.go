// config/db.go

package config

import (
            "fmt"
            "os"
            "log"
            "github.com/jinzhu/gorm"
            _ "github.com/jinzhu/gorm/dialects/postgres"
            _ "github.com/lib/pq"
)

var db *gorm.DB

func SetupDB() *gorm.DB {

    var (
            connectionName = "actup-273804:us-central1:postgres-actup" //mustGetenv("CLOUDSQL_CONNECTION")
            user           = "postgres" //mustGetenv("CLOUDSQL_USER")
            dbName         = "advocacy_db" //os.Getenv("CLOUDSQL_DATABASE")
            password       = "postgres" //os.Getenv("CLOUDSQL_PASSWORD")
            socket         = "cloudsql" //os.Getenv("CLOUDSQL_SOCKET_PREFIX")
            sslmode        = "disable"
        )



    // connection string format: user=USER password=PASSWORD host=/cloudsql/PROJECT_ID:REGION_ID:INSTANCE_ID/[ dbname=DB_NAME]
    //dbURI := fmt.Sprintf("user=%s password=%s host=/%s/%s dbname=%s", user, password, socket, connectionName, dbName)
    dbURI := fmt.Sprintf("user=%s password=%s host=/%s/%s dbname=%s sslmode=%s", user, password, socket, connectionName, dbName, sslmode)
    conn, err := gorm.Open("postgres", dbURI)

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
