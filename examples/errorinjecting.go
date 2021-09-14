package examples

import (
	"fmt"
	"math/rand"
	"reflect"
)

type ErrorInjectingInterceptor struct {
	Random             *rand.Rand
	FailureProbability float32
}

var (
	ErrInjectedFailure = fmt.Errorf("injected error")
)

func canReturnError(method reflect.Value) bool {
	if method.Type().NumOut() == 0 {
		return false
	}
	lastReturnType := method.Type().Out(method.Type().NumOut() - 1)
	return lastReturnType.Name() == "error" && lastReturnType.PkgPath() == ""
}

func (this *ErrorInjectingInterceptor) Intercept(receiver interface{}, method string, args []interface{}, delegate func([]interface{}) []interface{}) []interface{} {
	r := reflect.ValueOf(receiver)
	m := r.MethodByName(method)
	if canReturnError(m) && this.Random.Float32() < this.FailureProbability {
		rets := []interface{}{}
		for i := 0; i < m.Type().NumOut()-1; i += 1 {
			rets = append(rets, nil)
		}
		rets = append(rets, ErrInjectedFailure)
		return rets
	}
	return delegate(args)
}
