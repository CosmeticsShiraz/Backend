package productdto

type CreateProductRequest struct {
	Name         string
	*Description string
	Price        int
	Inventory    int
	*CategoryID  uint
	*BrandID     uint
}

type EditProductRequest struct {
	ProductID   uint
	Name        *string
	Description *string
	Price       *int
	Inventory   *int
	CategoryID  *uint
	BrandID     *uint
}

type DeleteProductRequest struct {
	ProductID uint
}

type AddPictureRequest struct {
	ProductID uint
	Picture   *multipart.FileHeader
}

type DeletePictureRequest struct {
	ProductID uint
	Picture   string
}