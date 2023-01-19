package verification

import (
	"fmt"
	"regexp"

	"github.com/dongri/phonenumber"
	"github.com/nyaruka/phonenumbers"
)

////////////  NUMBER VERIFICATION

func MustCompileNumber(getphone string) bool {
	re := regexp.MustCompile(`((?:\+|00)[17](?: |\-)?|(?:\+|00)[1-9]\d{0,2}(?: |\-)?|(?:\+|00)1\-\d{3}(?: |\-)?)?(0\d|\([0-9]{3}\)|[1-9]{0,3})(?:((?: |\-)[0-9]{2}){4}|((?:[0-9]{2}){4})|((?: |\-)[0-9]{3}(?: |\-)[0-9]{4})|([0-9]{7}))`)
	matched := re.MatchString(getphone)

	if !matched {
		fmt.Printf("Oops, wrong number %v\n", matched)
		b := false
		return b
	}

	b := true
	fmt.Printf("MustCompile Phone: %v\t%v\n", getphone, re.MatchString(getphone))
	return b
}

// /////// CHANGING NUMBER TO OUR STANDARTS

func IsNumberCorrect(getphone string) string {
	isEight := regexp.MustCompile(`^[8]`)
	if isEight.MatchString(getphone) {
		getphone = ReplaceFirstRune(getphone, "7")
		fmt.Println("isEight: ", getphone)
	}

	isFirstLetter := regexp.MustCompile(`^[+8]`)
	if isFirstLetter.MatchString(getphone) {
		getphone = ReplaceFirstTwoRune(getphone, "+7")
		fmt.Println("isFirstLetter: ", getphone)
	}

	isPlusSignExists := regexp.MustCompile(`^[+]`)
	if !isPlusSignExists.MatchString(getphone) {
		addPlus := "+" + getphone
		getphone = addPlus
		fmt.Printf("IsPlusSignExists: %v\n", getphone)
	}
	return getphone
}

func ReplaceFirstTwoRune(str, replacement string) string {
	return string([]rune(str)[:0]) + replacement + string([]rune(str)[2:])
}

func ReplaceFirstRune(str, replacement string) string {
	return string([]rune(str)[:0]) + replacement + string([]rune(str)[1:])
}

///////////// VERIFICATION IF STANDART IS OK

func Verification(getphone string) (string, bool) {
	num, err := phonenumbers.Parse(getphone, "")

	/// international
	formattedNum := phonenumbers.Format(num, phonenumbers.INTERNATIONAL)
	if err != nil {
		fmt.Println("Something Wrong Here:", err)
	}
	fmt.Printf("International Version: %v\n", formattedNum)

	/// verification
	verify := phonenumbers.IsValidNumber(num)
	fmt.Printf("Verification: %v\n", verify)

	return formattedNum, verify
}

///////// GETTING THE NAME OF THE COUNTRY

func GetCountryName(getphone string) string {
	formattedPhone := phonenumber.Parse(getphone, "")

	includeLandLine := true
	country := phonenumber.GetISO3166ByNumber(formattedPhone, includeLandLine)
	fmt.Printf("Country: %v\n", country.CountryName)
	return country.CountryName
}
