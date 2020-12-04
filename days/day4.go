package days

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

//
// Create new days with:
// sed -e 's/4/2/g' template > day2.go; sed -e 's/4/2/g' template_test > day2_test.go
//

const (
	day4_byr = 0b0000_0001
	day4_iyr = 0b0000_0010
	day4_eyr = 0b0000_0100
	day4_hgt = 0b0000_1000
	day4_hcl = 0b0001_0000
	day4_ecl = 0b0010_0000
	day4_pid = 0b0100_0000
	day4_cid = 0b1000_0000
)

const day4ValidPassportChecksum = day4_byr | day4_iyr | day4_eyr | day4_hgt | day4_hcl | day4_ecl | day4_pid | day4_cid
const day5ValidNPCredentialsChecksum = day4_byr | day4_iyr | day4_eyr | day4_hgt | day4_hcl | day4_ecl | day4_pid

var day4CredentialMap = map[string]int{
	"byr": day4_byr,
	"iyr": day4_iyr,
	"eyr": day4_eyr,
	"hgt": day4_hgt,
	"hcl": day4_hcl,
	"ecl": day4_ecl,
	"pid": day4_pid,
	"cid": day4_cid,
}

func (r *Runner) Day4Part1(in string) string {

	in = strings.ReplaceAll(in, "\n\n", "___")
	in = strings.ReplaceAll(in, "\n", " ")
	in = strings.ReplaceAll(in, "___", "\n")

	info := utils.AllStringsFromRegex(in, `([a-z]{3}):[^ ]+`)

	validPassports := 0
	for _, testPassport := range info {
		checksum := 0
		for _, field := range testPassport {
			checksum += day4CredentialMap[field[1]]
		}
		if checksum == day4ValidPassportChecksum || checksum == day5ValidNPCredentialsChecksum {
			validPassports++
		}
	}

	return strconv.Itoa(validPassports)
}

func (r *Runner) Day4Part2(in string) string {

	in = strings.ReplaceAll(in, "\n\n", "___")
	in = strings.ReplaceAll(in, "\n", " ")
	in = strings.ReplaceAll(in, "___", "\n")

	info := utils.AllStringsFromRegex(in, `([a-z]{3}):([^ ]+)`)

	validPassports := 0
	for _, testPassport := range info {
		checksum := 0
		for _, field := range testPassport {
			if r.Day4IsCredentialValid(day4CredentialMap[field[1]], field[2]) {
				checksum += day4CredentialMap[field[1]]
			}
		}
		if checksum == day4ValidPassportChecksum || checksum == day5ValidNPCredentialsChecksum {
			validPassports++
		}
	}
	return strconv.Itoa(validPassports)
}

func (r *Runner) Day4IsCredentialValid(credType int, in string) bool {

	switch credType {
	case day4_byr:
		if year, err := strconv.Atoi(in); err == nil && len(in) == 4 {
			return year >= 1920 && year <= 2002
		}
	case day4_iyr:
		if year, err := strconv.Atoi(in); err == nil && len(in) == 4 {
			return year >= 2010 && year <= 2020
		}
	case day4_eyr:
		if year, err := strconv.Atoi(in); err == nil && len(in) == 4 {
			return year >= 2020 && year <= 2030
		}
	case day4_hgt:
		if len(in) == 5 && in[3:5] == "cm" {
			if height, err := strconv.Atoi(in[0:3]); err == nil {
				return height >= 150 && height <= 193
			}
		} else if len(in) == 4 && in[2:4] == "in" {
			if height, err := strconv.Atoi(in[0:2]); err == nil {
				return height >= 59 && height <= 76
			}
		}
	case day4_hcl:
		matched, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, in)
		return matched
	case day4_ecl:
		matched, _ := regexp.MatchString(`^amb|blu|brn|gry|grn|hzl|oth$`, in)
		return matched
	case day4_pid:
		matched, _ := regexp.MatchString(`^[0-9]{9}$`, in)
		return matched
	case day4_cid:
		return true

	}

	return false
}
