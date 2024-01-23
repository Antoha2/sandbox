package service

import (
	"context"

	"github.com/Antoha2/sandbox/internal/repository"
	"github.com/Antoha2/sandbox/pkg/logger/sl"
	"github.com/pkg/errors"
)

//get userS
func (s *servImpl) GetUsers(ctx context.Context, filter *QueryUsersFilter) ([]*User, error) {
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
		return nil, errors.Wrap(err, "runtime error GetUsers")
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
	return users, nil
}

//get user
func (s *servImpl) GetUser(ctx context.Context, id int) (*User, error) {
	repUser, err := s.rep.GetUser(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "runtime error GetUser")
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

//del user
func (s *servImpl) DeleteUser(ctx context.Context, id int) (*User, error) {
	repUser, err := s.rep.DeleteUser(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "runtime error DeleteUser")
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

//add user
func (s *servImpl) AddUser(ctx context.Context, user *User) (*User, error) {

	age, err := s.ageClient.GetAge(ctx, user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "occurate error for request provider get age")
	}
	s.log.Debug("by username got age", sl.Atr("age", age), sl.Atr("name", user.Name))

	gender, err := s.genderClient.GetGender(ctx, user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "occurate error for request provider get gender")
	}
	s.log.Debug("by username got gender", sl.Atr("gender", gender), sl.Atr("name", user.Name))

	nationality, err := s.nationalityClient.GetNationality(ctx, user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "occurate error for request provider get nationality")
	}
	s.log.Debug("by username got nationality", sl.Atr("nationality", nationality), sl.Atr("name", user.Name))

	repUser := &repository.RepUser{
		Name:        user.Name,
		SurName:     user.SurName,
		Patronymic:  user.Patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
	}

	repUser, err = s.rep.AddUser(ctx, repUser)
	if err != nil {
		return nil, errors.Wrap(err, "runtime error AddUser")
	}

	respUser := &User{
		Id:          repUser.Id,
		Name:        repUser.Name,
		SurName:     repUser.SurName,
		Patronymic:  repUser.Patronymic,
		Age:         repUser.Age,
		Gender:      repUser.Gender,
		Nationality: repUser.Nationality,
	}
	return respUser, nil
}

//update user
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
		return nil, errors.Wrap(err, "runtime error UpdateUser")
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
