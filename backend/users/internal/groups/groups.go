package groups

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"gitlab.com/volgaIt/packages/postgres"
)

type Repo struct {
	db postgres.PostgresQuerier
}

func NewRepo(db postgres.PostgresQuerier) *Repo {
	return &Repo{db}
}

var (
	ErrGroupIsNotExists = errors.New("Указанной группы не существует. Убедитесь, что группа с указанным курсом и группой действительно существует")
)

type StudyGroup struct {
	Id     int    `json:"id" db:"id"`
	Course int    `json:"course" db:"course"`
	Title  string `json:"title" db:"title"`
}

func (r *Repo) InsertStudyGroup(ctx context.Context, course int, title string) (int, error) {
	var id int
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO study_groups (course, title, created_at) VALUES($1, $2, $3) RETURNING id`,
		course,
		title,
		time.Now().Unix()).Scan(&id)

	return id, err
}

func (r *Repo) UpdateStudyGroup(ctx context.Context, id int, course int, title string) error {
	_, err := r.db.Exec(ctx, `UPDATE study_groups SET course = $1, title = $2 WHERE id = $3`, course, title, id)
	return err
}

func (r *Repo) DeleteStudyGroup(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM study_groups WHERE id = $1`, id)
	return err
}

func (r *Repo) GetStudyGroups(ctx context.Context, limit *int, offset *int) ([]StudyGroup, error) {
	sql := `SELECT id, course, title FROM study_groups`
	if limit != nil && offset != nil {
		sql += fmt.Sprintf(" LIMIT %d OFFSET %d", *limit, *offset)
	}

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[StudyGroup])
}

func (r *Repo) GetStudyGroup(ctx context.Context, id int) (StudyGroup, error) {
	sql := `SELECT id, course, title FROM study_groups WHERE id = $1`

	rows, err := r.db.Query(ctx, sql, id)
	if err != nil {
		return StudyGroup{}, err
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[StudyGroup])
}

func (r *Repo) GetStudyGroupByCourseAndTitle(ctx context.Context, course int, group string) (StudyGroup, error) {
	sql := `SELECT id, course, title FROM study_groups WHERE course = $1 AND title = $2`

	rows, err := r.db.Query(ctx, sql, course, group)
	if err != nil {
		return StudyGroup{}, err
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[StudyGroup])
}

func (r *Repo) InsertStudentIntoStudyGroup(ctx context.Context, studentId int, groupId int) error {
	sql := `
	INSERT INTO student_groups (student_id, group_id, created_at)
	VALUES($1, $2, $3)
	`
	_, err := r.db.Exec(ctx, sql, studentId, groupId, time.Now().Unix())
	return err
}

func (r *Repo) DeleteStudentFromGroup(ctx context.Context, studentId, groupId int) error {
	sql := `
	DELETE FROM student_groups WHERE student_id = $1 AND group_id = $2  
	`

	_, err := r.db.Exec(ctx, sql, studentId, groupId)
	return err
}
