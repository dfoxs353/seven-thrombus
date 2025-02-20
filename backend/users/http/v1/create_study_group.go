package v1

import (
	"encoding/json"
	"main/internal/groups"
	"net/http"

	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

type AdminStudyGroup struct {
	Course int    `json:"course"`
	Group  string `json:"group"`
}

// @summary Создание группы обучения
// @tags studyGroups
// @description Создание группы обучения
// @id createStudyGroup
// @accept plain
// @produce json
// @Param reqBody body AdminStudyGroup true "Запрос на создание группы обучения"
// @Router /api/groups [post]
// @Success 201 {object} int
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func CreateStudyGroup(repo *groups.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req AdminStudyGroup
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errorx.BadRequest(err.Error())
		}

		id, err := repo.InsertStudyGroup(r.Context(), req.Course, req.Group)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		return json.NewEncoder(w).Encode(id)
	}
}
