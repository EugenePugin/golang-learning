package practice_interfaces

import (
	"errors"
	"fmt"
)

type SmallDrinker interface {
	SmallDrink() error
}

type LiquidAsset struct {
	Volume_in_liters float32
	Pack_material    string
	Price            float32
}

type Juice struct {
	Fruit_name string
	LiquidAsset
}

type Milk struct {
	LiquidAsset
}

type Alcohol struct {
	Alcohol_percentage float32
	LiquidAsset
}

func NewAlcoholItem(alcohol_percentage float32, volume_in_liters float32,
	pack_material string,
	price float32) (*Alcohol, error) {
	fmt.Println("Hey! I a Alcohol constructor")

	if pack_material == "" {
		return nil, errors.New("pack_material can't be empty")
	}
	if alcohol_percentage < 0 || alcohol_percentage >= 100 {
		return nil, errors.New("alcohol_percentage needs to be 0...100")
	}
	if volume_in_liters <= 0 {
		return nil, errors.New("volume_in_liters needs to be positive")
	}
	if price <= 0 {
		return nil, errors.New("price needs to be positive")
	}

	return &Alcohol{
		Alcohol_percentage: alcohol_percentage,
		LiquidAsset: LiquidAsset{
			Volume_in_liters: volume_in_liters,
			Pack_material:    pack_material,
			Price:            price,
		},
	}, nil
}

func NewMilkItem(volume_in_liters float32,
	pack_material string,
	price float32) (*Milk, error) {
	fmt.Println("Hey! I a Milk constructor")

	if pack_material == "" {
		return nil, errors.New("pack_material can't be empty")
	}
	if volume_in_liters <= 0 {
		return nil, errors.New("volume_in_liters needs to be positive")
	}
	if price <= 0 {
		return nil, errors.New("price needs to be positive")
	}

	return &Milk{
		LiquidAsset: LiquidAsset{
			Volume_in_liters: volume_in_liters,
			Pack_material:    pack_material,
			Price:            price,
		},
	}, nil
}

func NewJuiceItem(fruit_name string,
	volume_in_liters float32,
	pack_material string,
	price float32) (*Juice, error) {
	fmt.Println("Hey! I am Juice constructor")
	if fruit_name == "" {
		return nil, errors.New("fruit_name can't be empty")
	}
	if pack_material == "" {
		return nil, errors.New("pack_material can't be empty")
	}
	if volume_in_liters <= 0 {
		return nil, errors.New("volume_in_liters needs to be positive")
	}
	if price <= 0 {
		return nil, errors.New("price needs to be positive")
	}

	return &Juice{
		Fruit_name: fruit_name,
		LiquidAsset: LiquidAsset{
			Volume_in_liters: volume_in_liters,
			Pack_material:    pack_material,
			Price:            price,
		},
	}, nil
}

func ShowJuiceDetails(obj *Juice) error {
	if obj == nil {
		return errors.New("The argument is not initialized")
	}
	s_volume_in_liters := fmt.Sprintf("%.2f", obj.Volume_in_liters)
	fmt.Println("Object details:", obj.Fruit_name, obj.Pack_material, s_volume_in_liters, obj.Price)
	return nil
}

func ShowMilkDetails(obj *Milk) error {
	if obj == nil {
		return errors.New("The argument is not initialized")
	}
	s_volume_in_liters := fmt.Sprintf("%.2f", obj.Volume_in_liters)
	fmt.Println("Object details:", obj.Pack_material, s_volume_in_liters, obj.Price)
	return nil
}

func ShowAlcoholDetails(obj *Alcohol) error {
	if obj == nil {
		return errors.New("The argument is not initialized")
	}
	s_volume_in_liters := fmt.Sprintf("%.2f", obj.Volume_in_liters)
	fmt.Println("Object details:", obj.Alcohol_percentage, obj.Pack_material, s_volume_in_liters, obj.Price)
	return nil
}

// 	const SmallDrinkSize = 0.05
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("panic ", r, " has been addressed")
// 			obj.Volume_in_liters = 0
// 		}
// 	}()
// 	if obj == nil {
// 		return errors.New("The referring object is not initialized")
// 	}
// 	// if obj.Volume_in_liters < SmallDrinkSize {
// 	// 	obj.Volume_in_liters = 0
// 	// 	return nil
// 	// }
// 	obj.Volume_in_liters -= 0.05
// 	if obj.Volume_in_liters < 0 {
// 		panic("wow, volume can't be negative!")
// 	}
// 	return nil
// }

func (obj *LiquidAsset) SmallDrink() error { //decrement by 50ml
	const SmallDrinkSize = 0.05
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("panic ", r, " has been addressed")
	// 		obj.Volume_in_liters = 0
	// 	}
	// }()
	if obj == nil {
		return errors.New("The referring object is not initialized")
	}
	if obj.Volume_in_liters < SmallDrinkSize {
		obj.Volume_in_liters = 0
		return nil
	}
	obj.Volume_in_liters -= 0.05
	// if obj.Volume_in_liters < 0 {
	// 	panic("wow, volume can't be negative!")
	// }
	return nil
}

func TakeAGulp(item SmallDrinker) error {
	err := item.SmallDrink()
	if err != nil {
		// fmt.Println("woops, sorry a gulp could not be taken")
		return errors.New("woops, sorry a gulp could not be taken")
	}
	fmt.Println("I am done!")
	return nil
}

func structures() {
	// create structure Juice with fields like fruit_name, volume, pack_material, price
	// implement methods: small_drink, big_drink, show_volume
	juice_sample_ptr, err := NewJuiceItem("apple", 0.03, "paper", 56.5)
	if err != nil {
		fmt.Println("Error on struct construction:", err)
		return
	}
	fmt.Println("Juice object created successfully")
	err = ShowJuiceDetails(juice_sample_ptr)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = juice_sample_ptr.SmallDrink()
	if err != nil {
		fmt.Println("Error:", err)
	}
	ShowJuiceDetails(juice_sample_ptr)
	TakeAGulp(juice_sample_ptr)
	TakeAGulp(juice_sample_ptr)

	// juice_sample_ptr.SmallDrink()
	// juice_sample_ptr.SmallDrink()
	ShowJuiceDetails(juice_sample_ptr)

	milk_sample_ptr, err := NewMilkItem(0.5, "glass", 95)
	if err != nil {
		fmt.Println("Error on struct construction:", err)
		return
	}
	fmt.Println("Milk object created successfully")
	ShowMilkDetails(milk_sample_ptr)
	TakeAGulp(milk_sample_ptr)

	// milk_sample_ptr.SmallDrink()
	ShowMilkDetails(milk_sample_ptr)

	alcohol_sample_ptr, err := NewAlcoholItem(45, 0.5, "glass", 95)
	if err != nil {
		fmt.Println("Error on struct construction:", err)
		return
	}
	fmt.Println("Alcohol object created successfully")
	ShowAlcoholDetails(alcohol_sample_ptr)
	TakeAGulp(alcohol_sample_ptr)
	ShowAlcoholDetails(alcohol_sample_ptr)

}
