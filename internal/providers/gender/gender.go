package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type genderImpl struct {
	URL string
}

func NewGetGender(URL string) *genderImpl {
	return &genderImpl{
		URL: URL,
	}
}

type response struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Count       int     `json:"count"`
	Probability float32 `json:"probability"`
}

func (s *genderImpl) GetGender(ctx context.Context, name string) (string, error) {

	restResponse := new(response)
	query := fmt.Sprintf("%s?name=%s", s.URL, name)
	resp, err := http.Get(query)
	if err != nil {
		return "", errors.Wrap(err, "cant get resp Gender ")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "cant read Gender ")
	}

	err = json.Unmarshal(body, restResponse)
	if err != nil {
		return "", errors.Wrap(err, "cant unmarshall Gender ")
	}
	return restResponse.Gender, nil
}
