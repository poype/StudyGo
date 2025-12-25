package main

import (
	"fmt"
)

type GraphNode struct {
	val     string
	friends []*GraphNode
}

func friendCircle(relations [][]string) {
	graph := buildGraph(relations)
	visited := make(map[string]struct{})
	totalCircle := 0
	circleList := make([][]string, 0)
	for k, v := range *graph {
		if _, ok := visited[k]; ok {
			continue
		}
		totalCircle++
		circle := make([]string, 0)
		dfs(v, visited, &circle)
		circleList = append(circleList, circle)
	}
	fmt.Println(totalCircle)
	fmt.Println(circleList)
}

func buildGraph(relations [][]string) *map[string]*GraphNode {
	graph := make(map[string]*GraphNode)
	for _, relation := range relations {
		if _, ok := graph[relation[0]]; !ok {
			friendList := make([]*GraphNode, 0)
			graph[relation[0]] = &GraphNode{val: relation[0], friends: friendList}
		}
		if _, ok := graph[relation[1]]; !ok {
			friendList := make([]*GraphNode, 0)
			graph[relation[1]] = &GraphNode{val: relation[1], friends: friendList}
		}

		graph[relation[0]].friends = append(graph[relation[0]].friends, graph[relation[1]])
		graph[relation[1]].friends = append(graph[relation[1]].friends, graph[relation[0]])

	}
	return &graph
}

func dfs(node *GraphNode, visited map[string]struct{}, circleP *[]string) {
	if _, ok := visited[node.val]; ok {
		return
	}
	*circleP = append(*circleP, node.val)
	visited[node.val] = struct{}{}
	for _, friend := range node.friends {
		dfs(friend, visited, circleP)
	}
}
