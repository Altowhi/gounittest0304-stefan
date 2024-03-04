package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStefansTest(t *testing.T) {
	//t.Fatal("Failed")
}

// ett unittest är en funktion som testar en specifik setup/parameteriseing av en viss funktion
func TestWhen_Krogen_And_19_years_old_is_not_drunk_Then_I_should_get_beer(t *testing.T) {
	// Arrange
	location := "K"
	age := 19
	promille := 0.0

	//Act
	canBuy, err := CanBuyBeer(location, age, float32(promille))

	//Assert
	assert.True(t, canBuy)
	assert.Nil(t, err)
	// if canBuy == false || err != nil {
	// 	t.Fatal("Not allowed but should be")
	// }

}

func TestWhen_Krogen_And_17_years_old_is_not_drunk_Then_I_should_not_get_beer(t *testing.T) {
	// Arrange
	location := "K"
	age := 17
	promille := 0.0

	//Act
	canBuy, err := CanBuyBeer(location, age, float32(promille))

	//Assert
	assert.False(t, canBuy)
	assert.Nil(t, err)
	// if canBuy == false || err != nil {
	// 	t.Fatal("Not allowed but should be")
	// }

}

func TestWhen_Systemenet_And_20_years_old_is_not_drunk_Then_I_should_get_beer(t *testing.T) {
	// Arrange
	location := "S"
	age := 20
	promille := 0.0

	//Act
	canBuy, err := CanBuyBeer(location, age, float32(promille))

	//Assert
	assert.True(t, canBuy)
	assert.Nil(t, err)
	// if canBuy == false || err != nil {
	// 	t.Fatal("Not allowed but should be")
	// }

}
