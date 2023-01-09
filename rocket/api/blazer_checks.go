package api

import (
	"net/http"

	"rocket/dao"
	"rocket/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configBlazerChecksRouter(router *httprouter.Router) {
	router.GET("/blazerchecks", GetAllBlazerChecks)
	router.POST("/blazerchecks", AddBlazerChecks)
	router.GET("/blazerchecks/:argID", GetBlazerChecks)
	router.PUT("/blazerchecks/:argID", UpdateBlazerChecks)
	router.DELETE("/blazerchecks/:argID", DeleteBlazerChecks)
}

func configGinBlazerChecksRouter(router gin.IRoutes) {
	router.GET("/blazerchecks", ConverHttprouterToGin(GetAllBlazerChecks))
	router.POST("/blazerchecks", ConverHttprouterToGin(AddBlazerChecks))
	router.GET("/blazerchecks/:argID", ConverHttprouterToGin(GetBlazerChecks))
	router.PUT("/blazerchecks/:argID", ConverHttprouterToGin(UpdateBlazerChecks))
	router.DELETE("/blazerchecks/:argID", ConverHttprouterToGin(DeleteBlazerChecks))
}

// GetAllBlazerChecks is a function to get a slice of record(s) from blazer_checks table in the rocket_development database
// @Summary Get list of BlazerChecks
// @Tags BlazerChecks
// @Description GetAllBlazerChecks is a handler to get a slice of record(s) from blazer_checks table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlazerChecks}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerchecks [get]
// http "https://xinqi.dev:8080/blazerchecks?page=0&pagesize=20" X-Api-User:user123
func GetAllBlazerChecks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "blazer_checks", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlazerChecks(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlazerChecks is a function to get a single record from the blazer_checks table in the rocket_development database
// @Summary Get record from table BlazerChecks by  argID
// @Tags BlazerChecks
// @ID argID
// @Description GetBlazerChecks is a function to get a single record from the blazer_checks table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BlazerChecks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blazerchecks/{argID} [get]
// http "https://xinqi.dev:8080/blazerchecks/1" X-Api-User:user123
func GetBlazerChecks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_checks", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlazerChecks(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlazerChecks add to add a single record to blazer_checks table in the rocket_development database
// @Summary Add an record to blazer_checks table
// @Description add to add a single record to blazer_checks table in the rocket_development database
// @Tags BlazerChecks
// @Accept  json
// @Produce  json
// @Param BlazerChecks body model.BlazerChecks true "Add BlazerChecks"
// @Success 200 {object} model.BlazerChecks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerchecks [post]
// echo '{"id": 93,"creator_id": 67,"query_id": 41,"state": "NRtsAmUUOIVdqZDctKKJOxedB","schedule": "sKHumyriFOWpyVwLZbKwCTQJM","emails": "ABWemDyAsADqwNffUjxXkWOtp","slack_channels": "cDyWyIkTsfltphXOwBhvJwBgb","check_type": "PFjpEkWShpelMJvMVkVbCUhfP","message": "PHIXavaIMwBoOfMrHSgJutmZd","last_run_at": "2150-03-25T17:01:27.465368917-04:00","created_at": "2038-01-11T09:04:29.452009613-05:00","updated_at": "2248-05-31T05:21:41.470212697-04:00"}' | http POST "https://xinqi.dev:8080/blazerchecks" X-Api-User:user123
func AddBlazerChecks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blazerchecks := &model.BlazerChecks{}

	if err := readJSON(r, blazerchecks); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerchecks.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerchecks.Prepare()

	if err := blazerchecks.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_checks", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blazerchecks, _, err = dao.AddBlazerChecks(ctx, blazerchecks)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerchecks)
}

// UpdateBlazerChecks Update a single record from blazer_checks table in the rocket_development database
// @Summary Update an record in table blazer_checks
// @Description Update a single record from blazer_checks table in the rocket_development database
// @Tags BlazerChecks
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BlazerChecks body model.BlazerChecks true "Update BlazerChecks record"
// @Success 200 {object} model.BlazerChecks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerchecks/{argID} [put]
// echo '{"id": 93,"creator_id": 67,"query_id": 41,"state": "NRtsAmUUOIVdqZDctKKJOxedB","schedule": "sKHumyriFOWpyVwLZbKwCTQJM","emails": "ABWemDyAsADqwNffUjxXkWOtp","slack_channels": "cDyWyIkTsfltphXOwBhvJwBgb","check_type": "PFjpEkWShpelMJvMVkVbCUhfP","message": "PHIXavaIMwBoOfMrHSgJutmZd","last_run_at": "2150-03-25T17:01:27.465368917-04:00","created_at": "2038-01-11T09:04:29.452009613-05:00","updated_at": "2248-05-31T05:21:41.470212697-04:00"}' | http PUT "https://xinqi.dev:8080/blazerchecks/1"  X-Api-User:user123
func UpdateBlazerChecks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerchecks := &model.BlazerChecks{}
	if err := readJSON(r, blazerchecks); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerchecks.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerchecks.Prepare()

	if err := blazerchecks.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_checks", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerchecks, _, err = dao.UpdateBlazerChecks(ctx,
		argID,
		blazerchecks)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerchecks)
}

// DeleteBlazerChecks Delete a single record from blazer_checks table in the rocket_development database
// @Summary Delete a record from blazer_checks
// @Description Delete a single record from blazer_checks table in the rocket_development database
// @Tags BlazerChecks
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BlazerChecks
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blazerchecks/{argID} [delete]
// http DELETE "https://xinqi.dev:8080/blazerchecks/1" X-Api-User:user123
func DeleteBlazerChecks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_checks", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlazerChecks(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
