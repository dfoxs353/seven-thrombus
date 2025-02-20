package v1

import (
	"errors"
	"main/internal/disciplines"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgconn"
	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Удаление дисциплины из базы
// @tags disciplines
// @description Удаление дисциплины из базы
// @id deleteDiscipline
// @accept plain
// @produce plain
// @Param id path int true "id дисциплины"
// @Router /api/disciplines/{id} [delete]
// @Success 200
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func DeleteDiscipline(repo *disciplines.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return errorx.BadRequest("invalid id path param")
		}

		err = repo.DeleteDiscipline(r.Context(), id)
		if err != nil {
			// Проверяем, является ли ошибка нарушением внешнего ключа

			var pqErr *pgconn.PgError
			if errors.As(err, &pqErr) {
				if pqErr.Code == "23503" {
					return errorx.BadRequest("Невозможно удалить дисциплину, т.к. на нее установлены занятия")
				}
			}
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}
}
