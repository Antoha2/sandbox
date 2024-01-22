CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    surname varchar(255) NOT NULL,
    patronymic varchar(255),
    age int NOT NULL,
    gender varchar(255) NOT NULL,
    nationality varchar(255) NOT NULL
)