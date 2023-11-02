package main

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kisukegremory/ninabank/internal/db"
)

func init() {
	ctx := context.Background()
	dbcon, err := sql.Open("mysql", "example_user:password@tcp(localhost:3306)/production")
	if err != nil {
		panic(err.Error())
	}
	queries := db.New(dbcon)
	defer dbcon.Close()

	result, err := queries.CreateAccount(ctx, db.CreateAccountParams{
		Name:     "Gustavo de Souza",
		Email:    "nina@gmail.com",
		Username: "nina.gustavo",
		Password: "123456",
	})

	if err != nil {
		panic(err)
	}
	lastId, _ := result.LastInsertId()

	slog.Info("Last inserted User %s", lastId)

}

func main() {
	router := gin.Default()
	router.POST("/v1/account/create")
	router.POST("/v1/account/login")
	router.POST("/v1/bank/deposit")
	router.POST("/v1/bank/withdraw")
	router.POST("/v1/bank/transfer")
	router.GET("/v1/bank/balance")
	router.GET("/v1/bank/statement")
	router.Run()
}
