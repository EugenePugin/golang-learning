// showcase using interfaces for unit-test coverage of sin function
package interface_practice

import (
	"fmt"
	"math"
)

// TrigProvider описывает поведение вычисления синуса
type TrigProvider interface {
	SinDegrees(degrees float64) (float64, error)
}

// RealTrigImplementation — реальная реализация для продакшена
type RealTrigImplementation struct{}

func (r *RealTrigImplementation) SinDegrees(degrees float64) (float64, error) {
	return math.Sin(degrees * math.Pi / 180), nil
}

// MockTrigImplementation — моковая реализация для продакшена
type MockTrigImplementation struct{}

func (m *MockTrigImplementation) SinDegrees(degrees float64) (float64, error) {
	if degrees == 90 {
		return 1, nil
	}
	return 0, fmt.Errorf("this value is not supported")
}

// Calculator — структура, которая использует интерфейс
type Calculator struct {
	provider TrigProvider
}

func (c *Calculator) CalculateSin(degrees float64) (float64, error) {
	return c.provider.SinDegrees(degrees)
}

func main() {
	// 1. Инициализируем реальную реализацию
	realImplementation := &RealTrigImplementation{}

	// 2. Внедряем её в калькулятор (Dependency Injection)
	calc := Calculator{
		provider: realImplementation,
	}

	// 3. Используем калькулятор
	angle := 90.0
	if result, err := calc.CalculateSin(angle); err != nil {
		fmt.Printf("sin(%.2f) = %.2f\n", angle, result)
	}

	mockImplementation := &MockTrigImplementation{}
	calc_m := Calculator{
		provider: mockImplementation,
	}
	if result_m, err_m := calc_m.CalculateSin(angle); err_m != nil {
		fmt.Printf("sin(%.2f) = %.2f\n", angle, result_m)
	}
}
