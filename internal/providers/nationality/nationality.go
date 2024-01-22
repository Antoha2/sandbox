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
	URL string
}

func NewGetNat(URL string) *natImpl {
	return &natImpl{
		URL: URL,
	}
}

type response struct {
	Name    string               `json:"name"`
	Count   int                  `json:"count"`
	Country []NationalityCountry `json:"country"`
}

type NationalityCountry struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

func (s *natImpl) GetNationality(ctx context.Context, name string) (string, error) {

	restResponse := new(response)
	query := fmt.Sprintf("%s?name=%s", s.URL, name)
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

	//choosing one option with the highest probability
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
