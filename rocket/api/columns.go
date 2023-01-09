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

func configColumnsRouter(router *httprouter.Router) {
	router.GET("/columns", GetAllColumns)
	router.POST("/columns", AddColumns)
	router.GET("/columns/:argID", GetColumns)
	router.PUT("/columns/:argID", UpdateColumns)
	router.DELETE("/columns/:argID", DeleteColumns)
}

func configGinColumnsRouter(router gin.IRoutes) {
	router.GET("/columns", ConverHttprouterToGin(GetAllColumns))
	router.POST("/columns", ConverHttprouterToGin(AddColumns))
	router.GET("/columns/:argID", ConverHttprouterToGin(GetColumns))
	router.PUT("/columns/:argID", ConverHttprouterToGin(UpdateColumns))
	router.DELETE("/columns/:argID", ConverHttprouterToGin(DeleteColumns))
}

// GetAllColumns is a function to get a slice of record(s) from columns table in the rocket_development database
// @Summary Get list of Columns
// @Tags Columns
// @Description GetAllColumns is a handler to get a slice of record(s) from columns table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Columns}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /columns [get]
// http "http://localhost:8080/columns?page=0&pagesize=20" X-Api-User:user123
func GetAllColumns(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "columns", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllColumns(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetColumns is a function to get a single record from the columns table in the rocket_development database
// @Summary Get record from table Columns by  argID
// @Tags Columns
// @ID argID
// @Description GetColumns is a function to get a single record from the columns table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Columns
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /columns/{argID} [get]
// http "http://localhost:8080/columns/1" X-Api-User:user123
func GetColumns(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "columns", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetColumns(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddColumns add to add a single record to columns table in the rocket_development database
// @Summary Add an record to columns table
// @Description add to add a single record to columns table in the rocket_development database
// @Tags Columns
// @Accept  json
// @Produce  json
// @Param Columns body model.Columns true "Add Columns"
// @Success 200 {object} model.Columns
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /columns [post]
// echo '{"battery_id": 86,"id": 97,"type": "PxrsYgZIcgcTxeXDvkUlLXTuL","num_of_floors_served": 92,"status": "GjkFXAlWapXkuopYwXQpwyyJs","information": "fJsNjkFIvMfvhSUkXeJVjmZdU","notes": "GbTYAlQKWShJIMeNnmsHnMjrV","created_at": "2234-12-01T20:46:52.822820591-05:00","updated_at": "2129-10-11T00:57:53.705769736-04:00"}' | http POST "http://localhost:8080/columns" X-Api-User:user123
func AddColumns(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	columns := &model.Columns{}

	if err := readJSON(r, columns); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := columns.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	columns.Prepare()

	if err := columns.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "columns", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	columns, _, err = dao.AddColumns(ctx, columns)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, columns)
}

// UpdateColumns Update a single record from columns table in the rocket_development database
// @Summary Update an record in table columns
// @Description Update a single record from columns table in the rocket_development database
// @Tags Columns
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Columns body model.Columns true "Update Columns record"
// @Success 200 {object} model.Columns
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /columns/{argID} [put]
// echo '{"battery_id": 86,"id": 97,"type": "PxrsYgZIcgcTxeXDvkUlLXTuL","num_of_floors_served": 92,"status": "GjkFXAlWapXkuopYwXQpwyyJs","information": "fJsNjkFIvMfvhSUkXeJVjmZdU","notes": "GbTYAlQKWShJIMeNnmsHnMjrV","created_at": "2234-12-01T20:46:52.822820591-05:00","updated_at": "2129-10-11T00:57:53.705769736-04:00"}' | http PUT "http://localhost:8080/columns/1"  X-Api-User:user123
func UpdateColumns(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	columns := &model.Columns{}
	if err := readJSON(r, columns); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := columns.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	columns.Prepare()

	if err := columns.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "columns", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	columns, _, err = dao.UpdateColumns(ctx,
		argID,
		columns)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, columns)
}

// DeleteColumns Delete a single record from columns table in the rocket_development database
// @Summary Delete a record from columns
// @Description Delete a single record from columns table in the rocket_development database
// @Tags Columns
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Columns
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /columns/{argID} [delete]
// http DELETE "http://localhost:8080/columns/1" X-Api-User:user123
func DeleteColumns(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "columns", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteColumns(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
