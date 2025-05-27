package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// ShortResponse: 짧은 키
type ShortResponse struct {
	A string `json:"a"`
	B int    `json:"b"`
	C bool   `json:"c"`
}

// LongResponse: 긴 키
type LongResponse struct {
	CustomerFullName           string `json:"customer_full_name"`
	TotalTransactionAmountUSD  int    `json:"total_transaction_amount_usd"`
	IsEligibleForPromotionalAd bool   `json:"is_eligible_for_promotional_ad"`
}

func handleShortJSON(c echo.Context) error {
	start := time.Now()
	resp := ShortResponse{
		A: "Alice",
		B: 12345,
		C: true,
	}
	elapsed := time.Since(start).Microseconds()
	c.Response().Header().Set("X-Elapsed-μs", fmt.Sprint(elapsed))
	return c.JSON(http.StatusOK, resp)
}

func handleLongJSON(c echo.Context) error {
	start := time.Now()
	resp := LongResponse{
		CustomerFullName:           "Alice Johnson",
		TotalTransactionAmountUSD:  12345,
		IsEligibleForPromotionalAd: true,
	}
	elapsed := time.Since(start).Microseconds()
	c.Response().Header().Set("X-Elapsed-μs", fmt.Sprint(elapsed))
	return c.JSON(http.StatusOK, resp)
}
