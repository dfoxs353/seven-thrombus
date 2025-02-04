package v1

import (
	"encoding/json"
	"main/internal/users"
	"net/http"

	errorx "gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

type RefreshReq struct {
	Token string `json:"refreshToken"`
}

// @summary Получение новой пары токенов
// @tags users
// @description Получение новой токенов.
// @id refresh
// @accept json
// @produce json
// @Param reqBody body RefreshReq true "Запрос на получение новой пары токенов"
// @Router /api/refresh [post]
// @Success 201 {object} jwt.TokenPair
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 404 {object} errorx.ResponseError
func Refresh(service *users.Service) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req RefreshReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errorx.BadRequest(err.Error())
		}

		tokens, err := service.RefreshTokens(
			r.Context(),
			req.Token,
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
