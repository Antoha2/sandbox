package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type natImpl struct {
	addr string
}

func NewGetNat(addr string) *natImpl {
	return &natImpl{
		addr: addr,
	}
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

func (s *natImpl) GetNationality(ctx context.Context, name string) (string, error) {

	restResponse := new(Nat)
	query := fmt.Sprintf("%s?name=%s", s.addr, name)
	resp, err := http.Get(query)
	if err != nil {
		return "", errors.Wrap(err, "cant get resp Nationality ")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "cant read Nationality ")
	}

	err = json.Unmarshal(body, restResponse)
	if err != nil {
		return "", errors.Wrap(err, "cant Unmarshal Nationality ")
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
