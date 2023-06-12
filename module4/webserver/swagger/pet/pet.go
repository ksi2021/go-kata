package pet

import "encoding/json"

func UnmarshalPet(data []byte) (Pet, error) {
	var r Pet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Pet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Pet struct {
	ID        int        `json:"id"`
	Category  Category   `json:"category"`
	Name      string     `json:"name"`
	PhotoUrls []string   `json:"photoUrls"`
	Tags      []Category `json:"tags"`
	Status    string     `json:"status"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
