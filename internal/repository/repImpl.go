package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

//add user
func (r *Rep) AddUser(ctx context.Context, user *RepUser) (*RepUser, error) {

	repUser := RepUser{}

	query := "INSERT INTO users (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, surname, patronymic, age, gender, nationality"
	row := r.DB.QueryRowContext(ctx, query, user.Name, user.SurName, user.Patronymic, user.Age, user.Gender, user.Nationality)
	if err := row.Scan(&repUser.Id, &repUser.Name, &repUser.SurName, &repUser.Patronymic, &repUser.Age, &repUser.Gender, &repUser.Nationality); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql insert User failed, query: %s", query))
	}

	return &repUser, nil
}

//delete user
func (r *Rep) DeleteUser(ctx context.Context, id int) (*RepUser, error) {
	repUser := RepUser{}

	query := "DELETE FROM users WHERE id = $1 RETURNING id, name, surname, patronymic, age, gender, nationality"
	row := r.DB.QueryRowContext(ctx, query, id)
	if err := row.Scan(&repUser.Id, &repUser.Name, &repUser.SurName, &repUser.Patronymic, &repUser.Age, &repUser.Gender, &repUser.Nationality); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql delete User failed, query: %s", query))
	}
	return &repUser, nil
}

//get userS
func (r *Rep) GetUsers(ctx context.Context, filter *RepQueryFilter) ([]*RepUser, error) {

	users := make([]*RepUser, 0)
	queryConstrain, args := buildQueryConstrain(filter)

	query := fmt.Sprintf("SELECT id, name, surname, patronymic, age, gender, nationality FROM users%s LIMIT %d OFFSET %d", queryConstrain, filter.Limit, filter.Offset)

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql select Users failed, query: %s", query))
	}

	for rows.Next() {
		user := RepUser{}
		err := rows.Scan(&user.Id, &user.Name, &user.SurName, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality)
		if err != nil {
			return nil, errors.Wrap(err, "sql scan Users failed")
		}
		users = append(users, &user)

	}

	return users, nil
}

//get user
func (r *Rep) GetUser(ctx context.Context, id int) (*RepUser, error) {

	user := RepUser{}

	query := "SELECT id, name, surname, patronymic, age, gender, nationality FROM users WHERE id = $1"
	row := r.DB.QueryRowContext(ctx, query, id)
	if err := row.Scan(&user.Id, &user.Name, &user.SurName, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql select User failed, query: %s", query))
	}
	return &user, nil
}

//update user
func (r *Rep) UpdateUser(ctx context.Context, user *RepUser) (*RepUser, error) {
	repUser := RepUser{}
	query := "UPDATE users SET name=$1, surname=$2, patronymic=$3, age=$4, gender=$5, nationality=$6 WHERE id=$7 RETURNING id, name, surname, patronymic, age, gender, nationality"

	row := r.DB.QueryRowContext(ctx, query, user.Name, user.SurName, user.Patronymic, user.Age, user.Gender, user.Nationality, user.Id)
	if err := row.Scan(&repUser.Id, &repUser.Name, &repUser.SurName, &repUser.Patronymic, &repUser.Age, &repUser.Gender, &repUser.Nationality); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql update User failed, query: %s", query))
	}
	return &repUser, nil
}

//build query string
func buildQueryConstrain(filter *RepQueryFilter) (string, []any) {
	log.Println("2!!!!!!!!!!!!!!!!!!!!!!!! ", filter)

	i := 1
	constrains := make([]string, 0, 6)
	args := make([]any, 0, 6)
	if filter.Name != "" {
		s := fmt.Sprintf("name=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Name)
	}
	if filter.SurName != "" {
		s := fmt.Sprintf("surname=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.SurName)
	}
	if filter.Patronymic != "" {
		s := fmt.Sprintf("patronymic=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Patronymic)
	}
	if filter.Age != 0 {
		s := fmt.Sprintf("age=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Age)
	}
	if filter.Gender != "" {

		s := fmt.Sprintf("gender=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Gender)
	}
	if filter.Nationality != "" {
		s := fmt.Sprintf("nationality=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Nationality)
	}

	queryConstrain := strings.Join(constrains, " AND ")
	if queryConstrain != "" {

		queryConstrain = fmt.Sprintf(" WHERE %s ORDER BY id ASC", queryConstrain)
	}
	return queryConstrain, args
}
