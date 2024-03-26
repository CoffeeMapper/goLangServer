package controller

import (
	"context"
	"database/sql"
	"log"
	"strconv"
)

type CoffeePoint struct {
	CoffeePointId int    `json:"coffee_point_id"`
	Logo          string `json:"logo"`
	Description   string `json:"description"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	BrandId       int    `json:"brand_id"`
	WorkTime      string `json:"work_time"`
	OrderId       int    `json:"order_id"`
	AvgRating     int    `json:"avg_rating"`
	FeedbackId    int    `json:"feedback_id"`
}

func GetAllCoffeePoint(ctx context.Context, db *sql.DB) ([]CoffeePoint, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM coffee_point")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coffeePoints []CoffeePoint
	for rows.Next() {
		var coffeePoint CoffeePoint

		err := rows.Scan(&coffeePoint.CoffeePointId, &coffeePoint.Logo, &coffeePoint.Description, &coffeePoint.Address, &coffeePoint.Email,
			&coffeePoint.BrandId, &coffeePoint.WorkTime, &coffeePoint.OrderId, &coffeePoint.AvgRating, &coffeePoint.FeedbackId)
		if err != nil {
			return nil, err
		}
		coffeePoints = append(coffeePoints, coffeePoint)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return coffeePoints, nil
}

func GetCoffeePointById(ctx context.Context, db *sql.DB, coffeePointID string) (*CoffeePoint, error) {
	row := db.QueryRowContext(ctx, "SELECT * FROM coffee_point where coffee_point_id = $1", coffeePointID)

	var coffeePoint CoffeePoint
	err := row.Scan(&coffeePoint.CoffeePointId, &coffeePoint.Logo, &coffeePoint.Description, &coffeePoint.Address, &coffeePoint.Email,
		&coffeePoint.BrandId, &coffeePoint.WorkTime, &coffeePoint.OrderId, &coffeePoint.AvgRating, &coffeePoint.FeedbackId)
	if err != nil {
		return &CoffeePoint{}, err
	}
	return &coffeePoint, nil
}

func CreateCoffeePoint(db *sql.DB, cp *CoffeePoint) error {
	//var id string

	stmt, err := db.Prepare("INSERT INTO coffee_point(logo,description,address,email,brand_id,work_time,order_id,avg_rating,feedback_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING coffee_point_id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	//Вставка нового пользователя и получение его айдишника
	//здесь Scan, с помощью которого можно считать все полученные данные в переменные
	//,после запроса возвращаем в id значение
	err = stmt.QueryRow(cp.Logo, cp.Description, cp.Address, cp.Email, cp.BrandId, cp.WorkTime, cp.OrderId, cp.AvgRating, cp.FeedbackId).Scan(&cp.CoffeePointId)

	log.Println("Запрос прошёл")

	if err != nil {
		return err
	}

	return nil
}

// user_id из запроса
func UpdateCoffeePoint(db *sql.DB, cp *CoffeePoint, coffeePointID string) (*CoffeePoint, error) {
	res, err := db.Exec("UPDATE coffee_point SET logo = $1, description = $2, address = $3, email=$4, brand_id = $5, work_time = $6, order_id = $7, avg_rating=$8, feedback_id=$9  where coffee_point_id = $10", cp.Logo, cp.Description, cp.Address, cp.Email, cp.BrandId, cp.WorkTime, cp.OrderId, cp.AvgRating, cp.FeedbackId, coffeePointID)
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

	cp.CoffeePointId, _ = strconv.Atoi(coffeePointID)

	return cp, nil
}

func DeleteCoffeePoint(db *sql.DB, coffeePointID string) error {
	_, err := db.Exec("DELETE FROM coffee_point where coffee_point_id = $1", coffeePointID)
	if err != nil {
		return err
	}
	return nil
}
