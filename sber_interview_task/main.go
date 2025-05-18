package main

import (
	"fmt"
	"time"
)

/*
1. Сделать структуры Base и Child
2. Структура Base должна содержать строковое поле name
3. Структура Child должна содержать строковое поле lastName
4. Сделать функцию Say у структуры Base, которая распечатывает на экране "Hello, <name>!"
5. Пронаследовать Child от Base
6. Инициализировать экземпляр b1 Base
	присвоить name значение Parent
7. Инициализировать экземпляр c1 Child
	привсвоить name значение Child
	присвоить lastName значение inherited
8. Вызывать у обоих экземпляров метод Say
9. Переопределить метод Say для структуры Child, чтобы он выводил на экран: Hello, %lastName %name!
10. Сделать массив, содержащий b1 и c1
11. Вызывать Say у всех элементов массива из шага 10
12. Сделать метод NewObject для создания экземпляров Base и Child в зависимости от входного параметра
13. Написать юнит тесты для метода NewObject
14. Сделать генератор объектов Base и Child такой, чтобы:
	Объекты Base создавались в фоновом потоке с задержкой в 1 секунду
	Объекты Child создавались в фоновом потоке с задержкой в 2 секунды
	Общее время генерации объектов не превышало 11 секунд
15. Сделать асинхронный обработчик сгенерированных объектов такой, чтобы:
	Метод Say вызывался в порядке генерации объектов
	Не приводил к утечкам памяти
*/

type Sayer interface {
	Say()
}

type Base struct {
	name string
}

type Child struct {
	Base
	lastName string
}

func (b Base) Say() {
	fmt.Printf("Hello, %s\n", b.name)
}

func (c Child) Say() {
	fmt.Printf("Hello, %s %s\n", c.lastName, c.name)
}

func NewObject(kind string) Sayer {
	switch kind {
	case "child":
		return Child{
			Base:     Base{name: "Parent"},
			lastName: "Inherited",
		}
	case "base":
		return Base{name: "Parent"}
	}
	return nil
}

func generator() <-chan Sayer {
	out := make(chan Sayer)
	t := time.NewTicker(time.Second * 11)
	tChild := time.NewTicker(time.Second * 2)
	tBase := time.NewTicker(time.Second)
	go func() {
		defer t.Stop()
		defer tChild.Stop()
		defer tBase.Stop()
		defer close(out)
		for {
			select {
			case <-t.C:
				return
			case <-tChild.C:
				out <- NewObject("child")
				break
			case <-tBase.C:
				out <- NewObject("base")
				break
			}
		}
	}()
	return out
}

func worker(in <-chan Sayer) {
	for v := range in {
		v.Say()
	}
}

func main() {
	//b1 := Base{name: "Parent"}
	//c1 := Child{
	//	Base:     b1,
	//	lastName: "Inherited",
	//}
	//
	//b1.Say()
	//c1.Say()
	//
	//arr := [2]Sayer{b1, c1}
	//for _, v := range arr {
	//	v.Say()
	//}

	in := generator()
	worker(in)
}
