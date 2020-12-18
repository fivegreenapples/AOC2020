package days

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

func (r *Runner) Day18Part1(in string) string {
	calcs := utils.Lines(in)

	totalSum := 0
	for _, c := range calcs {
		totalSum += r.d18Part1DoCalculation(c)
	}

	return strconv.Itoa(totalSum)
}

func (r *Runner) Day18Part2(in string) string {
	calcs := utils.Lines(in)

	totalSum := 0
	for _, c := range calcs {
		totalSum += r.d18Part2DoCalculation(c)
	}

	return strconv.Itoa(totalSum)
}

func (r *Runner) d18Part1DoCalculation(c string) int {

	scnr := bufio.NewScanner(strings.NewReader(c))
	scnr.Split(bufio.ScanBytes)

	type stackItem struct {
		total    int
		multiply bool
	}
	stack := []stackItem{}
	stack = append(stack, stackItem{})
	for scnr.Scan() {
		tok := scnr.Text() // will always return single byte string

		switch tok {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			// assumes single digit numbers
			num := int(tok[0]) - '0'
			if stack[len(stack)-1].multiply {
				stack[len(stack)-1].total *= num
			} else {
				stack[len(stack)-1].total += num
			}
		case "+":
			stack[len(stack)-1].multiply = false
		case "*":
			stack[len(stack)-1].multiply = true
		case " ":
			// ignore
		case "(":
			stack = append(stack, stackItem{})
		case ")":
			if stack[len(stack)-2].multiply {
				stack[len(stack)-2].total *= stack[len(stack)-1].total
			} else {
				stack[len(stack)-2].total += stack[len(stack)-1].total
			}
			stack = stack[:len(stack)-1]
		default:
			fmt.Println("unhandled token in calculation: " + tok)
			return 0
		}

	}
	return stack[0].total
}

func (r *Runner) d18Part2DoCalculation(c string) int {

	scnr := bufio.NewScanner(strings.NewReader(c))
	scnr.Split(bufio.ScanBytes)

	tokenisedInput := r.d18Part2Tokenise(scnr)
	result := r.d18Part2CalculateExpression(tokenisedInput)

	return result
}

type d18TokenType int

const (
	number d18TokenType = iota
	plus
	multiply
	expression
)

type d18Token struct {
	typ    d18TokenType
	val    int
	tokens []d18Token
}

func (r *Runner) d18Part2Tokenise(scnr *bufio.Scanner) []d18Token {

	tokens := []d18Token{}

	for scnr.Scan() {
		tok := scnr.Text() // will always return single byte string

		switch tok {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			// assumes single digit numbers
			tokens = append(tokens, d18Token{
				typ: number,
				val: int(tok[0]) - '0',
			})
		case "+":
			tokens = append(tokens, d18Token{
				typ: plus,
			})
		case "*":
			tokens = append(tokens, d18Token{
				typ: multiply,
			})
		case " ":
			// ignore
		case "(":
			tokens = append(tokens, d18Token{
				typ:    expression,
				tokens: r.d18Part2Tokenise(scnr),
			})
		case ")":
			return tokens
		default:
			fmt.Println("unhandled token in calculation: " + tok)
			return nil
		}

	}

	return tokens
}

func (r *Runner) d18Part2CalculateExpression(tokens []d18Token) int {

	toMultiply := []int{}
	tokIdx := 0
	for tokIdx < len(tokens) {

		switch tokens[tokIdx].typ {
		case number:
			toMultiply = append(toMultiply, tokens[tokIdx].val)
		case plus:
			// look ahead to next token. grab value and add to top of multiplication stack
			tokIdx++
			switch tokens[tokIdx].typ {
			case number:
				toMultiply[len(toMultiply)-1] += tokens[tokIdx].val
			case expression:
				toMultiply[len(toMultiply)-1] += r.d18Part2CalculateExpression(tokens[tokIdx].tokens)
			default:
				panic(fmt.Sprintf("unhandled token type: %d", tokens[tokIdx].typ))
			}
		case multiply:
			// ignore
		case expression:
			toMultiply = append(toMultiply, r.d18Part2CalculateExpression(tokens[tokIdx].tokens))
		default:
			panic(fmt.Sprintf("unhandled token type: %d", tokens[tokIdx].typ))
		}
		tokIdx++
	}

	result := 1
	for _, val := range toMultiply {
		result *= val
	}

	return result
}
