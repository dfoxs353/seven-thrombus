package v1

import (
	"encoding/json"
	"main/internal/users"
	"net/http"

	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Получение профиля пользователя
// @tags users
// @description Получение профиля пользователя
// @id me
// @accept plain
// @produce json
// @Router /api/accounts/me [get]
// @Success 200 {object} users.User
// @Failure 400 {object} errorx.ResponseError
// @Security ApiKeyAuth
func Profile(repo *users.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		token, ok := middleware.TokenFromContext(r.Context())
		if !ok {
			return errorx.Unauthorized
		}

		user, err := repo.GetUserById(r.Context(), token.Uid)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(user)
	}
}
