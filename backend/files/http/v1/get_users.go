package v1

import (
	"encoding/json"
	"main/internal/users"
	"net/http"
	"strconv"

	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Получение списка пользователей
// @tags admin
// @description Получение списка пользователей
// @id getUsers
// @accept plain
// @produce json
// @Param count query int false "Размер выборки. По умолчанию 20"
// @Param from query int false "Начало выборки. По умолчанию 1"
// @Router /api/accounts [get]
// @Success 200 {array} users.User
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func GetUsers(repo *users.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var (
			from  = 1
			count = 20
		)

		queryFrom := r.URL.Query().Get("from")
		if queryFrom != "" {
			fromVal, _ := strconv.Atoi(queryFrom)
			if fromVal <= 0 {
				return errorx.BadRequest("invalid from query param. Should be integer greater or equal 1")
			}
			from = fromVal
		}

		queryCount := r.URL.Query().Get("count")
		if queryCount != "" {
			countVal, _ := strconv.Atoi(queryCount)
			if countVal <= 0 {
				return errorx.BadRequest("invalid count query param. Should be integer greater or equal 1")
			}
			count = countVal
		}

		// from-1 т.к. offset на 1 меньше, чем мы хотим сделать выборку
		users, err := repo.GetUsers(r.Context(), from-1, count, nil, nil)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(users)
	}
}
