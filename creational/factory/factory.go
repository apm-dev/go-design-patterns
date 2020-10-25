package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(float642 float32) string
}

const (
	Cash     = 1
	ZarinPal = 2
)

func GetPaymentMethod(method int) (PaymentMethod, error) {
	switch method {
	case Cash:
		return new(CashPM), nil
	case ZarinPal:
		return new(ZarinPalPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not exists", method))
	}
}

type CashPM struct {
}

type ZarinPalPM struct {
}

func (*CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash", amount)
}

func (*ZarinPalPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using zarinpal", amount)
}
