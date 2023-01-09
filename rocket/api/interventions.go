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

func configInterventionsRouter(router *httprouter.Router) {
	router.GET("/interventions", GetAllInterventions)
	router.POST("/interventions", AddInterventions)
	router.GET("/interventions/:argID", GetInterventions)
	router.PUT("/interventions/:argID", UpdateInterventions)
	router.DELETE("/interventions/:argID", DeleteInterventions)
}

func configGinInterventionsRouter(router gin.IRoutes) {
	router.GET("/interventions", ConverHttprouterToGin(GetAllInterventions))
	router.POST("/interventions", ConverHttprouterToGin(AddInterventions))
	router.GET("/interventions/:argID", ConverHttprouterToGin(GetInterventions))
	router.PUT("/interventions/:argID", ConverHttprouterToGin(UpdateInterventions))
	router.DELETE("/interventions/:argID", ConverHttprouterToGin(DeleteInterventions))
}

// GetAllInterventions is a function to get a slice of record(s) from interventions table in the rocket_development database
// @Summary Get list of Interventions
// @Tags Interventions
// @Description GetAllInterventions is a handler to get a slice of record(s) from interventions table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Interventions}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /interventions [get]
// http "https://xinqi.dev:8080/interventions?page=0&pagesize=20" X-Api-User:user123
func GetAllInterventions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "interventions", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllInterventions(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetInterventions is a function to get a single record from the interventions table in the rocket_development database
// @Summary Get record from table Interventions by  argID
// @Tags Interventions
// @ID argID
// @Description GetInterventions is a function to get a single record from the interventions table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Interventions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /interventions/{argID} [get]
// http "https://xinqi.dev:8080/interventions/1" X-Api-User:user123
func GetInterventions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "interventions", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetInterventions(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddInterventions add to add a single record to interventions table in the rocket_development database
// @Summary Add an record to interventions table
// @Description add to add a single record to interventions table in the rocket_development database
// @Tags Interventions
// @Accept  json
// @Produce  json
// @Param Interventions body model.Interventions true "Add Interventions"
// @Success 200 {object} model.Interventions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /interventions [post]
// echo '{"id": 12,"author": "ULXeAoeOWHaVQIcntLRAhNcqT","customer_id": 79,"building_id": 57,"battery_id": 22,"column_id": 22,"elevator_id": 91,"employee_id": 66,"start_datetime": "2214-02-19T15:08:08.150883615-05:00","end_datetime": "2075-07-06T00:41:23.667801662-04:00","result": "waesoTBRXtTJhAaixBbqPuONs","report": "BQFuKyFDyiswYQSdanEplLQsP","status": "ZUOMEQLWemSyWbgowWyIGRaGO","created_at": "2081-06-15T15:33:32.12739989-04:00","updated_at": "2209-02-01T03:47:35.324875894-05:00"}' | http POST "https://xinqi.dev:8080/interventions" X-Api-User:user123
func AddInterventions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	interventions := &model.Interventions{}

	if err := readJSON(r, interventions); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := interventions.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	interventions.Prepare()

	if err := interventions.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "interventions", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	interventions, _, err = dao.AddInterventions(ctx, interventions)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, interventions)
}

// UpdateInterventions Update a single record from interventions table in the rocket_development database
// @Summary Update an record in table interventions
// @Description Update a single record from interventions table in the rocket_development database
// @Tags Interventions
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Interventions body model.Interventions true "Update Interventions record"
// @Success 200 {object} model.Interventions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /interventions/{argID} [put]
// echo '{"id": 12,"author": "ULXeAoeOWHaVQIcntLRAhNcqT","customer_id": 79,"building_id": 57,"battery_id": 22,"column_id": 22,"elevator_id": 91,"employee_id": 66,"start_datetime": "2214-02-19T15:08:08.150883615-05:00","end_datetime": "2075-07-06T00:41:23.667801662-04:00","result": "waesoTBRXtTJhAaixBbqPuONs","report": "BQFuKyFDyiswYQSdanEplLQsP","status": "ZUOMEQLWemSyWbgowWyIGRaGO","created_at": "2081-06-15T15:33:32.12739989-04:00","updated_at": "2209-02-01T03:47:35.324875894-05:00"}' | http PUT "https://xinqi.dev:8080/interventions/1"  X-Api-User:user123
func UpdateInterventions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	interventions := &model.Interventions{}
	if err := readJSON(r, interventions); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := interventions.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	interventions.Prepare()

	if err := interventions.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "interventions", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	interventions, _, err = dao.UpdateInterventions(ctx,
		argID,
		interventions)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, interventions)
}

// DeleteInterventions Delete a single record from interventions table in the rocket_development database
// @Summary Delete a record from interventions
// @Description Delete a single record from interventions table in the rocket_development database
// @Tags Interventions
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Interventions
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /interventions/{argID} [delete]
// http DELETE "https://xinqi.dev:8080/interventions/1" X-Api-User:user123
func DeleteInterventions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "interventions", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteInterventions(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
