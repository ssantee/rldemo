package fib

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var pregeneratedSequence = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040}

const maxPregenerated = 1000000

func Handler(w http.ResponseWriter, r *http.Request) {
	userProvidedNumber := r.URL.Query().Get("n")
	if userProvidedNumber == "" {
		http.Error(w, "Please provide a number in the query parameter 'n'", http.StatusBadRequest)
		return
	}

	optionalStartX := r.URL.Query().Get("startx")
	optionalStartY := r.URL.Query().Get("starty")

	var startX, startY int
	if optionalStartX != "" {
		var err error
		startX, err = strconv.Atoi(optionalStartX)
		if err != nil {
			http.Error(w, "Invalid startx provided", http.StatusBadRequest)
			return
		}
	} else {
		startX = 0 // Default value for startx
	}
	if optionalStartY != "" {
		var err error
		startY, err = strconv.Atoi(optionalStartY)
		if err != nil {
			http.Error(w, "Invalid starty provided", http.StatusBadRequest)
			return
		}
	} else {
		startY = 1 // Default value for starty
	}

	// Convert user-provided number to integer
	n, err := strconv.Atoi(userProvidedNumber)
	if err != nil {
		http.Error(w, "Invalid number provided", http.StatusBadRequest)
		return
	}

	result, err := fibonacci(n, startX, startY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

// fibonacci generates a Fibonacci sequence up to positive integer
// n, inclusive, starting with the specified startx and starty values.
// For startx=0 and starty=1, it uses pregenerated Fibonacci numbers.
func fibonacci(n int, startx int, starty int) ([]int, error) {
	if n < 0 || startx < 0 || starty < 0 {
		return nil, errors.New("invalid input, n, startx, and starty must be non-negative integers")
	}

	// For standard start location, just return a chunk of the pregenerated sequence.
	if startx == 0 && starty == 1 {
		if n >= maxPregenerated {
			return pregeneratedSequence, nil
		}

		for i, e := range pregeneratedSequence {
			if e > n {
				return pregeneratedSequence[:i], nil
			}
		}
	}

	var result []int

	// deal with out of sequence startx and starty
	if startx < starty {
		result = append(result, startx, starty)
	} else {
		result = append(result, starty, startx)
	}

	// start at 2 because the first two already exist
	for i := 2; i <= n; i++ {
		next := result[i-1] + result[i-2]
		if next > n || next > maxPregenerated {
			// If the next Fibonacci number exceeds n or maxPregenerated stop.
			break
		}
		result = append(result, next)
	}

	return result, nil
}
