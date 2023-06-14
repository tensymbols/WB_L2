package main

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
## Паттерн Visitor относится к поведенческим паттернам уровня объекта.

 Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
 а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс этого
 объекта.

## Применение
	Когда необходимо изменить поведене структуры, не меняя при этом структуру
*/

import (
	"fmt"
)

// структура квадрат
type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s) //  вызываем метод вычисления у посетителя
}

func (s *square) getType() string {
	return "Square"
}

// структура круг
type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

// структура треугольник
type triangle struct {
	h    int
	base int
}

func (t *triangle) accept(v visitor) {
	v.visitForTriangle(t)
}

func (t *triangle) getType() string {
	return "triangle"
}

// интерфейс посетителя
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForTriangle(*triangle)
}

// структура вычисления площади фигуры в зависимости от фигуры
type areaCalculator struct {
	area float64
}

func (a *areaCalculator) visitForSquare(s *square) {
	a.area = float64(s.side * s.side)
	fmt.Println("Calculating area for", s.getType(), "=", a.area)
}

func (a *areaCalculator) visitForCircle(c *circle) {
	a.area = 3.1415926 * float64(c.radius*c.radius)
	fmt.Println("Calculating area for", c.getType(), "=", a.area)
}

func (a *areaCalculator) visitForTriangle(t *triangle) {
	a.area = 0.5 * float64(t.base*t.h)
	fmt.Println("Calculating area for", t.getType(), "=", a.area)
}

func main() {
	//  создание фигур
	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &triangle{h: 2, base: 3}

	// создание струтуры для вычисления площади фигуры в зависимости от фигуры
	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

}
