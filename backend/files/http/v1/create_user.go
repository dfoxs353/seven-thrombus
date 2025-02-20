package v1

import (
	"encoding/json"
	"main/internal/users"
	"net/http"

	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
	"golang.org/x/exp/maps"
)

type AdminUserReq struct {
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Roles     []string `json:"roles"`
}

// @summary Создание пользователя
// @tags admin
// @description Создание пользователя
// @id createUser
// @accept plain
// @produce json
// @Param reqBody body AdminUserReq true "Запрос на создание пользователя"
// @Router /api/accounts [post]
// @Success 201 {object} int
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func CreateUser(service *users.Service) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req AdminUserReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errorx.BadRequest(err.Error())
		}

		rolesMap := make(map[users.Role]struct{})
		for _, role := range req.Roles {
			rolesMap[users.StringToRole(role)] = struct{}{}
		}

		id, err := service.SignUp(
			r.Context(),
			req.Username,
			req.Password,
			req.FirstName,
			req.LastName,
			maps.Keys(rolesMap),
		)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		return json.NewEncoder(w).Encode(id)
	}
}
