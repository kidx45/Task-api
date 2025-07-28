package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
	"task-api/internal/adapter/inbound/http"
	"task-api/internal/adapter/outbound"
	"task-api/internal/application"
	
)

func main() {
	dsn := "root:manyahle1234$@@tcp(127.0.0.1:3306)/taskdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("DB unreachable: %v", err)
	}

	repo := persistence.NewMySQLRepo(db)
	svc := service.NewTaskService(repo)

	router := gin.Default()
	http.NewTaskHandler(router, svc)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
