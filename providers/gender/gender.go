package provider

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Antoha2/sandbox/service"
)

type genderImpl struct {
	//ageClient service.AgeProvider
}

func NewGetGender() *genderImpl {
	return &genderImpl{}
}

type Gender struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Count       int     `json:"count"`
	Probability float32 `json:"probability"`
}

func (s *genderImpl) GetGender(r *service.Query) (string, error) {

	resp, err := http.Get(r.Addr + r.Name)
	if err != nil {
		log.Println("http.Get() - ", err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() -", err)
		return "", err
	}

	restResponse := new(Gender)
	err = json.Unmarshal(body, restResponse)
	if err != nil {
		log.Println("json.Unmarshal() -", err)
		return "", err
	}
	return restResponse.Gender, nil
}
