package pattern

import "fmt"

type BProduct struct {
	PartA string
	PartB string
	PartC string
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetResult() *BProduct
}

type ConcreteBuilder struct {
	product *BProduct
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{product: &BProduct{}}
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.PartA = "PartA"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.PartB = "PartB"
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.PartC = "PartC"
}

func (b *ConcreteBuilder) GetResult() *BProduct {
	return b.product
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() *BProduct {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
	return d.builder.GetResult()
}

func builder() {
	builder := NewConcreteBuilder()

	director := NewDirector(builder)

	product := director.Construct()

	fmt.Printf("Product parts: %s, %s, %s\n", product.PartA, product.PartB, product.PartC)
}

/*
### Применимость паттерна "Строитель":

1. Создание сложных объектов:
   - Применяется, когда процесс создания объекта состоит из множества шагов, и различные конфигурации объекта могут быть созданы с использованием одного и того же процесса строительства.

2. Конфигурирование объектов:
   - Применяется, когда клиент должен иметь гибкий способ конфигурирования сложных объектов, а также когда необходимо избежать "телескопического" конструктора.

3. Создание объектов с разными представлениями:
   - Применяется, когда необходимо создавать объекты с различными представлениями, но с использованием одной и той же основной логики построения.

### Плюсы паттерна "Строитель":

1. Отделение сложности конструирования от клиентского кода:
   - Клиентский код работает с директором и строителем, не заботясь о деталях конструирования объекта.

2. Гибкость и повторное использование:
   - Возможность пошагового конструирования объекта и переиспользования тех же шагов для создания различных конфигураций.

3. Изоляция кода конструирования:
   - Код конструирования объекта изолирован в отдельном классе строителя, что упрощает его изменение и поддержку.

### Минусы паттерна "Строитель":

1. Увеличение числа классов:
   - Может привести к созданию большого числа классов, особенно если у объекта много параметров.

2. Зависимость между строителем и продуктом:
   - Клиент должен знать о конкретных строителях и продуктах, что может нарушить инкапсуляцию.

### Примеры использования на практике:

1. Построение DOM-дерева в библиотеках для парсинга HTML:
   - Строитель может представлять собой процесс пошагового создания DOM-дерева при парсинге HTML-страницы.

2. Конструирование SQL-запросов:
   - Строитель может использоваться для пошагового формирования сложных SQL-запросов с различными условиями.

3. Конфигурирование объектов в GUI-библиотеках:
   - Строитель может помочь в пошаговом создании конфигурации элементов пользовательского интерфейса.
*/
