package examples

import (
	"fmt"
)

type LoggingInterceptor struct{}

func (this *LoggingInterceptor) Intercept(receiver interface{}, method string, args []interface{}, delegate func([]interface{}) []interface{}) []interface{} {
	fmt.Printf("ENTER: receiver = %+v, method = %s, args = %+v\n", receiver, method, args)
	rets := delegate(args)
	fmt.Printf("EXIT: receiver = %+v, method = %s, args = %+v, retvals = %+v\n", receiver, method, args, rets)
	return rets
}
