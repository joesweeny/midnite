package app

type Event struct {
	Type   string `json:"type"`
	Amount string `json:"amount"`
	UserID int64  `json:"user_id"`
	T      int    `json:"t"`
}

type Report struct {
	Alert      bool  `json:"alert"`
	AlertCodes []int `json:"alert_codes"`
	UserID     int64 `json:"user_id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
