package v1

import (
	"main/internal/groups"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Удаление группы обучения
// @tags studyGroups
// @description Удаление группы обучения
// @id deleteStudyGroup
// @accept plain
// @produce plain
// @Param id path int true "id группы обучения"
// @Router /api/groups/{id} [delete]
// @Success 200
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func DeleteStudyGroup(repo *groups.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return errorx.BadRequest("invalid id path param")
		}

		err = repo.DeleteStudyGroup(r.Context(), id)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}
}
