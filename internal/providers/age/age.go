package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type ageImpl struct {
	URL string
}

func NewGetAge(URL string) *ageImpl {
	return &ageImpl{
		URL: URL,
	}
}

type response struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

func (s *ageImpl) GetAge(ctx context.Context, name string) (int, error) {

	query := fmt.Sprintf("%s?name=%s", s.URL, name)
	resp, err := http.Get(query)
	if err != nil {
		return 0, errors.Wrap(err, "cant get resp Age ")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.Wrap(err, "cant read Age ")
	}

	defer resp.Body.Close()

	res := response{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return 0, errors.Wrap(err, "cant unmarshal Age ")
	}
	return res.Age, nil
}
