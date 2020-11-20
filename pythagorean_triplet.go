package pythagorean

type Triplet [3]int

func Range(min, max int) (t []Triplet) {

	var a, b, c int
	for a = min; a <= max; a++ {
		for b = a; b <= max; b++ {
			for c = b; c <= max; c++ {
				if a*a+b*b == c*c {
					t = append(t, Triplet{a, b, c})
				}
			}
		}
	}
	return t
}

func Sum(n int) (t []Triplet) {

	var a, b, c int
	for a = 1; a <= n; a++ {
		for b = a; b <= n; b++ {
			if c = n - a - b; a*a+b*b == c*c {
				t = append(t, Triplet{a, b, c})
			}
		}
	}

	return t
}
