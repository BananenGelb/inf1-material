package aufgabe6

// import "slices"

// Duplicates erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die mehr als einmal in l1 vorkommen.
// Im Ergebnis kommt jedes Element nur einmal vor.
// Die Elemente stehen im Ergebnis in der Reihenfolge ihres ersten Auftretens.
func Duplicates(list []int) []int {
	k := []int{}
	for i, el1 := range list {
		if contains(list[i+1:], el1) {
			if !contains(k, el1) {
				k = append(k, el1)
			}

		}
	}
	return k
}
func contains(l []int, el int) bool {
	for _, e := range l {
		if e == el {
			return true
		}
	}
	return false
}

//slices.Sort(list)
//for i := 0; i < len(list); i++ {
//	if list [i] < list [i] {
//		result = append(result, list[i])
//	}
// }
