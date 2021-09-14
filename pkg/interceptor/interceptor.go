package interceptor

import (
	"reflect"
)

// Interceptor

type Interceptor interface {
	Intercept(receiver interface{}, method string, args []interface{}, delegate func([]interface{}) []interface{}) (rets []interface{})
}

type InterceptingInvocationHandler struct {
	Delegate    interface{}
	Interceptor Interceptor
}

func (this *InterceptingInvocationHandler) Invoke(method string, args []interface{}) []interface{} {
	delegate := func(args []interface{}) []interface{} {
		r := reflect.ValueOf(this.Delegate)
		m := r.MethodByName(method)

		vargs := []reflect.Value{}
		for _, arg := range args {
			vargs = append(vargs, reflect.ValueOf(arg))
		}

		vrets := m.Call(vargs)

		rets := []interface{}{}
		for _, vret := range vrets {
			rets = append(rets, vret.Interface())
		}

		return rets
	}
	rets := this.Interceptor.Intercept(this.Delegate, method, args, delegate)
	return rets
}

// Interceptors

type Interceptors []Interceptor

func (this Interceptors) Intercept(receiver interface{}, method string, args []interface{}, delegate func([]interface{}) []interface{}) []interface{} {
	if len(this) == 0 {
		return delegate(args)
	}

	head := this[0]
	tails := Interceptors(this[1:])

	recursive := func(args []interface{}) []interface{} {
		return tails.Intercept(receiver, method, args, delegate)
	}

	return head.Intercept(receiver, method, args, recursive)
}
