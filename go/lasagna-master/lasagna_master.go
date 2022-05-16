package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (noodels int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodels += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, scaledQuantities int) (scaled []float64) {
	for _, item := range quantities {
		scaled = append(scaled, ((item / 2) * float64(scaledQuantities)))
	}
	return
}
