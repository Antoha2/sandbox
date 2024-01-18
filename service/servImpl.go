package service

import (
	"log"

	"github.com/Antoha2/sandbox/repository"
)

func (s *servImpl) GetUsers(filter GetQueryFilter) ([]*User, error) {

	readFilter := repository.RepQueryFilter{
		Id:          filter.Id,
		Name:        filter.Name,
		SurName:     filter.SurName,
		Patronymic:  filter.Patronymic,
		Age:         filter.Age,
		Gender:      filter.Gender,
		Nationality: filter.Nationality,
		Limit:       filter.Limit,
		Offset:      filter.Offset,
	}
	repUsers, err := s.rep.GetUsers(readFilter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	users := make([]*User, len(repUsers))
	for index, user := range repUsers {
		t := &User{
			Id:          user.Id,
			Name:        user.Name,
			SurName:     user.SurName,
			Patronymic:  user.Patronymic,
			Age:         user.Age,
			Gender:      user.Gender,
			Nationality: user.Nationality,
		}
		users[index] = t
		// log.Println("!!!!!!!", user)
		//log.Println(index, users[index])
	}
	log.Println(users)
	return users, nil
}

func (s *servImpl) GetUser(id int) (User, error) {
	//repUser := new(repository.RepUser)

	repUser, err := s.rep.GetUser(id)
	if err != nil {
		log.Println(err)
		//return err
	}
	user := &User{
		Id:          repUser.Id,
		Name:        repUser.Name,
		SurName:     repUser.SurName,
		Patronymic:  repUser.Patronymic,
		Age:         repUser.Age,
		Gender:      repUser.Gender,
		Nationality: repUser.Nationality,
	}
	log.Println("!!!!!!!!!!!!!!!!!!!!!!", user)
	return *user, nil
}

func (s *servImpl) DelUser(id int) error {

	err := s.rep.DelUser(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *servImpl) AddUser(user User) (User, error) {
	var err error
	var respUser User
	var reposUser repository.RepUser
	reposUser.Name = user.Name
	reposUser.SurName = user.SurName
	reposUser.Patronymic = user.Patronymic

	req := &Query{
		Name: user.Name,
		Addr: s.cfg.AddrAge,
	}

	reposUser.Age, err = s.ageClient.GetAge(req)
	if err != nil {
		log.Println(err)
		return respUser, err
	}
	req.Addr = s.cfg.AddrGender
	reposUser.Gender, err = s.genderClient.GetGender(req)
	if err != nil {
		log.Println(err)
		return respUser, err
	}
	req.Addr = s.cfg.AddrNationality
	reposUser.Nationality, err = s.nationalityClient.GetNationality(req)
	if err != nil {
		log.Println(err)
		return respUser, err
	}

	id, err := s.rep.AddUser(reposUser)
	if err != nil {
		log.Println(err)
		return respUser, err
	}

	respUser.Id = id
	respUser.Name = user.Name
	respUser.SurName = user.SurName
	respUser.Patronymic = user.Patronymic
	respUser.Age = reposUser.Age
	respUser.Gender = reposUser.Gender
	respUser.Nationality = reposUser.Nationality

	return respUser, nil
}

func (s *servImpl) UpdateUser(user User) (User, error) {

	reposUser := &repository.RepUser{
		Id:          user.Id,
		Name:        user.Name,
		SurName:     user.SurName,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}

	reposUser, err := s.rep.UpdateUser(reposUser)
	if err != nil {
		log.Println(err)
		return user, err
	}

	respUser := &User{
		Id:          reposUser.Id,
		Name:        reposUser.Name,
		SurName:     reposUser.SurName,
		Patronymic:  reposUser.Patronymic,
		Age:         reposUser.Age,
		Gender:      reposUser.Gender,
		Nationality: reposUser.Nationality,
	}
	return *respUser, nil
}
