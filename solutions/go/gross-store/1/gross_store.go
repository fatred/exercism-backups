package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := map[string]int{}
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	NewBill := make(map[string]int)
    return NewBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	myunit, exists := units[unit]
    if exists == false {
        return exists
    } else {
    	bill[item] += myunit
        return exists
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    _, exists := bill[item]
        if exists == false {
            return exists
        } else {
        	myunit, exists := units[unit]
            if exists == false {
                return exists
            } else {
            	if units[unit] > bill[item] {
                    return false
                } else if units[unit] == bill[item] {
                	delete(bill, item)
                    return true
                } else {
                	bill[item] = bill[item] - myunit
                    return true
                }
            }
        
        }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, exists := bill[item]
    if exists == false {
        return 0, exists
    } else {
		return bill[item], exists
    }
}
