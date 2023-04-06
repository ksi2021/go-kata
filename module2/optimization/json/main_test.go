package main

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type Pets []Pet

func UnmarshalPets(data []byte) (Pets, error) {
	var r Pets
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalPets2(data []byte) (Pets, error) {
	var r Pets
	err := jsoniter.Unmarshal(data, &r)
	return r, err
}

func (r *Pets) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Pets) Marshal2() ([]byte, error) {
	return jsoniter.Marshal(r)
}

type Pet struct {
	ID        int64      `json:"id"`
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

var jsonData = []byte(`[
	{
	  "id": 0,
	  "category": {
		"id": 0,
		"name": "string"
	  },
	  "name": "doggie",
	  "photoUrls": [
		"string"
	  ],
	  "tags": [
		{
		  "id": 0,
		  "name": "string"
		}
	  ],
	  "status": "available"
	}
  ]`)

func BenchmarkJson(b *testing.B) {
	var (
		pets Pets
		err  error
		data []byte
	)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pets, err = UnmarshalPets(jsonData)
		if err != nil {
			panic(err)
		}
		data, err = pets.Marshal()
		if err != nil {
			panic(err)
		}
		_ = data
	}
}

func BenchmarkJsonIter(b *testing.B) {
	var (
		pets Pets
		err  error
		data []byte
	)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pets, err = UnmarshalPets2(jsonData)
		if err != nil {
			panic(err)
		}
		data, err = pets.Marshal2()
		if err != nil {
			panic(err)
		}
		_ = data
	}

}
