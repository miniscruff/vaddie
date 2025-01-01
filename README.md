# Vaddy

[![Go Packge](https://img.shields.io/badge/Go-Reference-grey?style=for-the-badge&logo=go&logoColor=white&label=%20&labelColor=007D9C)](https://pkg.go.dev/github.com/miniscruff/vaddy)
[![GitHub release](https://img.shields.io/github/v/release/miniscruff/vaddy?style=for-the-badge&logo=github)](https://github.com/miniscruff/vaddy/releases)
[![GitHub License](https://img.shields.io/github/license/miniscruff/vaddy?style=for-the-badge)](https://github.com/miniscruff/vaddy/blob/main/LICENSE)

Validation library using type safe, developer friendly and extendible functions.

> [!WARNING]
> This library is strictly around validations, parsing or building is otherwise left
> up to you.
> Validations that would otherwise fail parsing, are not handled or expected to be supported.

## Support

Currently experimental, valdiations will be added, changed or removed.
New validations may be added with `Exp` prefixes indicating they are still experimental.

## Examples

* [User](./example_user_test.go)

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

## Need help?
Use the [discussions page](https://github.com/miniscruff/vaddy/discussions) for help requests and how-to questions.

Please open [GitHub issues](https://github.com/miniscruff/vaddy/issues) for bugs and feature requests.
File an issue before creating a pull request, unless it is something simple like a typo.

## Want to Contribute?
If you want to contribute through code or documentation, the [Contributing guide](CONTRIBUTING.md) is the place to start.
If you need additional help create an issue or post on discussions.
