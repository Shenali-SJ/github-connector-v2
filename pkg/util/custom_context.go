package util

import (
"context"
"github.com/labstack/echo/v4"
)


type CustomContext struct {
	Echo echo.Context
	GoContext context.Context
	AppGoContext context.Context
}

func (a *CustomContext)GetContext(val string) interface{} {

	if a.GoContext != nil{
		result :=  a.GoContext.Value(val)
		if result != nil{
			return result
		}else {
			return a.AppGoContext.Value(val)
		}

	}else if a.Echo != nil{
		result := a.Echo.Get(val)
		if result != nil{
			return result
		}else {
			return a.AppGoContext.Value(val)
		}
	}
	return nil
}
