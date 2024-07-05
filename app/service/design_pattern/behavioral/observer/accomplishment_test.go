package observer

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test(t *testing.T) {

	shirtItem := newItem("Nike Shirt")

	observerService(shirtItem)

	assert.Equal(t, shirtItem.inStock, true)
}
