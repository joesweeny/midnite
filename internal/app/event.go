package app

type Event struct {
	Type   string `json:"type"`
	Amount string `json:"amount"`
	UserID int    `json:"user_id"`
	T      int    `json:"t"`
}

type Report struct {
	Alert      bool  `json:"alert"`
	AlertCodes []int `json:"alert_codes"`
	UserID     int   `json:"user_id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type EventRepository interface {
	// Insert receives a pointer to an Event struct and stores within the underlying data store.
	Insert(e *Event) error
	// ByUserID receives a user id and returns a slice of Event struct associated to a user.
	ByUserID(userID int) ([]*Event, error)
}
