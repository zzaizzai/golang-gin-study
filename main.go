
package main

import (
    "github.com/gin-gonic/gin"
    "gogo/controllers" // ProductController가 있는 패키지를 import
    // "gogo/models"      // 모델 패키지 import
    // "net/http"
)

func main() {
    r := gin.Default()

    // 제품 관련 핸들러를 생성하고 라우팅에 등록
    productController := controllers.NewProductController()
    r.GET("/products", productController.GetProducts)
    r.GET("/products/:id", productController.GetProductByID)

    // 다른 핸들러 및 라우팅 등록

    r.Run(":8080")
}