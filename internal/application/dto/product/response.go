package productdto

type ProductInfoResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description`
	Price       int    `json:"price"`
	Inventory   int    `json:"inventory"`
	CategoryID  uint   `json:"categoryID"`
	BrandID     uint   `json:"brandID"`
	Picture     string `json:"picture"`
}