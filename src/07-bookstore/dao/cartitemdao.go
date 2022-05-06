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

func GetCartItemByBookID(booID string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ?"
	row := util.Db.QueryRow(sqlStr, booID)
	cartItem := &model.CartItem{}
	err := row.Scan(cartItem.CartItemID, cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil,err
	}
	return cartItem, nil
}

func GetCartItemByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where cart_id = ?"
	rows, err := util.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil,err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		cartItem := &model.CartItem{}
		err := rows.Scan(cartItem.CartItemID, cartItem.Count, &cartItem.Amount, &cartItem.CartID)
		if err != nil {
			return nil,err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

