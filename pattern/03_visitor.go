package pattern

import "fmt"

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor is visiting ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor is visiting ConcreteElementB")
}

func Client(elements []Element, visitor Visitor) {
	for _, element := range elements {
		element.Accept(visitor)
	}
}

func visitor() {
	elementA := &ConcreteElementA{}
	elementB := &ConcreteElementB{}

	visitor := &ConcreteVisitor{}

	Client([]Element{elementA, elementB}, visitor)
}

/*

## Применимость паттерна "Посетитель":

Применяется, когда необходимо выполнить какую-то операцию над всеми элементами сложной структуры

### Плюсы паттерна "Посетитель":

1. Упрощает добавление операций, работающих со сложными структурами объектов.
2. Объединяет родственные операции в одном классе.
3. Посетитель может накапливать состояние при обходе структуры элементов.

### Минусы паттерна "Посетитель":

1. Паттерн не оправдан, если иерархия элементов часто меняется.
2. Может привести к нарушению инкапсуляции элементов.

### Примеры использования на практике:

1. Обработка сложных древовидных структур.
2. Обработка XML-документов.
*/
