package main

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

// const passwordLength int = 2
// const useNumbers bool = false
// const useLowercaseLetters bool = true
// const useUppercaseLetters bool = true

const letters string = "abcdefghijklmnopqrstuvwxyz"

var passwordHistory []string

// func main() {
// 	fmt.Println(generatePassword(0))
// 	fmt.Println(generatePassword(0))
// 	fmt.Println(generatePassword(0))
// 	fmt.Println(generatePassword(0))
// 	fmt.Println(generatePassword(0))
// 	fmt.Println(generatePassword(0))
// 	fmt.Println(generatePassword(0))
// 	fmt.Println(generatePassword(0))
// 	fmt.Println(generatePassword(0))
// 	fmt.Println(generatePassword(0))
// 	fmt.Println(generatePassword(0))
// }

func generateRandomSymbol(sets []func() string) string {
	return sets[rand.Intn(len(sets))]()
}

func generatePassword(attempt int, passwordLength int, useNumbers, useLowercaseLetters, useUppercaseLetters bool) (string, error) {
	sets := make([]func() string, 0)
	password := ""

	if useNumbers {
		sets = append(sets, func() string { return strconv.Itoa(rand.Intn(10)) })
	}

	if useLowercaseLetters {
		sets = append(sets, func() string { return string(letters[rand.Intn(len(letters))]) })
	}

	if useUppercaseLetters {
		sets = append(sets, func() string { return strings.ToUpper(string(letters[rand.Intn(len(letters))])) })
	}

	if useNumbers && (!useUppercaseLetters && !useLowercaseLetters) && passwordLength > 10 {
		return "", errors.New("Can't generate 11 unique numbers!")
	}

	if passwordLength < len(sets) {
		return "", errors.New("Can't use all sets with this little password length!")
	}

	for i := 0; i < len(sets); i++ {
		password += sets[i]()
	}

	for i := 0; i < passwordLength-len(sets); i++ {

		// re-generate symbol if already exists in password
		symbol := generateRandomSymbol(sets)
		for strings.ContainsAny(password, symbol) {
			symbol = generateRandomSymbol(sets)
		}

		password += symbol
	}

	// // check if symbols from all sets are used in password, regenerate if not
	// var numberRegex = regexp.MustCompile(`[0-9]+`)
	// var lowerLetterRegex = regexp.MustCompile(`[a-z]+`)
	// var upperLetterRegex = regexp.MustCompile(`[A-Z]+`)

	// if useNumbers && !numberRegex.MatchString(password) {
	// 	password, _ = generatePassword()
	// }

	// if useLowercaseLetters && !lowerLetterRegex.MatchString(password) {
	// 	password, _ = generatePassword()
	// }

	// if useUppercaseLetters && !upperLetterRegex.MatchString(password) {
	// 	password, _ = generatePassword()
	// }

	// shuffle string
	shuff := []rune(password)
	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	password = string(shuff)

	// check if password differ from previous ones in this application cycle
	if attempt >= 999 {
		return "", errors.New("Can't generate more unique passwords with the same configuration after 999 attempts!")
	}

	for _, p := range passwordHistory {
		if p == password {

			// regenerate password
			pass, err := generatePassword(attempt+1, passwordLength, useNumbers, useLowercaseLetters, useUppercaseLetters)
			if err != nil {
				return "", errors.New("Can't generate more unique passwords with the same configuration after 999 attempts!")
			}
			password = pass
		}
	}

	passwordHistory = append(passwordHistory, password)
	return password, nil
}
