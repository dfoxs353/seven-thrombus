package v1

import (
	"encoding/json"
	"main/internal/users"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	errorx "gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
	"golang.org/x/exp/maps"
)

// @summary Обновление аккаунта пользователя.
// @tags admin
// @description Обновление данных об аккаунте.
// @id updateUser
// @accept json
// @produce json
// @Param id path int true "id пользователя"
// @Param reqBody body AdminUserReq true "Запрос на аккаунта пользователя"
// @Router /api/accounts/{id} [put]
// @Success 200
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func UpdateUser(service *users.Service) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		uid, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return errorx.BadRequest("invalid id path param")
		}

		var req AdminUserReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errorx.BadRequest(err.Error())
		}

		rolesMap := make(map[users.Role]struct{})
		for _, role := range req.Roles {
			rolesMap[users.StringToRole(role)] = struct{}{}
		}

		err = service.UpdateUser(
			r.Context(),
			users.User{
				Id:        uid,
				Username:  req.Username,
				Password:  req.Password,
				FirstName: req.FirstName,
				LastName:  req.LastName,
				Roles:     maps.Keys(rolesMap),
			},
		)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}
}
