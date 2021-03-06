package kiteconnect

import "testing"

func (ts *TestSuite) TestGetOrderMargins(t *testing.T) {
	t.Parallel()

	params := OrderMarginParam{
		Exchange:        "NSE",
		Tradingsymbol:   "INFY",
		TransactionType: "BUY",
		Variety:         "regular",
		Product:         "CNC",
		OrderType:       "MARKET",
		Quantity:        1,
		Price:           0,
		TriggerPrice:    0,
	}

	orderResponse, err := ts.KiteConnect.GetOrderMargins([]OrderMarginParam{params})
	if err != nil {
		t.Errorf("Error while getting order margins: %v", err)
	}

	if len(orderResponse)!= 1 {
		t.Errorf("Incorrect response, expected len(orderResponse) to be 0, got: %v", len(orderResponse))
	}

	if orderResponse[0].Total != 961.45 {
		t.Errorf("Incorrect total, expected 961.45, got: %v", orderResponse[0].Total)
	}
}
