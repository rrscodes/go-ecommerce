package database

import "errors"

var (
	ErrCantFindProduct = errors.New("Can't fint the product")
	ErrCantDecodeProducts = errors.New("Can't fint the product")
	ErrUserIdIsNotValid = errors.New("User not valid")
	ErrCantUpdateCart = errors.New("Can't add item to cart")
	ErrCantRemoveItemCart = errors.New("Can't remove item from cart")
	ErrCantGetItem = errors.New("Unable to get item from cart")
	ErrCantBuyCartItem = errors.New("Can't update the purchase")
)

func AddProductToCart(){

}

func RemoveCartItem(){

}

func BuyItemFromCart(){

}

func InstantBuyer(){

}