package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		inDegree[i] = 0
	}

	graph := buildCourseGraph(prerequisites, &inDegree)

	for course := range inDegree {
		if inDegree[course] == 0 {
			dfsCourseGraph(&graph, &inDegree, course)
		}
	}

	for course := range inDegree {
		if inDegree[course] != 0 {
			return false
		}
	}
	return true
}

func buildCourseGraph(prerequisites [][]int, inDegree *map[int]int) map[int][]int {
	graph := make(map[int][]int)
	for _, prerequisite := range prerequisites {
		if _, ok := graph[prerequisite[1]]; !ok {
			graph[prerequisite[1]] = make([]int, 0)
		}
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])

		if _, ok := (*inDegree)[prerequisite[0]]; !ok {
			(*inDegree)[prerequisite[0]] = 0
		}
		(*inDegree)[prerequisite[0]]++
	}

	return graph
}

func dfsCourseGraph(graph *map[int][]int, inDegree *map[int]int, studyCourse int) {
	postCourses := (*graph)[studyCourse]
	for _, postCourse := range postCourses {
		if (*inDegree)[postCourse] == 0 {
			continue
		}
		(*inDegree)[postCourse]--
		if (*inDegree)[postCourse] == 0 {
			dfsCourseGraph(graph, inDegree, postCourse)
		}
	}
}
