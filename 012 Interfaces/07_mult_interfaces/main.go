package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) string
}

type CardInfoProvider interface {
	CardInfo() string
}

type GiftCard struct {
	Code    string
	Balance float64
}

func (g GiftCard) Pay(amount float64) string {
	if amount > g.Balance {
		// if true, вернется этот return
		return "Not enough balance!"
	}
	// if amount > g.Balance false, вернется этот return
	return fmt.Sprintf("Paid $%.2f using gift card", amount)
}

func (g GiftCard) CardInfo() string {
	return fmt.Sprintf("Gift card code: %s | Balance: $%.2f", g.Code, g.Balance)
}

func main() {
	card := GiftCard{
		Balance: 125.0,
		Code:    "GC0001",
	}
	var pay PaymentMethod = card
	var info CardInfoProvider = card

	fmt.Println(info.CardInfo())
	fmt.Println(pay.Pay(235.50))
}
