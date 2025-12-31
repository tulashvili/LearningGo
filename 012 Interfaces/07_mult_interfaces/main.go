package main

import "fmt"

type PaymnentMethod interface {
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
		return "Not enough balance"
	}
	return fmt.Sprintf("Paid $%.2f using gift card", amount)

}

func (g GiftCard) CardInfo() string {
	return fmt.Sprintf("Gift card code: %s | Balance: $%.2f", g.Code, g.Balance)
}

func main() {
	card := GiftCard{
		Code:    "GC00001",
		Balance: 125.0,
	}

	var pay PaymnentMethod = card
	var info CardInfoProvider = card

	fmt.Println(pay.Pay(35.50))
	fmt.Println(info.CardInfo())

}
