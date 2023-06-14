package main

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
## Паттерн Facade относится к структурным паттернам уровня объекта.

  Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс в виде набора имен методов к набору взаимосвязанных
  классов или объектов некоторой подсистемы, что облегчает ее использование.

  Разбиение сложной системы на подсистемы позволяет упростить процесс разработки, а также помогает максимально снизить зависимости
  одной подсистемы от другой. Однако использовать такие подсистемы становиться довольно сложно. Один из способов решения этой
  проблемы является паттерн Facade. Наша задача, сделать простой, единый интерфейс, через который можно было бы взаимодействовать
  с подсистемами.

## Преимущества и недостатки
+  Изолирует клиентов от компонентов сложной подсистемы.
-  Фасад рискует стать объектом, привязанным ко всем классам программы.

## Применимость
 Когда вам нужно представить простой или урезанный интерфейс к сложной подсистеме.
 Когда вы хотите разложить подсистему на отдельные слои.
*/

// В качестве фасада выступает структура SmartHome, которая содержит в себе объекты, имеющие сложную логику
type SmartHome struct {
	va  *VoiceAssistant
	rvc *RobotVacuumCleaner
	c   *Camera
}

func NewSmartHome() *SmartHome {
	return &SmartHome{
		va:  NewVoiceAssistant(),
		rvc: NewRobotVacuumCleaner(),
		c:   NewCamera()}
}

// Функции On и Off упрощают работу с объектами структуры SmartHome
func (sh *SmartHome) On() {
	sh.va.On()
	sh.rvc.On()
	sh.c.On()
}

func (sh *SmartHome) Off() {
	sh.va.Off()
	sh.rvc.Off()
	sh.c.Off()
}

// структура VoiceAssistant, имеющая свою сложную логику
type VoiceAssistant struct {
}

func NewVoiceAssistant() *VoiceAssistant {
	return &VoiceAssistant{}
}

func (va *VoiceAssistant) On() {
	fmt.Println("VoiceAssistant turned on")
}

func (va *VoiceAssistant) Off() {
	fmt.Println("VoiceAssistant turned off")
}

// структура RobotVacuumCleaner, имеющая свою сложную логику
type RobotVacuumCleaner struct {
}

func NewRobotVacuumCleaner() *RobotVacuumCleaner {
	return &RobotVacuumCleaner{}
}

func (va *RobotVacuumCleaner) On() {
	fmt.Println("RobotVacuumCleaner turned on")
}

func (va *RobotVacuumCleaner) Off() {
	fmt.Println("RobotVacuumCleaner turned off")
}

// структура Camera, имеющая свою сложную логику
type Camera struct {
}

func NewCamera() *Camera {
	return &Camera{}
}

func (va *Camera) On() {
	fmt.Println("Camera turned on")
}

func (va *Camera) Off() {
	fmt.Println("Camera turned off")
}

func main() {
	smartHome := NewSmartHome()
	smartHome.On() // простой интерфейс работы с сложными объектами
	time.Sleep(time.Second)
	smartHome.Off()
}
