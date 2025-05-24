package main

import "fmt"

func main() {

	interviewtTask()
	//anotherExample()
}

func interviewtTask() {
	foo := make([]int, 0, 4)
	foo = append(foo, 1, 2, 3)
	bar := append(foo, 4)
	baz := append(foo, 5)

	fmt.Println(bar)
	fmt.Println(baz)
}

func anotherExample() {
	exampleSlice := make([]int, 0, 4)
	// Вот тут создался слайс
	// с полями cap = 4 len = 0, ptr==&xxx

	differentSlice := append(exampleSlice, 5)
	fmt.Println(exampleSlice)
	fmt.Println(differentSlice)

	// Когда мы передаём сюда этот слайс, то его структура передаётся по копии.
	// append принимает этот слайс, копирует у него поля len и cap и создаёт новый слайс,
	// а затем вставляет элемент в конец len.
	// Причём элемент вставится по ссылке на массив

	// В итоге мы получим
	// exampleSlice   = []  len = 0, cap = 4
	// differentSlice = [5] len = 1, cap = 4

	// Причём exampleSlice и differentSlice ссылаются на один массив.
	// Но т.к. в exampleSlice len = 0 то мы не увидим в нём элемент 5

	exampleSlice = append(exampleSlice, 1000)

	fmt.Println(exampleSlice)
	fmt.Println(differentSlice)

	// Но если мы переполним cap и произойдёт реаллокация, то массив у каждого слайса он будет свой
	exampleSlice = append(exampleSlice, 2, 3, 4, 5) // len = 5, cap = 8
	differentSlice[0] = 1

	// В итоге мы получим
	// exampleSlice   = [1000, 2, 3, 4, 5] 	len = 5, cap = 8, ptr = &xxx1
	// differentSlice = [1] 				len = 1, cap = 4, ptr = &xxx2
	fmt.Println(exampleSlice)
	fmt.Println(differentSlice)
}
