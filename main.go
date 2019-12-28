package main

import (
	"errors"
	"fmt"
)

// infrastructure/repository.go
type repository interface {
	save() bool
}

func SaveToDB(r repository) {
	r.save()
}

//EOF infrastructure/repository.go

// domain/user.go
var ErrEmptyOrderID = errors.New("empty order id")

type ID string

type User struct {
	ID   ID
	Name string
}

func NewUser(id ID, name string) (*User, error) {
	if len(id) < 0 {
		return nil, ErrEmptyOrderID
	}
	return &User{id, name}, nil
}

//EOF domain/user.go

// infrastructure/user.go
func (u User) save() bool {
	fmt.Println("Save to user database", u)
	return true
}

//EOF infrastructure/user.go

func main() {
	// interface/user.go
	user := User{ID: "123", Name: "PAR1234H1"}
	//EOF interface/user.go

	// application/user.go
	req, err := NewUser(user.ID, user.Name)
	if err != nil {
		panic("Error")
	}
	SaveToDB(req)
	//EOF In application/user.go

}
