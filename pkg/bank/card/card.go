package card

import (
	"bank/pkg/bank/types"
)

//PaymentSource из слайс карт сделать слайс источников оплаты
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var operations []types.PaymentSource
	for _, card := range cards {
		if (card.Balance > 0) && (card.Active == true) {

			operations = append(operations, types.PaymentSource{Type: "card", Number: string(card.PAN), Balance: card.Balance})
		}
	}
	return operations
}
