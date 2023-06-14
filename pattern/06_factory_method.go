package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
## Паттерн Фабричный метод
 порождающий паттерн проектирования, который определяет общий интерфейс
 для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

## Преимущества и недостатки
+ Избавляет класс от привязки к конкретным классам продуктов.
+ Выделяет код производства продуктов в одно место, упрощая поддержку кода.
+ Упрощает добавление новых продуктов в программу.
+ Реализует принцип открытости/закрытости.
- Может привести к созданию больших параллельных иерархий классов, так как для каждого класса
  продукта надо создать свой подкласс создателя.

## Применение
 Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
 Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки.
 Когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо
 порождения новых.
*/

// IToy - интерфейс продукта
type IToy interface {
	setName(name string)
	setPrice(power int)
	getName() string
	getPrice() int
}

// Toy - обобщенный продукт
type Toy struct {
	name  string
	price int
}

func (g *Toy) setName(name string) {
	g.name = name
}

func (g *Toy) getName() string {
	return g.name
}

func (g *Toy) setPrice(price int) {
	g.price = price
}

func (g *Toy) getPrice() int {
	return g.price
}

// Doll - конкретный продукт (кукла)
type Doll struct {
	Toy
}

func newDoll() IToy {
	return &Doll{
		Toy: Toy{
			name:  "Barbie doll",
			price: 2000,
		},
	}
}

// Doll - конкретный продукт (солдатик)
type Soldier struct {
	Toy
}

func newSoldier() IToy {
	return &Soldier{
		Toy: Toy{
			name:  "Soldier",
			price: 1500,
		},
	}
}

func getToy(gunType string) (IToy, error) {
	if gunType == "doll" {
		return newDoll(), nil
	}
	if gunType == "soldier" {
		return newSoldier(), nil
	}
	return nil, fmt.Errorf("Wrong toy type passed")
}

func main() {
	doll, _ := getToy("doll")
	soldier, _ := getToy("soldier")

	printDetails(doll)
	printDetails(soldier)
}

func printDetails(g IToy) {
	fmt.Printf("Toy: %s", g.getName())
	fmt.Println()
	fmt.Printf("Price: %d", g.getPrice())
	fmt.Println()
}
