package response

type GenereateLinkSpecializationResponse struct {
	Link string `json:"link"`
}

func NewGenerateLinkSpecializationResponse(link string) GenereateLinkSpecializationResponse {
	return GenereateLinkSpecializationResponse{
		Link: link,
	}
}
