package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
)

// пакет используется для проверки ответа, не удаляйте его

// пакет используется для проверки ответа, не удаляйте его

func main() {
	// Записывать данные, а в дальнейшем читать их мы будем из буфера,
	// но его можно заменить любым другим объектом, удовлетворяющим
	// интерфейсу io.ReadWriter
	buf := bytes.NewBuffer(nil)

	w := csv.NewWriter(buf)

	for i := 1; i <= 3; i++ {
		// Запись данных может производится поэтапно, например в цикле
		val1 := fmt.Sprintf("row %d col 1", i)
		val2 := fmt.Sprintf("row %d col 2", i)
		val3 := fmt.Sprintf("row %d col 3", i)
		if err := w.Write([]string{val1, val2, val3}); err != nil { // Аргументом Write является срез строк
			// ...
		}
	}
	w.Flush() // Этот метод приведет к фактической записи данных из буфера csv.Writer в buf

	// Либо данные можно записать за один раз
	w.WriteAll([][]string{ // Аргументом WriteAll является срез срезов строк
		{"row 4 col 1", "row 4 col 2", "row 4 col 3"},
		{"row 5 col 1", "row 5 col 2", "row 5 col 3"},
	})

	r := csv.NewReader(buf)

	for i := 1; i <= 2; i++ {
		// Читать данные мы тоже можем построчно, получая срез строк за каждую итерацию
		row, err := r.Read()
		if err != nil && err != io.EOF { // Здесь тоже нужно учитывать конец файла
			// ...
		}
		fmt.Println(row)
	}

	// Либо прочитать данные за один раз
	data, err := r.ReadAll()
	if err != nil {
		// Когда мы читаем данные до конца файла io.EOF не возвращается, а служит сигналом к завершению чтения
		// ...
	}

	for _, row := range data {
		fmt.Println(row)
	}

	// [row 1 col 1 row 1 col 2 row 1 col 3]
	// [row 2 col 1 row 2 col 2 row 2 col 3]
	// [row 3 col 1 row 3 col 2 row 3 col 3]
	// [row 4 col 1 row 4 col 2 row 4 col 3]
	// [row 5 col 1 row 5 col 2 row 5 col 3]
}
