package main

type OrderItem struct {
	ProductID int
	Quantity  int
	Price     int
}

type Order struct {
	ID       int
	Items    []OrderItem
	Discount float64
	Total    float64
	Status   string
}

type DiscountRule struct {
	MinAmount       float64
	DiscountPercent float64
	Description     string
}

func CalculateSubtotal(order Order) float64 {
	sum := 0.0  
	for _, item := range order.Items {
		sum += float64(item.Quantity * item.Price)
	}
	return sum
}

func ApplyDiscountRules(order Order, rules []DiscountRule) Order {
	price := CalculateSubtotal(order)
	bestDiscount := 0.0
	for _, r := range rules {
		if price >= r.MinAmount {
			if bestDiscount < r.DiscountPercent {
				bestDiscount = r.DiscountPercent
			}
		}
	}

	order.Discount = bestDiscount
	order.Total = price - (price * bestDiscount / 100)
	return order
}

func ProcessOrders(orders []Order, rules []DiscountRule) []Order {
	for i := 0; i < len(orders); i++ {
		orders[i].Status = "Processing"
		orders[i] = ApplyDiscountRules(orders[i], rules)
		if orders[i].Total < 0 {
			orders[i].Status = "cancelled"
		} else {
			orders[i].Status = "completed"
		}
	}
	return orders
}

func FilterOrdersByStatus(orders []Order, status string) []Order {
	var newList []Order
	for _, v := range orders {
		if v.Status == status {
			newList = append(newList, v)
		}
	}
	return newList
}

func CalculateOrderStatistics(orders []Order) map[string]interface{} {
	totalOrders := len(orders)
	totalRevenue := 0.0
	completedOrders := 0
	totalDiscount := 0.0

	for _, v := range orders {
		totalRevenue += v.Total
		if v.Status == "completed" {
			completedOrders++
		}
		subtotal := CalculateSubtotal(v)
		totalDiscount += subtotal - v.Total
	}

	averageOrderValue := 0.0
	if totalOrders > 0 {
		averageOrderValue = totalRevenue / float64(totalOrders)
	}

	stats := make(map[string]interface{})
	stats["totalOrders"] = totalOrders
	stats["totalRevenue"] = totalRevenue
	stats["averageOrderValue"] = averageOrderValue
	stats["completedOrders"] = completedOrders
	stats["totalDiscount"] = totalDiscount

	return stats
}
