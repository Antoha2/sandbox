package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

//add user
func (r *Rep) AddUser(ctx context.Context, user *RepUser) (*RepUser, error) {

	respUser := RepUser{}
	query := "INSERT INTO users (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, surname, patronymic, age, gender, nationality"
	row := r.DB.QueryRow(query, user.Name, user.SurName, user.Patronymic, user.Age, user.Gender, user.Nationality)
	if err := row.Scan(&respUser.Id, &respUser.Name, &respUser.SurName, &respUser.Patronymic, &respUser.Age, &respUser.Gender, &respUser.Nationality); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql AddUser request failed %s", query))
	}
	return &respUser, nil
}

//delete user
func (r *Rep) DeleteUser(ctx context.Context, id int) (*RepUser, error) {
	respUser := RepUser{}
	query := "DELETE FROM users WHERE id = $1 RETURNING id, name, surname, patronymic, age, gender, nationality"
	row := r.DB.QueryRow(query, id)
	if err := row.Scan(&respUser.Id, &respUser.Name, &respUser.SurName, &respUser.Patronymic, &respUser.Age, &respUser.Gender, &respUser.Nationality); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql DelUser request failed %s", query))
	}
	return &respUser, nil
}

//get userS
func (r *Rep) GetUsers(ctx context.Context, filter *RepQueryFilter) ([]*RepUser, error) {

	users := make([]*RepUser, 0)
	buildQuery, args := buildQueryConstrain(filter)

	query := fmt.Sprintf("SELECT id, name, surname, patronymic, age, gender, nationality FROM users%s", buildQuery)
	stmtGet, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql GetUsers query failed %s", query))
	}

	for stmtGet.Next() {
		user := new(RepUser)
		err := stmtGet.Scan(&user.Id, &user.Name, &user.SurName, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality)
		if err != nil {
			return nil, errors.Wrap(err, "sql GetUsers scan failed users")
		}
		users = append(users, user)
	}
	return users, nil
}

//get user
func (r *Rep) GetUser(ctx context.Context, id int) (*RepUser, error) {
	user := RepUser{}
	query := "SELECT id, name, surname, patronymic, age, gender, nationality FROM users WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	if err := row.Scan(&user.Id, &user.Name, &user.SurName, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality); err != nil {
		return nil, errors.Wrap(err, "sql GetUsers scan failed users")
	}
	return &user, nil
}

//update user
func (r *Rep) UpdateUser(ctx context.Context, user *RepUser) (*RepUser, error) {
	respUser := RepUser{}
	query := "UPDATE users SET name=$1, surname=$2, patronymic=$3, age=$4, gender=$5, nationality=$6 WHERE id=$7 RETURNING id, name, surname, patronymic, age, gender, nationality"

	row := r.DB.QueryRow(query, user.Name, user.SurName, user.Patronymic, user.Age, user.Gender, user.Nationality, user.Id)
	if err := row.Scan(&respUser.Id, &respUser.Name, &respUser.SurName, &respUser.Patronymic, &respUser.Age, &respUser.Gender, &respUser.Nationality); err != nil {
		return nil, errors.Wrap(err, "sql UpdateUser scan failed users")
	}
	return &respUser, nil
}

//build query string
func buildQueryConstrain(filter *RepQueryFilter) (string, []any) {
	i := 1
	constrain := make([]string, 0, 6)
	args := make([]any, 0, 6)
	if filter.Name != "" {
		s := fmt.Sprintf("name=$%d", i)
		constrain = append(constrain, s)
		args = append(args, filter.Name)
		i++
	}
	if filter.SurName != "" {
		s := fmt.Sprintf("surname=$%d", i)
		constrain = append(constrain, s)
		args = append(args, filter.SurName)
		i++
	}
	if filter.Patronymic != "" {
		s := fmt.Sprintf("patronymic=$%d", i)
		constrain = append(constrain, s)
		args = append(args, filter.Patronymic)
		i++
	}
	if filter.Age != 0 {
		s := fmt.Sprintf("age=$%d", i)
		constrain = append(constrain, s)
		args = append(args, filter.Age)
		i++
	}
	if filter.Gender != "" {
		s := fmt.Sprintf("gender=$%d", i)
		constrain = append(constrain, s)
		args = append(args, filter.Gender)
		i++
	}
	if filter.Nationality != "" {
		s := fmt.Sprintf("nationality=$%d", i)
		constrain = append(constrain, s)
		args = append(args, filter.Nationality)
		i++
	}

	query := fmt.Sprintf(" WHERE %s ORDER BY id ASC LIMIT %d OFFSET %d", strings.Join(constrain, " AND "), filter.Limit, filter.Offset)

	return query, args
}
