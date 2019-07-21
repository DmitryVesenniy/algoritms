// dijkstra_algorithm
package main

import (
	"fmt"
	"math"
)

func in_element(element string, processed *[]string) bool {
	for _, el := range *processed {
		if element == el {
			return true
		}
	}
	return false
}

func find_lowest_cost_node(costs *map[string]int, processed *[]string) string {
	var lowest_cost float64 = math.Inf(1)
	var lowest_cost_node string
	for node, cost := range *costs {
		if float64(cost) < lowest_cost && !in_element(node, processed) {
			lowest_cost = float64(cost)
			lowest_cost_node = node
		}
	}
	return lowest_cost_node
}

func dijkstra(graph *map[string]interface{}, costs *map[string]int, parents *map[string]string) {
	var processed []string
	var node string = find_lowest_cost_node(costs, &processed)

	for node != "" {
		var cost, ok = (*costs)[node]
		if !ok {
			continue
		}

		if (*graph)[node] == nil {
			break
		}

		// Обозначение x.(T) называется Утверждение типа.
		var neighbors map[string]int = (*graph)[node].(map[string]int) // приводим интерфейс к карте

		for key, val := range neighbors {
			var new_cost int = cost + val

			if (*costs)[key] > new_cost {
				(*costs)[key] = new_cost
				(*parents)[key] = node
			}
		}

		processed = append(processed, node)

		node = find_lowest_cost_node(costs, &processed)

	}
}

func test() {
	var graph map[string]interface{} = map[string]interface{}{
		"a": map[string]int{
			"fin": 1,
		},
		"b": map[string]int{
			"a":   3,
			"fin": 5,
		},
		"fin": nil,
	}

	var costs map[string]int = map[string]int{
		"a":   6,
		"b":   2,
		"fin": math.MaxInt64,
	}

	var parents map[string]string = map[string]string{
		"a": "start",
		"b": "start",
		"c": "",
	}

	fmt.Println("data: ", costs, "   ", parents)

	dijkstra(&graph, &costs, &parents)

	fmt.Println("data: ", costs, "   ", parents)
}
