// greedy_algorithm
package main

// Задача о покрытии множества

func problemOfCoveringSet(inp *[]string, data map[string]Set) Set {
	var inpSet Set = Set{}
	inpSet.AddSlice(inp)

	var resultSet Set = Set{}

	for inpSet.Size() > 0 {
		var best_value string
		var set_covered Set = Set{}

		for key, setValue := range data {
			var covered Set = *inpSet.Intersection(&setValue)

			if covered.Size() > set_covered.Size() {
				best_value = key
				set_covered = covered
			}
		}
		inpSet = *inpSet.Difference(&set_covered)
		resultSet.Add(best_value)
	}
	return resultSet
}
