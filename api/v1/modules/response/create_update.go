package response

type CreateUpdateModuleResponse struct {
	ID string `json:"id"`
}

func NewCreateUpdateModuleResponse(id string) CreateUpdateModuleResponse {
	return CreateUpdateModuleResponse{
		ID: id,
	}
}
