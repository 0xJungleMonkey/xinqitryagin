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

func configBlazerDashboardQueriesRouter(router *httprouter.Router) {
	router.GET("/blazerdashboardqueries", GetAllBlazerDashboardQueries)
	router.POST("/blazerdashboardqueries", AddBlazerDashboardQueries)
	router.GET("/blazerdashboardqueries/:argID", GetBlazerDashboardQueries)
	router.PUT("/blazerdashboardqueries/:argID", UpdateBlazerDashboardQueries)
	router.DELETE("/blazerdashboardqueries/:argID", DeleteBlazerDashboardQueries)
}

func configGinBlazerDashboardQueriesRouter(router gin.IRoutes) {
	router.GET("/blazerdashboardqueries", ConverHttprouterToGin(GetAllBlazerDashboardQueries))
	router.POST("/blazerdashboardqueries", ConverHttprouterToGin(AddBlazerDashboardQueries))
	router.GET("/blazerdashboardqueries/:argID", ConverHttprouterToGin(GetBlazerDashboardQueries))
	router.PUT("/blazerdashboardqueries/:argID", ConverHttprouterToGin(UpdateBlazerDashboardQueries))
	router.DELETE("/blazerdashboardqueries/:argID", ConverHttprouterToGin(DeleteBlazerDashboardQueries))
}

// GetAllBlazerDashboardQueries is a function to get a slice of record(s) from blazer_dashboard_queries table in the rocket_development database
// @Summary Get list of BlazerDashboardQueries
// @Tags BlazerDashboardQueries
// @Description GetAllBlazerDashboardQueries is a handler to get a slice of record(s) from blazer_dashboard_queries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlazerDashboardQueries}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboardqueries [get]
// http "http://localhost:8080/blazerdashboardqueries?page=0&pagesize=20" X-Api-User:user123
func GetAllBlazerDashboardQueries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "blazer_dashboard_queries", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlazerDashboardQueries(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlazerDashboardQueries is a function to get a single record from the blazer_dashboard_queries table in the rocket_development database
// @Summary Get record from table BlazerDashboardQueries by  argID
// @Tags BlazerDashboardQueries
// @ID argID
// @Description GetBlazerDashboardQueries is a function to get a single record from the blazer_dashboard_queries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BlazerDashboardQueries
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blazerdashboardqueries/{argID} [get]
// http "http://localhost:8080/blazerdashboardqueries/1" X-Api-User:user123
func GetBlazerDashboardQueries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboard_queries", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlazerDashboardQueries(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlazerDashboardQueries add to add a single record to blazer_dashboard_queries table in the rocket_development database
// @Summary Add an record to blazer_dashboard_queries table
// @Description add to add a single record to blazer_dashboard_queries table in the rocket_development database
// @Tags BlazerDashboardQueries
// @Accept  json
// @Produce  json
// @Param BlazerDashboardQueries body model.BlazerDashboardQueries true "Add BlazerDashboardQueries"
// @Success 200 {object} model.BlazerDashboardQueries
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboardqueries [post]
// echo '{"id": 16,"dashboard_id": 25,"query_id": 25,"position": 33,"created_at": "2083-06-27T00:34:44.350960304-04:00","updated_at": "2108-09-07T00:04:05.593812459-04:00"}' | http POST "http://localhost:8080/blazerdashboardqueries" X-Api-User:user123
func AddBlazerDashboardQueries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blazerdashboardqueries := &model.BlazerDashboardQueries{}

	if err := readJSON(r, blazerdashboardqueries); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerdashboardqueries.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerdashboardqueries.Prepare()

	if err := blazerdashboardqueries.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboard_queries", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blazerdashboardqueries, _, err = dao.AddBlazerDashboardQueries(ctx, blazerdashboardqueries)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerdashboardqueries)
}

// UpdateBlazerDashboardQueries Update a single record from blazer_dashboard_queries table in the rocket_development database
// @Summary Update an record in table blazer_dashboard_queries
// @Description Update a single record from blazer_dashboard_queries table in the rocket_development database
// @Tags BlazerDashboardQueries
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BlazerDashboardQueries body model.BlazerDashboardQueries true "Update BlazerDashboardQueries record"
// @Success 200 {object} model.BlazerDashboardQueries
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboardqueries/{argID} [put]
// echo '{"id": 16,"dashboard_id": 25,"query_id": 25,"position": 33,"created_at": "2083-06-27T00:34:44.350960304-04:00","updated_at": "2108-09-07T00:04:05.593812459-04:00"}' | http PUT "http://localhost:8080/blazerdashboardqueries/1"  X-Api-User:user123
func UpdateBlazerDashboardQueries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerdashboardqueries := &model.BlazerDashboardQueries{}
	if err := readJSON(r, blazerdashboardqueries); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerdashboardqueries.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerdashboardqueries.Prepare()

	if err := blazerdashboardqueries.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboard_queries", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerdashboardqueries, _, err = dao.UpdateBlazerDashboardQueries(ctx,
		argID,
		blazerdashboardqueries)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerdashboardqueries)
}

// DeleteBlazerDashboardQueries Delete a single record from blazer_dashboard_queries table in the rocket_development database
// @Summary Delete a record from blazer_dashboard_queries
// @Description Delete a single record from blazer_dashboard_queries table in the rocket_development database
// @Tags BlazerDashboardQueries
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BlazerDashboardQueries
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blazerdashboardqueries/{argID} [delete]
// http DELETE "http://localhost:8080/blazerdashboardqueries/1" X-Api-User:user123
func DeleteBlazerDashboardQueries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboard_queries", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlazerDashboardQueries(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
