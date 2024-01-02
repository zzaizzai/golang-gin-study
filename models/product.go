// models/product.go

package models

type Product struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

// 임시 데이터 저장소로 사용하기 위한 간단한 데이터 슬라이스
var products []Product

func init() {
    // 초기 데이터 로드 또는 데이터베이스 연결 설정 등 초기화 로직 수행
    // 여기에서는 간단한 예제를 위해 몇 가지 제품을 추가합니다.
    products = append(products, Product{ID: "1", Name: "Product 1", Price: 19.99})
    products = append(products, Product{ID: "2", Name: "Product 2", Price: 29.99})
    products = append(products, Product{ID: "3", Name: "Product 3", Price: 39.99})
}

func GetAllProducts() []Product {
    // 모든 제품을 반환
    return products
}

func GetProductByID(id string) *Product {
    // 제품 ID로 제품을 찾아 반환, 존재하지 않을 경우 nil 반환
    for _, p := range products {
        if p.ID == id {
            return &p
        }
    }
    return nil
}