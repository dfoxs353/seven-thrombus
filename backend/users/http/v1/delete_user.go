package v1

import (
	"main/internal/users"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Мягкое удаление пользователя из базы
// @tags admin
// @description Мягкое удаление пользователя из базы
// @id deleteUser
// @accept plain
// @produce plain
// @Param id path int true "id пользователя"
// @Router /api/accounts/{id} [delete]
// @Success 200
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func DeleteUserSoft(repo *users.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		uid, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return errorx.BadRequest("invalid id path param")
		}

		err = repo.DeleteUserSoft(r.Context(), uid)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}
}
