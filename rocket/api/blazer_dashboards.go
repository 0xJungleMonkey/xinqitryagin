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

func configBlazerDashboardsRouter(router *httprouter.Router) {
	router.GET("/blazerdashboards", GetAllBlazerDashboards)
	router.POST("/blazerdashboards", AddBlazerDashboards)
	router.GET("/blazerdashboards/:argID", GetBlazerDashboards)
	router.PUT("/blazerdashboards/:argID", UpdateBlazerDashboards)
	router.DELETE("/blazerdashboards/:argID", DeleteBlazerDashboards)
}

func configGinBlazerDashboardsRouter(router gin.IRoutes) {
	router.GET("/blazerdashboards", ConverHttprouterToGin(GetAllBlazerDashboards))
	router.POST("/blazerdashboards", ConverHttprouterToGin(AddBlazerDashboards))
	router.GET("/blazerdashboards/:argID", ConverHttprouterToGin(GetBlazerDashboards))
	router.PUT("/blazerdashboards/:argID", ConverHttprouterToGin(UpdateBlazerDashboards))
	router.DELETE("/blazerdashboards/:argID", ConverHttprouterToGin(DeleteBlazerDashboards))
}

// GetAllBlazerDashboards is a function to get a slice of record(s) from blazer_dashboards table in the rocket_development database
// @Summary Get list of BlazerDashboards
// @Tags BlazerDashboards
// @Description GetAllBlazerDashboards is a handler to get a slice of record(s) from blazer_dashboards table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlazerDashboards}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboards [get]
// http "https://xinqi.dev:8080/blazerdashboards?page=0&pagesize=20" X-Api-User:user123
func GetAllBlazerDashboards(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "blazer_dashboards", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlazerDashboards(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlazerDashboards is a function to get a single record from the blazer_dashboards table in the rocket_development database
// @Summary Get record from table BlazerDashboards by  argID
// @Tags BlazerDashboards
// @ID argID
// @Description GetBlazerDashboards is a function to get a single record from the blazer_dashboards table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BlazerDashboards
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blazerdashboards/{argID} [get]
// http "https://xinqi.dev:8080/blazerdashboards/1" X-Api-User:user123
func GetBlazerDashboards(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboards", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlazerDashboards(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlazerDashboards add to add a single record to blazer_dashboards table in the rocket_development database
// @Summary Add an record to blazer_dashboards table
// @Description add to add a single record to blazer_dashboards table in the rocket_development database
// @Tags BlazerDashboards
// @Accept  json
// @Produce  json
// @Param BlazerDashboards body model.BlazerDashboards true "Add BlazerDashboards"
// @Success 200 {object} model.BlazerDashboards
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboards [post]
// echo '{"id": 45,"creator_id": 31,"name": "SvUNqABgiDvKxMoOjamrEDNPW","created_at": "2226-04-23T16:26:38.487757247-04:00","updated_at": "2063-10-14T19:51:05.570973021-04:00"}' | http POST "https://xinqi.dev:8080/blazerdashboards" X-Api-User:user123
func AddBlazerDashboards(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blazerdashboards := &model.BlazerDashboards{}

	if err := readJSON(r, blazerdashboards); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerdashboards.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerdashboards.Prepare()

	if err := blazerdashboards.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboards", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blazerdashboards, _, err = dao.AddBlazerDashboards(ctx, blazerdashboards)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerdashboards)
}

// UpdateBlazerDashboards Update a single record from blazer_dashboards table in the rocket_development database
// @Summary Update an record in table blazer_dashboards
// @Description Update a single record from blazer_dashboards table in the rocket_development database
// @Tags BlazerDashboards
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BlazerDashboards body model.BlazerDashboards true "Update BlazerDashboards record"
// @Success 200 {object} model.BlazerDashboards
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazerdashboards/{argID} [put]
// echo '{"id": 45,"creator_id": 31,"name": "SvUNqABgiDvKxMoOjamrEDNPW","created_at": "2226-04-23T16:26:38.487757247-04:00","updated_at": "2063-10-14T19:51:05.570973021-04:00"}' | http PUT "https://xinqi.dev:8080/blazerdashboards/1"  X-Api-User:user123
func UpdateBlazerDashboards(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerdashboards := &model.BlazerDashboards{}
	if err := readJSON(r, blazerdashboards); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazerdashboards.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazerdashboards.Prepare()

	if err := blazerdashboards.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboards", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazerdashboards, _, err = dao.UpdateBlazerDashboards(ctx,
		argID,
		blazerdashboards)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazerdashboards)
}

// DeleteBlazerDashboards Delete a single record from blazer_dashboards table in the rocket_development database
// @Summary Delete a record from blazer_dashboards
// @Description Delete a single record from blazer_dashboards table in the rocket_development database
// @Tags BlazerDashboards
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BlazerDashboards
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blazerdashboards/{argID} [delete]
// http DELETE "https://xinqi.dev:8080/blazerdashboards/1" X-Api-User:user123
func DeleteBlazerDashboards(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_dashboards", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlazerDashboards(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
