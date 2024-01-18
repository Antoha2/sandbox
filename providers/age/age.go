package provider

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Antoha2/sandbox/service"
)

type ageImpl struct {
	//ageClient service.AgeProvider
	// сюда нужно передавать юрл при инциализации
}

func NewGetAge() *ageImpl {
	return &ageImpl{}
}

type Age struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

func (s *ageImpl) GetAge(r *service.Query) (int, error) {

	resp, err := http.Get(r.Addr + r.Name)
	if err != nil {
		log.Println("http.Get() - ", err)
		return 0, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() -", err) // ошибки плохо сделан, не надо код вставлять
		return 0, err
	}

	restResponse := new(Age)
	err = json.Unmarshal(body, restResponse)
	if err != nil {
		log.Println("json.Unmarshal() -", err)
		return 0, err
	}
	return restResponse.Age, nil
}
