package user

import (
	"reflect"
	"testing"
)

var users = []User{
	{
		Email:     "test1@mail.com",
		FirstName: "test1", ID: 0, LastName: "test1", Password: "pass1", Phone: "phone1", UserStatus: 0, Username: "user1",
	},
	{
		Email:     "test2@mail.com",
		FirstName: "test2", ID: 0, LastName: "test2", Password: "pass2", Phone: "phone2", UserStatus: 0, Username: "user2",
	},
}

func TestUserStorage_Create(t *testing.T) {
	type args struct {
		User User
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "Test Create 1",
			args: args{
				User: User{
					Email:     "test1@mail.com",
					FirstName: "test1", ID: 0, LastName: "test1", Password: "pass1", Phone: "phone1", UserStatus: 0, Username: "user1",
				},
			},
			want: User{
				Email:     "test1@mail.com",
				FirstName: "test1", ID: 0, LastName: "test1", Password: "pass1", Phone: "phone1", UserStatus: 0, Username: "user1",
			},
		},
		{
			name: "Test Create 2",
			args: args{
				User: User{
					Email:     "test2@mail.com",
					FirstName: "test2", ID: 1, LastName: "test2", Password: "pass2", Phone: "phone2", UserStatus: 0, Username: "user2",
				},
			},
			want: User{
				Email:     "test2@mail.com",
				FirstName: "test2", ID: 1, LastName: "test2", Password: "pass2", Phone: "phone2", UserStatus: 0, Username: "user2",
			},
		},
	}
	p := NewUserStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.Create(tt.args.User); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStorage_CreateWithList(t *testing.T) {

	type args struct {
		users []User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Create with list 1",
			args: args{
				users: users,
			},
			wantErr: false,
		},
	}
	p := NewUserStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.CreateWithList(tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("CreateWithList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserStorage_Delete(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test delete 1",
			args:    args{username: "user1"},
			wantErr: false,
		}, {
			name:    "Test delete 2",
			args:    args{username: "test9999"},
			wantErr: true,
		},
	}
	p := NewUserStorage()

	for _, v := range users {
		p.Create(v)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.Delete(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserStorage_GetByUsername(t *testing.T) {
	type args struct {
		Username string
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name:    "Test get 1",
			args:    args{Username: "user1"},
			want:    users[0],
			wantErr: false,
		}, {
			name:    "Test get 2",
			args:    args{Username: "user2"},
			want:    users[1],
			wantErr: false,
		}, {
			name:    "Test get 3",
			args:    args{Username: "test9999"},
			want:    users[0],
			wantErr: true,
		},
	}
	p := NewUserStorage()
	for _, v := range users {
		p.Create(v)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.GetByUsername(tt.args.Username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStorage_Update(t *testing.T) {
	type args struct {
		username string
		user     User
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "Test update 1",
			args: args{username: "user1", user: User{
				Email: "test1_update@mail.com", FirstName: "test1_update", ID: 0, LastName: "test1_update",
				Password: "pass1_update", Phone: "phone1_update", UserStatus: 0, Username: "user1_update",
			}},
			want: User{
				Email: "test1_update@mail.com", FirstName: "test1_update", ID: 2, LastName: "test1_update",
				Password: "pass1_update", Phone: "phone1_update", UserStatus: 0, Username: "user1_update",
			},
			wantErr: false,
		}, {
			name: "Test update 2",
			args: args{username: "user2", user: User{
				Email: "test2_update@mail.com", FirstName: "test2_update", ID: 1, LastName: "test2_update",
				Password: "pass2_update", Phone: "phone2_update", UserStatus: 0, Username: "user2_update",
			}},
			want: User{
				Email: "test2_update@mail.com", FirstName: "test2_update", ID: 3, LastName: "test2_update",
				Password: "pass2_update", Phone: "phone2_update", UserStatus: 0, Username: "user2_update",
			},
			wantErr: false,
		}, {
			name: "Test get 3",
			args: args{username: "test9999", user: User{
				Email: "test1_update@mail.com", FirstName: "test1_update", ID: 0, LastName: "test1_update",
				Password: "pass1_update", Phone: "phone1_update", UserStatus: 0, Username: "user1_update",
			}},
			want: User{
				Email: "test1_update@mail.com", FirstName: "test1_update", ID: 0, LastName: "test1_update",
				Password: "pass1_update", Phone: "phone1_update", UserStatus: 0, Username: "user1_update",
			},
			wantErr: true,
		},
	}
	p := NewUserStorage()
	for _, v := range users {
		p.Create(v)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.Update(tt.args.username, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
