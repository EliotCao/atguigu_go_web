package dao

import (
	"atguigu_go_web/src/07-bookstore/model"
	util "atguigu_go_web/src/07-bookstore/utils"
)

func AddCartItem(item *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	_, err := util.Db.Exec(sqlStr, item.Count, item.Amount, item.Book.Id, item.CartID)
	if err != nil {
		return err
	}
	return nil
}
