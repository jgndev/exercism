package lasagna

// PreparationTime accepts a slice of string for the layers of the lasagna and an int
// representing the number of average minutes per layer that preparation is expected
// to take. Returns the number of layers multiplied by the minutes per layer.
func PreparationTime(layers []string, minutesPerLayer int) int {
    // Use a default minutes per layer of 2 if the value passed in was 0.
    if minutesPerLayer == 0 {
        minutesPerLayer = 2
    }

    return len(layers) * minutesPerLayer
}

// Quantities takes a slice of string for the layers of the lasagna. Determines the
// quantity of noodles and sauce needed to create the lasagna. Noodles returned
// as an int, sauce returned as a float64.
func Quantities(layers []string) (int, float64) {
    var noodles int
    var sauce float64

    for _, v := range layers {
        switch {
        case v == "sauce":
            sauce += 0.2

        case v == "noodles":
            noodles += 50
        }
    }

    return noodles, sauce
}

// AddSecretIngredient takes two lists of string, the first representing a friend's lasagna
// recipe and the second representing my lasagna recipe. The secret ingredient is the last
// item in the friend's lasagna recipe. Modify my recipe in place to replace the last item
// of the recipe with the last item from the friend's recipe.
func AddSecretIngredient(friendsRecipe, myRecipe []string) {
    secretIngredient := friendsRecipe[len(friendsRecipe)-1]
    myRecipe[len(myRecipe)-1] = secretIngredient
}

// ScaleRecipe takes a slice of float64 representing quantities of ingredients to make two portions.
// Returns a slice of float64 representing the quantities scaled to make the number of servings
// passed in as the argument for servings. One serving is two portions. The number of ingredients
// needed to scale the recipe is each ingredient scaled to half the number of servings/portions.
func ScaleRecipe(amounts []float64, servings int) []float64 {
    scaledAmounts := make([]float64, len(amounts))

    for i := 0; i < len(amounts); i++ {
        // Correct for a serving being two portions
        onePortion := amounts[i] / 2
        scaledAmounts[i] = onePortion * float64(servings)
    }

    return scaledAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
