// main.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/chiraggoti2016/go-web-service/config"
    "github.com/chiraggoti2016/go-web-service/routes"
    "github.com/gorilla/mux"
)

func main() {
    // Initialize database
    config.ConnectDB()

    // Initialize router
    r := mux.NewRouter()
    routes.RegisterBookRoutes(r)

    fmt.Println("Server is running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
