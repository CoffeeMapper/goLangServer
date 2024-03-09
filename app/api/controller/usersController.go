package controller

import (
	"context"
	"database/sql"
	"log"
	"strconv"
)

//Создать CRUD для таблицы с пользователей
//GET /api/v1/users - получает всех пользователей
//GET /api/v1/users/:id - получает пользователя по id
//POST /api/v1/users - создает пользователя
//PUT /api/v1/users/:id - редактирует данные пользователя по id
//DELETE /api/v1/users/:id - удаляет пользователя по id

type User struct {
	UserId   int    `json:"user_id" db:"user_id"`
	Login    string `json:"login" db:"login"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	SubId    int    `json:"sub_id" db:"sub_id"`
}

func GetAllUsers(ctx context.Context, db *sql.DB) ([]User, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User

		err := rows.Scan(&user.UserId, &user.Login, &user.Email, &user.Password, &user.SubId)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil

}

func GetUserById(ctx context.Context, db *sql.DB, userID string) (*User, error) {
	row := db.QueryRowContext(ctx, "SELECT * FROM users where user_id = $1", userID)

	var user User
	err := row.Scan(&user.UserId, &user.Login, &user.Email, &user.Password, &user.SubId)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func CreateUser(db *sql.DB, u *User) error {
	//var id string

	stmt, err := db.Prepare("INSERT INTO users(login,email,password,sub_id) VALUES ($1,$2,$3,$4) RETURNING user_id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	//Вставка нового пользователя и получение его айдишника
	//здесь Scan, с помощью которого можно считать все полученные данные в переменные
	//,после запроса возвращаем в id значение
	err = stmt.QueryRow(u.Login, u.Email, u.Password, u.SubId).Scan(&u.UserId)

	log.Println("Запрос прошёл")

	if err != nil {
		return err
	}

	return nil
}

// user_id из запроса
func UpdateUser(db *sql.DB, u *User, user_id string) (*User, error) {
	res, err := db.Exec("UPDATE users SET login = $1, email = $2, password = $3, sub_id=$4 where user_id = $5", u.Login, u.Email, u.Password, u.SubId, user_id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, err
	}

	log.Println("Запрос прошёл")

	u.UserId, _ = strconv.Atoi(user_id)

	return u, nil
}

func DeleteUser(db *sql.DB, userID string) error {
	_, err := db.Exec("DELETE FROM users where user_id = $1", userID)
	if err != nil {
		return err
	}
	return nil
}
