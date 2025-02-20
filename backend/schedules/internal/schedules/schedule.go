package schedules

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"gitlab.com/volgaIt/packages/postgres"
)

type Repo struct {
	db postgres.PostgresQuerier
}

func NewRepo(db postgres.PostgresQuerier) *Repo {
	return &Repo{db: db}
}

type Lesson struct {
	Id               int       `json:"id" db:"id"`
	DisciplineId     int       `json:"disciplineId" db:"discipline_id"`
	DisciplineTitle  string    `json:"discipline" db:"discipline_title"`
	TeacherId        int       `json:"teacherId" db:"teacher_id"`
	TeacherFirstName string    `json:"teacherFirstName" db:"teacher_first_name"`
	TeacherLastName  string    `json:"teacherLastName" db:"teacher_last_name"`
	GroupId          int       `json:"groupId" db:"group_id"`
	Date             time.Time `json:"date" db:"date"`
	IsOnline         bool      `json:"isOnline" db:"is_online"`
	FileIds          []int     `json:"fileIds" db:"file_ids"`
}

func (r *Repo) InsertLesson(ctx context.Context, lesson Lesson) (int, error) {
	sql := `
	INSERT INTO lessons
	(discipline_id, teacher_id, group_id, "date", is_online, file_ids, created_at)
	VALUES($1, $2, $3, $4, $5, $6, $7)
	RETURNING id;`

	filesBytes := []byte("[]")
	if len(lesson.FileIds) > 0 {
		filesBytes, _ = json.Marshal(lesson.FileIds)
	}

	var id int
	err := r.db.QueryRow(ctx, sql,
		lesson.DisciplineId,
		lesson.TeacherId,
		lesson.GroupId,
		lesson.Date,
		lesson.IsOnline,
		filesBytes,
		time.Now().Unix(),
	).Scan(&id)

	return id, err
}

func (r *Repo) UpdateLesson(ctx context.Context, lesson Lesson) error {
	sql := `
	UPDATE lessons
	SET discipline_id=$2, teacher_id=$3, group_id=$4, "date"=$5, is_online=$6, file_ids=$7
	WHERE id=$1`

	filesBytes := []byte("[]")
	if len(lesson.FileIds) > 0 {
		filesBytes, _ = json.Marshal(lesson.FileIds)
	}

	_, err := r.db.Exec(ctx, sql,
		lesson.Id,
		lesson.DisciplineId,
		lesson.TeacherId,
		lesson.GroupId,
		lesson.Date,
		lesson.IsOnline,
		filesBytes,
	)

	return err
}

func (r *Repo) DeleteLesson(ctx context.Context, id int) error {
	sql := `DELETE FROM lessons WHERE id=$1;`

	_, err := r.db.Exec(ctx, sql, id)
	return err
}

func (r *Repo) GetLesson(ctx context.Context, id int) (Lesson, error) {
	sql := `
	SELECT 
		l.id, 
		discipline_id, 
		teacher_id, 
		group_id, 
		"date", 
		is_online, 
		file_ids, 
		u.first_name as teacher_first_name, 
		u.last_name as teacher_last_name,
		d.title as "discipline_title"
	FROM lessons l
	LEFT JOIN users u ON u.id = l.teacher_id  
	LEFT JOIN disciplines d ON d.id = l.discipline_id 
	WHERE l.id = $1`

	row, err := r.db.Query(ctx, sql, id)
	if err != nil {
		return Lesson{}, err
	}
	defer row.Close()

	return pgx.CollectOneRow(row, pgx.RowToStructByNameLax[Lesson])
}

func (r *Repo) GetLessons(ctx context.Context, dateFrom, dateUntil *time.Time, groupId *int) ([]Lesson, error) {
	sql := `
	SELECT 
		l.id, 
		discipline_id, 
		teacher_id, 
		group_id, 
		"date", 
		is_online, 
		file_ids, 
		u.first_name as teacher_first_name, 
		u.last_name as teacher_last_name,
		d.title as "discipline_title"
	FROM lessons l
	LEFT JOIN users u ON u.id = l.teacher_id  
	LEFT JOIN disciplines d ON d.id = l.discipline_id 
	`

	var (
		args       []any
		conditions string
	)

	if dateFrom != nil {
		args = append(args, *dateFrom)
		conditions += fmt.Sprintf(" AND date >= $%d", len(args))
	}
	if dateUntil != nil {
		args = append(args, *dateUntil)
		conditions += fmt.Sprintf(" AND date <= $%d", len(args))
	}
	if groupId != nil {
		args = append(args, *groupId)
		conditions += fmt.Sprintf(" AND group_id = $%d", len(args))
	}

	if len(args) > 0 {
		conditions = strings.Replace(conditions, "AND", "", 1)
		sql += " WHERE " + conditions
	}

	sql += " ORDER BY l.created_at"

	row, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	return pgx.CollectRows(row, pgx.RowToStructByNameLax[Lesson])
}
