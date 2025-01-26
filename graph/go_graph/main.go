package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}

	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("A", "B") // Menambahkan hubungan dari A ke B
	addEdge("B", "C") // Menambahkan hubungan dari B ke C
	addEdge("A", "C") // Menambahkan hubungan dari A ke C

	fmt.Println(hasEdge("A", "B")) // true (Ada hubungan dari A ke B)
	fmt.Println(hasEdge("A", "C")) // true (Ada hubungan dari A ke C)
	fmt.Println(hasEdge("B", "A")) // false (Tidak ada hubungan dari B ke A)
}
