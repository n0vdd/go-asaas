package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/n0vdd/go-asaas/asaas"
)

func main() {
	// Get access token from environment
	accessToken := os.Getenv("ASAAS_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("ASAAS_ACCESS_TOKEN environment variable not set")
	}

	// Create charge service
	chargeService := asaas.NewCharge(asaas.EnvSandbox, accessToken)

	// Test date filtering with different statuses
	ctx := context.Background()

	// Create date range for testing
	startDate := asaas.Date(time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC))
	endDate := asaas.Date(time.Date(2025, 7, 31, 0, 0, 0, 0, time.UTC))

	fmt.Println("Testing date filtering with bracket parameters...")
	fmt.Printf("Date range: %s to %s\n\n", startDate, endDate)

	// Test with RECEIVED status
	fmt.Println("Testing with status=RECEIVED:")
	request := asaas.GetAllChargesRequest{
		Status:    asaas.ChargeStatusReceived,
		DueDateGe: &startDate,
		DueDateLe: &endDate,
		Limit:     10,
	}

	response, err := chargeService.GetAll(ctx, request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success! Found %d charges\n", response.TotalCount)
		if response.TotalCount > 0 && len(response.Data) > 0 {
			fmt.Printf("First charge ID: %s\n", response.Data[0].Id)
		}
	}

	// Test with CONFIRMED status
	fmt.Println("\nTesting with status=CONFIRMED:")
	request.Status = asaas.ChargeStatusConfirmed

	response, err = chargeService.GetAll(ctx, request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success! Found %d charges\n", response.TotalCount)
		if response.TotalCount > 0 && len(response.Data) > 0 {
			fmt.Printf("First charge ID: %s\n", response.Data[0].Id)
		}
	}

	// Test with PENDING status
	fmt.Println("\nTesting with status=PENDING:")
	request.Status = asaas.ChargeStatusPending

	response, err = chargeService.GetAll(ctx, request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success! Found %d charges\n", response.TotalCount)
		if response.TotalCount > 0 && len(response.Data) > 0 {
			fmt.Printf("First charge ID: %s\n", response.Data[0].Id)
		}
	}

	fmt.Println("\nDate filtering test complete!")
}
