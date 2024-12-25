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
		v.Is(u.FirstName, "first_name", v.StrMin(2), v.StrMax(64)),
		v.Is(u.LastName, "last_name", v.StrMin(2), v.StrMax(64)),
		v.Is(u.Age, "age", v.OrderedGte(uint8(0)), v.OrderedLte(uint8(130))),
		v.Is(u.Email, "email", v.StrNotEmpty()), // no email check yet
		v.Is(u.FavoriteColor, "favorite_color",
			v.StrNotEmpty(),
			//v.Or(
			// hex
			// rgb
			// rgba
			//),
		),
		// For some reason SliceMinLength cannot infer *Address
		v.All(u.Addresses, "addresses", v.SliceMinLength[*Address](1)),
		v.All(u.Hobbies, "hobbies",
			v.SliceMinLength[string](1),
			v.Dive(v.StrMin(3), v.StrMax(64)),
		),
	)
}

// Address houses a users address information
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Planet string `json:"planet"`
	Phone  string `json:"phone"`
}

func (a *Address) Validate() error {
	return v.Join(
		v.Is(a.Street, "street", v.StrNotEmpty()),
		v.Is(a.City, "city", v.StrNotEmpty()),
		v.Is(a.Planet, "planet", v.StrNotEmpty()),
		v.Is(a.Phone, "phone", v.StrNotEmpty()),
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
