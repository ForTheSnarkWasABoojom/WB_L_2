package pattern

import "fmt"

type Product interface {
	Use() string
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using ConcreteProductA"
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using ConcreteProductB"
}

type Creator interface {
	CreateProduct() Product
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func factory_method() {
	creatorA := &ConcreteCreatorA{}
	productA := creatorA.CreateProduct()
	fmt.Println(productA.Use())

	creatorB := &ConcreteCreatorB{}
	productB := creatorB.CreateProduct()
	fmt.Println(productB.Use())
}

/*
### Применимость паттерна "Фабричный метод":

Применяется, когда базовый класс предоставляет интерфейс для создания объектов, но тип создаваемого объекта оставляется для наследующих классов.

### Плюсы паттерна "Фабричный метод":

1. Обеспечивает гибкость и расширяемость, позволяя добавлять новые классы продуктов и создателей без изменения существующего кода.

2. Разделяет процесс создания объекта от его использования, что способствует более низкой связанности в системе.

### Минусы паттерна "Фабричный метод":

1. Может привести к усложнению структуры кода из-за большого числа интерфейсов и классов.

2. При добавлении нового продукта необходимо создавать новый подкласс создателя, что может увеличивать количество классов в системе.

### Примеры использования на практике:

1. Приложения, содержащие мини-игры

2. Приложения, содержащие множество инструментов, например, графические редакторы
*/
