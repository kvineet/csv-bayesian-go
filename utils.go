package main

// Unique given slice of string elements returns slice of unique elements
func Unique(strings ...string) []string {
	keys := make(map[string]bool)
	uniques := []string{}
	for _, entry := range strings {
		if _, val := keys[entry]; !val {
			keys[entry] = true
			uniques = append(uniques, entry)
		}
	}
	return uniques
}

//EPSILON denotes smallest quantity
var EPSILON float64 = 0.00000001

//FloatEquals checks for equality of floating points with defined EPSILON value
func FloatEquals(a, b float64) bool {
	return (a-b) < EPSILON && (b-a) < EPSILON
}
