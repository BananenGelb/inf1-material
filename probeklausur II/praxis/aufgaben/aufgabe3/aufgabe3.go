package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountOdd.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein.
*/

// CountOdd erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen darin.
func CountOdd(list []int) int {
	// TODO
	if len(list) == 0 {
		return 0
	}
	s, result := list[0], CountOdd(list[1:])
	if s%2 != 0 {
		result++
	}
	return result

}
