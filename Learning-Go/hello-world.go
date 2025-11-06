package main

import "fmt"

type Person struct {
	On          bool
	Ammo, Power int
}

func (p *Person) Shoot() bool {
	if !p.On {
		return false
	}
	if p.Ammo != 0 {
		// fmt.Println("!")
		p.Ammo--
		return true
	} else {
		return false
	}
}

func (p *Person) RideBike() bool {
	if !p.On {
		return false
	}
	if p.Power != 0 {
		p.Power--
		return true
	} else {
		return false
	}
}

func main() {
	testStruct := Person{On: true, Ammo: 1, Power: 2}
	testStructPtr := &testStruct
	// fmt.Println(testStruct.On, testStruct.Power, testStruct.Ammo)
	fmt.Println(testStructPtr.On, testStructPtr.Power, testStructPtr.Ammo)
	result := testStructPtr.Shoot()
	fmt.Println("After Shoot:", result, "===", testStructPtr.On, testStructPtr.Power, testStructPtr.Ammo)
	result = testStructPtr.RideBike()
	fmt.Println("After RideBike:", result, "===", testStructPtr.On, testStructPtr.Power, testStructPtr.Ammo)

}
