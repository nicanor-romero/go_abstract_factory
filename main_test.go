package main

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a - b) <= 1e-7
}

func TestVoiceSpainGermanyCdr(t *testing.T) {
	cdr := &CDR{
		recordType:    "voice",
		originId:      "SPAIN",
		destinationId: "GERMANY",
		usage:         30,
	}
	expectedCost := 0.75
	actualCost := rateCDR(cdr)
	if !almostEqual(expectedCost, actualCost) {
		t.Errorf("Actual cost (%f) differs from expected cost (%f)", actualCost, expectedCost)
	}
}

func TestVoiceSpainArgentinaCdr(t *testing.T) {
	cdr := &CDR{
		recordType:    "voice",
		originId:      "SPAIN",
		destinationId: "ARGENTINA",
		usage:         50,
	}
	expectedCost := 2.7
	actualCost := rateCDR(cdr)
	if !almostEqual(expectedCost, actualCost) {
		t.Errorf("Actual cost (%f) differs from expected cost (%f)", actualCost, expectedCost)
	}
}

func TestSmsSpainGermanyCdr(t *testing.T) {
	cdr := &CDR{
		recordType:    "sms",
		originId:      "SPAIN",
		destinationId: "GERMANY",
		usage:         1,
	}
	expectedCost := 0.10
	actualCost := rateCDR(cdr)
	if !almostEqual(expectedCost, actualCost) {
		t.Errorf("Actual cost (%f) differs from expected cost (%f)", actualCost, expectedCost)
	}
}

func TestSmsSpainArgentinaCdr(t *testing.T) {
	cdr := &CDR{
		recordType:    "sms",
		originId:      "SPAIN",
		destinationId: "ARGENTINA",
		usage:         2,
	}
	expectedCost := 0.50
	actualCost := rateCDR(cdr)
	if !almostEqual(expectedCost, actualCost) {
		t.Errorf("Actual cost (%f) differs from expected cost (%f)", actualCost, expectedCost)
	}
}