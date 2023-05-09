package main

import (
	"testing"
)

func TestShift(t *testing.T) {

	tests := []struct {
		name string
		want string
	}{
		{
			name: "Shift 1",
			want: "Synthesizing the interface won't do anything, we need to copy the cross-platform IB feed!",
		},
		{
			name: "Shift 2",
			want: "We need to generate the redundant HTTP microchip!",
		},
		{
			name: "Shift 3",
			want: "If we transmit the monitor, we can get to the HTTP program through the virtual COM port!",
		},
	}

	list := &DoubleLinkedList{}
	_ = list.LoadData("./test.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.Shift(); got.data.Message != tt.want {
				t.Errorf("Shift() = %v, want %v", got.data.Message, tt.want)
			}
		})
	}
}

func TestPop(t *testing.T) {

	tests := []struct {
		name string
		want string
	}{
		{
			name: "Pop 1",
			want: "I'Ll hack the redundant AI monitor, that should program the SAS card!",
		},
		{
			name: "Pop 2",
			want: "We need to transmit the back-end USB application!",
		},
		{
			name: "Pop 3",
			want: "You can't parse the monitor without indexing the open-source SQL program!",
		},
	}
	list := &DoubleLinkedList{}
	_ = list.LoadData("./test.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.Pop(); got.data.Message != tt.want {
				t.Errorf("Pop() = %v, want %v", got.data.Message, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	type args struct {
		position int
		data     *Commit
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Insert #1",
			args: args{100000, &Commit{}},
			want: false,
		},
		{
			name: "Insert #2",
			args: args{-100, &Commit{}},
			want: false,
		},
		{
			name: "Insert #3",
			args: args{25, &Commit{}},
			want: true,
		},
	}
	list := &DoubleLinkedList{}
	_ = list.LoadData("./test.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.Insert(tt.args.position, *tt.args.data); got != tt.want {
				t.Errorf("Inser postion(%d) = %v, want %v", tt.args.position, got, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Search #1",
			args: args{message: "Use the bluetooth HTTP protocol, then you can parse the cross-platform pixel!"},
			want: "6956ec68-875b-11ed-8150-acde48001122",
		},
		{
			name: "Search #2",
			args: args{message: "We need to back up the open-source GB alarm!"},
			want: "6956e1a0-875b-11ed-8150-acde48001122",
		},
		{
			name: "Search #3",
			args: args{message: "The RAM feed is down, input the haptic capacitor so we can navigate the ADP bus!"},
			want: "6956e84e-875b-11ed-8150-acde48001122",
		},
	}
	list := &DoubleLinkedList{}
	_ = list.LoadData("./test.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.Search(tt.args.message); got == nil || got.data.UUID != tt.want {
				t.Errorf("Search = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchUUID(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "SearchUUID #1",
			args: args{uuid: "6956ec68-875b-11ed-8150-acde48001122"},
			want: "Use the bluetooth HTTP protocol, then you can parse the cross-platform pixel!",
		},
		{
			name: "SearchUUID #2",
			args: args{uuid: "6956e1a0-875b-11ed-8150-acde48001122"},
			want: "We need to back up the open-source GB alarm!",
		},
		{
			name: "SearchUUID #3",
			args: args{uuid: "6956e84e-875b-11ed-8150-acde48001122"},
			want: "The RAM feed is down, input the haptic capacitor so we can navigate the ADP bus!",
		},
	}
	list := &DoubleLinkedList{}
	_ = list.LoadData("./test.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.SearchUUID(tt.args.uuid); got == nil || got.data.Message != tt.want {
				t.Errorf("SearchUUID = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		position int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Delete #1",
			args: args{position: -1},
			want: false,
		},
		{
			name: "Delete #2",
			args: args{position: 25},
			want: true,
		},
		{
			name: "Delete #3",
			args: args{position: 100000},
			want: false,
		},
		{
			name: "Delete #4",
			args: args{position: 151},
			want: false,
		},
		{
			name: "Delete #5",
			args: args{position: 0},
			want: true,
		},
	}

	list := &DoubleLinkedList{}
	_ = list.LoadData("./test.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.Delete(tt.args.position); got != tt.want {
				t.Errorf("SearchUUID = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Reverse #1",
			want: "I'Ll hack the redundant AI monitor, that should program the SAS card!",
		},
		{
			name: "Reverse #2",
			want: "We need to transmit the back-end USB application!",
		},
		{
			name: "Reverse #3",
			want: "You can't parse the monitor without indexing the open-source SQL program!",
		},
	}

	list := &DoubleLinkedList{}
	_ = list.LoadData("./test.json")
	list.Reverse()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := list.Next(); got == nil && got.data.Message != tt.want {
				t.Errorf("SearchUUID = %+v, want %v", got, tt.want)
			}
			list.Next()
		})

	}
}
