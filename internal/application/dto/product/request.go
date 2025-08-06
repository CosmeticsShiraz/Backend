package productdto

type CreateProductRequest struct {
	Name        string
	Description string
	Price       int
	Inventory   int
	CategoryID  uint
	BrandID     uint
	Pictures    []*multipart.FileHeader
}

type EditProductRequest struct {
	ProductID   uint
	Name        *string
	Description *string
	Price       int
	Inventory   int
	CategoryID  uint
	BrandID     uint
	Pictures    []*multipart.FileHeader
}

type DeleteProductRequest struct {
	ProductID   uint
}