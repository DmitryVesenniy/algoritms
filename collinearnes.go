// algoritms project main.go
package main

type Point struct {
	X float64
	Y float64
}

func collinearnesPoint(a *Point, b *Point, c *Point) (float64, string) {
	var status string
	var result float64 = Round((b.X-a.X)*(c.Y-a.Y)-(b.Y-a.Y)*(c.X-a.X), 4)

	switch {
	case result == 0.0:
		status = "Коллиенарны"
		break
	case result < 0.0:
		status = "Левый поворот"
		break
	case result > 0.0:
		status = "Правый поворот"
	}
	return result, status
}
