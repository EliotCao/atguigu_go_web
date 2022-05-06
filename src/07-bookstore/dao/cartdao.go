package dao

import (
	"atguigu_go_web/src/04/util"
	"atguigu_go_web/src/07-bookstore/model"
)

func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := util.Db.Exec(sqlStr, cart.CartID, cart.TotalCount, cart.TotalAmount, cart.UserID)
	if err != nil {
		return err
	}
	cartItems := cart.CartItem
	for _, v := range cartItems{
		AddCartItem(v)
	}
	return nil
}
