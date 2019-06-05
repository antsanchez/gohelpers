package helpers

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var r *rand.Rand // Rand for this package
func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// CheckError check if err is not nil and print it to log
func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// ContainsOr checks if any of the substrings is contained in the string
// It is case sensitive
func ContainsOr(s string, substr []string) bool {

	for _, value := range substr {
		if strings.Contains(s, value) {
			return true
		}
	}

	return false
}

// StartsWith checks if the string s starts exactly with the substr
// It is case sensitive
func StartsWith(s string, substr string) bool {

	lenString := len(s)
	lenSubstring := len(substr)

	if lenString < lenSubstring {
		return false
	}

	fragment := string(s[0:lenSubstring])

	if strings.Compare(fragment, substr) == 0 {
		return true
	}

	return false
}

func CreateSlug(name string) (slug string) {

	type Replacement struct {
		Original string
		Replace  string
	}

	replacements := []Replacement{
		Replacement{Original: "ü", Replace: "ue"},
		Replacement{Original: "ö", Replace: "eo"},
		Replacement{Original: "ä", Replace: "ae"},
		Replacement{Original: "ß", Replace: "ss"},
		Replacement{Original: "ñ", Replace: "n"},
	}

	for _, repl := range replacements {
		slug = strings.Replace(name, repl.Original, repl.Replace, -1)
	}

	slug = strings.ToLower(slug)

	var slugRegexp = regexp.MustCompile("[^a-z0-9 ]+")
	slug = slugRegexp.ReplaceAllString(slug, "")
	slug = strings.Replace(slug, " ", "-", -1)

	return slug
}

// RemoveRedundantWhiteSpaces remove all leading/trailing whitespace as well as all two or more whitespace symbols inside a string
func RemoveRedundantWhiteSpaces(s string) string {
	s = regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`).ReplaceAllString(s, "")
	return regexp.MustCompile(`[\s\p{Zs}]{2,}`).ReplaceAllString(s, " ")
}

// RandomString generate a random string of the given length
func RandomString(strlen int) string {

	if strlen == 0 {
		strlen = 32
	}

	const lower = "abcdefghijklmnopqrstuvwxyz"
	const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const numbers = "0123456789"
	const nonalpha = "#@-.$*()+;~:'/%_?,=&!"

	chars := lower + upper + numbers + nonalpha

	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}

	return string(result)
}

// FileExists checks if a given file exits
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
