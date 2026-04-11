// Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>2.
//  Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
//  Напишите функцию, которая по указанному натуральному n возвращает φn.

package fib

func getFibonacchiListItemByOrderNumber(order_num uint) int {
	// fmt.Println("order_num:", order_num)
	// special cases
	// formula
	if order_num < 1 {
		return -1 // incorrect input
	}
	if order_num == 1 || order_num == 2 {
		return 1 // special cases
	}

	result := getFibonacchiListItemByOrderNumber(order_num-1) + getFibonacchiListItemByOrderNumber(order_num-2)
	// switch order_num {

	// }
	// switch order_num {
	// case <1: result = -1	// incorrect input
	// case 1 || 2:	result = 1

	// }
	return result
}
