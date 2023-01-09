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

func configBlazerQueriesRouter(router *httprouter.Router) {
	router.GET("/blazerqueries", GetAllBlazerQueries)
	router.POST("/blazerqueries", AddBlazerQueries)
	router.GET("/blazerqueries/:argID", GetBlazerQueries)
	router.PUT("/blazerqueries/:argID", UpdateBlazerQueries)
	router.DELETE("/blazerqueries/:argID", DeleteBlazerQueries)
}

func configGinBlazerQueriesRouter(router gin.IRoutes) {
	router.GET("/blazerqueries", ConverHttprouterToGin(GetAllBlazerQueries))
	router.POST("/blazerqueries", ConverHttprouterToGin(AddBlazerQueries))
	router.GET("/blazerqueries/:argID", ConverHttprouterToGin(GetBlazerQueries))
	router.PUT("/blazerqueries/:argID", ConverHttprouterToGin(UpdateBlazerQueries))
	router.DELETE("/blazerqueries/:argID", ConverHttprouterToGin(DeleteBlazerQueries))
}

// GetAllBlazerQueries is a function to get a slice of record(s) from blazer_queries table in the rocket_development database
// @Summary Get list of BlazerQueries
// @Tags BlazerQueries
// @Description GetAllBlazerQueries is a handler to get a slice of record(s) from blazer_queries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlazerQueries}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerqueries [get]
// http "http://localhost:8080/blazerqueries?page=0&pagesize=20" X-Api-User:user123
func GetAllBlazerQueries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "blazer_queries", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlazerQueries(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlazerQueries is a function to get a single record from the blazer_queries table in the rocket_development database
// @Summary Get record from table BlazerQueries by  argID
// @Tags BlazerQueries
// @ID argID
// @Description GetBlazerQueries is a function to get a single record from the blazer_queries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BlazerQueries
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blazerqueries/{argID} [get]
// http "http://localhost:8080/blazerqueries/1" X-Api-User:user123
func GetBlazerQueries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_queries", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlazerQueries(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlazerQueries add to add a single record to blazer_queries table in the rocket_development database
// @Summary Add an record to blazer_queries table
// @Description add to add a single record to blazer_queries table in the rocket_development database
// @Tags BlazerQueries
// @Accept  json
// @Produce  json
// @Param BlazerQueries body model.BlazerQueries true "Add BlazerQueries"
// @Success 200 {object} model.BlazerQueries
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerqueries [post]
// echo '{"id": 24,"creator_id": 63,"name": "NKJULBgGyTRTRBeQwnNCFhZny","description": "lqrjhYBOgbkVZvZvkiMEXSQVu","statement": "WnsrTBkdhcfPjyNEyZixmYKRS","data_source": "FICQVMFiXtpGrIDPIGeBMraPu","status": "jfoLkbpiruRwrJFlUBJhXxyOg","created_at": "2039-06-21T02:48:05.900665277-04:00","updated_at": "2038-06-04T17:32:44.805071809-04:00"}' | http POST "http://localhost:8080/blazerqueries" X-Api-User:user123
func AddBlazerQueries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blazerqueries := &model.BlazerQueries{}

	if err := readJSON(r, blazerqueries); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerqueries.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerqueries.Prepare()

	if err := blazerqueries.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_queries", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blazerqueries, _, err = dao.AddBlazerQueries(ctx, blazerqueries)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerqueries)
}

// UpdateBlazerQueries Update a single record from blazer_queries table in the rocket_development database
// @Summary Update an record in table blazer_queries
// @Description Update a single record from blazer_queries table in the rocket_development database
// @Tags BlazerQueries
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BlazerQueries body model.BlazerQueries true "Update BlazerQueries record"
// @Success 200 {object} model.BlazerQueries
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerqueries/{argID} [put]
// echo '{"id": 24,"creator_id": 63,"name": "NKJULBgGyTRTRBeQwnNCFhZny","description": "lqrjhYBOgbkVZvZvkiMEXSQVu","statement": "WnsrTBkdhcfPjyNEyZixmYKRS","data_source": "FICQVMFiXtpGrIDPIGeBMraPu","status": "jfoLkbpiruRwrJFlUBJhXxyOg","created_at": "2039-06-21T02:48:05.900665277-04:00","updated_at": "2038-06-04T17:32:44.805071809-04:00"}' | http PUT "http://localhost:8080/blazerqueries/1"  X-Api-User:user123
func UpdateBlazerQueries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerqueries := &model.BlazerQueries{}
	if err := readJSON(r, blazerqueries); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerqueries.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerqueries.Prepare()

	if err := blazerqueries.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_queries", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerqueries, _, err = dao.UpdateBlazerQueries(ctx,
		argID,
		blazerqueries)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerqueries)
}

// DeleteBlazerQueries Delete a single record from blazer_queries table in the rocket_development database
// @Summary Delete a record from blazer_queries
// @Description Delete a single record from blazer_queries table in the rocket_development database
// @Tags BlazerQueries
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BlazerQueries
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blazerqueries/{argID} [delete]
// http DELETE "http://localhost:8080/blazerqueries/1" X-Api-User:user123
func DeleteBlazerQueries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_queries", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlazerQueries(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
