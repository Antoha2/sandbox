package provider

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Antoha2/sandbox/service"
)

type natImpl struct {
}

func NewGetNat() *natImpl {
	return &natImpl{}
}

type Nat struct {
	Name    string       `json:"name"`
	Count   int          `json:"count"`
	Country []natCountry `json:"country"`
}

type natCountry struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

func (s *natImpl) GetNationality(r *service.Query) (string, error) {

	resp, err := http.Get(r.Addr + r.Name)
	if err != nil {
		log.Println("client.Do() - ", err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() -", err)
		return "", err
	}

	restResponse := new(Nat)
	err = json.Unmarshal(body, restResponse)
	if err != nil {
		log.Println("json.Unmarshal() -", err)
		return "", err
	}

	//выбор одного варианта с наибольшей вероятностью
	var probability float32
	var country string
	for _, v := range restResponse.Country {
		if v.Probability > probability {
			probability = v.Probability
			country = v.CountryId
		}
	}
	return country, nil
}
