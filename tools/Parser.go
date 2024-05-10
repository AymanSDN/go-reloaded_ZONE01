package tools

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func HandlePunct(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == ' ' && IsPunct(rune(s[i+1])) {
			s = s[:i] + s[i+1:]
		}
	}
	for i := 0; i < len(s)-1; i++ {
		if IsPunct(rune(s[i])) && (!IsPunct(rune(s[i+1])) && s[i+1] != '\'') {
			s = s[:i+1] + " " + s[i+1:]
		}
	}
	return strings.Join(RemoveEmptyString(strings.Split(s, " ")), " ")
}

func GetApostIdx(s string) []int {
	s_size := len(s)
	var tab []int
	if s[0] == '\'' {
		tab = append(tab, 0)
	}
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '\'' && (!unicode.IsLetter(rune(s[i+1])) || !unicode.IsLetter(rune(s[i-1]))) {
			tab = append(tab, i)
		}
	}
	if s[s_size-1] == '\'' {
		tab = append(tab, s_size-1)
	}
	return tab
}

func HandleSingleQuote(s string) string {
	lines := strings.Split(s, "\n")
	for l := 0; l < len(lines); l++ {
		var tmp []string
		if len(lines[l]) > 3 {
			tab := GetApostIdx(lines[l])
			size := len(tab)
			var lasttxt string
			if size > 0 {
				lasttxt = lines[l][tab[len(tab)-1]+1:]
				lines[l] = lines[l][:tab[len(tab)-1]+1]
			}
			if size == 0 {
				return s
			}
			i := 1
			tmp = append(tmp, strings.TrimSpace(lines[l][:tab[0]]))
			for ; i < size; i++ {
				if i%2 == 1 {
					tmp = append(tmp, "'"+strings.TrimSpace(lines[l][tab[i-1]+1:tab[i]])+"'")
				} else {
					tmp = append(tmp, strings.TrimSpace(lines[l][tab[i-1]+1:tab[i]]))
				}
			}
			if size%2 == 1 {
				tmp = append(tmp, strings.TrimSpace(lines[l][tab[i-1]:]))
			}
			lines[l] = strings.Join(tmp, " ")
			tab = GetApostIdx(lines[l])
			if lines[l][tab[len(tab)-1]] == '\'' && lines[l][tab[len(tab)-1]-1] == ' ' {
				lines[l] = lines[l][:tab[len(tab)-1]-1] + lines[l][tab[len(tab)-1]:]
			}
			lines[l] += lasttxt
		}
	}
	return strings.Join(lines, "\n")
}

func HandleA(s string) string {
	size := len(s)
	if (s[size-1] == 'a' || s[size-1] == 'A') && (size == 1 || !unicode.IsLetter(rune(s[size-2]))) {
		s += "n"
	}
	return s
}

func HandleVowels(splitted []string) string {
	for i := 1; i < len(splitted); i++ {
		if IsVowel(rune(splitted[i][0])) {
			splitted[i-1] = HandleA(splitted[i-1])
		}
	}
	return strings.Join(RemoveEmptyString(splitted), " ")
}

func GetNumber(s string, rule string) (int, bool) {
	size := len(s)
	for i := 0; i < size; i++ {
		if unicode.IsNumber(rune(s[i])) && s[size-1] == ')' {
			num, err := strconv.Atoi(strings.TrimSpace(s[:size-1]))
			CheckError(err, "Error: failed to get second argument: "+rule)
			return num, true
		} else {
			log.Fatal("Error: cannot find a proper flag " + rule + "\n")
		}
	}
	return 1, false
}

func Bin(i int, splitted []string) []string {
	if strings.Contains(splitted[i], "(bin)") && i != 0 {
		x := 1
		if splitted[i] != "(bin)" {
			log.Fatal("Error: make sure that you have entred \"(bin)\"")
		}
		if splitted[i-x] == "" && i-x-1 >= 0 {
			if splitted[i-x-1] == "" {
				x++
			}
			x++
		}
		decimal, err := strconv.ParseInt(splitted[i-x], 2, 64)
		CheckError(err, "Error: make sure that ["+splitted[i-x]+"] is a binary number.")
		splitted[i-x] = strconv.Itoa(int(decimal))
		splitted[i] = ""
	}
	return splitted
}

func Hex(i int, splitted []string) []string {
	if strings.Contains(splitted[i], "(hex)") && i != 0 {
		x := 1
		if splitted[i] != "(hex)" {
			log.Fatal("Error: make sure that you have entred \"(hex)\"")
		}
		if splitted[i-x] == "" && i-x-1 >= 0 {
			if splitted[i-x-1] == "" {
				x++
			}
			x++
		}
		decimal, err := strconv.ParseInt(splitted[i-x], 16, 64)
		CheckError(err, "Error: make sure that ["+splitted[i-x]+"] is an Hexadecimal number.")
		splitted[i-x] = strconv.Itoa(int(decimal))
		splitted[i] = ""
	}
	return splitted
}

func Up(i int, splitted []string) []string {
	size := len(splitted)
	num, flag := 1, false
	if strings.Contains(splitted[i], "(up") {
		if i+1 < size && splitted[i] != "(up)" && splitted[i] == "(up," {
			num, flag = GetNumber(splitted[i+1], "(up, <number>)")
		} else if splitted[i] != "(up)" {
			log.Fatal("Error: make sure that you have entred \"(up)\" or \"(up, <number>)\"")
		}
		if flag {
			splitted[i], splitted[i+1] = "", ""
		}
		splitted[i] = ""
		for j := 1; j <= num && j <= i; j++ {
			if splitted[i-j] == "" {
				num++
			} else if !ContainsLetter(splitted[i-j]) {
				log.Fatal("Error: [" + splitted[i-j] + "] cannot be uppercase")
			}
			splitted[i-j] = strings.ToUpper(splitted[i-j])
		}
	}
	return splitted
}

func Low(i int, splitted []string) []string {
	size := len(splitted)
	num, flag := 1, false

	if strings.Contains(splitted[i], "(low") {
		if i+1 < size && splitted[i] != "(low)" && splitted[i] == "(low," {
			num, flag = GetNumber(splitted[i+1], "(low, <number>)")
		} else if splitted[i] != "(low)" {
			log.Fatal("Error: make sure that you have entred \"(low)\" or \"(low, <number>)\"")
		}
		if flag {
			splitted[i], splitted[i+1] = "", ""
		}
		splitted[i] = ""
		for j := 1; j <= num && j <= i; j++ {
			if splitted[i-j] == "" {
				num++
			} else if !ContainsLetter(splitted[i-j]) {
				log.Fatal("Error: [" + splitted[i-j] + "] cannot be lowercase")
			}
			splitted[i-j] = strings.ToLower(splitted[i-j])
		}
	}
	return splitted
}

func Cap(i int, splitted []string) []string {
	size := len(splitted)
	num, flag := 1, false
	if strings.Contains(splitted[i], "(cap") {
		if i+1 < size && splitted[i] != "(cap)" && splitted[i] == "(cap," {
			num, flag = GetNumber(splitted[i+1], "(cap, <number>)")
		} else if splitted[i] != "(cap)" {
			log.Fatal("Error: make sure that you have entred \"(cap)\" or \"(cap, <number>)\"")
		}
		if flag {
			splitted[i], splitted[i+1] = "", ""
		}
		splitted[i] = ""
		for j := 1; j <= num && j <= i; j++ {
			if splitted[i-j] == "" {
				num++
			} else if !ContainsLetter(splitted[i-j]) {
				log.Fatal("Error: [" + splitted[i-j] + "] cannot be capitalized")
			}
			splitted[i-j] = strings.Title(strings.ToLower(splitted[i-j]))
		}
	}

	return splitted
}

func HandleWordSingleQuote(i int, splitted []string) []string {
	size := len(splitted)
	if splitted[i] != "" && i+1 < size {
		if size > 2 && splitted[i] == "'" && IsWordApost(splitted[i+1]) {
			splitted[i+1] = splitted[i-1] + splitted[i] + splitted[i+1]
			splitted[i-1], splitted[i] = "", ""
		} else if size > 1 && splitted[i][0] == '\'' && IsWordApost(splitted[i][1:]) {
			splitted[i] = splitted[i-1] + splitted[i]
			splitted[i-1] = ""
		} else if size > 1 && splitted[i][len(splitted[i])-1] == '\'' && IsWordApost(splitted[i+1]) {
			splitted[i+1] = splitted[i] + splitted[i+1]
			splitted[i] = ""
		}
	}
	return splitted
}

func ApplyRules(text string) []string {
	splitted := RemoveEmptyString(strings.Split(text, " "))
	size := len(splitted)
	for i := 0; i < size; i++ {
		// check if the string contains "bin" and convert it into decimal
		splitted = Bin(i, splitted)
		// check if the string contains "hex" and convert it into decimal
		splitted = Hex(i, splitted)
		// check if the string contains "up" and convert it into uppercase
		splitted = Up(i, splitted)
		// check if the string contains "low" and convert it into lowercase
		splitted = Low(i, splitted)
		// check if the string contains "cap" and capitalized it
		splitted = Cap(i, splitted)
		splitted = HandleWordSingleQuote(i, splitted)
	}
	return RemoveEmptyString(splitted)
}

func Parse_Text(text string) string {
	splitted := ApplyRules(strings.ReplaceAll(text, "\n", " \n"))
	text = HandleVowels(splitted)
	return HandlePunct(HandleSingleQuote(text))
}
