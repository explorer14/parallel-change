package parallelchange

import "testing"

func TestShoppingCartMayHaveJustOneItem(t *testing.T) {
	shoppingCart := NewShoppingCart()
	shoppingCart.Add(10)

	expected := 1
	if actual := shoppingCart.NumberOfProducts(); actual != expected {
		t.Errorf("Expected number of products to be %d, but got %d", expected, actual)
	}
}

func TestShoppingCartMayHaveMultipleItems(t *testing.T) {
	shoppingCart := NewShoppingCart()
	shoppingCart.Add(10)
	shoppingCart.Add(40)

	expected := 2
	if actual := shoppingCart.NumberOfProducts(); actual != expected {
		t.Errorf("Expected number of products to be %d, but got %d", expected, actual)
	}
}

func TestShoppingCartHasTotalPriceEqualToTotalPriceOfItsContents(t *testing.T) {
	shoppingCart := NewShoppingCart()
	shoppingCart.Add(10)

	expected := 10
	if actual := shoppingCart.CalculateTotalPrice(); actual != expected {
		t.Errorf("Expected total price to be %d, but got %d", expected, actual)
	}
}

func TestShoppingCartWithMultipleItemsHasTotalPriceEqualToTotalPriceOfItsContents(t *testing.T) {
	shoppingCart := NewShoppingCart()
	shoppingCart.Add(10)
	shoppingCart.Add(30)

	expected := 40
	if actual := shoppingCart.CalculateTotalPrice(); actual != expected {
		t.Errorf("Expected total price to be %d, but got %d", expected, actual)
	}
}

func TestShoppingCartHasDiscountWhenContainsAtLeastOnePremiumItem(t *testing.T) {
	shoppingCart := NewShoppingCart()
	shoppingCart.Add(100)

	expected := true
	if actual := shoppingCart.HasDiscount(); actual != expected {
		t.Errorf("Expected 'has discount' to be %v, but got %v", expected, actual)
	}
}

func TestShoppingCartWithMultipleItemsHasDiscountWhenContainsAtLeastOnePremiumItem(t *testing.T) {
	shoppingCart := NewShoppingCart()
	shoppingCart.Add(100)
	shoppingCart.Add(101)
	shoppingCart.Add(20)
	shoppingCart.Add(99)

	expected := true
	if actual := shoppingCart.HasDiscount(); actual != expected {
		t.Errorf("Expected 'has discount' to be %v, but got %v", expected, actual)
	}
}

func TestShoppingCartDoesNotHaveDiscountWhenAllItemsAreCheap(t *testing.T) {
	shoppingCart := NewShoppingCart()
	shoppingCart.Add(10)

	expected := false
	if actual := shoppingCart.HasDiscount(); actual != expected {
		t.Errorf("Expected 'has discount' to be %v, but got %v", expected, actual)
	}
}

func TestShoppingCartWithMultipleItemsDoesNotHaveDiscountWhenAllItemsAreCheap(t *testing.T) {
	shoppingCart := NewShoppingCart()
	shoppingCart.Add(10)
	shoppingCart.Add(30)
	shoppingCart.Add(99)

	expected := false
	if actual := shoppingCart.HasDiscount(); actual != expected {
		t.Errorf("Expected 'has discount' to be %v, but got %v", expected, actual)
	}
}
