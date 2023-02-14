// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    users, err := UnmarshalUsers(bytes)
//    bytes, err = users.Marshal()

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	users := getUsers()            // получаем пользователей
	products := getProducts()      // достаем продукты пользователей
	m := make(map[int64][]Product) // создаем мапу, чтобы не итерировать цикл в цикле
	var i int
	for i = range products { // проходимся по продуктам 1 раз
		m[products[i].UserID] = append(m[products[i].UserID], products[i]) // записываем продукты в map по UserID

	}

	for i = range users { // итерируем пользователей один раз
		if v, ok := m[users[i].ID]; ok { // проверяем есть ли продукты с таким user.ID, оптимизировали, чтобы не проходить в цикле
			users[i].Products = v // если ok, записываем продукты пользователя в поле Products
		}
	}

	usersToJson := Users(users) // преобразовываем в готовый тип Users для вывода в JSON (type casting)
	checkSolution(usersToJson)
	b, err := usersToJson.Marshal() // получаем байты от данных json
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b)) // преобразовываем байты в string (type casting), выводим данные в консоль
}

func checkSolution(users Users) {
	if len(users[0].Products) != 2 ||
		len(users[1].Products) != 3 ||
		len(users[3].Products) != 3 ||
		len(users[4].Products) != 2 {
		panic("task not solved")
	}
}

func getUsers() []User {
	users, err := UnmarshalUsers(usersData)
	if err != nil {
		panic(err)
	}

	return users
}

func getProducts() []Product {
	products, err := UnmarshalProducts(productsData)
	if err != nil {
		panic(err)
	}

	return products
}

type Users []User

func UnmarshalUsers(data []byte) (Users, error) {
	var r Users
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Users) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type User struct {
	ID         int64     `json:"id"`
	IsActive   bool      `json:"isActive"`
	Balance    string    `json:"balance"`
	Age        int64     `json:"age"`
	Name       string    `json:"name"`
	Gender     string    `json:"gender"`
	Company    string    `json:"company"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	About      string    `json:"about"`
	Registered string    `json:"registered"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Products   []Product `json:"products"`
}

type Products []Product

func UnmarshalProducts(data []byte) (Products, error) {
	var r Products
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Products) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Product struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
}

var usersData = []byte(`[
  {
    "id": 0,
    "isActive": false,
    "balance": "$3,877.94",
    "age": 37,
    "name": "Lillie Curtis",
    "gender": "female",
    "company": "HOPELI",
    "email": "lilliecurtis@hopeli.com",
    "phone": "+1 (965) 545-2271",
    "address": "332 Provost Street, Levant, Alabama, 1800",
    "about": "Ad reprehenderit aliqua qui esse consectetur do laborum ad. Fugiat labore anim sunt ex enim. Minim amet est cillum ea fugiat laboris reprehenderit dolor deserunt laborum sit. Qui esse veniam veniam reprehenderit. Tempor magna magna reprehenderit ipsum aliqua tempor commodo ex.\r\n",
    "registered": "2020-02-10T05:13:09 -06:00",
    "latitude": 80.249338,
    "longitude": 18.31463,
    "products": [
    ]
  },
  {
    "id": 1,
    "isActive": true,
    "balance": "$3,100.68",
    "age": 32,
    "name": "Ava Cross",
    "gender": "female",
    "company": "SYBIXTEX",
    "email": "avacross@sybixtex.com",
    "phone": "+1 (880) 531-3662",
    "address": "730 Hunts Lane, Turah, New Hampshire, 7698",
    "about": "Tempor officia amet esse anim cillum. Occaecat sit qui irure incididunt mollit amet consequat elit aute ad reprehenderit qui laboris anim. Irure labore mollit dolore non. Cupidatat sit incididunt sint ullamco et excepteur minim ipsum incididunt culpa eu. Adipisicing sunt eu aute esse officia nostrud excepteur nostrud exercitation pariatur elit laborum. Consectetur magna tempor ipsum nisi.\r\n",
    "registered": "2017-12-21T03:50:45 -06:00",
    "latitude": 40.098561,
    "longitude": 94.227292,
    "products": [
    ]
  },
  {
    "id": 2,
    "isActive": false,
    "balance": "$2,933.00",
    "age": 21,
    "name": "Snow Mooney",
    "gender": "male",
    "company": "ZAPPIX",
    "email": "snowmooney@zappix.com",
    "phone": "+1 (811) 591-3228",
    "address": "293 Kent Street, Crumpler, Massachusetts, 1768",
    "about": "Cupidatat sint dolore in pariatur mollit reprehenderit consectetur quis ipsum. Adipisicing est commodo mollit deserunt esse cillum culpa consequat quis occaecat velit. Officia exercitation in Lorem ut excepteur irure deserunt esse consequat magna nostrud.\r\n",
    "registered": "2020-07-25T10:25:48 -06:00",
    "latitude": 10.262894,
    "longitude": 117.312007,
    "products": [
    ]
  },
  {
    "id": 3,
    "isActive": true,
    "balance": "$1,351.67",
    "age": 26,
    "name": "Cook Dean",
    "gender": "male",
    "company": "ONTAGENE",
    "email": "cookdean@ontagene.com",
    "phone": "+1 (995) 568-3974",
    "address": "850 Channel Avenue, Kraemer, Georgia, 1622",
    "about": "Do reprehenderit ea proident adipisicing. Minim adipisicing ad qui tempor ut laborum. Excepteur elit pariatur mollit irure reprehenderit Lorem officia ad. Eiusmod ipsum nulla nostrud occaecat culpa tempor. Et do nisi culpa anim.\r\n",
    "registered": "2021-11-03T05:28:08 -06:00",
    "latitude": -10.887827,
    "longitude": 70.053254,
    "products": [
    ]
  },
  {
    "id": 4,
    "isActive": true,
    "balance": "$1,248.34",
    "age": 35,
    "name": "Letha Larsen",
    "gender": "female",
    "company": "DIGIPRINT",
    "email": "lethalarsen@digiprint.com",
    "phone": "+1 (817) 495-3533",
    "address": "967 Division Avenue, Groveville, Illinois, 3913",
    "about": "Elit anim pariatur Lorem in velit mollit deserunt incididunt ipsum amet esse et ullamco in. Qui ea ut mollit aute reprehenderit ipsum sit consequat ad. Laboris quis est exercitation et laboris officia sunt. Nostrud aliquip deserunt veniam deserunt reprehenderit velit eiusmod officia dolor sit aliquip. Aute irure incididunt eiusmod adipisicing ad deserunt sint labore.\r\n",
    "registered": "2014-01-02T03:09:35 -06:00",
    "latitude": 38.785346,
    "longitude": -134.565844,
    "products": [
    ]
  }
]`)

var productsData = []byte(`[
      {
        "id": 0,
        "name": "Page",
        "user_id": 5
      },
      {
        "id": 1,
        "name": "Mcintyre",
        "user_id": 5
      },
      {
        "id": 2,
        "name": "Black",
        "user_id": 3
      },
      
      {
        "id": 0,
        "name": "Hatfield",
        "user_id": 3
      },
      {
        "id": 1,
        "name": "Ochoa",
        "user_id": 0
      },
      {
        "id": 2,
        "name": "Mcgee",
        "user_id": 2
      },
      {
        "id": 0,
        "name": "Davidson",
        "user_id": 0
      },
      {
        "id": 1,
        "name": "Manning",
        "user_id": 1
      },
      {
        "id": 2,
        "name": "Melendez",
        "user_id": 2
      },
      {
        "id": 0,
        "name": "Molina",
        "user_id": 3
      },
      {
        "id": 1,
        "name": "Cooper",
        "user_id": 4
      },
      {
        "id": 2,
        "name": "Stephens",
        "user_id": 4
      },
      {
        "id": 0,
        "name": "Rivas",
        "user_id": 1
      },
      {
        "id": 1,
        "name": "Stein",
        "user_id": 1
      },
      {
        "id": 2,
        "name": "Atkins",
        "user_id": 5
      }
]`)
