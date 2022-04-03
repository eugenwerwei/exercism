package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		averagePreparationTime = 2
	}

	return len(layers) * averagePreparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	var noodlesCount int
	var sauceCount int

	for _, v := range layers {
		if v == "noodles" {
			noodlesCount++
		}

		if v == "sauce" {
			sauceCount++
		}
	}

	return noodlesCount * 50, float64(sauceCount) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	var desiredAmounts []float64
	desiredAmounts = append(desiredAmounts, amounts...)

	for i, v := range desiredAmounts {
		desiredAmounts[i] = v / 2 * float64(portions)
	}

	return desiredAmounts
}
