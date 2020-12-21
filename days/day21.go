package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2020/utils"
)

type d21Ingredient string
type d21Allergen string
type d21Food struct {
	ingredients []d21Ingredient
	allergens   []d21Allergen
}

func (r *Runner) Day21Part1(in string) string {

	totalAppearances, _ := r.d21Solve(in)
	return strconv.Itoa(totalAppearances)

}

func (r *Runner) Day21Part2(in string) string {

	_, solvedAllergens := r.d21Solve(in)

	allergens := []string{}
	for a := range solvedAllergens {
		allergens = append(allergens, string(a))
	}
	sort.Strings(allergens)

	ingredients := []string{}
	for _, a := range allergens {
		ingredients = append(ingredients, string(solvedAllergens[d21Allergen(a)]))
	}
	return strings.Join(ingredients, ",")
}

func (r *Runner) d21Solve(in string) (int, map[d21Allergen]d21Ingredient) {
	foodDetails := utils.StringsFromRegex(in, `^([a-z]+( [a-z]+)*) \(contains ([a-z ,]+)\)$`)

	allFoods := []d21Food{}
	allAllergens := map[d21Allergen]int{}
	allIngredients := map[d21Ingredient]int{}
	for _, f := range foodDetails {

		var ingredients []d21Ingredient
		var allergens []d21Allergen

		for _, i := range strings.Split(f[1], " ") {
			ingredients = append(ingredients, d21Ingredient(i))
		}
		for _, a := range utils.CsvToStrings(strings.ReplaceAll(f[3], " ", "")) {
			allergens = append(allergens, d21Allergen(a))
		}

		for _, i := range ingredients {
			allIngredients[i]++
		}
		for _, a := range allergens {
			allAllergens[a]++
		}

		allFoods = append(allFoods, d21Food{
			ingredients: ingredients,
			allergens:   allergens,
		})
	}

	allergensByIngredient := map[d21Ingredient]map[d21Allergen]bool{}
	ingredientsByAllergen := map[d21Allergen]map[d21Ingredient]bool{}

	for i := range allIngredients {
		allergensByIngredient[i] = map[d21Allergen]bool{}
		for a := range allAllergens {
			allergensByIngredient[i][a] = true
		}
	}
	for a := range allAllergens {
		ingredientsByAllergen[a] = map[d21Ingredient]bool{}
		for i := range allIngredients {
			ingredientsByAllergen[a][i] = true
		}
	}

	for _, f := range allFoods {

		for _, a := range f.allergens {
			newMap := map[d21Ingredient]bool{}
			for _, i := range f.ingredients {
				if ingredientsByAllergen[a][i] {
					newMap[i] = true
				}
			}
			ingredientsByAllergen[a] = newMap
		}

	}
	if r.verbose {
		fmt.Println("ingredientsByAllergen", ingredientsByAllergen)
	}

	solvedAllergens := map[d21Allergen]d21Ingredient{}
	changeMade := true
	for changeMade {
		changeMade = false

		for a, ingredients := range ingredientsByAllergen {

			if len(ingredients) == 1 && solvedAllergens[a] == "" {
				for i := range ingredients {
					solvedAllergens[a] = i
				}
			}

		}

		for a, ingredients := range ingredientsByAllergen {
			origLen := len(ingredients)
			for solvedA, solvedI := range solvedAllergens {
				if a == solvedA {
					continue
				}

				delete(ingredients, solvedI)
			}
			if origLen != len(ingredients) {
				ingredientsByAllergen[a] = ingredients
				changeMade = true
			}

		}

	}

	if r.verbose {
		fmt.Println("ingredientsByAllergen after solving", ingredientsByAllergen)
	}

	nonAllergenicIngredients := []d21Ingredient{}
	for i, allergens := range allergensByIngredient {
		for solvedA, solvedI := range solvedAllergens {

			if i == solvedI {
				continue
			}
			delete(allergens, solvedA)
		}
		allergensByIngredient[i] = allergens
		if len(allergens) == 0 {
			nonAllergenicIngredients = append(nonAllergenicIngredients, i)
		}
	}

	if r.verbose {
		fmt.Println("nonAllergenicIngredients", nonAllergenicIngredients)
	}

	totalAppearances := 0
	for _, i := range nonAllergenicIngredients {
		totalAppearances += allIngredients[i]
	}

	return totalAppearances, solvedAllergens
}
