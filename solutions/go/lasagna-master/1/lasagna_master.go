package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageTimePerLayer int) int {
    if averageTimePerLayer == 0 {
        averageTimePerLayer = 2
    }
    return len(layers) * averageTimePerLayer
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var noodles int
    var sauce float64
    for _, layer := range layers {
        if layer == "sauce" {
            sauce += 0.2
        } else if layer == "noodles" {
            noodles += 50
        }
    }
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsIngredients []string, myIngredients []string) {
    myIngredients[len(myIngredients) - 1] = friendsIngredients[len(friendsIngredients) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
   	scalledQuantities := make([]float64,len(quantities))
    for index, quantity := range quantities {
		scalledQuantities[index] = (quantity / 2.0) * float64(portions)
    }
    return scalledQuantities
}

