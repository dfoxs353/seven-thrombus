package v1

import (
	"encoding/json"
	"main/internal/users"
	"net/http"

	errorx "gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
	"golang.org/x/exp/maps"
)

type SignUpReq struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// @summary Регистрация нового пользователя
// @tags users
// @description Регистрация нового пользователя.
// @description Подразумевается, что username является уникальным
// @id singUp
// @accept json
// @produce json
// @Param reqBody body SignUpReq true "Запрос на создание пользователя"
// @Router /api/signup [post]
// @Success 201 {object} int
// @Failure 400 {object} errorx.ResponseError
func SignUp(service *users.Service, defaultRoles []string) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req SignUpReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errorx.BadRequest(err.Error())
		}

		rolesMap := make(map[users.Role]struct{})
		for _, role := range defaultRoles {
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
