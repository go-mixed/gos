package main

func add(i, j int) int {
	return i + j
}

func addS(i, j string) int {
	return addSafe(i, i)
}
