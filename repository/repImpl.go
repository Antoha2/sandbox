package repository

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

//add user
func (r *Rep) AddUser(user RepUser) (int, error) {

	query := "INSERT INTO users (name, surname, patronymic, age, gender, nationality) values ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := r.DB.QueryRow(query, user.Name, user.SurName, user.Patronymic, user.Age, user.Gender, user.Nationality)
	if err := row.Scan(&user.Id); err != nil {
		return 0, err
	}
	log.Println("создана запись - ", user)
	return user.Id, nil
}

//del user
func (r *Rep) DelUser(id int) error {

	query := "DELETE FROM users WHERE id = $1"
	stmtDel, err := r.DB.Exec(query, id)
	if err == nil {
		count, err := stmtDel.RowsAffected()
		if count == 0 || err != nil {
			return errors.New("id not fined")
		}
	}
	log.Printf("удалена запись с id - %v\n", id)
	return nil
}

//get users
func (r *Rep) GetUsers(filter RepQueryFilter) ([]*RepUser, error) {

	users := make([]*RepUser, 0)
	buildQuery, args := buildQueryConstrain(filter)

	query := fmt.Sprintf("SELECT * FROM users%s", buildQuery)
	stmtGet, err := r.DB.Query(query, args...)
	if err != nil {
		panic(err)
	}

	for stmtGet.Next() {
		user := new(RepUser)
		err := stmtGet.Scan(&user.Id, &user.Name, &user.SurName, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
		fmt.Println("считана запись -", user)
	}
	return users, nil
}

//get user
func (r *Rep) GetUser(id int) (RepUser, error) {
	user := new(RepUser)
	query := "SELECT * FROM users WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	if err := row.Scan(&user.Id, &user.Name, &user.SurName, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality); err != nil {
		panic(err)
	}
	log.Printf("получена запись с id - %v\n", id)
	return *user, nil
}

//update user
func (r *Rep) UpdateUser(user *RepUser) (*RepUser, error) {
	query := "update users set name=$2, surname=$3, patronymic=$4, age=$5, gender=$6, nationality=$7 where id=$1"
	stmtUp, err := r.DB.Exec(query, user.Id, user.Name, user.SurName, user.Patronymic, user.Age, user.Gender, user.Nationality)
	if err == nil {
		count, err := stmtUp.RowsAffected()
		if count == 0 || err != nil {
			return user, err
		}
	}
	fmt.Println("изменена запись c id -", user.Id)
	return user, nil
}

//build query string
func buildQueryConstrain(filter RepQueryFilter) (string, []any) {
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
	query := strings.Join(constrain, " AND ")
	if query != "" {
		query = fmt.Sprintf(" WHERE %s ORDER BY id ASC LIMIT %d OFFSET %d", query, filter.Limit, filter.Offset)
	}
	return query, args
}
