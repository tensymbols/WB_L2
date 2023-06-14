package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
## Паттерн строситель - порождающий шаблон проектирования предоставляет способ создания составного объекта

## Преимущества и недстатки
+  Позволяет создавать продукты пошагово.
+  Изолирует сложный код сборки продукта от его основной бизнес-логики.
+  Позволяет использовать один и тот же код для создания различных продуктов.
-  Усложняет код программы из-за введения дополнительных классов

## Применимость
 Когда нужен сложный продукт, который требует нескольких шагов для построения.
В таких случаях несколько конструкторных методов подойдут лучше, чем один громадный конструктор.

При использовании пошагового построения объектов потенциальной проблемой является выдача клиенту частично построенного
нестабильного продукта. Паттерн строитель скрывает объект до тех пор, пока он не построен до конца
*/

// Структура, порождающаяся от интерфейса ComputerBuilderI

type Computer struct {
	CPU string
	RAM int
	MB  string
}

// Интерфейс ComputerBuilderI предоставляющий все необходимые методы

type ComputerBuilderI interface {
	CPU(val string) ComputerBuilderI
	RAM(val int) ComputerBuilderI
	MB(val string) ComputerBuilderI

	Build() Computer
}

// Структура, имеющая все методы интерфейса ComputerBuilderI

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

// Конструктор данной струтуры

func NewComputerBuilder() ComputerBuilderI {
	return computerBuilder{}
}

// Методы структуры. Возвращаем ComputerBuilderI, чтобы при вызоые можно было сделать конструкцию типа объект.CPU("").RAM().MB("").Build()

func (cb computerBuilder) CPU(val string) ComputerBuilderI {
	cb.cpu = val
	return cb
}
func (cb computerBuilder) RAM(val int) ComputerBuilderI {
	cb.ram = val
	return cb
}
func (cb computerBuilder) MB(val string) ComputerBuilderI {
	cb.mb = val
	return cb
}

func (cb computerBuilder) Build() Computer {
	return Computer{
		CPU: cb.cpu,
		RAM: cb.ram,
		MB:  cb.mb,
	}
}

func main() {
	compBuilder := NewComputerBuilder()
	computer := compBuilder.CPU("core i3").RAM(8).MB("gigabyte").Build()
	fmt.Println(computer)
}
