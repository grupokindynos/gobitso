package models

type NoCredentials struct{}

func (nc *NoCredentials) Error() string{
	return "you cant use Bitso's private api without an api secret and key. go to '' o learn how to get them and use the SetAuth() method to configure your profile."
}