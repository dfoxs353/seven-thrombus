package v1

import (
	"encoding/json"
	"main/internal/schedules"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Обновление занятия
// @tags lessons
// @description Обновление занятия
// @id updateLesson
// @accept json
// @produce plain
// @param id path int true "id занятия"
// @Param reqBody body LessonReq true "Запрос на создание занятия. Поле date должно быть в следующем формате: 2006-01-02T15:04:05Z07:00 (ФОРМАТ ISO 8601)"
// @Router /api/lessons/{id} [put]
// @Success 200
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func UpdateLesson(repo *schedules.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req LessonReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errorx.BadRequest(err.Error())
		}

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return errorx.BadRequest("invalid id path param")
		}

		t, err := time.Parse(time.RFC3339, req.Date)
		if err != nil {
			return errorx.BadRequest(err.Error())
		}

		lesson := schedules.Lesson{
			Id:           id,
			DisciplineId: req.DisciplineId,
			TeacherId:    req.TeacherId,
			GroupId:      req.GroupId,
			Date:         t,
		}

		err = repo.UpdateLesson(
			r.Context(),
			lesson,
		)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}
}
