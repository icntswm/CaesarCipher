package encr_src

type alphabet struct {
	eng_lower []rune
	eng_upper []rune
	rus_lower []rune
	rus_upper []rune
}

func initialAlphabet() alphabet {
	var structAlph = alphabet{
		eng_lower: []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
		eng_upper: []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'},
		rus_lower: []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'},
		rus_upper: []rune{'А', 'Б', 'В', 'Г', 'Д', 'Е', 'Ё', 'Ж', 'З', 'И', 'Й', 'К', 'Л', 'М', 'Н', 'О', 'П', 'Р', 'С', 'Т', 'У', 'Ф', 'Х', 'Ц', 'Ч', 'Ш', 'Щ', 'Ъ', 'Ы', 'Ь', 'Э', 'Ю', 'Я'},
	}
	return structAlph
}

func encryption(str []rune, alpha []rune, i, shift int) []rune {
	for j, _ := range alpha {
		if str[i] == alpha[j] {
			if j+shift < len(alpha) {
				str[i] = alpha[j+shift]
			} else {
				letterFromDict := j + shift - len(alpha)
				for letterFromDict > len(alpha) {
					letterFromDict -= len(alpha)
				}
				str[i] = alpha[letterFromDict]
			}
			break
		}
	}
	return str
}

func CaesarCipher(str []rune, shift int) []rune {

	primer := initialAlphabet()

	for i, _ := range str {
		if str[i] >= 'a' && str[i] <= 'z' {
			str = encryption(str, primer.eng_lower, i, shift)
		} else if str[i] >= 'A' && str[i] <= 'Z' {
			str = encryption(str, primer.eng_upper, i, shift)
		} else if str[i] >= 'а' && str[i] <= 'я' {
			str = encryption(str, primer.rus_lower, i, shift)
		} else if str[i] >= 'А' && str[i] <= 'Я' {
			str = encryption(str, primer.rus_upper, i, shift)
		}
	}
	return str
}
