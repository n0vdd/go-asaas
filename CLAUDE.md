# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is **go-asaas**, a Go SDK for interacting with the Asaas payment gateway (Brazilian fintech). The library provides comprehensive coverage of Asaas API functionality including payments, subscriptions, customers, transfers, PIX operations, and more.

## Architecture

### Core Structure
- **`asaas/`** - Main package containing all SDK functionality
- **`internal/util/`** - Internal utilities for string manipulation, JSON handling, and reflection
- **`_example/`** - Example implementations for each feature
- **Environment Support**: Production (`EnvProduction`) and Sandbox (`EnvSandbox`) environments

### Key Components

#### Request Layer (`request.go`)
- Generic HTTP client with type safety using Go generics: `Request[T]`
- Handles both JSON and multipart form requests
- Automatic query parameter conversion for GET/DELETE requests
- Built-in logging and error handling
- Access token authentication via headers

#### Domain Objects
Each business domain has dedicated files:
- **Customer** (`customer.go`) - Customer management
- **Charge** (`charge.go`) - Payment charges (credit card, bank slip, PIX, undefined)
- **Subscription** (`subscription.go`) - Recurring payments
- **Transfer** (`transfer.go`) - Money transfers
- **PIX** (`pix.go`) - Brazilian instant payment system
- **Invoice** (`invoice.go`) - Fiscal invoices (NFSe)
- **Webhook** (`webhook.go`) - Event notifications
- And many more...

#### Type System
- **`enum.go`** - Contains all enumerated types (BillingType, ChargeStatus, etc.)
- **`date.go` & `datetime.go`** - Custom date/time handling for Asaas API formats
- **`error.go`** - Error response structures
- **`response.go`** - Base response types with success/failure methods

## Development Commands

### Testing
```bash
# Run all tests with coverage
go test -race -covermode atomic -coverprofile=covprofile ./asaas/...

# Run tests for specific package
go test ./asaas/

# Run specific test
go test -run TestCustomer ./asaas/
```

### Building
```bash
# Download dependencies
go mod download

# Build (if needed)
go build ./asaas/...

# Verify module integrity
go mod verify
```

### Environment Setup
Tests require Asaas API credentials via environment variables:
- `ASAAS_ACCESS_TOKEN` - Primary sandbox token
- `ASAAS_ACCESS_TOKEN_SECONDARY` - Secondary token for advanced features
- `ASAAS_WALLET_ID_SECONDARY` - Wallet ID for transfer operations

## Key Patterns

### Client Initialization
```go
// All services follow this pattern
customerService := asaas.NewCustomer(asaas.EnvSandbox, accessToken)
chargeService := asaas.NewCharge(asaas.EnvSandbox, accessToken)
```

### Response Handling
All responses implement standard methods:
- `.IsSuccess()` - Check if operation succeeded
- `.IsFailure()` - Check if operation failed
- `.IsNoContent()` - Check if no data returned
- `.Errors` - Array of error responses

### Request Context
All operations accept `context.Context` as first parameter for timeout/cancellation control.

## Testing Architecture

The test suite uses a sophisticated setup in `main_test.go`:
- **Environment-based testing** with automatic resource cleanup
- **Initialization functions** (`init*()`) create test data
- **Cleanup functions** (`clear*()`) remove test resources
- **TestMain()** orchestrates full cleanup on exit
- **Comprehensive test coverage** across all domains

Example test data flows:
1. `initCustomer()` creates test customer
2. `initCreditCardCharge()` creates charge for that customer  
3. `initAnticipation()` creates anticipation for that charge
4. Cleanup runs in reverse order

## Important Implementation Details

### Billing Types
- **Credit Card** - `BillingTypeCreditCard`
- **Bank Slip** - `BillingTypeBankSlip` 
- **PIX** - `BillingTypePix`
- **Undefined** - `BillingTypeUndefined` (customer chooses payment method)

### Date Handling
- Custom `Date` and `DateTime` types handle Asaas API format requirements
- Use `NewDate()` and `DateNow()` helper functions

### File Uploads
- Multipart form handling for document uploads
- Supports single files (`*os.File`) and file arrays (`[]*os.File`)

### Environment Switching
- `EnvSandbox` - Testing environment (default for development)
- `EnvProduction` - Live environment (requires production tokens)

This SDK emphasizes type safety, comprehensive testing, and follows Brazilian payment processing conventions.