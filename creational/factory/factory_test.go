package factory

import (
	"strings"
	"testing"
)

func TestCachePaymentMethod(t *testing.T) {
	pm, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("Payment method of type 'Cash' must exist")
	}

	msg := pm.Pay(10.0)
	if !strings.Contains(strings.ToLower(msg), "paid using cash") {
		t.Error("The cache payment method message was not correct")
	}
	t.Log("LOG:", msg)
}

func TestZarinPalPaymentMethod(t *testing.T) {
	pm, err := GetPaymentMethod(ZarinPal)
	if err != nil {
		t.Fatal("Payment method of type 'ZarinPal' must exist")
	}

	msg := pm.Pay(10.0)
	if !strings.Contains(strings.ToLower(msg), "paid using zarinpal") {
		t.Error("The zarin-pal payment method message was not correct")
	}
	t.Log("LOG:", msg)
}

func TestNotExistentPaymentMethod(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Fatal("Payment method of not existent type must return error")
	}
	t.Log("LOG:", err)
}
