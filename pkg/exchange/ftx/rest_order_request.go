package ftx

import (
	"context"
	"encoding/json"
	"fmt"
)

type orderRequest struct {
	*restRequest
}

/*
{
  "market": "XRP-PERP",
  "side": "sell",
  "price": 0.306525,
  "type": "limit",
  "size": 31431.0,
  "reduceOnly": false,
  "ioc": false,
  "postOnly": false,
  "clientId": null
}
*/
type PlaceOrderPayload struct {
	Market     string
	Side       string
	Price      float64
	Type       string
	Size       float64
	ReduceOnly bool
	IOC        bool
	PostOnly   bool
	ClientID   string
}

func (r *orderRequest) PlaceOrder(ctx context.Context, p PlaceOrderPayload) (orderResponse, error) {
	resp, err := r.
		Method("POST").
		ReferenceURL("api/orders").
		Payloads(map[string]interface{}{
			"market":     p.Market,
			"side":       p.Side,
			"price":      p.Price,
			"type":       p.Type,
			"size":       p.Size,
			"reduceOnly": p.ReduceOnly,
			"ioc":        p.IOC,
			"postOnly":   p.PostOnly,
			"clientId":   p.ClientID,
		}).
		DoAuthenticatedRequest(ctx)

	if err != nil {
		return orderResponse{}, err
	}
	var o orderResponse
	if err := json.Unmarshal(resp.Body, &o); err != nil {
		return orderResponse{}, fmt.Errorf("failed to unmarshal order response body to json: %w", err)
	}

	return o, nil
}

func (r *orderRequest) OpenOrders(ctx context.Context, market string) (ordersResponse, error) {
	resp, err := r.
		Method("GET").
		ReferenceURL("api/orders").
		Payloads(map[string]interface{}{"market": market}).
		DoAuthenticatedRequest(ctx)

	if err != nil {
		return ordersResponse{}, err
	}

	var o ordersResponse
	if err := json.Unmarshal(resp.Body, &o); err != nil {
		return ordersResponse{}, fmt.Errorf("failed to unmarshal open orders response body to json: %w", err)
	}

	return o, nil
}
