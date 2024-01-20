package service

import (
	"context"

	"github.com/Antoha2/sandbox/internal/repository"
	"github.com/pkg/errors"
)

// get userS
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
		return nil, err // wrap
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
	// logs!!
	return users, nil
}

// get user
func (s *servImpl) GetUser(ctx context.Context, id int) (*User, error) {
	repUser, err := s.rep.GetUser(ctx, id)
	if err != nil {
		return nil, err // wrap
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
	// logs!!
	return user, nil
}

// del user
func (s *servImpl) DelUser(ctx context.Context, id int) (*User, error) {
	repUser, err := s.rep.DelUser(ctx, id)
	if err != nil {
		return nil, err
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
	// logs!!
	return user, nil
}

// add user
func (s *servImpl) AddUser(ctx context.Context, user *User) (*User, error) {
	var err error

	age, err := s.ageClient.GetAge(ctx, user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error for request provider get age")
	}
	s.log.Debug("by username got age", user.Name, age)

	gender, err := s.genderClient.GetGender(ctx, user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error for request provider get gender")
	}
	s.log.Debug("by username got gender", user.Name, gender)

	nationality, err := s.nationalityClient.GetNationality(ctx, user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error for request provider get nationality")
	}
	s.log.Debug("by username got nationality", user.Name, nationality)

	// плохая практика держать наполовину инициализированную структуру
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
		return nil, err
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

// update user
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
		return nil, err
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
