package service

import (
	"context"
	"log"

	"github.com/Antoha2/sandbox/repository"
)

func (s *servImpl) GetUsers(ctx context.Context, filter *GetQueryFilter) ([]*User, error) {

	readFilter := &repository.RepQueryFilter{
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
	repUsers, err := s.rep.GetUsers(ctx, readFilter)
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
	}
	log.Println(users)
	return users, nil
}

func (s *servImpl) GetUser(ctx context.Context, id int) (*User, error) {

	repUser, err := s.rep.GetUser(ctx, id)
	if err != nil {
		log.Println(err)
		//return , err
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
	return user, nil
}

func (s *servImpl) DelUser(ctx context.Context, id int) error {

	err := s.rep.DelUser(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *servImpl) AddUser(ctx context.Context, user *User) (*User, error) {
	var err error

	reposUser := &repository.RepUser{
		Name:       user.Name,
		SurName:    user.SurName,
		Patronymic: user.Patronymic,
	}

	req := &Query{
		Name: user.Name,
		Addr: s.cfg.AddrAge,
	}

	reposUser.Age, err = s.ageClient.GetAge(req)
	if err != nil {
		log.Println(err)
		return user, err
	}
	req.Addr = s.cfg.AddrGender
	reposUser.Gender, err = s.genderClient.GetGender(req)
	if err != nil {
		log.Println(err)
		return user, err
	}
	req.Addr = s.cfg.AddrNationality
	reposUser.Nationality, err = s.nationalityClient.GetNationality(req)
	if err != nil {
		log.Println(err)
		return user, err // произошла ошибка, не надо юзера возвращать, nil надо
	}

	id, err := s.rep.AddUser(ctx, reposUser)
	if err != nil {
		log.Println(err)
		return user, err // произошла ошибка, не надо юзера возвращать, nil надо
	}

	respUser := &User{
		Id:          id,
		Name:        user.Name,
		SurName:     user.SurName,
		Patronymic:  user.Patronymic,
		Age:         reposUser.Age,
		Gender:      reposUser.Gender,
		Nationality: reposUser.Nationality,
	}
	return respUser, nil
}

func (s *servImpl) UpdateUser(ctx context.Context, user *User) (*User, error) {

	reposUser := &repository.RepUser{
		Id:          user.Id,
		Name:        user.Name,
		SurName:     user.SurName,
		Patronymic:  user.Patronymic,
		Age:         user.Age,
		Gender:      user.Gender,
		Nationality: user.Nationality,
	}

	reposUser, err := s.rep.UpdateUser(ctx, reposUser)
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
	return respUser, nil
}
