package models

type BizCommonResponse struct {
	// 业务返回码 1000 为成功，其它为失败
	StatusCode string `json:"statusCode"`
	// 返回信息
	Msg string `json:"msg,omitempty"`
}

func (s *BizCommonResponse) SetStatusCode(v string) *BizCommonResponse {
	s.StatusCode = v
	return s
}

func (s *BizCommonResponse) SetMsg(v string) *BizCommonResponse {
	s.Msg = v
	return s
}
