package bitso

type Bitso struct {
	ApiKey		string 	`json:"api_key"`
	ApiSecret 	string 	`json:"api_secret"`
	Url 		string 	`json:"bitso_url"`
}

func NewBitso() {

}

func SignRequest() (string, error){
	return "", nil
}