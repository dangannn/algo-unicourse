package main

func generateParenthesis(n int) (res []string) {
	// Создаем слайс байтов для хранения текущей комбинации скобок
	cur := make([]byte, 2*n)
	var l, r int // Переменные для отслеживания количества открывающих и закрывающих скобок

	// Объявляем рекурсивную функцию генерации комбинаций
	var generator func()
	generator = func() {
		// Если количество открывающих и закрывающих скобок достигло максимума, добавляем текущую комбинацию в результат
		if l+r == 2*n {
			res = append(res, string(cur))
			return
		}
		// Если количество открывающих скобок меньше n, добавляем '(' и рекурсивно вызываем генератор
		if l < n {
			cur[l+r] = '('
			l++
			generator() // Рекурсивный вызов
			l--         // Откатываем изменение для других вариантов
		}
		// Если количество закрывающих скобок меньше количества открывающих, добавляем ')' и рекурсивно вызываем генератор
		if r < l {
			cur[l+r] = ')'
			r++
			generator() // Рекурсивный вызов
			r--         // Откатываем изменение для других вариантов
		}
	}
	generator() // Начинаем генерацию
	return      // Возвращаем результат
}