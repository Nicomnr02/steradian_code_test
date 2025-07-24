package orderrepository

import (
	"context"
	"steradian_code_test/domain"
	"steradian_code_test/helper/date"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(c context.Context, data domain.Order) error
	Update(c context.Context, data domain.Order) error
	Delete(c context.Context, ID int) error
	GetAll(c context.Context) ([]domain.Order, error)
	GetByIDs(c context.Context, IDs []int) ([]domain.Order, error)
	GetByCarID(c context.Context, ID int) ([]domain.Order, error)
}
type RepositoryImpl struct {
	DB *pgxpool.Pool
}

func New(DB *pgxpool.Pool) Repository {
	return &RepositoryImpl{
		DB: DB,
	}
}

func (r *RepositoryImpl) Create(c context.Context, data domain.Order) error {
	sql := `insert into orders (
			car_id,
			order_date,
			pickup_date,
			drop_off_date,
			pick_up_location,
			drop_off_location
		) values ($1,$2,$3,$4,$5,$6)`

	_, err := r.DB.Exec(c, sql,
		data.CarID,
		time.Now().Format(date.ShortDateLayout),
		data.PickupDate,
		data.DropOffDate,
		data.PickupLocation,
		data.DropOffLocation,
	)

	return err
}
func (r *RepositoryImpl) Update(c context.Context, data domain.Order) error {
	sql := `UPDATE orders
	SET
	  car_id = $1,
	  pickup_date = $2,
	  drop_off_date = $3,
	  pick_up_location = $4,
	  drop_off_location = $5
	WHERE id = $6
	`

	_, err := r.DB.Exec(c, sql,
		data.CarID,
		data.PickupDate,
		data.DropOffDate,
		data.PickupLocation,
		data.DropOffLocation,
		data.ID,
	)
	return err
}

func (r *RepositoryImpl) Delete(c context.Context, ID int) error {
	sql := `DELETE FROM orders WHERE id = $1`
	_, err := r.DB.Exec(c, sql, ID)
	return err
}

func (r *RepositoryImpl) GetAll(c context.Context) ([]domain.Order, error) {
	sql := `
	SELECT
		id,
		car_id,
		order_date,
		pickup_date,
		drop_off_date,
		pick_up_location,
		drop_off_location
	FROM orders
	`

	rows, err := r.DB.Query(c, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Order])
}

func (r *RepositoryImpl) GetByIDs(c context.Context, IDs []int) ([]domain.Order, error) {
	sql := `
	SELECT
		id,
		car_id,
		order_date,
		pickup_date,
		drop_off_date,
		pick_up_location,
		drop_off_location
	FROM orders
	WHERE id = any($1)
	`

	rows, err := r.DB.Query(c, sql, IDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Order])
}

func (r *RepositoryImpl) GetByCarID(c context.Context, ID int) ([]domain.Order, error) {
	sql := `
	SELECT
		id,
		car_id,
		order_date,
		pickup_date,
		drop_off_date,
		pick_up_location,
		drop_off_location
	FROM orders
	WHERE car_id = $1
	`

	rows, err := r.DB.Query(c, sql, ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Order])
}

