package v1

import (
	"main/internal/schedules"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Удаление занятия по id
// @tags lessons
// @description Удаление занятия по id
// @id deleteLesson
// @accept plain
// @produce plain
// @param id path int true "id занятия"
// @Router /api/lessons/{id} [delete]
// @Success 200
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func DeleteLesson(repo *schedules.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return errorx.BadRequest("invalid id path param")
		}
		// from-1 т.к. offset на 1 меньше, чем мы хотим сделать выборку
		err = repo.DeleteLesson(r.Context(), id)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}
}
