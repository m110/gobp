package controllers

type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
}
