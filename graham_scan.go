package main

import (
	"fmt"
	"math"
	"sort"
)

func compute(points []Point) []Point {
	var result []Point
	var n int = len(points)
	if n < 3 {
		return points
	}

	// Находим наинизшую точку Y и меняем ее с последней
	// точкой массива points[], если она не является таковой
	var lowest int = 0
	var lowestY = points[0].Y

	for i := 1; i < n; i++ {
		if points[i].Y < lowestY {
			lowestY = points[i].Y
			lowest = i
		}
	}

	if lowest != n-1 {
		var temp Point = points[n-1]
		points[n-1] = points[lowest]
		points[lowest] = temp
	}

	fmt.Println("Last point: ", points[len(points)-1])

	// Сортируем points[0..п-2] в порядке убывания полярного
	// угла по отношению к наинизшей точке points[п-1].
	sort.Slice(points[:n-2], func(i, j int) bool {
		var one = points[i]
		var two = points[j]

		if one == two {
			return false
		}
		var baseX float64 = points[len(points)-1].X
		var baseY float64 = points[len(points)-1].Y

		var oneY float64 = one.Y
		var twoY float64 = two.Y

		var oneAngle float64 = math.Atan2(oneY-baseY, one.X-baseX)
		var twoAngle float64 = math.Atan2(twoY-baseY, two.X-baseX)

		if oneAngle > twoAngle {
			return false
		} else if oneAngle < twoAngle {
			return true
		}

		if oneY > twoY {
			return false
		} else if oneY < twoY {
			return true
		}
		return false
	})

	//result = append(result, points[0])
	result = append(result, points[len(points)-2])
	result = append(result, points[len(points)-1])

	// Если все точки коллинеарны, обрабатываем их, чтобы
	// избежать проблем позже
	var firstAngle float64 = math.Atan2(points[0].Y-lowestY, points[0].X-points[n-1].X)
	var lastAngle float64 = math.Atan2(points[n-2].Y-lowestY, points[n-2].X-points[n-1].X)

	if firstAngle == lastAngle {
		return []Point{points[n-1], points[0]}
	}

	// Последовательно посещаем каждую точку, удаляя точки,
	// приводящие к ошибке. Поскольку у нас всегда есть как
	// минимум один "правый поворот", внутренний цикл всегда
	// завершается

	for i := 0; i < n-1; i++ {
		for len(result) > 1 && isLeftTurn(result[len(result)-2], result[len(result)-1], points[i]) {
			result = result[:len(result)-1]
		}
		result = append(result, points[i])
	}

	// Последняя точка дублируется, так что мы берем
	// п-1 точек начиная с наинизшей

	result[len(result)-1] = points[n-1]

	return result
}

func isLeftTurn(p1 Point, p2 Point, p3 Point) bool {
	return (p2.X-p1.X)*(p3.Y-p1.Y)-(p2.Y-p1.Y)*(p3.X-p1.X) < 0
}
