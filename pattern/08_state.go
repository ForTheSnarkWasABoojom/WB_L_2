package pattern

import "fmt"

type State interface {
	Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("Handling in ConcreteStateA")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("Handling in ConcreteStateB")
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle()
}

func state() {
	context := &Context{}

	context.SetState(&ConcreteStateA{})
	context.Request()

	context.SetState(&ConcreteStateB{})
	context.Request()
}

/*
### Применимость паттерна "Состояние":

Этот паттерн полезен, когда объект может изменять свое поведение в зависимости от своего внутреннего состояния, и эти состояния могут изменяться динамически.

### Плюсы паттерна "Состояние":

1. Каждое состояние инкапсулировано в отдельном классе, что обеспечивает изолированность и понятность кода.

2. Добавление новых состояний или изменение поведения существующих становится более простым без изменения остальной части кода.

### Минусы паттерна "Состояние":

1. Может привести к увеличению числа классов в системе, особенно если есть много различных состояний.

2. Если контекст имеет множество переходов между состояниями, это может сделать контекст сложным для понимания.

### Примеры использования на практике:

1. Редактор текста и различные состояние текста

2. Автомат с продажей напитков
*/
