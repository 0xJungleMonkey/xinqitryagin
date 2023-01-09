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

func configArInternalMetadataRouter(router *httprouter.Router) {
	router.GET("/arinternalmetadata", GetAllArInternalMetadata)
	router.POST("/arinternalmetadata", AddArInternalMetadata)
	router.GET("/arinternalmetadata/:argKey", GetArInternalMetadata)
	router.PUT("/arinternalmetadata/:argKey", UpdateArInternalMetadata)
	router.DELETE("/arinternalmetadata/:argKey", DeleteArInternalMetadata)
}

func configGinArInternalMetadataRouter(router gin.IRoutes) {
	router.GET("/arinternalmetadata", ConverHttprouterToGin(GetAllArInternalMetadata))
	router.POST("/arinternalmetadata", ConverHttprouterToGin(AddArInternalMetadata))
	router.GET("/arinternalmetadata/:argKey", ConverHttprouterToGin(GetArInternalMetadata))
	router.PUT("/arinternalmetadata/:argKey", ConverHttprouterToGin(UpdateArInternalMetadata))
	router.DELETE("/arinternalmetadata/:argKey", ConverHttprouterToGin(DeleteArInternalMetadata))
}

// GetAllArInternalMetadata is a function to get a slice of record(s) from ar_internal_metadata table in the rocket_development database
// @Summary Get list of ArInternalMetadata
// @Tags ArInternalMetadata
// @Description GetAllArInternalMetadata is a handler to get a slice of record(s) from ar_internal_metadata table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ArInternalMetadata}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /arinternalmetadata [get]
// http "http://localhost:8080/arinternalmetadata?page=0&pagesize=20" X-Api-User:user123
func GetAllArInternalMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "ar_internal_metadata", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllArInternalMetadata(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetArInternalMetadata is a function to get a single record from the ar_internal_metadata table in the rocket_development database
// @Summary Get record from table ArInternalMetadata by  argKey
// @Tags ArInternalMetadata
// @ID argKey
// @Description GetArInternalMetadata is a function to get a single record from the ar_internal_metadata table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argKey path string true "key"
// @Success 200 {object} model.ArInternalMetadata
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /arinternalmetadata/{argKey} [get]
// http "http://localhost:8080/arinternalmetadata/hello world" X-Api-User:user123
func GetArInternalMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argKey, err := parseString(ps, "argKey")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "ar_internal_metadata", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetArInternalMetadata(ctx, argKey)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddArInternalMetadata add to add a single record to ar_internal_metadata table in the rocket_development database
// @Summary Add an record to ar_internal_metadata table
// @Description add to add a single record to ar_internal_metadata table in the rocket_development database
// @Tags ArInternalMetadata
// @Accept  json
// @Produce  json
// @Param ArInternalMetadata body model.ArInternalMetadata true "Add ArInternalMetadata"
// @Success 200 {object} model.ArInternalMetadata
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /arinternalmetadata [post]
// echo '{"key": "dKUlQbIKZeTnxsSQlctjCTAjC","value": "MXUJDIUtCiRjnnFWOTrrRReod","created_at": "2122-09-13T01:38:12.718168922-04:00","updated_at": "2128-11-09T15:14:40.438021339-05:00"}' | http POST "http://localhost:8080/arinternalmetadata" X-Api-User:user123
func AddArInternalMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	arinternalmetadata := &model.ArInternalMetadata{}

	if err := readJSON(r, arinternalmetadata); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := arinternalmetadata.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	arinternalmetadata.Prepare()

	if err := arinternalmetadata.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "ar_internal_metadata", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	arinternalmetadata, _, err = dao.AddArInternalMetadata(ctx, arinternalmetadata)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, arinternalmetadata)
}

// UpdateArInternalMetadata Update a single record from ar_internal_metadata table in the rocket_development database
// @Summary Update an record in table ar_internal_metadata
// @Description Update a single record from ar_internal_metadata table in the rocket_development database
// @Tags ArInternalMetadata
// @Accept  json
// @Produce  json
// @Param  argKey path string true "key"
// @Param  ArInternalMetadata body model.ArInternalMetadata true "Update ArInternalMetadata record"
// @Success 200 {object} model.ArInternalMetadata
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /arinternalmetadata/{argKey} [put]
// echo '{"key": "dKUlQbIKZeTnxsSQlctjCTAjC","value": "MXUJDIUtCiRjnnFWOTrrRReod","created_at": "2122-09-13T01:38:12.718168922-04:00","updated_at": "2128-11-09T15:14:40.438021339-05:00"}' | http PUT "http://localhost:8080/arinternalmetadata/hello world"  X-Api-User:user123
func UpdateArInternalMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argKey, err := parseString(ps, "argKey")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	arinternalmetadata := &model.ArInternalMetadata{}
	if err := readJSON(r, arinternalmetadata); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := arinternalmetadata.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	arinternalmetadata.Prepare()

	if err := arinternalmetadata.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "ar_internal_metadata", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	arinternalmetadata, _, err = dao.UpdateArInternalMetadata(ctx,
		argKey,
		arinternalmetadata)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, arinternalmetadata)
}

// DeleteArInternalMetadata Delete a single record from ar_internal_metadata table in the rocket_development database
// @Summary Delete a record from ar_internal_metadata
// @Description Delete a single record from ar_internal_metadata table in the rocket_development database
// @Tags ArInternalMetadata
// @Accept  json
// @Produce  json
// @Param  argKey path string true "key"
// @Success 204 {object} model.ArInternalMetadata
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /arinternalmetadata/{argKey} [delete]
// http DELETE "http://localhost:8080/arinternalmetadata/hello world" X-Api-User:user123
func DeleteArInternalMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argKey, err := parseString(ps, "argKey")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "ar_internal_metadata", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteArInternalMetadata(ctx, argKey)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
