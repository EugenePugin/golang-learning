package practice_interfaces

import (
	"errors"
	"fmt"
)

// let's have an interface Charger with the only method Charge
// structures SmartPhone and Watch, implementing that interface, making their power incremented
type Charger interface {
	Charge(int) error
}

type PowerModel struct {
	power int
	model string
}

func (obj *PowerModel) Charge(chargeValue int) error {
	if chargeValue <= 0 {
		return errors.New("Charge value must be positive")
	}
	obj.power += chargeValue
	return nil
}

type SmartPhone struct {
	PowerModel
	cameraPresent bool
}

type Watch struct {
	PowerModel
}

func NewWatch(p int, m string) (*Watch, error) {
	if m == "" {
		return nil, errors.New("Model name can't be empty")
	}
	var obj Watch
	obj.power = p
	obj.model = m
	return &obj, nil
}

func NewSmartPhone(p int, m string, cPresent bool) (*SmartPhone, error) {
	if m == "" {
		return nil, errors.New("Model name can't be empty")
	}
	var obj SmartPhone
	obj.power = p
	obj.model = m
	obj.cameraPresent = cPresent
	return &obj, nil
}

func (obj *SmartPhone) Expose() {
	fmt.Println(obj.power, obj.model, obj.cameraPresent)
}

func (obj *Watch) Expose() {
	fmt.Println(obj.power, obj.model)
}

// func (obj *SmartPhone) Charge(chargeValue int) error {
// 	if chargeValue <= 0 {
// 		return errors.New("Charge value must be positive")
// 	}
// 	obj.power += chargeValue
// 	return nil
// }

// func (obj *Watch) Charge(chargeValue int) error {
// 	if chargeValue <= 0 {
// 		return errors.New("Charge value must be positive")
// 	}
// 	obj.power += chargeValue
// 	return nil
// }

func Charge(obj Charger, chargeValue int) error {
	err := obj.Charge(chargeValue)
	return err
}

func interaces_and_structures() {
	var smartPhoneSample *SmartPhone
	var err error
	if smartPhoneSample, err = NewSmartPhone(10, "Apple", true); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	smartPhoneSample.Expose()
	if err = smartPhoneSample.Charge(10); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	smartPhoneSample.Expose()

	var watchSample *Watch
	if watchSample, err = NewWatch(5, "iWatch"); err != nil {
		fmt.Println("Error:", err)
		return
	}
	watchSample.Expose()

	Charge(smartPhoneSample, 5)
	smartPhoneSample.Expose()
	Charge(watchSample, 10)
	watchSample.Expose()
}
