package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	assert.EqualError(t, order.Validate(), "id is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}
	assert.EqualError(t, order.Validate(), "invalid price")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.EqualError(t, order.Validate(), "invalid tax")
}

func Test_WithAllValidParams(t *testing.T) {
	order := Order{ID: "123", Price: 10, Tax: 12}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 12.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 22.0, order.FinalPrice)
}
