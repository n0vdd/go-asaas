# Agent Guidelines for the go-asaas Project

This document provides guidelines for AI agents working on the `go-asaas` codebase. Adhering to these conventions will help maintain code consistency and quality.

## 1. General Go Coding Conventions

*   **Formatting:** All Go code **MUST** be formatted using `gofmt` or `goimports` before committing.
*   **Error Handling:**
    *   Functions that can encounter errors **MUST** return an `error` as their last return value.
    *   Errors **MUST** be checked immediately after a function call. Do not ignore errors.
    *   Use `fmt.Errorf` with `w%` to wrap errors when adding context, or create custom error types if appropriate.
*   **Context:**
    *   `context.Context` **MUST** be the first parameter for functions that involve network requests, database calls, or other long-running/cancellable operations.
*   **Godoc Comments:**
    *   All public types, functions, and methods **MUST** have clear and comprehensive godoc comments.
    *   Comments should explain what the function does, its parameters, and its return values.
    *   For API client methods, document expected API responses and error conditions (see API Client Specifics).

## 2. Naming Conventions

*   **Packages:** `lowercase` (e.g., `asaas`, `util`).
*   **Files:** `snake_case.go` (e.g., `charge.go`, `charge_test.go`).
*   **Structs & Interfaces:** `CamelCase` (e.g., `CreateChargeRequest`, `ChargeResponse`, `Charge`).
*   **Methods & Public Functions:** `CamelCase` (e.g., `Create`, `GetById`).
*   **Variables (local & private struct fields):** `camelCase` (e.g., `accessToken`, `chargeId`).
*   **Constants/Enums:** `CamelCase` (e.g., `BillingTypeUndefined`, `EnvSandbox`).

## 3. API Client Specifics

This library is a client for the Asaas API.

*   **Request/Response Structs:**
    *   Request structs should generally end with the suffix `Request` (e.g., `CreateChargeRequest`).
    *   Response structs should generally end with the suffix `Response` (e.g., `ChargeResponse`, `DeleteResponse`).
    *   Struct fields **MUST** be annotated with `json:"fieldName,omitempty"` tags for correct JSON serialization/deserialization. `omitempty` should be used for optional fields.
    *   For fields in request objects that are optional and where a zero value (e.g., empty string, 0, false) is a valid value that needs to be distinguished from an unset field, use pointers (e.g., `Description *string`).
*   **Service Interfaces:**
    *   Each Asaas API resource (e.g., Charges, Customers) **SHOULD** have a corresponding Go interface (e.g., `Charge`, `Customer`).
    *   These interfaces define the available operations for that resource.
*   **Constructor Functions:**
    *   Provide a public constructor function for each service interface, typically named `New<ServiceName>(env Env, accessToken string) <ServiceInterface>` (e.g., `NewCharge(env Env, accessToken string) Charge`).
    *   This function initializes the service client with the environment (e.g., sandbox, production) and the API access token.
*   **Method Signatures:**
    *   Methods that perform API calls **MUST** accept `context.Context` as their first argument.
    *   Methods that create, retrieve, or update a single resource typically return `(*<ResponseStruct>, error)`.
    *   Methods that delete a resource typically return `(*DeleteResponse, error)`. The `DeleteResponse` struct usually contains information about whether the deletion was successful or if errors occurred.
    *   Methods that list resources (collections) typically return `(*Pageable[<ResponseStruct>], error)`, where `Pageable` is a generic struct handling pagination.
*   **Documentation Comments for API Methods:**
    *   Godoc comments for API methods **MUST** be detailed.
    *   **Response Handling:** Clearly describe the expected structure of the response struct on success.
    *   **Error Conditions:**
        *   Explain how API errors are represented (e.g., within the `Errors` field of the response struct like `ChargeResponse.Errors`).
        *   Distinguish between errors returned by the library itself (network issues, marshalling problems - where the `error` return value is non-nil) and API-level errors (where `error` might be nil, but the response struct indicates a failure, e.g., `response.IsFailure()`).
        *   List common HTTP status codes the API might return and how they are handled or can be interpreted by the caller.
    *   **API Docs Link:** Include a direct link to the relevant page in the official Asaas API documentation using a `# DOCS` heading (e.g., `# DOCS\n\nListar cobranças: https://docs.asaas.com/reference/listar-cobrancas`).

## 4. Testing

*   **Test Files:** Test files **MUST** be named `<filename>_test.go` (e.g., `charge_test.go`).
*   **Test Functions:** Test functions **MUST** be named `Test<StructName><MethodName>` or `Test<FunctionName>` (e.g., `TestChargeCreate`).
*   **Coverage:** Strive for high test coverage. New features or bug fixes **MUST** include corresponding tests.
*   **Setup & Preconditions:**
    *   Use helper functions (e.g., `initCustomer()`, `initCreditCardCharge()`) to set up necessary preconditions or test data. These are often found in `main_test.go` or specific `_test.go` files.
    *   Tests often rely on environment variables for configuration (e.g., API tokens, entity IDs for testing). Refer to existing tests for patterns.
*   **Context with Timeout:** Tests involving network requests **MUST** use `context.WithTimeout` to prevent tests from hanging indefinitely.
*   **Assertions:** Use the provided assertion helpers (e.g., `assertResponseSuccess(t, resp, err)`) for common checks. Add new assertion helpers if needed for new common patterns.
*   **Real API Calls:** Tests in this project appear to make real calls to a sandbox environment. Ensure tests clean up created resources if possible, or use specific test accounts/data that can be reset.

## 5. Dependencies and Tooling

*   **Go Modules:** Dependencies **MUST** be managed using Go Modules. Update `go.mod` and `go.sum` as necessary (`go get`, `go mod tidy`).
*   **Go Version:** Refer to the `go.mod` file for the minimum Go version required.
*   **Linters:** While not explicitly stated, assume standard Go linters (like `golangci-lint`) might be used. Write clean, idiomatic Go.

By following these guidelines, you will help ensure that the `go-asaas` library remains maintainable, robust, and easy to use.
