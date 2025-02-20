package disciplines

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"gitlab.com/volgaIt/packages/postgres"
)

type Discipline struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

type Repo struct {
	db postgres.PostgresQuerier
}

func NewRepo(db postgres.PostgresQuerier) *Repo {
	return &Repo{db}
}

func (r *Repo) InsertDiscipline(ctx context.Context, title string) (int, error) {
	var id int
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO disciplines (title, created_at) VALUES($1, $2)
		RETURNING id
		`,
		title,
		time.Now().Unix()).Scan(&id)

	return id, err
}

func (r *Repo) UpdateDiscipline(ctx context.Context, id int, title string) error {
	_, err := r.db.Exec(ctx, `UPDATE disciplines SET title = $1 WHERE id = $2`, title, id)
	return err
}

func (r *Repo) DeleteDiscipline(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM disciplines WHERE id = $1`, id)
	return err
}

func (r *Repo) GetDisciplines(ctx context.Context, limit *int, offset *int) ([]Discipline, error) {
	sql := `SELECT id, title FROM disciplines`
	if limit != nil && offset != nil {
		sql += fmt.Sprintf(" LIMIT %d OFFSET %d", *limit, *offset)
	}

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[Discipline])
}

func (r *Repo) GetDiscipline(ctx context.Context, id int) (Discipline, error) {
	sql := `SELECT id, title FROM disciplines WHERE id = $1`

	rows, err := r.db.Query(ctx, sql, id)
	if err != nil {
		return Discipline{}, err
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Discipline])
}
