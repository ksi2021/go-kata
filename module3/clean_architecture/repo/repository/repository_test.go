package repository

import (
	"log"
	"os"
	"testing"
)

func TestFind(t *testing.T) {

	tests := []struct {
		name     string
		want     bool
		searchID int
	}{
		{
			name:     "Find 1",
			want:     true,
			searchID: 01,
		},
		{
			name:     "Find 2",
			want:     false,
			searchID: 6666,
		},
		{
			name:     "Find 3",
			want:     true,
			searchID: 02,
		},
	}
	f, err := os.OpenFile("user.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	defer func() { f.Close() }()

	if err != nil {
		log.Fatal(err)
	}
	test := NewUserRepository(f)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := test.Find(tt.searchID); err != tt.want {
				t.Errorf("Find = %v, result = %v , want %v", tt.searchID, err, tt.want)
			}
			f.Seek(0, 0)
		})
	}
}

func TestSave(t *testing.T) {
	tests := []User{
		{
			ID: 01, Name: "TEST1 SAVE",
		},
		{
			ID: 02, Name: "TEST2 SAVE",
		},
		{
			ID: 03, Name: "TEST3 SAVE",
		},
	}
	f, err := os.OpenFile("user.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	defer func() { f.Close() }()

	if err != nil {
		log.Fatal(err)
	}
	test := NewUserRepository(f)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := test.Save(tt); got != nil {
				t.Errorf("Save = %v, result = %v , want %v", tt, got, nil)
			}
		})
	}
}
