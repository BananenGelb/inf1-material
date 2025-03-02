package aufgabe3

// CountX erwartet eine Liste von Zahlen sowie eine Zahl x.
// CountX liefert die Anzahl der Vorkommen von x in der Liste.
func CountX(list []int, x int) int {
	counter := 0
	for _, pos := range list {
		if pos == x {
			counter++
		}
	}
	return counter
}
