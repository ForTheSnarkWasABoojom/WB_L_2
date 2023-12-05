package pattern

import "fmt"

type Handler interface {
	Handle(request int)
	SetSuccessor(successor Handler)
}

type ConcreteHandlerA struct {
	successor Handler
}

func (h *ConcreteHandlerA) Handle(request int) {
	if request < 10 {
		fmt.Println("ConcreteHandlerA handled the request")
	} else if h.successor != nil {
		h.successor.Handle(request)
	}
}

func (h *ConcreteHandlerA) SetSuccessor(successor Handler) {
	h.successor = successor
}

type ConcreteHandlerB struct {
	successor Handler
}

func (h *ConcreteHandlerB) Handle(request int) {
	if request >= 10 && request < 20 {
		fmt.Println("ConcreteHandlerB handled the request")
	} else if h.successor != nil {
		h.successor.Handle(request)
	}
}

func (h *ConcreteHandlerB) SetSuccessor(successor Handler) {
	h.successor = successor
}

func CORClient(handler Handler, requests []int) {
	for _, request := range requests {
		handler.Handle(request)
	}
}

func chain_of_resp() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetSuccessor(handlerB)

	CORClient(handlerA, []int{5, 12, 18, 25})
}

/*

Этот паттерн полезен, например, при создании системы фильтрации веб-запросов, где каждый фильтр может обрабатывать запрос и передавать его следующему фильтру в цепочке.

### Применимость паттерна "Цепочка обязанностей":

Этот паттерн полезен, например, при создании системы фильтрации веб-запросов, где каждый фильтр может обрабатывать запрос и передавать его следующему фильтру в цепочке.

### Плюсы паттерна "Цепочка обязанностей":

1. Гибкость и расширяемость:

2. Избежание жесткой привязки к конкретным обработчикам

3. Поддержка принципа единственной обязанности (Single Responsibility Principle), т.к.
каждый обработчик отвечает за свою часть логики, что соблюдает принцип единственной обязанности.

### Минусы паттерна "Цепочка обязанностей":

1. Нет гарантии, что запрос будет обработан каким-либо обработчиком.

2. Если цепочка не настроена правильно, запрос может "провалиться" до конца цепочки без обработки.

### Примеры использования на практике:

1. Системы фильтрации в веб-разработке

2. Системы логгирования
*/
