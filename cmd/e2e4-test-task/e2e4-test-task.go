package main

import (
	"fmt"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/exchange-rates", handleExchangeRates)

	err := e.Start(":8080")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
}
