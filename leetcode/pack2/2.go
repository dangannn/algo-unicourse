package main

import "strings"

func solveNQueens(n int) [][]string {
	// Создаем доску размера n x n, заполненную точками
	board := make([][]rune, n)
	for r := 0; r < n; r++ {
		board[r] = make([]rune, n)
		for c := 0; c < n; c++ {
			board[r][c] = '.'
		}
	}

	// Инициализируем переменную для хранения результатов
	res := [][]string{}
	// Размещаем королев на доске, начиная с первого столбца
	placeQueen(0, n, board, &res)
	return res // Возвращаем результаты
}

// Рекурсивная функция для размещения королев на доске
func placeQueen(qId, n int, board [][]rune, res *[][]string) {
	// Если все королевы успешно размещены, добавляем текущее состояние доски в результат
	if qId == n {
		addToResult(board, res)
		return
	}

	// Перебираем все строки в текущем столбце
	for row := 0; row < n; row++ {
		// Проверяем, можно ли разместить королеву в данной ячейке
		if isValidCell(board, row, qId) {
			// Размещаем королеву в текущей ячейке
			board[row][qId] = 'Q'
			// Рекурсивно размещаем следующую королеву
			placeQueen(qId+1, n, board, res)
			// Возвращаем доску в исходное состояние для перебора других вариантов размещения
			board[row][qId] = '.'
		}
	}
}

// Функция для проверки, можно ли разместить королеву в данной ячейке
func isValidCell(board [][]rune, row, col int) bool {
	n := len(board)

	// Проверяем, нет ли королев в этой строке до текущего столбца
	for c := 0; c < col; c++ {
		if board[row][c] == 'Q' {
			return false
		}
	}

	// Проверяем, нет ли королев в диагонали слева вверху
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}

	// Проверяем, нет ли королев в диагонали слева внизу
	for r, c := row+1, col-1; r < n && c >= 0; r, c = r+1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}

	return true
}

// Функция для добавления текущего состояния доски в результат
func addToResult(board [][]rune, res *[][]string) {
	n := len(board)
	rows := []string{}

	// Преобразуем доску в строки и добавляем их в результат
	for r := 0; r < n; r++ {
		row := strings.Builder{}
		for c := 0; c < n; c++ {
			row.WriteRune(board[r][c])
		}
		rows = append(rows, row.String())
	}

	// Добавляем текущее состояние доски в результат
	*res = append(*res, rows)
}
