package main

func numIslands(grid [][]byte) int {
	// Счетчик для подсчета количества островов
	count := 0

	// Количество строк и столбцов в сетке
	ROW := len(grid)
	COL := len(grid[0])

	// Определяем функцию DFS для поиска по сетке
	var dfs func(grid [][]byte, row, col int)
	dfs = func(grid [][]byte, row, col int) {
		// Проверяем выход за границы сетки или нахождение на воде ('0')
		if row < 0 || col < 0 || row >= ROW || col >= COL || grid[row][col] == '0' {
			return
		}

		// Помечаем текущую клетку как посещенную (заменяем '1' на '0')
		grid[row][col] = '0'

		// Рекурсивно вызываем DFS для всех соседних клеток (вверх, вниз, вправо, влево)
		dfs(grid, row+1, col)
		dfs(grid, row-1, col)
		dfs(grid, row, col+1)
		dfs(grid, row, col-1)
	}

	// Проходим по всей сетке
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			// Если находим землю ('1'), это начало нового острова
			if grid[i][j] == '1' {
				// Увеличиваем счетчик островов
				count++

				// Запускаем DFS для пометки всех клеток, принадлежащих текущему острову
				dfs(grid, i, j)
			}
		}
	}

	// Возвращаем количество найденных островов
	return count
}
