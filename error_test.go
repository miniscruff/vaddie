package vaddie

import (
	"errors"
	"testing"

	"github.com/miniscruff/vaddie/then"
)

func Test_JoinAndWithError(t *testing.T) {
	a := errors.New("a")
	then.Equals(t, a, JoinAnd(nil, a, nil))
}

func Test_JoinAndNoError(t *testing.T) {
	then.Nil(t, JoinAnd(nil, nil))
}

func Test_ValidationError_Message(t *testing.T) {
	for _, tc := range []struct {
		err error
		msg string
	}{
		{
			err: &ValidationError{
				Key:     "value",
				Message: "did not pass",
				Help:    "",
				Index:   nil,
			},
			msg: "value did not pass",
		},
		{
			err: &ValidationError{
				Key:     "anotherThing",
				Message: "is not the best",
				Help:    "try using a hammer",
				Index:   nil,
			},
			msg: "anotherThing is not the best ( try using a hammer )",
		},
		{
			err: &ValidationError{
				Key:     "whichThing",
				Message: "should be fixed",
				Help:    "keep it in check",
				Index:   new(5),
			},
			msg: "whichThing[5] should be fixed ( keep it in check )",
		},
	} {
		then.Equals(t, tc.msg, tc.err.Error())
	}
}
