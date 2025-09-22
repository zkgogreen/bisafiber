package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/zkgogreen/bisago/domain"
)

type customerRepository struct {
	db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: goqu.New("mysql", con),
	}
}

func (c *customerRepository) FindAll(ctx context.Context) (result []domain.Customer, err error) {
	dataset := c.db.From("customer").Select("id", "code", "name")
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (c *customerRepository) FindById(ctx context.Context, id int) (result *domain.Customer, err error) {
	result = &domain.Customer{}
	dataset := c.db.From("customer").Select("id", "code", "name").Where(goqu.C("deletedat").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, result)
	if err != nil {
		return nil, err
	}
	return
}

func (c *customerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	executor := c.db.Insert("customer").Rows(customer).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (c *customerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	executor := c.db.Update("customer").Set(customer).Where(goqu.C("id").Eq(customer.ID)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (c *customerRepository) Delete(ctx context.Context, customer *domain.Customer) error {
	executor := c.db.Update("customer").Where(goqu.C("id").Eq(customer.ID)).Set(goqu.Record{"deletedat": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
