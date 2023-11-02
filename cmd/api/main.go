package main

import "github.com/gin-gonic/gin"

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
