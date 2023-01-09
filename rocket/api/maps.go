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

func configMapsRouter(router *httprouter.Router) {
	router.GET("/maps", GetAllMaps)
	router.POST("/maps", AddMaps)
	router.GET("/maps/:argID", GetMaps)
	router.PUT("/maps/:argID", UpdateMaps)
	router.DELETE("/maps/:argID", DeleteMaps)
}

func configGinMapsRouter(router gin.IRoutes) {
	router.GET("/maps", ConverHttprouterToGin(GetAllMaps))
	router.POST("/maps", ConverHttprouterToGin(AddMaps))
	router.GET("/maps/:argID", ConverHttprouterToGin(GetMaps))
	router.PUT("/maps/:argID", ConverHttprouterToGin(UpdateMaps))
	router.DELETE("/maps/:argID", ConverHttprouterToGin(DeleteMaps))
}

// GetAllMaps is a function to get a slice of record(s) from maps table in the rocket_development database
// @Summary Get list of Maps
// @Tags Maps
// @Description GetAllMaps is a handler to get a slice of record(s) from maps table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Maps}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /maps [get]
// http "http://localhost:8080/maps?page=0&pagesize=20" X-Api-User:user123
func GetAllMaps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "maps", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllMaps(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetMaps is a function to get a single record from the maps table in the rocket_development database
// @Summary Get record from table Maps by  argID
// @Tags Maps
// @ID argID
// @Description GetMaps is a function to get a single record from the maps table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Maps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /maps/{argID} [get]
// http "http://localhost:8080/maps/1" X-Api-User:user123
func GetMaps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "maps", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetMaps(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddMaps add to add a single record to maps table in the rocket_development database
// @Summary Add an record to maps table
// @Description add to add a single record to maps table in the rocket_development database
// @Tags Maps
// @Accept  json
// @Produce  json
// @Param Maps body model.Maps true "Add Maps"
// @Success 200 {object} model.Maps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /maps [post]
// echo '{"id": 78,"created_at": "2051-05-02T14:01:46.153259597-04:00","updated_at": "2192-06-03T01:32:49.142357321-04:00"}' | http POST "http://localhost:8080/maps" X-Api-User:user123
func AddMaps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	maps := &model.Maps{}

	if err := readJSON(r, maps); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := maps.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	maps.Prepare()

	if err := maps.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "maps", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	maps, _, err = dao.AddMaps(ctx, maps)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, maps)
}

// UpdateMaps Update a single record from maps table in the rocket_development database
// @Summary Update an record in table maps
// @Description Update a single record from maps table in the rocket_development database
// @Tags Maps
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Maps body model.Maps true "Update Maps record"
// @Success 200 {object} model.Maps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /maps/{argID} [put]
// echo '{"id": 78,"created_at": "2051-05-02T14:01:46.153259597-04:00","updated_at": "2192-06-03T01:32:49.142357321-04:00"}' | http PUT "http://localhost:8080/maps/1"  X-Api-User:user123
func UpdateMaps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	maps := &model.Maps{}
	if err := readJSON(r, maps); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := maps.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	maps.Prepare()

	if err := maps.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "maps", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	maps, _, err = dao.UpdateMaps(ctx,
		argID,
		maps)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, maps)
}

// DeleteMaps Delete a single record from maps table in the rocket_development database
// @Summary Delete a record from maps
// @Description Delete a single record from maps table in the rocket_development database
// @Tags Maps
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Maps
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /maps/{argID} [delete]
// http DELETE "http://localhost:8080/maps/1" X-Api-User:user123
func DeleteMaps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "maps", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteMaps(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
