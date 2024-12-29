package main

import (
	"fmt"

	v "github.com/miniscruff/vaddy"
)

type User struct {
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	Age           uint8      `json:"age"`
	Email         string     `json:"email"`
	FavoriteColor string     `json:"favoriteColor"`
	Hobbies       []string   `json:"hobbies"`
	Addresses     []*Address `json:"addresses"`
}

func (u *User) Validate() error {
	return v.Join(
		v.AllOf(u.FirstName, "first_name", v.StrMin(2), v.StrMax(64)),
		v.AllOf(u.LastName, "last_name", v.StrMin(2), v.StrMax(64)),
		v.AllOf(u.Age, "age", v.OrderedGte(uint8(0)), v.OrderedLte(uint8(130))),
		v.AllOf(u.Email, "email", v.StrNotEmpty()), // no email check
		v.AllOf(u.FavoriteColor, "favorite_color",
			v.StrNotEmpty(),
		),
		// For some reason SliceMinLength cannot infer *Address
		v.All(u.Addresses, "addresses", v.SliceMinLength[*Address](1)),
		v.All(u.Hobbies, "hobbies",
			v.SliceMinLength[string](1),
			v.Dive(v.StrMin(3), v.StrMax(64)),
		),
	)
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Planet string `json:"planet"`
	Phone  string `json:"phone"`
}

func (a *Address) Validate() error {
	return v.Join(
		v.AllOf(a.Street, "street", v.StrNotEmpty()),
		v.AllOf(a.City, "city", v.StrNotEmpty()),
		v.AllOf(a.Planet, "planet", v.StrNotEmpty()),
		v.AllOf(a.Phone, "phone", v.StrNotEmpty()),
	)
}

func main() {
	// Edit to see the output of the validations
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
		City:   "Unknown",
	}

	user := &User{
		FirstName:     "Badger",
		LastName:      "Smith",
		Age:           45,
		Email:         "Badger.Smith@gmail",
		FavoriteColor: "#000",
		Addresses:     []*Address{address},
		Hobbies:       []string{"RC Cars"},
	}

	fmt.Printf("validating user:\n%v\n", user.Validate())
}
