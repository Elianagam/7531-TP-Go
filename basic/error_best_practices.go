package main

import "fmt"

type ErrZeroDivision struct {
	message string
}

func NewErrZeroDivision(message string) *ErrZeroDivision {
	return &ErrZeroDivision{
		message: message,
	}
}

func (e *ErrZeroDivision) Error() string {
	return e.message
}

func main() {
	result, err := divide(1.0, 0.0)
	if err != nil {
		switch err.(type) {
		case *ErrZeroDivision:
			fmt.Println(err.Error())
		default:
			fmt.Println("¿Qué pasó?")
		}
	}
	fmt.Println(result)
}
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, NewErrZeroDivision("ZeroDivisionError")
	}
	return a / b, nil
}
