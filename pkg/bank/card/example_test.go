package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	payments := []types.Card{
		{PAN: "01",
			Active:  true,
			Balance: 10_000,
		},
		{PAN: "01",
			Active:  false,
			Balance: 10_000,
		},
		{PAN: "01",
			Active:  true,
			Balance: -10_000,
		},
	}
	maxes := PaymentSources(payments)
	for _, max := range maxes {
		fmt.Println(max.Number)
	}
	//Output:
	//01
}
