package categorydto

type CategoryInfoResponse struct {
	ID          uint                   `json:"id"`
	Name        string                 `json:"name"`
	Slug        string                 `json:"slug"`
	Description string                 `json:"description"`
	ParentID    uint                   `json:"parentID"`
	Children    []CategoryInfoResponse `json:"children"`
}