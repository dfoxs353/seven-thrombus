package v1

import (
	"main/internal/jwt"
	"net/http"

	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Проверка токена на валидность
// @tags users
// @description Регистрация нового пользователя.
// @description Подразумевается, что username является уникальным
// @id validate
// @accept plain
// @produce plain
// @Param accessToken query string true "AccessToken"
// @Router /api/validate [get]
// @Success 200
// @Failure 400 {object} errorx.ResponseError
func Validate(validator jwt.TokenManager) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		token := r.URL.Query().Get("accessToken")
		_, err := validator.ParseToken(token)
		if err != nil {
			return errorx.BadRequest(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}
}
