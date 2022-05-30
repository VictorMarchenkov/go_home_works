package main

import (
	"fmt"
	"time"
)

// CustomError error expands with time stamp
type CustomError struct {
	message  string
	dateTime time.Time
}

// Error modifies the method with time stamp to return
func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.dateTime, e.message)
}

func main() {
	array := []int{0, 1, 2, 3, 4}
	defer func() {
		if v := recover(); v != nil {
			//fmt.Println("recovered: ", v)
			err := CustomError{
				message:  fmt.Sprintf("recovered after: %v", v),
				dateTime: time.Now(),
			}
			fmt.Println(err)
		}
	}()
	for i := 0; i < 6; i++ {
		fmt.Println(array[i])
	}
}
