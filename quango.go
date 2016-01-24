package quango

import (
	"fmt"
	"reflect"
	"testing/quick"
)

type quangoMatcher struct {
	counterexample string
}

func Hold() *quangoMatcher {
	return &quangoMatcher{}
}

func (q *quangoMatcher) Match(userPredicate interface{}) (successfulMatch bool, matchErr error) {

	typeOfPredicate := reflect.TypeOf(userPredicate)

	if typeOfPredicate.Kind() != reflect.Func {
		return false, fmt.Errorf("Expected a function, not a %s", typeOfPredicate)
	}

	cleanedPredicate := userPredicate

	if typeOfPredicate.NumOut() == 0 {
		predicateValue := reflect.ValueOf(userPredicate)

		predicateWrapper := func(args []reflect.Value) (returnVals []reflect.Value) {
			returnVals = []reflect.Value{reflect.ValueOf(true)}

			defer func() {
				err := recover()
				if err != nil {
					returnVals = []reflect.Value{reflect.ValueOf(false)}
				}
			}()

			predicateValue.Call(args)

			return
		}

		inType := []reflect.Type{}
		for i := 0; i < typeOfPredicate.NumIn(); i++ {
			inType = append(inType, typeOfPredicate.In(i))
		}

		predicateType := reflect.FuncOf(inType, []reflect.Type{reflect.TypeOf(true)}, false)
		cleanedPredicate = reflect.MakeFunc(predicateType, predicateWrapper).Interface()
	}

	err := quick.Check(cleanedPredicate, nil)

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
