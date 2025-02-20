package v1

import (
	"encoding/json"
	"main/internal/schedules"
	"net/http"
	"time"

	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

type LessonReq struct {
	DisciplineId int    `json:"disciplineId"`
	TeacherId    int    `json:"teacherId"`
	GroupId      int    `json:"groupId"`
	Date         string `json:"date"`
}

// @summary Создание занятия
// @tags lessons
// @description Создание занятия
// @id createLesson
// @accept json
// @produce json
// @Param reqBody body LessonReq true "Запрос на создание занятия. Поле date должно быть в следующем формате: 2006-01-02T15:04:05Z07:00 (ФОРМАТ ISO 8601)"
// @Router /api/lessons [post]
// @Success 201 {object} int
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func CreateLesson(repo *schedules.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req LessonReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errorx.BadRequest(err.Error())
		}

		t, err := time.Parse(time.RFC3339, req.Date)
		if err != nil {
			return errorx.BadRequest(err.Error())
		}

		lesson := schedules.Lesson{
			DisciplineId: req.DisciplineId,
			TeacherId:    req.TeacherId,
			GroupId:      req.GroupId,
			Date:         t,
		}

		id, err := repo.InsertLesson(
			r.Context(),
			lesson,
		)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		return json.NewEncoder(w).Encode(id)
	}
}
