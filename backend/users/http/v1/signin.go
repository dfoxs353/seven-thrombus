package v1

import (
	"encoding/json"
	"main/internal/users"
	"net/http"

	errorx "gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

const refreshTokenCookieName = "refresh_token"

type SignInReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @summary Вход пользователя в аккаунт
// @tags users
// @description Вход пользоватлея в аккаунт и получение новой пары jwt токенов.
// @id singIn
// @accept json
// @produce json
// @Param reqBody body SignInReq true "Запрос на авторизацию"
// @Router /api/signin [post]
// @Success 201 {object} jwt.TokenPair
// @Failure 400 {object} errorx.ResponseError
// @Failure 404 {object} errorx.ResponseError
func SignIn(service *users.Service) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req SignInReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errorx.BadRequest(err.Error())
		}

		tokens, err := service.SignIn(
			r.Context(),
			req.Username,
			req.Password,
		)
		if err != nil {
			return err
		}

		http.SetCookie(w, &http.Cookie{
			Name:     refreshTokenCookieName,
			Value:    tokens.RefreshToken,
			MaxAge:   tokens.RefreshTokenExpiresIn,
			HttpOnly: true,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		return json.NewEncoder(w).Encode(tokens)
	}
}
