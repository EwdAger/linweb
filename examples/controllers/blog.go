package controllers

import (
	"fmt"
	"linweb/interfaces"
	"net/http"
)

type BlogController struct {
}

//[GET("/blog/:id")]
func (blog *BlogController) GetBlog(c interfaces.IContext) {
	fmt.Println(c.Param("id"))
	c.Response().String(http.StatusOK, "id=%s", c.Param("id"))
}