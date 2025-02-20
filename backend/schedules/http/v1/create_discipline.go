package v1

import (
	"encoding/json"
	"main/internal/disciplines"
	"net/http"

	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

type DisciplineReq struct {
	Title string `json:"title"`
}

// @summary Создание дисциплины
// @tags disciplines
// @description Создание дисциплины
// @id createDiscipline
// @accept json
// @produce json
// @Param reqBody body DisciplineReq true "Запрос на создание дисциплины"
// @Router /api/disciplines [post]
// @Success 201 {object} int
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func CreateDiscipline(repo *disciplines.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req DisciplineReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errorx.BadRequest(err.Error())
		}

		id, err := repo.InsertDiscipline(
			r.Context(),
			req.Title,
		)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		return json.NewEncoder(w).Encode(id)
	}
}
