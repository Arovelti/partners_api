package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"partners_api/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	startDB()

	r := gin.Default()

	handlers := handlers.PartnerHandler{}

	r.POST("/partners", handlers.CreatePartnerHandler)
	r.PUT("/partners/:id", handlers.SetPartnerStatusHandler)
	r.GET("/partners/:id", handlers.GetPartnerByIDHandler)
	r.GET("/partners", handlers.GetAllPartners)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Ыerver gracefully shutdown")
}

func startDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres dbname=partners sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
