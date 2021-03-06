package router

import (
	"github.com/girishkoundinya/SAC_Server/controller"
	"github.com/julienschmidt/httprouter"
)

func InitRouter() *httprouter.Router {
	router := httprouter.New()
	mapRoutes(router)
	return router
}

func mapRoutes(router *httprouter.Router) {
	router.GET("/", controller.Index)
	router.GET("/search", controller.Search)
	router.GET("/search_suggest", controller.SearchSuggestions)
	router.GET("/search_chrome_extension", controller.SearchChromeExtension)
	router.GET("/shop/:shopid", controller.ShopDetail)
	router.GET("/shop/:shopid/tags", controller.ShopTags)
	router.GET("/shop/:shopid/product/:productid", controller.ProductDetail)

	router.POST("/shop/:shopid", controller.AddTag)
	router.POST("/shop", controller.ShopCreate)
}
