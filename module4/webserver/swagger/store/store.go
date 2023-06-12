package store

import "encoding/json"

func UnmarshalOrder(data []byte) (Order, error) {
	var r Order
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Order) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Order struct {
	ID       int    `json:"id"`
	PetID    int    `json:"petId"`
	Quantity int    `json:"quantity"`
	ShipDate string `json:"shipDate"`
	Status   string `json:"status"`
	Complete bool   `json:"complete"`
}
