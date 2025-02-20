package files

import (
	"context"

	"github.com/jackc/pgx/v5"
	"gitlab.com/volgaIt/packages/postgres"
)

type Repo struct {
	db postgres.PostgresQuerier
}

func NewRepo(db postgres.PostgresQuerier) *Repo {
	return &Repo{db}
}

type File struct {
	Id      int    `json:"id" db:"id"`
	Path    string `json:"path" db:"path"`
	SrcName string `json:"srcName" db:"src_name"`
}

func (r *Repo) GetFile(ctx context.Context, id int) (File, error) {
	sql := `
	SELECT id, "path", src_name 
	FROM files
	`

	rows, err := r.db.Query(ctx, sql, id)
	if err != nil {
		return File{}, err
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[File])
}

func (r *Repo) InsertFile(ctx context.Context, path, srcName string) (int, error) {
	sql := `
	INSERT INTO files("path", src_name)
	VALUES($1, $2)
	RETURNING id
	`

	var id int
	err := r.db.QueryRow(ctx, sql, path, srcName).Scan(&id)
	return id, err
}
