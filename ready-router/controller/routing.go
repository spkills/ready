package controller

import (
	"github.com/buaazp/fasthttprouter"
)

// AddRouteing add routing path from go generate ready
func AddRouting(router *fasthttprouter.Router) {

	router.GET("/test", TestGetHandler)
	router.POST("/test", TestPostHandler)
	router.PUT("/test", TestPutHandler)
	router.DELETE("/test", TestDeleteHandler)
	router.HEAD("/test", TestHeadHandler)

	router.GET("/test/list", TestListGetHandler)
	router.POST("/test/list", TestListPostHandler)
	router.PUT("/test/list", TestListPutHandler)
	router.DELETE("/test/list", TestListDeleteHandler)
	router.HEAD("/test/list", TestListHeadHandler)

}
