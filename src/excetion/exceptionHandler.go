package excetion

import (
	"fmt"
)

func CommonExceptionHandler[T any](value T, err error) T {
	if err != nil {
		fmt.Printf("Exception Message: %+v", err.Error())
	}
	return value
}
