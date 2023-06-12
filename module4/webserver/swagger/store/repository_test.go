package store

import (
	"reflect"
	"testing"
)

var orders = []Order{
	{
		ID:       0,
		PetID:    34,
		Quantity: 353,
		Status:   "test",
		Complete: true,
		ShipDate: "12.06.2023",
	}, {
		ID:       1,
		PetID:    12,
		Quantity: 575,
		Status:   "sold",
		Complete: false,
		ShipDate: "12.06.2023",
	}, {
		ID:       2,
		PetID:    93,
		Quantity: 12,
		Status:   "available",
		Complete: false,
		ShipDate: "12.06.2023",
	},
}

func TestOrderStorage_Create(t *testing.T) {

	type args struct {
		Order Order
	}
	tests := []struct {
		name string
		args args
		want Order
	}{
		{
			name: "Test Create 1",
			args: args{
				Order: Order{
					ID:       0,
					PetID:    34,
					Quantity: 353,
					Status:   "test",
					Complete: true,
					ShipDate: "12.06.2023",
				},
			},
			want: Order{
				ID:       0,
				PetID:    34,
				Quantity: 353,
				Status:   "test",
				Complete: true,
				ShipDate: "12.06.2023",
			},
		},
	}

	p := NewOrderStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.Create(tt.args.Order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderStorage_Delete(t *testing.T) {
	type args struct {
		OrderID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test Create 1",
			args:    args{OrderID: 0},
			wantErr: false,
		}, {
			name:    "Test Delete 2",
			args:    args{OrderID: 1},
			wantErr: false,
		}, {
			name:    "Test Delete 3",
			args:    args{OrderID: 34},
			wantErr: true,
		},
	}

	p := NewOrderStorage()
	for _, v := range orders {
		p.Create(v)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.Delete(tt.args.OrderID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrderStorage_Get(t *testing.T) {
	type args struct {
		OrderID int
	}
	tests := []struct {
		name    string
		args    args
		want    Order
		wantErr bool
	}{
		{
			name: "Test get 1",
			args: args{OrderID: 0},
			want: Order{
				ID:       0,
				PetID:    34,
				Quantity: 353,
				Status:   "test",
				Complete: true,
				ShipDate: "12.06.2023",
			},
			wantErr: false,
		}, {
			name: "Test get 2",
			args: args{OrderID: 1},
			want: Order{
				ID:       1,
				PetID:    12,
				Quantity: 575,
				Status:   "sold",
				Complete: false,
				ShipDate: "12.06.2023",
			},
			wantErr: false,
		},
	}

	p := NewOrderStorage()
	for _, v := range orders {
		p.Create(v)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := p.Get(tt.args.OrderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
