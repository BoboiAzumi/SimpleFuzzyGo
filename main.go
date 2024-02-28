package main

import (
	"fmt"
	"simplefuzzy/Fuzzy"
	"strconv"
)

func main() {
	FL := Fuzzy.NewFuzzyLogic()

	FL.SetMinDomain("Brrrr", 0, 10.0)
	FL.SetMaxDomain("Bekasi", 80.0, 100.0)

	FL.AddTriangleFunction(5.0, 10.0, 15.0, "Dingin")
	FL.AddTriangleFunction(10.0, 15.0, 20.0, "Sejuk")
	FL.AddTriangleFunction(15.0, 20.0, 25.0, "Nyaman")
	FL.AddTriangleFunction(20.0, 25.0, 30.0, "Biasa")
	FL.AddTriangleFunction(25.0, 30.0, 35.0, "Hangat")
	FL.AddTriangleFunction(30.0, 35.0, 40.0, "Gerah")
	FL.AddTriangleFunction(35.0, 60.0, 90.0, "Panas")

	// Isi Temperature Disini
	var temp float32 = 13.5

	FL.CalculateWeight(temp)
	Result := FL.GetResult()

	fmt.Printf("Suhu %s°C kira kira \n\n", strconv.FormatFloat(float64(temp), 'f', -1, 32))

	for _, element := range Result {
		fmt.Println(element.Label, "\t: ", element.Weight)
	}

	MaxWeight := FL.GetMaxWeight()

	fmt.Printf("\nBerdasarkan bobot diatas dapat disimpulkan bahwa %s°C adalah %s\ndengan bobot %f\n", strconv.FormatFloat(float64(temp), 'f', -1, 32), MaxWeight.Label, MaxWeight.Weight)
}
