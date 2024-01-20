package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// add user
func (r *Rep) AddUser(ctx context.Context, user *RepUser) (*RepUser, error) {

	repUser := RepUser{}
	query := "INSERT INTO users (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, surname, patronymic, age, gender, nationality"
	// везде переделать на методы с контекстом
	row := r.DB.QueryRowxContext(ctx, query, user.Name, user.SurName, user.Patronymic, user.Age, user.Gender, user.Nationality)
	if err := row.Scan(&repUser.Id, &repUser.Name, &repUser.SurName, &repUser.Patronymic, &repUser.Age, &repUser.Gender, &repUser.Nationality); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql insert User failed, query: %s", query))
	}
	return &repUser, nil
}

// del user
// полное название как и вдругих методах delete
func (r *Rep) DelUser(ctx context.Context, id int) (*RepUser, error) {
	repUser := RepUser{}
	query := "DELETE FROM users WHERE id = $1 RETURNING id, name, surname, patronymic, age, gender, nationality"
	row := r.DB.QueryRow(query, id)
	if err := row.Scan(&repUser.Id, &repUser.Name, &repUser.SurName, &repUser.Patronymic, &repUser.Age, &repUser.Gender, &repUser.Nationality); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql delete User failed, query: %s", query))
	}
	return &repUser, nil
}

// get userS
func (r *Rep) GetUsers(ctx context.Context, filter *RepQueryFilter) ([]*RepUser, error) {
	users := make([]*RepUser, 0)

	queryConstrain, args := buildQueryConstrain(filter)
	// офсет и лимит должны быть всегда
	// в селекне не должно быть звездочки
	query := fmt.Sprintf("SELECT * FROM users %s LIMIT %d OFFSET %d", queryConstrain, filter.Limit, filter.Offset)
	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql select Users failed, query: %s", query))
	}

	for rows.Next() {
		user := RepUser{}
		err := rows.Scan(&user.Id, &user.Name, &user.SurName, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality)
		if err != nil {
			return nil, errors.Wrap(err, "sql GetUsers scan failed users")
		}
		users = append(users, &user)
	}

	return users, nil
}

// get user
func (r *Rep) GetUser(ctx context.Context, id int) (*RepUser, error) {
	user := RepUser{}
	// в селекне не должно быть звездочки
	query := "SELECT * FROM users WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	if err := row.Scan(&user.Id, &user.Name, &user.SurName, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality); err != nil {
		return nil, errors.Wrap(err, "sql GetUsers scan failed users")
	}

	return &user, nil
}

// update user
func (r *Rep) UpdateUser(ctx context.Context, user *RepUser) (*RepUser, error) {
	respUser := RepUser{}
	// нумеруй аргументы по порядку
	query := "UPDATE users SET name=$2, surname=$3, patronymic=$4, age=$5, gender=$6, nationality=$7 WHERE id=$1 RETURNING id, name, surname, patronymic, age, gender, nationality"

	row := r.DB.QueryRow(query, user.Id, user.Name, user.SurName, user.Patronymic, user.Age, user.Gender, user.Nationality)
	if err := row.Scan(&respUser.Id, &respUser.Name, &respUser.SurName, &respUser.Patronymic, &respUser.Age, &respUser.Gender, &respUser.Nationality); err != nil {
		// ошибки переделать по примеру выше
		return nil, errors.Wrap(err, "sql UpdateUser scan failed users")
	}
	return &respUser, nil
}

// build query string
func buildQueryConstrain(filter *RepQueryFilter) (string, []any) {
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
		queryConstrain = fmt.Sprintf("WHERE %s ORDER BY id ASC", queryConstrain)
	}
	return queryConstrain, args
}
