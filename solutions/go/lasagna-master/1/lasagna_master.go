package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    countLayers := len(layers)
    if time == 0 {
        return countLayers * 2
    } else {
    	return countLayers * time
    }
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noods := 0
	sawse := 0.0
	for layer := 0; layer < len(layers); layer++ {
		if layers[layer] == "noodles" {
			noods = noods + 50
		} else if layers[layer] == "sauce" {
			sawse = sawse + 0.2
		} else {
		}
	}
	return noods, sawse
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaledportion := make([]float64, len(quantities))
    for i := 0; i < len(quantities); i++ {
        scaledportion[i] = (quantities[i] / 2.0) * float64(portions)
    }
	return scaledportion    
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
