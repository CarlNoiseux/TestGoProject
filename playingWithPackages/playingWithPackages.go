package playingWithPackages

// Names are exported when they start with a capitalized letter, otherwise they are private to the package
func add(x int, y int) int {
	return x + y
}

func Add(x int, y int) int {
	return x + y
}
