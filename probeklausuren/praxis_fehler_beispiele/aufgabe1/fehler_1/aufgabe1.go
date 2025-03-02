package aufgabe1

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {

	shortestLen := -1
	shortestPos := 100

	for pos, val := range list {
		currentLen := len(val)
		if currentLen >= 3 && val[:3] == "abc" {
			if currentLen > shortestLen {
				shortestLen = currentLen
				shortestPos = pos
			}
		}
	}
	if shortestPos != 100 {
		return list[shortestPos]
	}
	return ""
}
