package response

type CreateUpdateCompanyResp struct {
	ID string `json:"id"`
}

func NewCreateNewCompanyResponse(id string) CreateUpdateCompanyResp {
	return CreateUpdateCompanyResp{
		ID: id,
	}
}

func NewUpdateCompanyProfileResponse(id string) CreateUpdateCompanyResp {
	return CreateUpdateCompanyResp{
		ID: id,
	}
}
