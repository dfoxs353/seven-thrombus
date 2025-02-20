package v1

import (
	"encoding/json"
	"main/internal/users"
	"net/http"

	errorx "gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

type UpdateProfileReq struct {
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// @summary Обновление данных об аккаунте.
// @tags users
// @description Обновление данных об аккаунте.
// @id profileUpdate
// @accept json
// @produce json
// @Param reqBody body UpdateProfileReq true "Запрос на обновление данных аккаунта"
// @Router /api/accounts/update [put]
// @Success 200
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Security ApiKeyAuth
func UpdateProfile(service *users.Service) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		token, ok := middleware.TokenFromContext(r.Context())
		if !ok {
			return errorx.Unauthorized
		}

		var req UpdateProfileReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errorx.BadRequest(err.Error())
		}

		err := service.UpdateProfile(
			r.Context(),
			token.Uid,
			req.Password,
			req.FirstName,
			req.LastName,
		)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}
}
