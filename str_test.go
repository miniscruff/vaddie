package vaddie

import (
	"regexp"
	"testing"
)

var strTests = []TestCase[string]{
	{
		Name:          "empty",
		ValidValues:   []string{""},
		InvalidValues: []string{"ab"},
		Validation:    StrEmpty(),
	},
	{
		Name:          "space",
		ValidValues:   []string{" \t\n"},
		InvalidValues: []string{"ab"},
		Validation:    StrSpace(),
	},
	{
		Name:          "min",
		ValidValues:   []string{"abcdefg", "🐈🐕🐢⏰🐦"},
		InvalidValues: []string{"ab"},
		Validation:    StrMin(5),
	},
	{
		Name:          "max",
		ValidValues:   []string{"abc"},
		InvalidValues: []string{"abcdefg", "🐈🐕🐢⏰🐦"},
		Validation:    StrMax(5),
	},
	{
		Name:          "unicode min",
		ValidValues:   []string{"abcdefg", "🐈🐕🐢⏰🐦🐈🐕🐢⏰🐦"},
		InvalidValues: []string{"ab", "🐈"},
		Validation:    StrUnicodeMin(5),
	},
	{
		Name:          "unicode max",
		ValidValues:   []string{"🐈🐕🐢⏰🐦"},
		InvalidValues: []string{"🐈🐕🐢⏰🐦🐈🐕🐢⏰🐦"},
		Validation:    StrUnicodeMax(8),
	},
	{
		Name:          "letters",
		ValidValues:   []string{"abcd"},
		InvalidValues: []string{"abcd1"},
		Validation:    StrLetters(),
	},
	{
		Name:          "ascii",
		ValidValues:   []string{"abcd"},
		InvalidValues: []string{"abcd😎"},
		Validation:    StrAscii(),
	},
	{
		Name:          "has prefix",
		ValidValues:   []string{"abcd"},
		InvalidValues: []string{"def"},
		Validation:    StrHasPrefix("abc"),
	},
	{
		Name:          "has suffix",
		ValidValues:   []string{"aaawxyz"},
		InvalidValues: []string{"aaabcde"},
		Validation:    StrHasSuffix("xyz"),
	},
	{
		Name:          "contains",
		ValidValues:   []string{"things.com"},
		InvalidValues: []string{"without a dot"},
		Validation:    StrContains("."),
	},
	{
		Name:          "contains any",
		ValidValues:   []string{"A", "ABCD", "alskdjflkasdfA"},
		InvalidValues: []string{"B", "UIOUP"},
		Validation:    StrContainsAny("A"),
	},
	{
		Name:          "match",
		ValidValues:   []string{"abc", "abcdefg"},
		InvalidValues: []string{"def", "ABCD", "01234"},
		Validation:    StrMatch("abc.*"),
	},
	{
		Name:          "match w/ invalid regex",
		InvalidValues: []string{"def", "ABCD", "01234"},
		Validation:    StrMatch("..**\\("),
	},
	{
		Name: "email",
		ValidValues: []string{
			"thing@gmail.com",
			"test.withdot@hotmail.com",
			"test+withplus@yahoo.com",
		},
		InvalidValues: []string{"def", "ABCD", "01234"},
		Validation:    StrEmail(),
	},
	{
		Name:          "regexp",
		ValidValues:   []string{"abz", "abd"},
		InvalidValues: []string{"zdf"},
		Validation:    StrRegexp(regexp.MustCompile("ab?")),
	},
}

func Test_Strings(t *testing.T) {
	for _, tc := range strTests {
		tc.Run(t)
	}
}
