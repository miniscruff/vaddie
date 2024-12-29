# Vaddy

_Badges coming soon_

Validation library using type safe, developer friendly and extendible functions.

> [!WARNING]
> This library is strictly around validations, parsing or building is otherwise left
> up to you.
> Validations that would otherwise fail parsing, are not handled or expected to be supported.

## Support

Currently experimental, valdiations will be added, changed or removed.
New validations may be added with `Exp` prefixes indicating they are still experimental.

## Example

```go
package main

import (
	"fmt"

	v "github.com/miniscruff/vaddy"
)

type User struct {
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	Age           uint8      `json:"age"`
}

func main() {
	user := &User{
		FirstName:     "Badger",
		LastName:      "Smith",
		Age:           45,
    }
    err := v.Join(
		v.AllOf(u.FirstName, "first_name", v.StrMin(2), v.StrMax(64)),
		v.AllOf(u.LastName, "last_name", v.StrMin(2), v.StrMax(64)),
		v.AllOf(u.Age, "age", v.OrderedGte(uint8(0)), v.OrderedLte(uint8(130))),
    )

	fmt.Printf("validating user:\n%v\n", err)
}
```

## Benchmarks

Coming soon

> [!NOTE]
> If you have a particularly large or complex validation requirement, please share.

## Contributing

Coming soon
