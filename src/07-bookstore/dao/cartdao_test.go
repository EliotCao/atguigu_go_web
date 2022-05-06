package dao

import (
	"atguigu_go_web/src/07-bookstore/model"
	"testing"
)

func TestCart(t *testing.T) {
	t.Run("AddCart", TestAddCart)
}
type CartItem struct {
	CartItemID int64
	Book       *model.Book
	Count      int64
	Amount     float64
	CartID     string
}

func TestAddCart(t *testing.T) {
	cart := model.Cart{
		CartID: "111",
		CartItem: []*model.CartItem{
			{
				CartItemID: 1,
				Book:&model.Book{
					Id: 1,
				},
				Count: 1,
				Amount: 12.32,
				CartID: "xxxx",
			},
		},
		TotalCount: 1,
		TotalAmount: 2,
		UserID: 1,
	}
	AddCart(&cart)
}
