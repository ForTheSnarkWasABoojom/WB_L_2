package pattern

import "fmt"

type Command interface {
	Execute()
}

type Receiver struct {
	Name string
}

func (r *Receiver) Action() {
	fmt.Printf("%s is performing action\n", r.Name)
}

type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func command() {
	receiver := &Receiver{Name: "Receiver"}

	command := NewConcreteCommand(receiver)

	invoker := &Invoker{}
	invoker.SetCommand(command)

	invoker.ExecuteCommand()
}

/*
### Применимость паттерна "Команда":

Применяется, когда необходимо отделить объект, который инициирует операцию (отправитель),
от объекта, который фактически выполняет операцию (получатель).

### Плюсы паттерна "Команда":

1. Позволяет легко добавлять новые команды и получателей без изменения существующего кода.

2. Логика выполнения операций сосредоточена в командах, что упрощает поддержку кода

### Минусы паттерна "Команда":

1. Код может стать очень сложным, если каждая операция требует создания нового класса команды.

2. Создание отдельных классов команд для каждой операции может привести к излишнему увеличению числа классов

### Примеры использования на практике:

1. Работа с текстовыми редакторами

2. Управление устройствами в умных домах
*/
