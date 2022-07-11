package request

import "github.com/Learning-Management-System-Kelompok-42/BE-LMS/business/requestFeat/spec"

type RequestFeatReq struct {
	UserID      string
	CompanyID   string `json:"company_id"`
	RequestType string `json:"request_type"`
	Title       string `json:"title"`
	Reason      string `json:"reason"`
	RequestDate string `json:"request_date"`
	RequestTime string `json:"request_time"`
}

func (req *RequestFeatReq) ToSpec() *spec.UpsertRequestFeat {
	return &spec.UpsertRequestFeat{
		UserID:      req.UserID,
		CompanyID:   req.CompanyID,
		RequestType: req.RequestType,
		Title:       req.Title,
		Reason:      req.Reason,
		RequestDate: req.RequestDate,
		RequestTime: req.RequestTime,
	}
}
