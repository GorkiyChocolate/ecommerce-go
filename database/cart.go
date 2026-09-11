package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Can`t find the product")
	ErrCantDecodeProducts = errors.New("Can`t decode the product")
	ErrUserIdIsNotValid   = errors.New("Users id is not valid")
	ErrCantUpdateUser     = errors.New("Cant update user")
	ErrCantRemoveItem     = errors.New("Cant remove item")
	ErrCantGetItem        = errors.New("Cant get the item")
	ErrCantBuyCartItem    = errors.New("cant buy cart item")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
