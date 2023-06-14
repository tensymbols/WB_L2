package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
## Состояние
 поведенческий паттерн проектирования, который позволяет объектам менять
 поведение в зависимости от своего состояния. Извне создаётся впечатление,
 что изменился класс объекта.

## Преимущества и недостатки
+ Избавляет от множества больших условных операторов машины состояний.
+ Концентрирует в одном месте код, связанный с определённым состоянием.
+ Упрощает код контекста.
- Может неоправданно усложнить код, если состояний мало и они редко меняются.

## Применение
 Когда у вас есть объект, поведение которого кардинально меняется в зависимости
 от внутреннего состояния, причём типов состояний много, и их код часто меняется.

 Когда код класса содержит множество больших, похожих друг на друга, условных операторов,
 которые выбирают поведения в зависимости от текущих значений полей класса.
*/

// MobileAlertStater реализует общий интерфейс для всех состояний
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert реализует оповещение в зависимости от состояния
type MobileAlert struct {
	state MobileAlertStater
}

func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// MobileAlertVibration вибрация на звонке
type MobileAlertVibration struct {
}

func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// MobileAlertSong песня на звонке
type MobileAlertSong struct {
}

func (a *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}

func main() {
	mobile := NewMobileAlert()

	result := mobile.Alert()
	fmt.Println(result)

	result = mobile.Alert()
	fmt.Println(result)

	mobile.SetState(&MobileAlertSong{})

	result = mobile.Alert()
	fmt.Println(result)
}
