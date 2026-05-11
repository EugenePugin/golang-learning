// fuzzing testnig practice

// have a function, reversing all the symbols of the given string
// and let's fuzzy-test it

package fuzzy_testing

import "github.com/rivo/uniseg"

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) //tbd
}

func reverseString_v2(s string) string {
	// Разбиваем строку на кластеры графем
	gr := uniseg.NewGraphemes(s)
	var clusters [][]rune

	for gr.Next() {
		// Каждая графема может состоять из нескольких рун
		// Копируем их, чтобы сохранить порядок внутри графемы
		clusters = append(clusters, gr.Runes())
	}

	// Разворачиваем слайс кластеров
	for i, j := 0, len(clusters)-1; i < j; i, j = i+1, j-1 {
		clusters[i], clusters[j] = clusters[j], clusters[i]
	}

	// Собираем всё обратно в строку
	var result []rune
	for _, c := range clusters {
		result = append(result, c...)
	}

	return string(result)
}
