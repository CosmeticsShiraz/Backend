package categorydto

type CreateCategoryRequest struct {
	Name        string
	Slug        *string
	Description *string
	ParentID    *uint
}

type EditCategoryRequest struct {
	CategoryID  uint
	Name        *string
	Slug        *string
	Description *string
	ParentID    *uint
}

type DeleteProductRequest struct {
	CategoryID   uint
}