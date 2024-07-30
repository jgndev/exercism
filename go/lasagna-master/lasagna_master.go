package lasagna

// PreparationTime calculates the number of minutes required to prepare the lasagna
// represented as an int. If minutes provided by the caller is passed as 0, the default
// preparation time per layer of 2 minutes is used.
func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		return len(layers) * 2
	}

	return len(layers) * minutes
}

// Quantities takes in a slice of layers for the Lasagna and returns an int representing
// the grams of noodles and a float64 representing the liters of sauce. Each layer of
// noodles in the layers slice adds 50g of noodles. Each layer in the slice adds 0.2
// liters of sauce.
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, l := range layers {
		if l == "noodles" {
			noodles += 50
		}

		if l == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// AddSecretIngredient takes the last element in friendList of ingredients
// and appends it to myList of ingredients. It is assumed that the last element
// in friendList is always the "secret ingredient" for a special lasagna.
func AddSecretIngredient(friendList, myList []string) {
	secretIngredient := friendList[len(friendList)-1]
	// Last element in myList is "?", and needs to be replaced with the last element
	// in friendList.
	myList[len(myList)-1] = secretIngredient
}

// ScaleRecipe returns a slice of []float64 with quantities of ingredients adjusted
// for the number of portions passed. The slice of []float64 passed in as quantities
// represents the amount of ingredients needed to make 2 portions. The adjusted
// quantities required for the number of portions is returned as a []float64.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := make([]float64, len(quantities))
	copy(newQuantities, quantities)

	// quantities are passed in for default values for 2 portions
	// { 1.2, 3.6, 10.5}
	// 1 = ( 0.6, 1.8, 5.25)
	// 2 = { 1.2, 3.6, 10.5 }
	// 3 = { 1.8, 5.4, 15.75 }
	// 4 = { 2.4, 7.2, 21.0 }

	sf := float64(portions) / 2.0

	for i, v := range quantities {
		newQuantities[i] = v * sf
	}

	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
