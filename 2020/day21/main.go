package main

import (
	"fmt"
	"gnodivad/advent-of-code/utils"
	"strings"
)

func main() {
	foods := utils.ReadStringsFromFile("2020/day21/input.txt")
	ingredients, definiteAllergens := identifyIngredientsAndAllergensFromFoods(foods)

	fmt.Println("--- Part One ---")
	fmt.Println(calculateIngredientsThatDefiniteNotAllergens(ingredients, definiteAllergens))

	fmt.Println("--- Part Two ---")
	fmt.Println(sortAllergens(definiteAllergens))
}

func calculateIngredientsThatDefiniteNotAllergens(ingredients []string, definiteAllergens map[string]string) (count int) {
	for _, ingredient := range ingredients {
		if _, ok := definiteAllergens[ingredient]; !ok {
			count++
		}
	}

	return
}

func sortAllergens(definiteAllergens map[string]string) string {
	sortedAllergens := make([]string, 0)

	for _, kv := range utils.SortByValue(definiteAllergens) {
		sortedAllergens = append(sortedAllergens, kv.Key)
	}

	return strings.Join(sortedAllergens, ",")
}

func identifyIngredientsAndAllergensFromFoods(foods []string) (allIngredients []string, definiteAllergens map[string]string) {
	allIngredients = make([]string, 0)
	definiteAllergens = make(map[string]string)

	possibleAllergens := make(map[string][]string)

	for _, food := range foods {
		food = strings.TrimSuffix(food, ")")
		foodSlice := strings.Split(food, " (contains ")
		ingredientList, allergenList := foodSlice[0], foodSlice[1]
		ingredients, allergens := strings.Split(ingredientList, " "), strings.Split(allergenList, ", ")

		for _, allergen := range allergens {
			if _, ok := possibleAllergens[allergen]; !ok {
				possibleAllergens[allergen] = ingredients
			} else {
				possibleAllergens[allergen] = utils.Intersect(possibleAllergens[allergen], ingredients)
			}
		}

		allIngredients = append(allIngredients, ingredients...)
	}

	for len(possibleAllergens) > 0 {
		for allergen, possibleIngredients := range possibleAllergens {
			if len(possibleIngredients) == 1 {
				ingredient := possibleIngredients[0]
				definiteAllergens[ingredient] = allergen

				delete(possibleAllergens, allergen)

				for otherAllergen, otherAllergenIngredients := range possibleAllergens {
					possibleAllergens[otherAllergen] = utils.RemoveItem(otherAllergenIngredients, ingredient)
				}

				break
			}
		}
	}

	return
}
