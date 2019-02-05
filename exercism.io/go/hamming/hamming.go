package hamming

import "fmt"

// Distance calculates the Hamming distance between 2 given strings.
func Distance(a, b string) (int, error) {
	count := 0

	if len(a) != len(b) {
		return count, fmt.Errorf("Not possible to calculate the Hamming distance between sequences of different lengths")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
