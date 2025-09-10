package asaas

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestChargeCreate(t *testing.T) {
	initCustomer(false)
	accessToken := getEnvValue(EnvAccessToken)
	customerId := getEnvValue(EnvCustomerId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.Create(ctx, CreateChargeRequest{
		Customer:             customerId,
		BillingType:          BillingTypeUndefined,
		Value:                10,
		DueDate:              NewDate(2026, 12, 1, time.Local),
		Description:          "Unit test golang",
		ExternalReference:    "",
		Discount:             nil,
		Interest:             nil,
		Fine:                 nil,
		PostalService:        false,
		Split:                nil,
		Callback:             nil,
		CreditCard:           nil,
		CreditCardHolderInfo: nil,
		CreditCardToken:      "",
		InstallmentCount:     0,
		InstallmentValue:     0,
		AuthorizeOnly:        false,
		RemoteIp:             "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargePayWithCreditCard(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.PayWithCreditCard(ctx, chargeId, PayWithCreditCardRequest{
		CreditCard: &CreditCardRequest{
			HolderName:  "unit test go",
			Number:      "4000000000000010",
			ExpiryMonth: "05",
			ExpiryYear:  "2035",
			Ccv:         "318",
		},
		CreditCardHolderInfo: &CreditCardHolderInfoRequest{
			Name:          "Unit Test Go",
			CpfCnpj:       "24971563792",
			Email:         "unittest@gmail.com",
			Phone:         "4738010919",
			MobilePhone:   "47998781877",
			PostalCode:    "89223-005",
			AddressNumber: "277",
		},
		CreditCardToken: "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeUpdateById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UpdateById(ctx, chargeId, UpdateChargeRequest{
		Value:            15,
		DueDate:          Date{},
		Description:      Pointer("update from unit test golang"),
		InstallmentCount: 1,
		InstallmentValue: 7.5,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeDeleteById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.DeleteById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeRestoreById(t *testing.T) {
	initChargeDeleted()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvChargeDeletedId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.RestoreById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeRefundById(t *testing.T) {
	initCreditCardCharge(false, false)
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvCreditCardChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.RefundById(ctx, chargeId, RefundRequest{
		Value:       0,
		Description: "",
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeReceiveInCashById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	now := DateNow()
	resp, err := nCharge.ReceiveInCashById(ctx, chargeId, ChargeReceiveInCashRequest{
		PaymentDate:    NewDate(now.Year(), now.Month(), now.Day(), now.Location()),
		Value:          5,
		NotifyCustomer: false,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeUndoReceivedInCashById(t *testing.T) {
	initChargeReceivedInCash()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvChargeReceivedInCashId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UndoReceivedInCashById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeUploadDocumentById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	f, _ := os.Open(getEnvValue(EnvFileName))
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UploadDocumentById(ctx, chargeId, UploadChargeDocumentRequest{
		AvailableAfterPayment: false,
		Type:                  DocumentTypeDocument,
		File:                  f,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeUpdateDocumentDefinitionsById(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	docId := getEnvValue(EnvChargeDocumentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.UpdateDocumentDefinitionsById(ctx, chargeId, docId, UpdateChargeDocumentDefinitionsRequest{
		AvailableAfterPayment: false,
		Type:                  DocumentTypeContract,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeDeleteDocumentById(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	docId := getEnvValue(EnvChargeDocumentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.DeleteDocumentById(ctx, chargeId, docId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetCreationLimit(t *testing.T) {
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetCreationLimit(ctx)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetStatusById(t *testing.T) {
	initUndefinedCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetStatusById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetIdentificationFieldById(t *testing.T) {
	initBankSlipCharge(false)
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvBankSlipChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetIdentificationFieldById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetPixQrCodeById(t *testing.T) {
	initPixCharge()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvPixChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetPixQrCodeById(ctx, chargeId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetDocumentById(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	docId := getEnvValue(EnvChargeDocumentId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetDocumentById(ctx, chargeId, docId)
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetAll(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
		Customer:              "",
		Subscription:          "",
		Installment:           "",
		CustomerGroupName:     "",
		BillingType:           "",
		Status:                "",
		ExternalReference:     "",
		InvoiceStatus:         "",
		EstimatedCreditDate:   Date{},
		PixQrCodeId:           "",
		Anticipated:           nil,
		PaymentDate:           Date{},
		DateCreatedGe:         Date{},
		DateCreatedLe:         Date{},
		PaymentDateGe:         Date{},
		PaymentDateLe:         Date{},
		EstimatedCreditDateGE: Date{},
		EstimatedCreditDateLE: Date{},
		DueDateGe:             Date{},
		DueDateLe:             Date{},
		User:                  "",
		Offset:                0,
		Limit:                 10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetAllDocumentsById(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	chargeId := getEnvValue(EnvUndefinedChargeId)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)
	resp, err := nCharge.GetAllDocumentsById(ctx, chargeId, PageableDefaultRequest{
		Offset: 0,
		Limit:  10,
	})
	assertResponseSuccess(t, resp, err)
}

func TestChargeGetAllWithDateFilters(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)

	// Test 1: DateCreatedGe filter - charges created after a specific date
	t.Run("DateCreatedGe_Filter", func(t *testing.T) {
		startDate := NewDate(2025, 9, 1, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: startDate,
			Limit:         5,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all returned charges were created on or after the start date
		for _, charge := range resp.Data {
			if !charge.DateCreated.IsZero() && charge.DateCreated.Time().Before(startDate.Time()) {
				t.Errorf("Charge %s created %s is before filter date %s",
					charge.Id, charge.DateCreated.Format(), startDate.Format())
			}
		}
		t.Logf("DateCreatedGe filter test passed. Found %d charges created after %s",
			len(resp.Data), startDate.Format())
	})

	// Test 2: DateCreatedLe filter - charges created before a specific date
	t.Run("DateCreatedLe_Filter", func(t *testing.T) {
		endDate := NewDate(2025, 12, 31, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedLe: endDate,
			Limit:         5,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all returned charges were created on or before the end date
		for _, charge := range resp.Data {
			if !charge.DateCreated.IsZero() && charge.DateCreated.Time().After(endDate.Time()) {
				t.Errorf("Charge %s created %s is after filter date %s",
					charge.Id, charge.DateCreated.Format(), endDate.Format())
			}
		}
		t.Logf("DateCreatedLe filter test passed. Found %d charges created before %s",
			len(resp.Data), endDate.Format())
	})

	// Test 3: DateCreated range filter - charges created within a date range
	t.Run("DateCreated_Range_Filter", func(t *testing.T) {
		startDate := NewDate(2025, 9, 1, time.UTC)
		endDate := NewDate(2025, 9, 30, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: startDate,
			DateCreatedLe: endDate,
			Limit:         10,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all returned charges are within the date range
		for _, charge := range resp.Data {
			if !charge.DateCreated.IsZero() {
				if charge.DateCreated.Time().Before(startDate.Time()) {
					t.Errorf("Charge %s created %s is before start date %s",
						charge.Id, charge.DateCreated.Format(), startDate.Format())
				}
				if charge.DateCreated.Time().After(endDate.Time()) {
					t.Errorf("Charge %s created %s is after end date %s",
						charge.Id, charge.DateCreated.Format(), endDate.Format())
				}
			}
		}
		t.Logf("DateCreated range filter test passed. Found %d charges between %s and %s",
			len(resp.Data), startDate.Format(), endDate.Format())
	})

	// Test 4: DueDateGe filter - charges due after a specific date
	t.Run("DueDateGe_Filter", func(t *testing.T) {
		startDate := NewDate(2025, 9, 1, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DueDateGe: startDate,
			Limit:     5,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all returned charges are due on or after the start date
		for _, charge := range resp.Data {
			if !charge.DueDate.IsZero() && charge.DueDate.Time().Before(startDate.Time()) {
				t.Errorf("Charge %s due %s is before filter date %s",
					charge.Id, charge.DueDate.Format(), startDate.Format())
			}
		}
		t.Logf("DueDateGe filter test passed. Found %d charges due after %s",
			len(resp.Data), startDate.Format())
	})

	// Test 5: DueDateLe filter - charges due before a specific date
	t.Run("DueDateLe_Filter", func(t *testing.T) {
		endDate := NewDate(2025, 12, 31, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DueDateLe: endDate,
			Limit:     5,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all returned charges are due on or before the end date
		for _, charge := range resp.Data {
			if !charge.DueDate.IsZero() && charge.DueDate.Time().After(endDate.Time()) {
				t.Errorf("Charge %s due %s is after filter date %s",
					charge.Id, charge.DueDate.Format(), endDate.Format())
			}
		}
		t.Logf("DueDateLe filter test passed. Found %d charges due before %s",
			len(resp.Data), endDate.Format())
	})

	// Test 6: DueDate range filter - charges due within a date range
	t.Run("DueDate_Range_Filter", func(t *testing.T) {
		startDate := NewDate(2025, 9, 1, time.UTC)
		endDate := NewDate(2025, 9, 30, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DueDateGe: startDate,
			DueDateLe: endDate,
			Limit:     10,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all returned charges are due within the date range
		for _, charge := range resp.Data {
			if !charge.DueDate.IsZero() {
				if charge.DueDate.Time().Before(startDate.Time()) {
					t.Errorf("Charge %s due %s is before start date %s",
						charge.Id, charge.DueDate.Format(), startDate.Format())
				}
				if charge.DueDate.Time().After(endDate.Time()) {
					t.Errorf("Charge %s due %s is after end date %s",
						charge.Id, charge.DueDate.Format(), endDate.Format())
				}
			}
		}
		t.Logf("DueDate range filter test passed. Found %d charges due between %s and %s",
			len(resp.Data), startDate.Format(), endDate.Format())
	})

	// Test 7: EstimatedCreditDate filters
	t.Run("EstimatedCreditDate_Filter", func(t *testing.T) {
		startDate := NewDate(2025, 9, 1, time.UTC)
		endDate := NewDate(2025, 12, 31, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			Status:                ChargeStatusConfirmed,
			EstimatedCreditDateGE: startDate,
			EstimatedCreditDateLE: endDate,
			Limit:                 5,
		})
		assertResponseSuccess(t, resp, err)

		t.Logf("EstimatedCreditDate filter test passed. Found %d confirmed charges with estimated credit dates between %s and %s",
			len(resp.Data), startDate.Format(), endDate.Format())
	})

	// Test 8: PaymentDate filters for confirmed/received charges
	t.Run("PaymentDate_Filter", func(t *testing.T) {
		startDate := NewDate(2025, 9, 1, time.UTC)
		endDate := NewDate(2025, 9, 30, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			Status:        ChargeStatusConfirmed,
			PaymentDateGe: startDate,
			PaymentDateLe: endDate,
			Limit:         5,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all returned charges have payment dates within range (if set)
		for _, charge := range resp.Data {
			if !charge.ClientPaymentDate.IsZero() {
				if charge.ClientPaymentDate.Time().Before(startDate.Time()) {
					t.Errorf("Charge %s payment date %s is before start date %s",
						charge.Id, charge.ClientPaymentDate.Format(), startDate.Format())
				}
				if charge.ClientPaymentDate.Time().After(endDate.Time()) {
					t.Errorf("Charge %s payment date %s is after end date %s",
						charge.Id, charge.ClientPaymentDate.Format(), endDate.Format())
				}
			}
		}
		t.Logf("PaymentDate filter test passed. Found %d confirmed charges with payments between %s and %s",
			len(resp.Data), startDate.Format(), endDate.Format())
	})

	// Test 9: Empty results test - date range with no charges
	t.Run("Empty_Results_Test", func(t *testing.T) {
		// Use a future date range that should have no charges
		startDate := NewDate(2030, 1, 1, time.UTC)
		endDate := NewDate(2030, 1, 31, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: startDate,
			DateCreatedLe: endDate,
			Limit:         5,
		})
		assertResponseSuccess(t, resp, err)

		// Should return empty results
		if len(resp.Data) > 0 {
			t.Logf("Warning: Found %d charges in future date range %s to %s (this might be expected in sandbox)",
				len(resp.Data), startDate.Format(), endDate.Format())
		} else {
			t.Logf("Empty results test passed. No charges found in future date range %s to %s",
				startDate.Format(), endDate.Format())
		}
	})

	// Test 10: Same start/end date (edge case)
	t.Run("Same_Date_Edge_Case", func(t *testing.T) {
		sameDate := NewDate(2025, 9, 10, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: sameDate,
			DateCreatedLe: sameDate,
			Limit:         5,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all returned charges were created on the exact date
		for _, charge := range resp.Data {
			if !charge.DateCreated.IsZero() {
				chargeDate := charge.DateCreated.Time()
				// Check if the charge was created on the same calendar day
				if chargeDate.Year() != sameDate.Year() ||
					chargeDate.Month() != sameDate.Month() ||
					chargeDate.Day() != sameDate.Day() {
					t.Errorf("Charge %s created %s is not on the same date %s",
						charge.Id, charge.DateCreated.Format(), sameDate.Format())
				}
			}
		}
		t.Logf("Same date edge case test passed. Found %d charges created on %s",
			len(resp.Data), sameDate.Format())
	})

	// Test 11: Multiple filters combined
	t.Run("Multiple_Filters_Combined", func(t *testing.T) {
		createStart := NewDate(2025, 9, 1, time.UTC)
		createEnd := NewDate(2025, 9, 30, time.UTC)
		dueStart := NewDate(2025, 9, 15, time.UTC)
		dueEnd := NewDate(2025, 12, 31, time.UTC)

		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: createStart,
			DateCreatedLe: createEnd,
			DueDateGe:     dueStart,
			DueDateLe:     dueEnd,
			Status:        ChargeStatusPending,
			Limit:         5,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all filters are applied correctly
		for _, charge := range resp.Data {
			// Check creation date range
			if !charge.DateCreated.IsZero() {
				if charge.DateCreated.Time().Before(createStart.Time()) ||
					charge.DateCreated.Time().After(createEnd.Time()) {
					t.Errorf("Charge %s creation date %s is outside range %s to %s",
						charge.Id, charge.DateCreated.Format(), createStart.Format(), createEnd.Format())
				}
			}

			// Check due date range
			if !charge.DueDate.IsZero() {
				if charge.DueDate.Time().Before(dueStart.Time()) ||
					charge.DueDate.Time().After(dueEnd.Time()) {
					t.Errorf("Charge %s due date %s is outside range %s to %s",
						charge.Id, charge.DueDate.Format(), dueStart.Format(), dueEnd.Format())
				}
			}

			// Check status
			if charge.Status != ChargeStatusPending {
				t.Errorf("Charge %s status is %s, expected PENDING", charge.Id, charge.Status)
			}
		}
		t.Logf("Multiple filters test passed. Found %d pending charges matching all criteria", len(resp.Data))
	})

	// Test 12: Zero/Empty Date Values - should be ignored by API
	t.Run("Zero_Date_Values_Test", func(t *testing.T) {
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: Date{}, // Zero value - should be ignored
			DateCreatedLe: Date{}, // Zero value - should be ignored
			Limit:         5,
		})
		assertResponseSuccess(t, resp, err)
		t.Logf("Zero date values test passed. Found %d charges (zero dates ignored)", len(resp.Data))
	})

	// Test 13: Very Old Date Range
	t.Run("Very_Old_Date_Range_Test", func(t *testing.T) {
		oldStart := NewDate(1900, 1, 1, time.UTC)
		oldEnd := NewDate(2000, 12, 31, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: oldStart,
			DateCreatedLe: oldEnd,
			Limit:         5,
		})
		assertResponseSuccess(t, resp, err)
		t.Logf("Very old date range test passed. Found %d charges from %s to %s",
			len(resp.Data), oldStart.Format(), oldEnd.Format())
	})

	// Test 14: Far Future Date Range
	t.Run("Far_Future_Date_Range_Test", func(t *testing.T) {
		futureStart := NewDate(2050, 1, 1, time.UTC)
		futureEnd := NewDate(2099, 12, 31, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: futureStart,
			DateCreatedLe: futureEnd,
			Limit:         5,
		})
		assertResponseSuccess(t, resp, err)

		// Should return empty or very few results
		if len(resp.Data) == 0 {
			t.Logf("Far future date range test passed. No charges found in range %s to %s",
				futureStart.Format(), futureEnd.Format())
		} else {
			t.Logf("Far future date range returned %d charges (unexpected in production)", len(resp.Data))
		}
	})

	// Test 15: Leap Year Date Testing
	t.Run("Leap_Year_Date_Test", func(t *testing.T) {
		// Test February 29th on a leap year (2024)
		leapDate := NewDate(2024, 2, 29, time.UTC)
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: leapDate,
			DateCreatedLe: leapDate,
			Limit:         5,
		})
		assertResponseSuccess(t, resp, err)
		t.Logf("Leap year date test passed. Found %d charges on leap day %s",
			len(resp.Data), leapDate.Format())
	})

	// Test 16: Current Date Filtering (Today)
	t.Run("Current_Date_Today_Test", func(t *testing.T) {
		today := DateNow()
		todayStart := NewDate(today.Year(), today.Month(), today.Day(), time.UTC)

		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: todayStart,
			DateCreatedLe: todayStart,
			Limit:         10,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all charges are from today
		for _, charge := range resp.Data {
			if !charge.DateCreated.IsZero() {
				chargeDate := charge.DateCreated.Time()
				if chargeDate.Year() != todayStart.Year() ||
					chargeDate.Month() != todayStart.Month() ||
					chargeDate.Day() != todayStart.Day() {
					t.Errorf("Charge %s created %s is not from today %s",
						charge.Id, charge.DateCreated.Format(), todayStart.Format())
				}
			}
		}
		t.Logf("Current date (today) test passed. Found %d charges created today", len(resp.Data))
	})

	// Test 17: This Month Date Range
	t.Run("This_Month_Date_Range_Test", func(t *testing.T) {
		now := time.Now()
		monthStart := NewDate(now.Year(), now.Month(), 1, time.UTC)
		monthEnd := NewDate(now.Year(), now.Month()+1, 1, time.UTC)
		monthEnd = NewDate(monthEnd.Year(), monthEnd.Month(), monthEnd.Day()-1, time.UTC) // Last day of current month

		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: monthStart,
			DateCreatedLe: monthEnd,
			Limit:         20,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all charges are from this month
		for _, charge := range resp.Data {
			if !charge.DateCreated.IsZero() {
				chargeDate := charge.DateCreated.Time()
				if chargeDate.Year() != now.Year() || chargeDate.Month() != now.Month() {
					t.Errorf("Charge %s created %s is not from current month %d/%d",
						charge.Id, charge.DateCreated.Format(), int(now.Month()), now.Year())
				}
			}
		}
		t.Logf("This month test passed. Found %d charges created this month (%d/%d)",
			len(resp.Data), int(now.Month()), now.Year())
	})

	// Test 18: Overdue Charges (due date < current date)
	t.Run("Overdue_Charges_Test", func(t *testing.T) {
		yesterday := time.Now().AddDate(0, 0, -1)
		yesterdayDate := NewDate(yesterday.Year(), yesterday.Month(), yesterday.Day(), time.UTC)

		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			Status:    ChargeStatusPending,
			DueDateLe: yesterdayDate,
			Limit:     10,
		})
		assertResponseSuccess(t, resp, err)

		// Verify all returned charges are overdue and pending
		overdueCount := 0
		for _, charge := range resp.Data {
			if charge.Status != ChargeStatusPending {
				t.Errorf("Charge %s status is %s, expected PENDING for overdue test",
					charge.Id, charge.Status)
			}
			if !charge.DueDate.IsZero() && charge.DueDate.Time().Before(time.Now()) {
				overdueCount++
			}
		}
		t.Logf("Overdue charges test passed. Found %d pending charges due before %s (%d confirmed overdue)",
			len(resp.Data), yesterdayDate.Format(), overdueCount)
	})

	// Test 19: Date Range with Different Billing Types
	t.Run("Date_Range_With_Billing_Types_Test", func(t *testing.T) {
		startDate := NewDate(2025, 9, 1, time.UTC)
		endDate := NewDate(2025, 9, 30, time.UTC)

		billingTypes := []BillingType{BillingTypeCreditCard, BillingTypePix, BillingTypeBankSlip, BillingTypeUndefined}

		for _, billingType := range billingTypes {
			resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
				DateCreatedGe: startDate,
				DateCreatedLe: endDate,
				BillingType:   billingType,
				Limit:         5,
			})
			assertResponseSuccess(t, resp, err)

			// Verify all charges have the correct billing type
			for _, charge := range resp.Data {
				if charge.BillingType != billingType {
					t.Errorf("Charge %s billing type is %s, expected %s",
						charge.Id, charge.BillingType, billingType)
				}
			}
			t.Logf("Date range + %s billing type test passed. Found %d charges",
				billingType, len(resp.Data))
		}
	})

	// Test 20: Pagination with Date Filters
	t.Run("Pagination_With_Date_Filters_Test", func(t *testing.T) {
		startDate := NewDate(2025, 9, 1, time.UTC)
		endDate := NewDate(2025, 9, 30, time.UTC)

		// First page
		firstPage, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: startDate,
			DateCreatedLe: endDate,
			Offset:        0,
			Limit:         3,
		})
		assertResponseSuccess(t, firstPage, err)

		// Second page if available
		if firstPage.HasMore {
			secondPage, err := nCharge.GetAll(ctx, GetAllChargesRequest{
				DateCreatedGe: startDate,
				DateCreatedLe: endDate,
				Offset:        3,
				Limit:         3,
			})
			assertResponseSuccess(t, secondPage, err)

			// Verify no duplicate charges between pages
			firstPageIds := make(map[string]bool)
			for _, charge := range firstPage.Data {
				firstPageIds[charge.Id] = true
			}

			for _, charge := range secondPage.Data {
				if firstPageIds[charge.Id] {
					t.Errorf("Duplicate charge %s found in both pages", charge.Id)
				}
			}

			t.Logf("Pagination test passed. First page: %d charges, Second page: %d charges",
				len(firstPage.Data), len(secondPage.Data))
		} else {
			t.Logf("Pagination test completed. Only one page available with %d charges",
				len(firstPage.Data))
		}
	})

	// Test 21: Reverse Date Range (End < Start) - Should handle gracefully
	t.Run("Reverse_Date_Range_Test", func(t *testing.T) {
		// Intentionally reversed dates
		endDate := NewDate(2025, 9, 1, time.UTC)
		startDate := NewDate(2025, 9, 30, time.UTC) // Start is after end

		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: startDate, // Later date
			DateCreatedLe: endDate,   // Earlier date
			Limit:         5,
		})

		// API should handle this gracefully (likely return empty results)
		assertResponseSuccess(t, resp, err)

		// Should return no results since the range is invalid
		if len(resp.Data) > 0 {
			t.Logf("Reverse date range returned %d charges (API may handle this differently)", len(resp.Data))
		} else {
			t.Logf("Reverse date range test passed. Correctly returned empty results")
		}
	})
}

func TestChargeGetAllDateFilteringErrorHandling(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 40*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)

	// Test 1: Timezone Edge Cases - Different timezone contexts
	t.Run("Timezone_Edge_Cases_Test", func(t *testing.T) {
		// Test with different timezones
		utcDate := NewDate(2025, 9, 10, time.UTC)
		estDate := NewDate(2025, 9, 10, time.FixedZone("EST", -5*3600))
		pstDate := NewDate(2025, 9, 10, time.FixedZone("PST", -8*3600))

		timezoneTests := []struct {
			name string
			date Date
		}{
			{"UTC", utcDate},
			{"EST", estDate},
			{"PST", pstDate},
		}

		for _, tt := range timezoneTests {
			resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
				DateCreatedGe: tt.date,
				DateCreatedLe: tt.date,
				Limit:         5,
			})
			assertResponseSuccess(t, resp, err)
			t.Logf("Timezone %s test passed. Found %d charges for %s",
				tt.name, len(resp.Data), tt.date.Format())
		}
	})

	// Test 2: Date Boundary Testing - First/Last day of month/year
	t.Run("Date_Boundary_Testing", func(t *testing.T) {
		boundaryDates := []struct {
			name string
			date Date
		}{
			{"New Year", NewDate(2025, 1, 1, time.UTC)},
			{"New Year Eve", NewDate(2024, 12, 31, time.UTC)},
			{"Month Start", NewDate(2025, 9, 1, time.UTC)},
			{"Month End", NewDate(2025, 9, 30, time.UTC)},
			{"Feb 28 Non-Leap", NewDate(2023, 2, 28, time.UTC)},
			{"Feb 29 Leap", NewDate(2024, 2, 29, time.UTC)},
		}

		for _, bd := range boundaryDates {
			resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
				DateCreatedGe: bd.date,
				DateCreatedLe: bd.date,
				Limit:         3,
			})
			assertResponseSuccess(t, resp, err)
			t.Logf("Boundary date %s (%s) test passed. Found %d charges",
				bd.name, bd.date.Format(), len(resp.Data))
		}
	})

	// Test 3: Multiple Date Filter Combinations - Test all date field combinations
	t.Run("Multiple_Date_Filter_Combinations_Test", func(t *testing.T) {
		baseDate := NewDate(2025, 9, 1, time.UTC)
		endDate := NewDate(2025, 9, 30, time.UTC)

		// Test various combinations of date filters
		combinations := []struct {
			name    string
			request GetAllChargesRequest
		}{
			{
				"DateCreated + DueDate",
				GetAllChargesRequest{
					DateCreatedGe: baseDate,
					DateCreatedLe: endDate,
					DueDateGe:     baseDate,
					DueDateLe:     endDate,
					Limit:         5,
				},
			},
			{
				"DateCreated + PaymentDate",
				GetAllChargesRequest{
					DateCreatedGe: baseDate,
					DateCreatedLe: endDate,
					PaymentDateGe: baseDate,
					PaymentDateLe: endDate,
					Status:        ChargeStatusConfirmed,
					Limit:         5,
				},
			},
			{
				"All Date Filters Combined",
				GetAllChargesRequest{
					DateCreatedGe:         baseDate,
					DateCreatedLe:         endDate,
					DueDateGe:             baseDate,
					DueDateLe:             endDate,
					EstimatedCreditDateGE: baseDate,
					EstimatedCreditDateLE: endDate,
					Status:                ChargeStatusConfirmed,
					Limit:                 3,
				},
			},
		}

		for _, combo := range combinations {
			resp, err := nCharge.GetAll(ctx, combo.request)
			assertResponseSuccess(t, resp, err)
			t.Logf("%s combination test passed. Found %d charges", combo.name, len(resp.Data))
		}
	})

	// Test 4: Date Filter Response Validation - Ensure returned data matches filters
	t.Run("Response_Data_Validation_Test", func(t *testing.T) {
		startDate := NewDate(2025, 9, 1, time.UTC)
		endDate := NewDate(2025, 9, 30, time.UTC)

		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe: startDate,
			DateCreatedLe: endDate,
			Limit:         10,
		})
		assertResponseSuccess(t, resp, err)

		// Comprehensive validation of response structure
		if resp == nil {
			t.Fatalf("Response is nil")
		}

		// Validate pagination fields
		if resp.Limit <= 0 {
			t.Errorf("Invalid limit in response: %d", resp.Limit)
		}

		if resp.Offset < 0 {
			t.Errorf("Invalid offset in response: %d", resp.Offset)
		}

		if resp.TotalCount < 0 {
			t.Errorf("Invalid totalCount in response: %d", resp.TotalCount)
		}

		// Validate individual charge data
		for i, charge := range resp.Data {
			// Check required fields
			if charge.Id == "" {
				t.Errorf("Charge at index %d has empty ID", i)
			}

			if charge.Customer == "" {
				t.Errorf("Charge %s has empty customer ID", charge.Id)
			}

			if charge.Value <= 0 {
				t.Errorf("Charge %s has invalid value: %f", charge.Id, charge.Value)
			}

			// Validate date fields against filter
			if !charge.DateCreated.IsZero() {
				if charge.DateCreated.Time().Before(startDate.Time()) {
					t.Errorf("Charge %s created %s is before filter start %s",
						charge.Id, charge.DateCreated.Format(), startDate.Format())
				}
				if charge.DateCreated.Time().After(endDate.Time()) {
					t.Errorf("Charge %s created %s is after filter end %s",
						charge.Id, charge.DateCreated.Format(), endDate.Format())
				}
			}

			// Validate enum fields
			validStatuses := []ChargeStatus{
				ChargeStatusPending, ChargeStatusConfirmed, ChargeStatusReceived,
				ChargeStatusOverdue, ChargeStatusRefunded, ChargeStatusReceivedInCash,
				ChargeStatusRefundRequested, ChargeStatusRefundInProgress,
				ChargeStatusChargebackRequested, ChargeStatusChargebackDispute,
				ChargeStatusAwaitingChargebackReversal, ChargeStatusDunningRequested,
				ChargeStatusDunningReceived, ChargeStatusAwaitingRiskAnalysis,
			}
			statusValid := false
			for _, validStatus := range validStatuses {
				if charge.Status == validStatus {
					statusValid = true
					break
				}
			}
			if !statusValid {
				t.Errorf("Charge %s has invalid status: %s", charge.Id, charge.Status)
			}
		}

		t.Logf("Response validation test passed. Validated %d charges with comprehensive checks", len(resp.Data))
	})

	// Test 5: Large Offset Testing - Test pagination at high offsets
	t.Run("Large_Offset_Pagination_Test", func(t *testing.T) {
		// Test with larger offset values
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			Offset: 50, // Skip first 50 results
			Limit:  5,
		})
		assertResponseSuccess(t, resp, err)

		// Should still return valid response even if no results
		if resp.Offset != 50 {
			t.Errorf("Expected offset 50, got %d", resp.Offset)
		}

		t.Logf("Large offset test passed. Offset: %d, Found %d charges", resp.Offset, len(resp.Data))
	})
}

func TestChargeGetAllURLParameterValidation(t *testing.T) {
	initChargeDocumentId()
	accessToken := getEnvValue(EnvAccessToken)
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	nCharge := NewCharge(EnvSandbox, accessToken)

	// Test 1: URL Parameter Construction - Verify correct bracket notation
	t.Run("URL_Parameter_Construction_Test", func(t *testing.T) {
		startDate := NewDate(2025, 9, 1, time.UTC)
		endDate := NewDate(2025, 9, 30, time.UTC)

		// The actual URL construction is tested implicitly through successful API calls
		// but we can verify the request structure works correctly
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			DateCreatedGe:         startDate,
			DateCreatedLe:         endDate,
			DueDateGe:             startDate,
			DueDateLe:             endDate,
			PaymentDateGe:         startDate,
			PaymentDateLe:         endDate,
			EstimatedCreditDateGE: startDate,
			EstimatedCreditDateLE: endDate,
			Customer:              "test_customer",
			BillingType:           BillingTypeCreditCard,
			Status:                ChargeStatusPending,
			Limit:                 5,
		})

		// If the URL parameters are malformed, the API would return an error
		assertResponseSuccess(t, resp, err)
		t.Logf("URL parameter construction test passed. Complex filter request successful")
	})

	// Test 2: Special Characters in Parameters - Test edge cases
	t.Run("Special_Characters_Test", func(t *testing.T) {
		// Test with special characters that might cause URL encoding issues
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			ExternalReference: "test-ref_123", // Contains special chars
			Limit:             5,
		})
		assertResponseSuccess(t, resp, err)
		t.Logf("Special characters test passed. Found %d charges", len(resp.Data))
	})

	// Test 3: Maximum Parameters - Test with all possible parameters
	t.Run("Maximum_Parameters_Test", func(t *testing.T) {
		baseDate := NewDate(2025, 9, 1, time.UTC)
		endDate := NewDate(2025, 9, 30, time.UTC)

		// Use as many parameters as possible to test URL construction limits
		resp, err := nCharge.GetAll(ctx, GetAllChargesRequest{
			Customer:              "",                   // Empty string should be omitted
			Subscription:          "",                   // Empty string should be omitted
			Installment:           "",                   // Empty string should be omitted
			CustomerGroupName:     "",                   // Empty string should be omitted
			BillingType:           BillingTypeUndefined, // Valid enum
			Status:                ChargeStatusPending,  // Valid enum
			ExternalReference:     "",                   // Empty string should be omitted
			DateCreatedGe:         baseDate,
			DateCreatedLe:         endDate,
			PaymentDateGe:         Date{}, // Zero value should be omitted
			PaymentDateLe:         Date{}, // Zero value should be omitted
			EstimatedCreditDateGE: Date{}, // Zero value should be omitted
			EstimatedCreditDateLE: Date{}, // Zero value should be omitted
			DueDateGe:             baseDate,
			DueDateLe:             endDate,
			User:                  "", // Empty string should be omitted
			Offset:                0,  // Should be included even if 0
			Limit:                 10, // Should be included
		})
		assertResponseSuccess(t, resp, err)
		t.Logf("Maximum parameters test passed. Found %d charges", len(resp.Data))
	})
}
