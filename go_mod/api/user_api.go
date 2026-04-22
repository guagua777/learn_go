package api

import (
	"go_web/service"
)

func GetResult() int {
	return service.Calc()
}
