package handler

type InvocationHandler interface {
	Invoke(method string, args []interface{}) (retvals []interface{})
}
