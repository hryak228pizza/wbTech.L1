package main

import "fmt"

type Human struct {
	Age uint8
	Name string
	Height uint16
	Weight uint16
}

func (h *Human) getAge() uint8 {
	return h.Age
}

func (h *Human) getName() string {
	return h.Name
}

func (h *Human) setHeight(height uint16) {
	h.Height = height
}

func (h *Human) getHeight() uint16 {
	return h.Height 
}

func (h *Human) startWork() {
	fmt.Println(h.Name + " начал работать")
}

func (h *Human) finishWork() {
	fmt.Println(h.Name + " закончил работать")
}


type Action struct {
	Human
}

func main() {
	// создаем Джона
	John := Human{Age: 20, Name: "John", Height: 180, Weight: 70} 
	// создаем переменную для управления Джоном
	actionOnJohn := Action{Human: John}

	// выводим возраст Джона
	fmt.Println(actionOnJohn.getAge())
	// выводим имя Джона
	fmt.Println(actionOnJohn.getName())

	// Джон подрос
	actionOnJohn.setHeight(185)
	// выводим рост Джона
	fmt.Println(actionOnJohn.getHeight())

	// заставляем Джона поработать
	actionOnJohn.startWork()
	// пока не закончится его рабочий день
	actionOnJohn.finishWork()
}