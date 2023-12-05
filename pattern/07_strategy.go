package pattern

import "fmt"

type Strategy interface {
	Execute()
}

type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute() {
	fmt.Println("Executing ConcreteStrategyA")
}

type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute() {
	fmt.Println("Executing ConcreteStrategyB")
}

type SContext struct {
	strategy Strategy
}

func (c *SContext) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *SContext) ExecuteStrategy() {
	c.strategy.Execute()
}

func strategy() {
	context := &SContext{}

	context.SetStrategy(&ConcreteStrategyA{})
	context.ExecuteStrategy()

	context.SetStrategy(&ConcreteStrategyB{})
	context.ExecuteStrategy()
}

/*
### Применимость паттерна "Стратегия":

Этот паттерн полезен, когда необходимо предоставить различные способы выполнения определенной задачи и
дать возможность клиенту выбирать стратегию выполнения в зависимости от контекста.

### Плюсы паттерна "Стратегия":

1. Позволяет легко добавлять новые стратегии или изменять существующие без изменения клиентского кода.

2. Улучшает читаемость кода, так как каждая стратегия инкапсулирована в отдельном классе.

3. Помогает избежать больших блоков условных операторов, заменяя их на вызовы соответствующих стратегий.

### Минусы паттерна "Стратегия":

1. Может привести к увеличению числа классов в системе, особенно если есть много различных стратегий.

2. Клиент должен знать о существовании различных стратегий и выбирать подходящую, что может быть сложно.

### Примеры использования на практике:

1. Функции сравнения, используемые в алгоритмах сортировки

2. Стратегии доставки в интернет-магазине
*/
