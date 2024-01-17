package service

import (
	"log"

	"github.com/Antoha2/sandbox/repository"
)

func (s *servImpl) GetUsers(user User) error {
	return nil
}

func (s *servImpl) DelUser(id int) error {
	return nil
}

func (s *servImpl) AddUser(user User) error {
	var err error

	repUser := new(repository.RepUser)
	repUser.Name = user.Name
	repUser.SurName = user.SurName
	repUser.Patronymic = user.Patronymic

	req := &Query{
		Name: user.Name,
		Addr: s.cfg.AddrAge,
	}

	repUser.Age, err = s.ageClient.GetAge(req)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Addr = s.cfg.AddrGender
	repUser.Gender, err = s.genderClient.GetGender(req)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Addr = s.cfg.AddrNationality
	repUser.Nationality, err = s.nationalityClient.GetNationality(req)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!! -", repUser)

	return nil
}

func (s *servImpl) UpdateUser(user User) error {
	return nil
}
