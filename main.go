// algoritms project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Set
	*/

	var stations map[string]Set = make(map[string]Set)
	// kone
	var kone = Set{}
	kone.AddSlice(&[]string{"id", "nv", "ut"})
	stations["kone"] = kone

	// ktwo
	var ktwo = Set{}
	ktwo.AddSlice(&[]string{"wa", "id", "mt"})
	stations["ktwo"] = ktwo

	// kthree
	var kthree = Set{}
	kthree.AddSlice(&[]string{"or", "nv", "ca"})
	stations["kthree"] = kthree

	// kfour
	var kfour = Set{}
	kfour.AddSlice(&[]string{"nv", "ut"})
	stations["kfour"] = kfour

	// kfive
	var kfive = Set{}
	kfive.AddSlice(&[]string{"ca", "az"})
	stations["kfive"] = kfive

	var res Set = problemOfCoveringSet(&[]string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}, stations)

	fmt.Println("Result: ", res.ToString())
	return

	/*
		dijkstra
	*/

	fmt.Println("Dijkstra: ")
	test()

	fmt.Println("Hello World!")
	var A *Point = &Point{2, 4}
	var B *Point = &Point{5, 6}
	var C *Point = &Point{7, 8}

	koeff, status := collinearnesPoint(A, B, C)
	fmt.Println("Результат вычислений по трем точкам: ", koeff, status)

	//Получим массив точек обрамляющий множество точек
	//Point{7, 4}, Point{4, 6}, Point{3, 11}, Point{4, 15}, Point{7, 18},
	//      Point{13, 18}, Point{17, 16}, Point{16, 11}, Point{11, 15}
	//**************************************************************
	//Point{5, 7}, Point{7, 7}, Point{6, 9}, Point{5, 12}, Point{8, 13}, Point{9, 15}
	//	    Point{11, 12}, Point{13, 14}, Point{13, 9}

	var points []Point = []Point{
		Point{5, 7}, Point{7, 4}, Point{4, 6}, Point{11, 12}, Point{8, 13}, Point{3, 11},
		Point{4, 15}, Point{7, 7}, Point{7, 18}, Point{6, 9}, Point{5, 12}, Point{13, 14},
		Point{13, 18}, Point{17, 16}, Point{13, 9}, Point{16, 11}, Point{11, 5}, Point{9, 15},
	}

	resultArr := compute(points)
	fmt.Println("result: ", resultArr) // результат правильный

	var vals []int = []int{111, 344, 56, 77, 190}
	fmt.Println(MaxElement(vals, 0, len(vals)))

	//minute hours count  Constructor
	//var counter *MinuteHourCounterConveer = NewMinuteHourCounterConveer()
	var counter *MinuteHourCounterInsert = Constructor()
	counter.Add(10)
	time.Sleep(time.Second * 20)
	counter.Add(30)
	//asset 40
	fmt.Println(counter.MinuteCount())
	time.Sleep(time.Second * 50)
	//asset 30
	fmt.Println(counter.MinuteCount())
	type myInt uint8
	var k myInt = 56
	fmt.Println(k)
}
