package v1

import (
	"encoding/json"
	"main/internal/schedules"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Получение занятия по id
// @tags lessons
// @description Получение занятия по id
// @id getLesson
// @accept plain
// @produce json
// @param id path int true "id занятия"
// @Router /api/lessons/{id} [get]
// @Success 200 {object} schedules.Lesson
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func GetLesson(repo *schedules.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return errorx.BadRequest("invalid id path param")
		}
		// from-1 т.к. offset на 1 меньше, чем мы хотим сделать выборку
		lesson, err := repo.GetLesson(r.Context(), id)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(lesson)
	}
}
