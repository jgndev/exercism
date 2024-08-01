package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, ok := units[unit]
	if !ok {
		return false
	}

	bill[item] += quantity
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantity, ok := units[unit]

	if !ok {
		return false
	}

	currentQuantity, exists := bill[item]
	if !exists {
		return false
	}

	if currentQuantity < quantity {
		return false
	}

	if quantity == 0 || currentQuantity == quantity {
		delete(bill, item)
	} else {
		bill[item] -= quantity
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, ok := bill[item]
	return quantity, ok
}
