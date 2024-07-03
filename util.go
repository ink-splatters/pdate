package main


func Permutations[T any](input []T) [][]T {
	result := [][]T{}

	if len(input) == 0 {
		return result
	}

	for i, elem := range input {
		rest := append(append([]T{}, input[:i]...), input[i+1:]...)

		perms := Permutations(rest)
		if len(perms) == 0 {
			result = append(result, []T{elem})
		} else {
			for _, perm := range perms {
				result = append(result, append([]T{elem}, perm...))
			}
		}
	}

	return result
}
