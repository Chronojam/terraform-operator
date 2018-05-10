package plural

var consonants = "bcdfghjklmnpqrsttvwxyz"

func Plural(singular string) string {
	var plural string
	if len(singular) < 2 {
		return plural
	}

	switch rune(singular[len(singular)-1]) {
	case 's', 'x', 'z':
		plural = esPlural(singular)
	case 'y':
		sl := rune(singular[len(singular)-2])
		if isConsonant(sl) {
			plural = iesPlural(singular)
		} else {
			plural = sPlural(singular)
		}
	case 'h':
		sl := rune(singular[len(singular)-2])
		if sl == 'c' || sl == 's' {
			plural = esPlural(singular)
		} else {
			plural = sPlural(singular)
		}
	case 'e':
		sl := rune(singular[len(singular)-2])
		if sl == 'f' {
			plural = vesPlural(singular[:len(singular)-1])
		} else {
			plural = sPlural(singular)
		}
	case 'f':
		plural = vesPlural(singular)
	default:
		plural = sPlural(singular)
	}
	return (plural)
}

func iesPlural(singular string) string {
	return singular[:len(singular)-1] + "ies"
}

func vesPlural(singular string) string {
	return singular[:len(singular)-1] + "ves"
}

func esPlural(singular string) string {
	return singular + "es"
}

func sPlural(singular string) string {
	return singular + "s"
}

func isConsonant(char rune) bool {
	for _, c := range consonants {
		if char == c {
			return true
		}
	}
	return false
}
