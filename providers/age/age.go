package provider

import (
	"github.com/Antoha2/sandbox/config"
)

type ageImpl struct {
	cfg *config.Config
}

func NewGetAge(cfg *config.Config) *ageImpl {
	return &ageImpl{
		cfg: cfg,
	}
}

type Age struct {
	Name  string `json:"name"`
	Age   string `json:"age"`
	Count int    `json:"count"`
}

func (s *ageImpl) GetParam(request string, cfg *config.Config) (string, error) {
	//log.Println(cfg.AddrAge + request)
	// client := &http.Client{}
	// req, err := http.NewRequest("GET", cfg.AddrAge+request, nil)
	// if err != nil {
	// 	log.Println("http.NewRequest() - ", err)
	// 	return "", err
	// }
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Println("client.Do() - ", err)
	// 	return "", err
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println("ioutil.ReadAll() -", err)
	// 	return "", err
	// }

	// restResponse := new(Age)
	// err = json.Unmarshal(body, restResponse)
	// if err != nil {
	// 	log.Println("json.Unmarshal() -", err)
	// 	return "", err
	// }
	// log.Println(restResponse)
	return "", nil
}
