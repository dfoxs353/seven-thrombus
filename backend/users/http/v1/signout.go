package v1

import (
	"main/internal/users"
	"net/http"

	errorx "gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Выход из аккаунта
// @tags users
// @description Refresh token берется из куки refreshToken и удаляется из базы.
// @description Access token автоматически испортится через время указанное в конфиге.
// @id singOut
// @accept plain
// @produce plain
// @Router /api/signout [put]
// @Success 200
// @Failure 400 {object} errorx.ResponseError
// @Security ApiKeyAuth
func SignOut(service *users.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		cookie, err := r.Cookie(refreshTokenCookieName)
		if err != nil {
			return errorx.Unauthorized
		}

		err = service.DeleteUserToken(
			r.Context(),
			cookie.Value,
		)
		if err != nil {
			return err
		}

		http.SetCookie(w, &http.Cookie{
			Name:   refreshTokenCookieName,
			MaxAge: -1,
		})

		w.WriteHeader(http.StatusOK)
		return nil
	}
}
