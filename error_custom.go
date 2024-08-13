package main

import (
	"fmt"
	//"errors"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "afif" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	
	// code pak eko
	// if err != nil {
	// 	//error
	// 	if validationErr, ok := err.(*validationError); ok {
	// 		fmt.Println("validation error:", validationErr.Error())
	// 	} else if notFoundErr, ok := err.(*notFoundError); ok {
	// 		fmt.Println("not found error:", notFoundErr.Error())
	// 	} else {
	// 		fmt.Println("unknown:", err.Error())
	// 	}
	// } else {
	// 	// success
	// 	fmt.Println("save data success")
	// }

	//code pak eko
	switch finalErr := err.(type) {
	case *validationError:
		fmt.Println("validation error:", finalErr.Error())
	case *notFoundError:
		fmt.Println("not found error:", finalErr.Error())
	default:
		fmt.Println("unknown:", finalErr.Error())
	}

	// code apip
	// if err == nil {
	// 	fmt.Println("inject data success")
	// } else {
	// 	fmt.Println(err.Error())
	// }


}