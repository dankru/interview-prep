package main

import "fmt"

type MyError struct{}

func (MyError) Error() string {
	return "MyError"
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	var err *MyError  // Вот тут мы создаём переменную типа Указатель. При создании переменной у неё будет zerovalue, поэтому это nil
	errorHandler(err) // nil потому что zeroValue для указателя это nil

	err = &MyError{}  // Вот тут мы уже создаём Пустую структуру
	errorHandler(err) // не nil, потому что поле интерфейса itab содержит тип MyError
}
