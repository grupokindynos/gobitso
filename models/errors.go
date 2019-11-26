package models


type ErrorResponse struct {
	Success bool `json:"success"`
	Error   struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"error"`
}

type NoCredentials struct{}

func (nc *NoCredentials) Error() string{
	return "you cant use Bitso's private api without an api secret and key. go to '' o learn how to get them and use the SetAuth() method to configure your profile."
}