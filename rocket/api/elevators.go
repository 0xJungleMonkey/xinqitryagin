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

func configElevatorsRouter(router *httprouter.Router) {
	router.GET("/elevators", GetAllElevators)
	router.POST("/elevators", AddElevators)
	router.GET("/elevators/:argID", GetElevators)
	router.PUT("/elevators/:argID", UpdateElevators)
	router.DELETE("/elevators/:argID", DeleteElevators)
}

func configGinElevatorsRouter(router gin.IRoutes) {
	router.GET("/elevators", ConverHttprouterToGin(GetAllElevators))
	router.POST("/elevators", ConverHttprouterToGin(AddElevators))
	router.GET("/elevators/:argID", ConverHttprouterToGin(GetElevators))
	router.PUT("/elevators/:argID", ConverHttprouterToGin(UpdateElevators))
	router.DELETE("/elevators/:argID", ConverHttprouterToGin(DeleteElevators))
}

// GetAllElevators is a function to get a slice of record(s) from elevators table in the rocket_development database
// @Summary Get list of Elevators
// @Tags Elevators
// @Description GetAllElevators is a handler to get a slice of record(s) from elevators table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Elevators}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /elevators [get]
// http "https://xinqi.dev:8080/elevators?page=0&pagesize=20" X-Api-User:user123
func GetAllElevators(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "elevators", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllElevators(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetElevators is a function to get a single record from the elevators table in the rocket_development database
// @Summary Get record from table Elevators by  argID
// @Tags Elevators
// @ID argID
// @Description GetElevators is a function to get a single record from the elevators table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Elevators
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /elevators/{argID} [get]
// http "https://xinqi.dev:8080/elevators/1" X-Api-User:user123
func GetElevators(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "elevators", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetElevators(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddElevators add to add a single record to elevators table in the rocket_development database
// @Summary Add an record to elevators table
// @Description add to add a single record to elevators table in the rocket_development database
// @Tags Elevators
// @Accept  json
// @Produce  json
// @Param Elevators body model.Elevators true "Add Elevators"
// @Success 200 {object} model.Elevators
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /elevators [post]
// echo '{"column_id": 30,"id": 45,"serial_number": 58,"model": "IKeegHIuvBoTUWkunJYeyTPwf","type": "TLlGeHVqPOQPjHouPoQTJFPEb","status": "GiigtMWNUudJhaihkAGIPPYbo","commision_date": "2146-01-14T15:28:09.265315357-05:00","last_inspection_date": "2106-10-12T16:30:03.220650278-04:00","inspection_cert": "vCuAIHJWUiihyrhbZXMKvOxTF","information": "giJMxGSmLyClPRQtRfpuipEAN","notes": "PwxYMjffUHsHoyJQsgxyYOmMM","created_at": "2123-10-22T17:09:06.165564197-04:00","updated_at": "2100-08-06T05:34:08.697236723-04:00"}' | http POST "https://xinqi.dev:8080/elevators" X-Api-User:user123
func AddElevators(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	elevators := &model.Elevators{}

	if err := readJSON(r, elevators); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := elevators.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	elevators.Prepare()

	if err := elevators.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "elevators", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	elevators, _, err = dao.AddElevators(ctx, elevators)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, elevators)
}

// UpdateElevators Update a single record from elevators table in the rocket_development database
// @Summary Update an record in table elevators
// @Description Update a single record from elevators table in the rocket_development database
// @Tags Elevators
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Elevators body model.Elevators true "Update Elevators record"
// @Success 200 {object} model.Elevators
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /elevators/{argID} [put]
// echo '{"column_id": 30,"id": 45,"serial_number": 58,"model": "IKeegHIuvBoTUWkunJYeyTPwf","type": "TLlGeHVqPOQPjHouPoQTJFPEb","status": "GiigtMWNUudJhaihkAGIPPYbo","commision_date": "2146-01-14T15:28:09.265315357-05:00","last_inspection_date": "2106-10-12T16:30:03.220650278-04:00","inspection_cert": "vCuAIHJWUiihyrhbZXMKvOxTF","information": "giJMxGSmLyClPRQtRfpuipEAN","notes": "PwxYMjffUHsHoyJQsgxyYOmMM","created_at": "2123-10-22T17:09:06.165564197-04:00","updated_at": "2100-08-06T05:34:08.697236723-04:00"}' | http PUT "https://xinqi.dev:8080/elevators/1"  X-Api-User:user123
func UpdateElevators(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	elevators := &model.Elevators{}
	if err := readJSON(r, elevators); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := elevators.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	elevators.Prepare()

	if err := elevators.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "elevators", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	elevators, _, err = dao.UpdateElevators(ctx,
		argID,
		elevators)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, elevators)
}

// DeleteElevators Delete a single record from elevators table in the rocket_development database
// @Summary Delete a record from elevators
// @Description Delete a single record from elevators table in the rocket_development database
// @Tags Elevators
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Elevators
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /elevators/{argID} [delete]
// http DELETE "https://xinqi.dev:8080/elevators/1" X-Api-User:user123
func DeleteElevators(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "elevators", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteElevators(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
