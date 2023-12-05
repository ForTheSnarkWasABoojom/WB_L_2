package pattern

import "fmt"

type SubsystemA struct {
}

func (s *SubsystemA) OperationA() {
	fmt.Println("SubsystemA: OperationA")
}

type SubsystemB struct {
}

func (s *SubsystemB) OperationB() {
	fmt.Println("SubsystemB: OperationB")
}

type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

func FacadeClient(facade *Facade) {
	fmt.Println("Client: Using the Facade to interact with subsystems.")
	facade.subsystemA.OperationA()
	facade.subsystemB.OperationB()
}

func facade() {
	facade := NewFacade()
	FacadeClient(facade)
}

/*
## Применимость паттерна "Фасад":

Данный паттерн применяется, когда необходимо предоставить единый интерфейс для взаимодействия с несколькими подсистемами.

Плюсы паттерна "Фасад":

1. Позволяет клиентам взаимодействовать с подсистемой, не заботясь о её сложной внутренней структуре.

2. Уменьшает зависимости клиентского кода от компонентов подсистемы

3. Пользователь взаимодействует только с фасадом, что снижает осведомленность пользователя о внутренней работе системы.

### Минусы паттерна "Фасад":

1. Может не предоставлять достаточной гибкости для некоторых случаев использования

2. Высокая степень зависимости от фасада

### Примеры использования на практике:

1. Графические библиотеки:
2. Библиотеки работы с базами данных
*/
