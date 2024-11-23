package anothermodule

import "main/parallelchange"

func AddItem() {
	shoppingCart := parallelchange.NewShoppingCart()
	shoppingCart.Add(10)
}
