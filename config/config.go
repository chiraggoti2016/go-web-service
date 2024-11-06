// config/config.go
package config

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
    var err error
    dsn := "root:1234@tcp(127.0.0.1:3306)/bookdb" // Update with your MySQL credentials
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Database connected successfully!")
}
