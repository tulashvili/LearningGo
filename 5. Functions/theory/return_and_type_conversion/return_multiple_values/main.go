package main

import "fmt"

// Client post money + his tip
func processPayment(clientMoney int, orderCost int, orderTip int) (int, int, int) {
	totalAmount := clientMoney + orderTip
	totalReturn := totalAmount - orderCost
	return totalAmount, totalReturn, orderTip
}

func main() {
	getFromClient, needClientReturn, clientTips := processPayment(10, 7, 1)
	fmt.Printf("Client give you $%d (tip is $%d).\nYou need return to client $%d\n", getFromClient, clientTips, needClientReturn)
}
