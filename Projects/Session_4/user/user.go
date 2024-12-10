package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Address struct {
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
	Text string  `json:"text"`
}

type User struct {
	Name    string  `json:"name"`
	Age     uint    `json:"age"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
}

func New(name string, age uint, email string, phone string, address Address) (*User, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if email == "" {
		return nil, errors.New("email is required")
	}

	if phone == "" {
		return nil, errors.New("phone is required")
	}

	return &User{
		Name:    name,
		Age:     age,
		Email:   email,
		Phone:   phone,
		Address: address,
	}, nil

}

func (user *User) Write() error {

	value, err := json.Marshal(user)

	if err != nil {
		return errors.New("json marshal fail")
	}

	fileName := fmt.Sprintf("%s.json", strings.ToLower(user.Name))
	err = os.WriteFile(fileName, value, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) Log() {
	fmt.Printf("Name: %s\n Age: %d\n Email: %s\n Phone: %s\n", user.Name, user.Age, user.Email, user.Phone)
}
