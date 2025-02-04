package v1

import (
	"encoding/json"
	"main/internal/users"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Получение доктора по id
// @tags doctors
// @description Получение доктора по id
// @id getDoctor
// @accept plain
// @produce plain
// @Param id path int true "id доктора"
// @Router /api/doctors/{id} [get]
// @Success 200
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 404 {object} errorx.ResponseError
// @Security ApiKeyAuth
func GetDoctor(repo *users.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		uid, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return errorx.BadRequest("invalid id path param")
		}

		user, err := repo.GetUserById(r.Context(), uid)
		if err != nil {
			return err
		}

		if !slices.Contains(user.Roles, users.Student) {
			return errorx.NotFound
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(user)
	}
}
