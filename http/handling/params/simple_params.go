package params

import (
	"errors"
	"reflect"
)

func getParametersFromURLValues(funcType reflect.Type, urlVars []string) (params []reflect.Value, err error) {
	if len(urlVars) == funcType.NumIn()-1 {
		params = make([]reflect.Value, funcType.NumIn()-1)
		for i := 0; i < len(urlVars); i++ {
			params[i], err = parseValueToType(funcType.In(i+1), urlVars[i])
			if err != nil {
				return
			}
		}
	} else {
		err = errors.New("Parameter number mismatch")
	}
	return
}
