package vaddy

import (
	"fmt"
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
	return Join(
		AllOf(u.FirstName, "first_name", StrMin(2), StrMax(64)),
		AllOf(u.LastName, "last_name", StrMin(2), StrMax(64)),
		AllOf(u.Age, "age", OrderedGte(0), OrderedLte(130)),
		AllOf(u.Email, "email", StrNotEmpty()), // no email check
		AllOf(u.FavoriteColor, "favorite_color",
			StrNotEmpty(),
		),
		// For some reason SliceMinLength cannot infer *Address
		All(u.Addresses, "addresses", SliceMinLength[*Address](1)),
		All(u.Hobbies, "hobbies",
			SliceMinLength[string](1),
			Dive(StrMin(3), StrMax(64)),
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
	return Join(
		AllOf(a.Street, "street", StrNotEmpty()),
		AllOf(a.City, "city", StrNotEmpty()),
		AllOf(a.Planet, "planet", StrNotEmpty()),
		AllOf(a.Phone, "phone", StrNotEmpty()),
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
