package payload

type Response struct {
	Message string      `json:"message"`
	Status string 		`json:"status"`
	Data    interface{} `json:"data"`
}