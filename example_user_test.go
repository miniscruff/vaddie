package vaddie_test

import (
	"fmt"

	v "github.com/miniscruff/vaddie"
)

type User struct {
	FirstName     string
	LastName      string
	Age           int
	Email         string
	FavoriteColor string
	Hobbies       []string
	Addresses     []*Address
}

func (u *User) Validate() error {
	return v.Join(
		v.AllOf(u.FirstName, "first_name", v.StrMin(2), v.StrMax(64)),
		v.AllOf(u.LastName, "last_name", v.StrMin(2), v.StrMax(64)),
		v.AllOf(u.Age, "age", v.OrderedGte(0), v.OrderedLte(130)),
		v.AllOf(u.Email, "email", v.StrEmail()),
		v.AllOf(u.FavoriteColor, "favorite_color",
			v.Not(v.StrEmpty()),
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
	Street string
	City   string
	Planet string
	Phone  string
}

func (a *Address) Validate() error {
	return v.Join(
		v.AllOf(a.Street, "street", v.Not(v.StrEmpty())),
		v.AllOf(a.City, "city", v.Not(v.StrEmpty())),
		v.AllOf(a.Planet, "planet", v.Not(v.StrEmpty())),
		v.AllOf(a.Phone, "phone", v.Not(v.StrEmpty())),
	)
}

func Example() {
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
	// Output
	// validating user:
	// <nil>
}
