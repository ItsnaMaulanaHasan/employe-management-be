package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(ctx)

	employeeIDs := seedEmployees(ctx, tx)
	seedAnnualReviews(ctx, tx, employeeIDs)

	if err := tx.Commit(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Seeder finished successfully ✅")
}

func seedEmployees(ctx context.Context, tx pgx.Tx) []int64 {
	query := `
		INSERT INTO employees (
			first_name,
			last_name,
			hire_date,
			termination_date,
			salary
		)
		VALUES
			('John', 'Smith', '2009-01-10', NULL, 30000),
			('Alice', 'Smith', '2010-03-15', NULL, 32000),
			('Bob', 'Brown', '2011-07-01', NULL, 28000),
			('Charlie', 'Doe', '2012-05-20', NULL, 35000),
			('Eve', 'Johnson', '2008-11-30', '2015-01-01', 40000)
		ON CONFLICT DO NOTHING
		RETURNING id
	`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		ids = append(ids, id)
	}

	log.Printf("Seeded employees: %d rows\n", len(ids))
	return ids
}

func seedAnnualReviews(ctx context.Context, tx pgx.Tx, employeeIDs []int64) {
	if len(employeeIDs) == 0 {
		log.Println("No employees inserted, skipping reviews")
		return
	}

	query := `
		INSERT INTO annual_reviews (
			employee_id,
			review_date
		)
		VALUES ($1, $2)
	`

	reviews := []struct {
		empIndex int
		date     string
	}{
		{0, "2010-01-01"},
		{0, "2011-01-01"},
		{1, "2011-01-01"},
		{1, "2012-01-01"},
		{3, "2013-01-01"},
	}

	for _, r := range reviews {
		_, err := tx.Exec(
			ctx,
			query,
			employeeIDs[r.empIndex],
			r.date,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seeded annual reviews")
}
