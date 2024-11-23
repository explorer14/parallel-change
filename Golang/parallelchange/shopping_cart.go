package parallelchange

type ShoppingCart struct {

	//
	// the goal is to remove this field, replacing with:
	// prices []int
	//

	prices []int
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (c *ShoppingCart) Add(price int) {
	c.prices = append(c.prices, price)
}

func (c *ShoppingCart) NumberOfProducts() int {
	return len(c.prices)
}

func (c *ShoppingCart) CalculateTotalPrice() int {
	var sum = 0
	for _, price := range c.prices {
		sum += price
	}
	return sum
}

func (c *ShoppingCart) HasDiscount() bool {
	var hasDiscount = false
	for _, price := range c.prices {
		if price >= 100 {
			hasDiscount = true
			break
		}
	}
	return hasDiscount
}
