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
	repUser := new(repository.RepUser)
	repUser.Name = user.Name
	repUser.SurName = user.SurName
	repUser.Patronymic = user.Patronymic

	log.Println(user.Name)

	reqAge := &Query{
		Name: user.Name,
		Addr: "https://api.agify.io/?name=",
	}
	log.Println(reqAge)
	age, err := s.getAge.GetParam(reqAge)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(age)

	// repUser.Age, err = strconv.Atoi(age)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }

	// repUser.Gender, err = s.getAge.GetParam(user.Name)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// repUser.Nationality, err = s.getAge.GetParam(user.Name)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	log.Println(repUser)
	return nil
}

func (s *servImpl) UpdateUser(user User) error {
	return nil
}
