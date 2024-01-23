package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type nationalityImpl struct {
	URL string
}

func NewGetNat(URL string) *nationalityImpl {
	return &nationalityImpl{
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

func (s *nationalityImpl) GetNationality(ctx context.Context, name string) (string, error) {

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

	res := response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", errors.Wrap(err, "cant Unmarshal Nationality ")
	}

	//choosing one option with the highest probability
	var probability float32
	var country string
	for _, v := range res.Country {
		if v.Probability > probability {
			probability = v.Probability
			country = v.CountryId
		}
	}
	return country, nil
}
