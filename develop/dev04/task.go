package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

//deleteRepeated удаляет повторные вхождения строк в слайс.
//Принимает слайс строк, в котором требуется удалить повторы.
//Возвращает слайс строк с удаленными повторами.
func deleteRepeated(in []string) []string {
	result := make([]string, 0)
	m := make(map[string]bool)

	for _, v := range in {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}

func AnagramDict(in []string) map[string][]string {
	for i := range in {
		in[i] = strings.ToLower(in[i])
	}
	woRepeated := deleteRepeated(in)      //удаляем повторяющиеся строки
	tempM := make(map[string][]string, 0) //промежуточная мапа, ключ - отсортированное слов

	for _, v := range woRepeated {
		sorted := []rune(v)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})
		sortedS := string(sorted) //отсортированное слово

		tempM[sortedS] = append(tempM[sortedS], v)
	}

	//результирующая мапа
	resultM := make(map[string][]string, 0)

	for _, v := range tempM {
		if len(v) > 1 { //если всего один элемент - в словрь не попадает
			resultM[v[0]] = v //нулевой элемент, это первый добавленный (первый просмотренный)
			sort.Strings(v)
		}
	}

	return resultM

}

func main() {
	input := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}

	fmt.Println(input)
	fmt.Println(AnagramDict(input))
}
