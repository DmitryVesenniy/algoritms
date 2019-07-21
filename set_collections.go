// greedy_algorithm
package main

import (
	"fmt"
	"strings"
)

type Set struct {
	hash map[string]int
	arr  []string
}

func (set *Set) Add(str string) {
	if set.hash == nil {
		set.hash = make(map[string]int)
		set.arr = make([]string, 0, 20)
	}
	if _, ok := set.hash[str]; !ok {
		set.arr = append(set.arr, str)
		set.hash[str] = len(set.arr) - 1
	}
}

func (set *Set) AddSlice(sl *[]string) {
	for _, val := range *sl {
		set.Add(val)
	}
}

func (set *Set) Del(str string) {
	if i, ok := set.hash[str]; !ok {
		set.arr[i] = set.arr[len(set.arr)-1]
		set.arr = set.arr[:len(set.arr)-1]
		delete(set.hash, str)
	}
}

func (set *Set) Size() int {
	return len(set.hash)
}

func (set *Set) Has(str string) bool {
	_, ok := set.hash[str]
	return ok
}

func (set *Set) Clear() {
	set.hash = make(map[string]int)
	set.arr = make([]string, 0, 20)
}

func (s *Set) Intersection(s2 *Set) *Set {
	s3 := Set{}

	for i := range s2.hash {
		_, ok := s.hash[i]
		if ok {
			s3.Add(i)
		}
	}
	return &s3
}

func (s *Set) Difference(s2 *Set) *Set {
	s3 := Set{}

	for _, i := range s.arr {
		_, ok := s2.hash[i]
		if !ok {
			if i == "" {
				continue
			}
			s3.Add(i)
		}
	}
	return &s3
}

func (s *Set) Subset(s2 *Set) bool {

	for _, i := range s.arr {
		_, ok := s2.hash[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *Set) ToString() string {
	return strings.Trim(strings.Join(s.arr, ", "), ", ")
}

func Test() {
	var set *Set = &Set{}
	var set2 *Set = &Set{}

	set.AddSlice(&[]string{"q", "w", "e"})
	set2.AddSlice(&[]string{"w", "e", "r"})

	fmt.Println("*** >>> SeT", set.Difference(set2))
}
