package v1

import (
	"encoding/json"
	"fmt"
	"main/internal/schedules"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/volgaIt/packages/errorx"
	middleware "gitlab.com/volgaIt/packages/middleware"
)

// @summary Получение списка занятий
// @tags lessons
// @description Получение списка дисциплин. Если требуется получить все занятия за 1 сентября 2025 года, то требуется указать параметр dateFrom=2025-09-01T00:00:00Z и dateTo=2025-09-02T00:00:00Z (для железобетонности можно указать 2025-09-01T023:59:59Z)
// @id getLessons
// @accept plain
// @produce json
// @Param dateFrom query string true "Начало временного промежутка в котором ищутся пары. ФОРМАТ ISO 8601"
// @Param dateTo query string true "Начало временного промежутка в котором ищутся пары. ФОРМАТ ISO 8601"
// @Param groupId query int true "Группа для которой требуется получить расписание"
// @Router /api/lessons [get]
// @Success 200 {array} schedules.Lesson
// @Failure 400 {object} errorx.ResponseError
// @Failure 401 {object} errorx.ResponseError
// @Failure 403 {object} errorx.ResponseError
// @Security ApiKeyAuth
func GetLessons(repo *schedules.Repo) middleware.ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		dateFrom, err := time.Parse(time.RFC3339, r.URL.Query().Get("dateFrom"))
		if err != nil {
			return errorx.BadRequest(fmt.Sprintln("dateFrom: ", err.Error()))
		}

		dateTo, err := time.Parse(time.RFC3339, r.URL.Query().Get("dateTo"))
		if err != nil {
			return errorx.BadRequest(fmt.Sprintln("dateTo: ", err.Error()))
		}

		groupId, err := strconv.Atoi(r.URL.Query().Get("groupId"))
		if err != nil || groupId <= 0 {
			return errorx.BadRequest("group id should be greater than 0")
		}

		lessons, err := repo.GetLessons(r.Context(), &dateFrom, &dateTo, &groupId)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return json.NewEncoder(w).Encode(lessons)
	}
}
