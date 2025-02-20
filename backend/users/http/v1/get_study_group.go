package v1

import (
	"encoding/json"
	"errors"
	"main/internal/groups"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Получение группы обучения по курсу и группе
// @tags studyGroups
// @description Получение группы обучения по курсу и группе
// @id getStudyGroupsByCourseAndGroup
// @Param course path int true "курс"
// @Param groupName path string true "название группы"
// @accept plain
// @produce json
// @Router /api/groups/{course}/{groupName} [get]
// @Success 200 {object} groups.StudyGroup
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func GetStudyGroup(repo *groups.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		course, err := strconv.Atoi(mux.Vars(r)["course"])
		if err != nil || course <= 0 {
			return errorx.BadRequest("course should be greater than 0")
		}

		group := mux.Vars(r)["groupName"]

		studyGroup, err := repo.GetStudyGroupByCourseAndTitle(r.Context(), course, group)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return errorx.BadRequest("group is not exists")
			}
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(studyGroup)
	}
}
