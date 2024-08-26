package defined

import "fmt"

type ActionError struct {
	ErrorCode    int
	ErrorMessage string
}

func (e ActionError) Error() string {
	return fmt.Sprintf("错误 %d: %s", e.ErrorCode, e.ErrorMessage)
}
