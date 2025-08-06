package branddto

type CreateBrandRequest struct {
	Name        string
	Slug        *string
	Description *string
	Logo        *multipart.FileHeader
	Website     *string
	Country     *string
}

type EditBrandRequest struct {
	BrandID     uint
	Name        *string
	Slug        *string
	Description *string
	Logo        *multipart.FileHeader
	Website     *string
	Country     *string
}

type DeleteBrandRequest struct {
	BrandID   uint
}