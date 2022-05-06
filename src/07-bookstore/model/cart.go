package model

type Cart struct {
	CartID      string
	CartItem    []*CartItem
	TotalCount  int64
	TotalAmount float64
	UserID      int
}

func (c Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range c.CartItem{
		totalCount = totalCount + v.Count
	}
	return totalCount
}

func (c Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range c.CartItem{
		totalAmount = totalAmount + v.Amount
	}
	return totalAmount
}