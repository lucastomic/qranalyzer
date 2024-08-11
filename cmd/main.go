package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/lucastomic/qranalyzer/controller"
	"github.com/lucastomic/qranalyzer/database"
	"github.com/lucastomic/qranalyzer/domain"
	"github.com/lucastomic/qranalyzer/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error loading .env file")
	}
	db, err := database.GetGormDB()
	db.AutoMigrate(domain.Visitor{})
	if err != nil {
		fmt.Errorf("Error connecting to database")
	}
	repo := database.NewGormVisitorRepository(db)
	service := service.NewVisitorService(repo)
	visitorController := controller.NewVisitorController(service)
	viewController := controller.NewVisitorViewController(service)
	middlewareContorller := controller.NewMiddlewareController()
	r := http.NewServeMux()

	fmt.Println(http.Dir("/public"))
	fileServer := http.FileServer(http.Dir("public"))
	r.Handle("GET /public/", http.StripPrefix("/public/", fileServer))

	r.HandleFunc("POST /api/visitor", visitorController.Create)
	r.HandleFunc("GET /health", healthCheck)

	r.Handle("GET /visitors", http.HandlerFunc(viewController.Visitors))

	r.Handle("GET /", http.HandlerFunc(middlewareContorller.LoggerMiddleware))
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("I'm alive!")
}
