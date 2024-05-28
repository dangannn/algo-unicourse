package main

// Node представляет узел графа.
type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph создает и возвращает глубокую копию графа.
func cloneGraph(node *Node) *Node {
	// Если исходный узел пустой, возвращаем nil.
	if node == nil {
		return nil
	}

	// Создаем слайс для хранения копий узлов графа.
	copies := make([]*Node, 101) // 101 предполагает, что значения узлов не превышают 100.

	// Запускаем рекурсивную функцию dfs для создания копий узлов.
	dfs(node, copies)

	// Возвращаем копию исходного узла.
	return copies[node.Val]
}

// dfs рекурсивно создает копии узлов графа.
func dfs(node *Node, copies []*Node) {
	// Создаем новый узел и устанавливаем его значение.
	newNode := new(Node)
	newNode.Val = node.Val

	// Сохраняем копию узла в слайсе copies.
	copies[node.Val] = newNode

	// Проходим по соседним узлам (ребрам графа).
	for _, neighbor := range node.Neighbors {
		// Если копия соседнего узла еще не создана, вызываем dfs рекурсивно.
		if copies[neighbor.Val] == nil {
			dfs(neighbor, copies)
		}

		// Добавляем копию соседнего узла к списку соседей нового узла.
		newNode.Neighbors = append(newNode.Neighbors, copies[neighbor.Val])
	}
}
