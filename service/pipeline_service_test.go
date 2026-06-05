package service

import (
	"testing"
	"backend/utils"
	)

func TestProcessPipelineValidation(t *testing.T) {

	records := []Record{
		{
			Name:  "John",
			Email: "john@example.com",
			Phone: "9876543210",
		},
		{
			Name:  "Jane",
			Email: "invalid-email",
			Phone: "123",
		},
	}

	// Simulate validation logic without checking DB operations
	for i := range records {
		record := &records[i]

		record.Status = "COMPLETED"

		if !utils.IsValidPhone(record.Phone) {
			record.Status = "FAILED"
			record.Error = "phone number must contain exactly 10 digits"
		}

		if !utils.IsValidEmail(record.Email) {
			record.Status = "FAILED"
			record.Error = "invalid email address"
		}
	}

	if records[0].Status != "COMPLETED" {
		t.Errorf("expected COMPLETED, got %s", records[0].Status)
	}

	if records[1].Status != "FAILED" {
		t.Errorf("expected FAILED, got %s", records[1].Status)
	}
}
