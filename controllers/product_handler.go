package controllers

import (
    "github.com/gin-gonic/gin"
    "gogo/models"
    "net/http"
)

type ProductController struct {
    // 필요한 의존성이 있다면 여기에 추가
}

func NewProductController() *ProductController {
    return &ProductController{}
}

func (pc *ProductController) GetProducts(c *gin.Context) {
    // 모든 제품을 가져와서 응답을 생성하는 로직
    products := models.GetAllProducts() // 모델 메서드 호출

    c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
    // 제품 ID를 가져와서 해당 제품을 찾아 응답을 생성하는 로직
    productID := c.Param("id")
    product := models.GetProductByID(productID) // 모델 메서드 호출

    c.JSON(http.StatusOK, product)
}
