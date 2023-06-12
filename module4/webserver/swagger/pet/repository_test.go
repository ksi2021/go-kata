package pet

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

var pets = []Pet{
	{
		ID: 0,
		Category: Category{
			ID:   0,
			Name: "test category",
		},
		Name: "Alma",
		PhotoUrls: []string{
			"https://sobakovod.club/uploads/posts/2021-12/1639834742_26-sobakovod-club-p-sobaki-laika-chernaya-29.jpg",
			"http://catfishes.ru/wp-content/uploads/2021/06/ruslaika1.jpg",
			"https://lapkins.ru/upload/iblock/ebd/ebd19f44cd131f425feed81f9b578198.jpg",
		},
		Tags:   []Category{},
		Status: "active",
	}, {
		ID: 0,
		Category: Category{
			ID:   0,
			Name: "test category",
		},
		Name: "Salma",
		PhotoUrls: []string{
			"https://sobakovod.club/uploads/posts/2021-12/1639834742_26-sobakovod-club-p-sobaki-laika-chernaya-29.jpg",
			"http://catfishes.ru/wp-content/uploads/2021/06/ruslaika1.jpg",
			"https://lapkins.ru/upload/iblock/ebd/ebd19f44cd131f425feed81f9b578198.jpg",
		},
		Tags:   []Category{},
		Status: "active",
	}}

func TestPetStorage_Create(t *testing.T) {
	type fields struct {
		data               []*Pet
		primaryKeyIDx      map[int64]*Pet
		autoIncrementCount int64
		Mutex              sync.Mutex
	}
	type args struct {
		pet Pet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Pet
	}{
		{
			name: "first test",
			args: args{
				pet: Pet{
					ID: 0,
					Category: Category{
						ID:   0,
						Name: "test category",
					},
					Name: "Alma",
					PhotoUrls: []string{
						"https://sobakovod.club/uploads/posts/2021-12/1639834742_26-sobakovod-club-p-sobaki-laika-chernaya-29.jpg",
						"http://catfishes.ru/wp-content/uploads/2021/06/ruslaika1.jpg",
						"https://lapkins.ru/upload/iblock/ebd/ebd19f44cd131f425feed81f9b578198.jpg",
					},
					Tags:   []Category{},
					Status: "active",
				},
			},
			want: Pet{
				ID: 0,
				Category: Category{
					ID:   0,
					Name: "test category",
				},
				Name: "Alma",
				PhotoUrls: []string{
					"https://sobakovod.club/uploads/posts/2021-12/1639834742_26-sobakovod-club-p-sobaki-laika-chernaya-29.jpg",
					"http://catfishes.ru/wp-content/uploads/2021/06/ruslaika1.jpg",
					"https://lapkins.ru/upload/iblock/ebd/ebd19f44cd131f425feed81f9b578198.jpg",
				},
				Tags:   []Category{},
				Status: "active",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPetStorage()
			var got Pet
			var err error
			got = p.Create(tt.args.pet)
			tt.want.ID = got.ID
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
			if got, err = p.GetByID(got.ID); err != nil || !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v, err %s", got, tt.want, err)
			}
		})
	}
}

func TestPetStorage_AddImage(t *testing.T) {
	type args struct {
		petID int
		img   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test delete 1",
			args:    args{petID: 0, img: "test.png"},
			wantErr: false,
		}, {
			name:    "Test delete 2",
			args:    args{petID: 100, img: "test.jpg"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPetStorage()
			for _, v := range pets {
				p.Create(v)
			}
			if err := p.AddImage(tt.args.petID, tt.args.img); (err != nil) != tt.wantErr {
				t.Errorf("AddImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPetStorage_Delete(t *testing.T) {
	type args struct {
		petID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test delete 1",
			args:    args{petID: 0},
			wantErr: false,
		}, {
			name:    "Test delete 2",
			args:    args{petID: 100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPetStorage()
			for _, v := range pets {
				p.Create(v)
			}
			if err := p.Delete(tt.args.petID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPetStorage_FindByStatus(t *testing.T) {
	type args struct {
		filters []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "findByStatus 1",
			args: args{filters: []string{"available", "pending"}},
			want: 0,
		}, {
			name: "findByStatus 2",
			args: args{filters: []string{"active"}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPetStorage()
			for _, v := range pets {
				p.Create(v)
			}
			fmt.Println(p.GetList())
			if got := p.FindByStatus(tt.args.filters); len(got) != tt.want {
				fmt.Println(len(got), "got")
				t.Errorf("FindByStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPetStorage_GetByID(t *testing.T) {
	type args struct {
		petID int
	}
	tests := []struct {
		name    string
		args    args
		want    Pet
		wantErr bool
	}{
		{
			name: "first test",
			args: args{petID: 0},
			want: Pet{
				ID: 0,
				Category: Category{
					ID:   0,
					Name: "test category",
				},
				Name: "Alma",
				PhotoUrls: []string{
					"https://sobakovod.club/uploads/posts/2021-12/1639834742_26-sobakovod-club-p-sobaki-laika-chernaya-29.jpg",
					"http://catfishes.ru/wp-content/uploads/2021/06/ruslaika1.jpg",
					"https://lapkins.ru/upload/iblock/ebd/ebd19f44cd131f425feed81f9b578198.jpg",
				},
				Tags:   []Category{},
				Status: "active",
			},
			wantErr: false,
		}, {
			name:    "second test",
			args:    args{petID: 1},
			wantErr: false,
			want: Pet{
				ID: 1,
				Category: Category{
					ID:   0,
					Name: "test category",
				},
				Name: "Salma",
				PhotoUrls: []string{
					"https://sobakovod.club/uploads/posts/2021-12/1639834742_26-sobakovod-club-p-sobaki-laika-chernaya-29.jpg",
					"http://catfishes.ru/wp-content/uploads/2021/06/ruslaika1.jpg",
					"https://lapkins.ru/upload/iblock/ebd/ebd19f44cd131f425feed81f9b578198.jpg",
				},
				Tags:   []Category{},
				Status: "active",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPetStorage()
			for _, v := range pets {
				p.Create(v)
			}
			got, err := p.GetByID(tt.args.petID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPetStorage_Update1(t *testing.T) {

	type args struct {
		name   string
		status string
		petID  int
	}
	tests := []struct {
		name    string
		fields  Pet
		args    args
		wantErr bool
	}{
		{
			name: "first test",
			fields: Pet{
				ID: 0,
				Category: Category{
					ID:   0,
					Name: "test category",
				},
				Name: "Alma",
				PhotoUrls: []string{
					"https://sobakovod.club/uploads/posts/2021-12/1639834742_26-sobakovod-club-p-sobaki-laika-chernaya-29.jpg",
					"http://catfishes.ru/wp-content/uploads/2021/06/ruslaika1.jpg",
					"https://lapkins.ru/upload/iblock/ebd/ebd19f44cd131f425feed81f9b578198.jpg",
				},
				Tags:   []Category{},
				Status: "active",
			},
			args:    args{name: "new name", status: "sold", petID: 0},
			wantErr: false,
		}, {
			name: "second test",
			fields: Pet{
				ID: 0,
				Category: Category{
					ID:   0,
					Name: "test category",
				},
				Name: "Lupa",
				PhotoUrls: []string{
					"https://sobakovod.club/uploads/posts/2021-12/1639834742_26-sobakovod-club-p-sobaki-laika-chernaya-29.jpg",
					"http://catfishes.ru/wp-content/uploads/2021/06/ruslaika1.jpg",
					"https://lapkins.ru/upload/iblock/ebd/ebd19f44cd131f425feed81f9b578198.jpg",
				},
				Tags:   []Category{},
				Status: "active",
			},
			args:    args{name: "new name", status: "sold", petID: 999},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPetStorage()
			p.Create(tt.fields)
			if err := p.Update(tt.args.name, tt.args.status, tt.args.petID); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
