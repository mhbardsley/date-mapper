package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Check that we have the correct number of arguments
	if len(os.Args) != 6 {
		fmt.Println("Usage: mapdate x a b c d")
		return
	}

	// Parse the arguments
	x, err := time.Parse("2006-01-02", os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	a, err := time.Parse("2006-01-02", os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := time.Parse("2006-01-02", os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	c, err := time.Parse("2006-01-02", os.Args[4])
	if err != nil {
		fmt.Println(err)
		return
	}
	d, err := time.Parse("2006-01-02", os.Args[5])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Map x to the equivalent date in the range (c,d)
	y, err := mapDate(x, a, b, c, d)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(y)
}

// Maps a date within the range of (a,b) to the equivalent one in (c,d) using a linear transformation.
// Returns an error if x is not within the range (a,b).
func mapDate(x, a, b, c, d time.Time) (time.Time, error) {
	if x.Before(a) || x.After(b) {
		return time.Time{}, fmt.Errorf("x is not within the range (a,b)")
	}

	// Calculate the difference between a and b in nanoseconds
	ab := b.Sub(a).Nanoseconds()

	// Calculate the difference between c and d in nanoseconds
	cd := d.Sub(c).Nanoseconds()

	// Calculate the ratio of the difference between x and a to the difference between b and a
	ratio := float64(x.Sub(a).Nanoseconds()) / float64(ab)

	// Map x to the equivalent date in the range (c,d) using the linear transformation
	return c.Add(time.Duration(ratio * float64(cd))), nil
}
