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

func (q *quangoMatcher) Match(actual interface{}) (bool, error) {
	typeOfActual := reflect.TypeOf(actual)
	if typeOfActual.Kind() != reflect.Func {
		return false, fmt.Errorf("Expected a function, not a %s", typeOfActual)
	}
	if typeOfActual.NumOut() == 0 {
		actualValue := reflect.ValueOf(actual)
		newActual := func(args []reflect.Value) []reflect.Value {
			defer func() []reflect.Value {
				recover()
				return []reflect.Value{reflect.ValueOf(false)}
			}()

			actualValue.Call(args)
			return []reflect.Value{reflect.ValueOf(true)}
		}

		inType := []reflect.Type{}
		for i := 0; i < typeOfActual.NumIn(); i++ {
			inType = append(inType, typeOfActual.In(i))
		}

		funcType := reflect.FuncOf(inType, []reflect.Type{reflect.TypeOf(true)}, false)
		actual = reflect.MakeFunc(funcType, newActual).Interface()
	}

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
