package main

import "fmt"

type Person struct {
	On          bool
	Ammo, Power int
}

// Shoot уменьшает количество патронов на 1 при выстреле, если возможно.
// Возвращает true, если выстрел произошёл, иначе false.
func (p *Person) Shoot() bool {
	if !p.On || p.Ammo <= 0 {
		return false
	}
	p.Ammo--
	return true
}

// RideBike уменьшает заряд на 1 при езде на велосипеде, если возможно.
// Возвращает true, если удалось поехать, иначе false.
func (p *Person) RideBike() bool {
	if !p.On || p.Power <= 0 {
		return false
	}
	p.Power--
	return true
}

func main() {
	testStruct := Person{On: true, Ammo: 1, Power: 2}
	testStructPtr := &testStruct

	// Вывод начального состояния
	fmt.Println("Before:", testStructPtr.On, testStructPtr.Ammo, testStructPtr.Power)

	// Проверка выстрела
	result := testStructPtr.Shoot()
	fmt.Println("After Shoot:", result, "===", testStructPtr.On, testStructPtr.Ammo, testStructPtr.Power)

	// Проверка поездки на велосипеде
	result = testStructPtr.RideBike()
	fmt.Println("After RideBike:", result, "===", testStructPtr.On, testStructPtr.Ammo, testStructPtr.Power)
}
