package vaddie

import (
	"errors"
	"testing"
)

func Test_JoinAndWithError(t *testing.T) {
	a := errors.New("a")
	if JoinAnd(nil, a, nil) != a {
		t.Fail()
	}
}

func Test_JoinAndNoError(t *testing.T) {
	if JoinAnd(nil, nil) != nil {
		t.Fail()
	}
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
				Index:   toPtr(5),
			},
			msg: "whichThing[5] should be fixed ( keep it in check )",
		},
	} {
		if tc.err.Error() != tc.msg {
			t.Errorf("%v != %v", tc.err.Error(), tc.msg)
		}
	}
}
