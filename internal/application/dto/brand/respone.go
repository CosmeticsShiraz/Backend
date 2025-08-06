package branddto

type BrandInfoResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Website     string `json:"website"`
	Country     string `json:"country"`
}