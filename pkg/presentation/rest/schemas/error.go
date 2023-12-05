package schemas

var NotFoundError = Error{"not found"}

type ErrorInfo struct {
	Error string `json:"error"`
}

type Error struct {
	Error string `json:"error"`
}
