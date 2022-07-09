package response

type UploadResp struct {
	URL string `json:"url"`
}

func NewUploadResp(url string) UploadResp {
	return UploadResp{
		URL: url,
	}
}
