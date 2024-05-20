package main

func NumberOfProvince(isConnected [][]int) int {
	// adjList := map[int]map[int]bool{}
	adjList := map[int][]int{}
	length := len(isConnected)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if isConnected[i][j] == 1 && i != j {
				adjList[i] = append(adjList[i], j)
				adjList[j] = append(adjList[j], i)
			}
		}
	}

	visited := make([]int, length)
	count := 0

	for i := 0; i < length; i++ {
		if visited[i] != 1 {
			bfs(i, adjList, visited)
			count++
		}
	}
	return count
}

func bfs(i int, adjList map[int][]int, visited []int) {
	queue := []int{i}
	visited[i] = 1
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		neighbors := adjList[current]
		for _, v := range neighbors {
			if visited[v] != 1 {
				queue = append(queue, v)
				visited[v] = 1
			}
		}
	}
}
