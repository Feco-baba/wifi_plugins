package excetion

import (
	"fmt"
)

func CommonExceptionHandler(value string, err error) string {
	if err != nil {
		fmt.Printf(value+" Exception Message: %+v", err.Error())
	}
	return value
}
