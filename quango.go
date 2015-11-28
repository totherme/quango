package quango

import (
	"fmt"
	"testing/quick"
)

type quangoMatcher struct {
	counterexample string
}

func Hold() *quangoMatcher {
	return &quangoMatcher{}
}

func (q *quangoMatcher) Match(actual interface{}) (bool, error) {
	err := quick.Check(actual, nil)

	q.counterexample = "False"
	if err != nil {
		cerr, ok := err.(*quick.CheckError)
		if !ok {
			return false, err
		}
		q.counterexample = fmt.Sprintf("%s", cerr)
		return false, nil
	}

	return true, nil
}

func (q *quangoMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Property to failed on test %s.", q.counterexample)
}

func (*quangoMatcher) NegatedFailureMessage(actual interface{}) string {
	return "Expected property not to hold. It did."
}
