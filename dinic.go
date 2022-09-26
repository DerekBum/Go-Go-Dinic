package main

import "fmt"

type edge struct {
	a, b, flow, cap int
}

const N = 1e6
const INF = 1e9

var graph [N][]int
var edges []edge
var dist [N]int
var query [N]int
var start, finish int

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func dfs(v, current int) int {
	if current == 0 || v == finish {
		return current
	}
	for to_index := range graph[v] {
		index := graph[v][to_index]
		to := edges[index].b
		if dist[to] != dist[v] + 1 {
			continue
		}
		update := dfs(to, min(current, edges[index].cap - edges[index].flow))
		if update > 0 {
			edges[index].flow += update
			edges[index^1].flow -= update
			return update
		}
	}
	return 0
}

func bfs() bool {
	q_prev := 0
	q_next := 0
	query[q_next] = start
	q_next++
	for i := range dist {
		dist[i] = -1
	}
	dist[start] = 0
	for q_prev < q_next && dist[finish] == -1 {
		v := query[q_prev]
		q_prev++
		for to_index := range graph[v] {
			index := graph[v][to_index]
			to := edges[index].b
			if dist[to] == -1 && edges[index].flow < edges[index].cap {
				query[q_next] = to
				q_next++
				dist[to] = dist[v] + 1
			}
		}
	}
	return dist[finish] != -1
}

func dinic() int {
	ans := 0
	for {
		if !bfs() {
			break
		}
		for update := dfs(start, INF); update != 0; update = dfs(start, INF) {
			ans += update
		}
	}
	return ans
}

func main() {
	var n, m int
    fmt.Scanf("%d %d", &n, &m)

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Scanf("%d %d %d", &a, &b, &c)

		edge_forward  := edge{a, b, 0, c}
		edge_backward := edge{b, a, 0, 0}
		graph[a] = append(graph[a], len(edges))
		edges = append(edges, edge_forward)
		graph[b] = append(graph[b], len(edges))
		edges = append(edges, edge_backward)
	}

	fmt.Scanf("%d %d", &start, &finish)

	fmt.Println(dinic())

}