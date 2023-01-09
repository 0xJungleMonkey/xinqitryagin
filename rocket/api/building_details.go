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

func configBuildingDetailsRouter(router *httprouter.Router) {
	router.GET("/buildingdetails", GetAllBuildingDetails)
	router.POST("/buildingdetails", AddBuildingDetails)
	router.GET("/buildingdetails/:argID", GetBuildingDetails)
	router.PUT("/buildingdetails/:argID", UpdateBuildingDetails)
	router.DELETE("/buildingdetails/:argID", DeleteBuildingDetails)
}

func configGinBuildingDetailsRouter(router gin.IRoutes) {
	router.GET("/buildingdetails", ConverHttprouterToGin(GetAllBuildingDetails))
	router.POST("/buildingdetails", ConverHttprouterToGin(AddBuildingDetails))
	router.GET("/buildingdetails/:argID", ConverHttprouterToGin(GetBuildingDetails))
	router.PUT("/buildingdetails/:argID", ConverHttprouterToGin(UpdateBuildingDetails))
	router.DELETE("/buildingdetails/:argID", ConverHttprouterToGin(DeleteBuildingDetails))
}

// GetAllBuildingDetails is a function to get a slice of record(s) from building_details table in the rocket_development database
// @Summary Get list of BuildingDetails
// @Tags BuildingDetails
// @Description GetAllBuildingDetails is a handler to get a slice of record(s) from building_details table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BuildingDetails}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildingdetails [get]
// http "https://xinqi.dev:443/buildingdetails?page=0&pagesize=20" X-Api-User:user123
func GetAllBuildingDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "building_details", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBuildingDetails(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBuildingDetails is a function to get a single record from the building_details table in the rocket_development database
// @Summary Get record from table BuildingDetails by  argID
// @Tags BuildingDetails
// @ID argID
// @Description GetBuildingDetails is a function to get a single record from the building_details table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BuildingDetails
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /buildingdetails/{argID} [get]
// http "https://xinqi.dev:443/buildingdetails/1" X-Api-User:user123
func GetBuildingDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "building_details", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBuildingDetails(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBuildingDetails add to add a single record to building_details table in the rocket_development database
// @Summary Add an record to building_details table
// @Description add to add a single record to building_details table in the rocket_development database
// @Tags BuildingDetails
// @Accept  json
// @Produce  json
// @Param BuildingDetails body model.BuildingDetails true "Add BuildingDetails"
// @Success 200 {object} model.BuildingDetails
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildingdetails [post]
// echo '{"building_id": 51,"id": 36,"information_key": "JDcukNVWiVBVRLAabTgfuuVnY","value": "HGdfSgMXNRgJFwmsImBCCKTsD","created_at": "2104-10-24T00:15:37.502813235-04:00","updated_at": "2206-05-21T03:09:44.430668259-04:00"}' | http POST "https://xinqi.dev:443/buildingdetails" X-Api-User:user123
func AddBuildingDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	buildingdetails := &model.BuildingDetails{}

	if err := readJSON(r, buildingdetails); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := buildingdetails.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	buildingdetails.Prepare()

	if err := buildingdetails.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "building_details", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	buildingdetails, _, err = dao.AddBuildingDetails(ctx, buildingdetails)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, buildingdetails)
}

// UpdateBuildingDetails Update a single record from building_details table in the rocket_development database
// @Summary Update an record in table building_details
// @Description Update a single record from building_details table in the rocket_development database
// @Tags BuildingDetails
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BuildingDetails body model.BuildingDetails true "Update BuildingDetails record"
// @Success 200 {object} model.BuildingDetails
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildingdetails/{argID} [put]
// echo '{"building_id": 51,"id": 36,"information_key": "JDcukNVWiVBVRLAabTgfuuVnY","value": "HGdfSgMXNRgJFwmsImBCCKTsD","created_at": "2104-10-24T00:15:37.502813235-04:00","updated_at": "2206-05-21T03:09:44.430668259-04:00"}' | http PUT "https://xinqi.dev:443/buildingdetails/1"  X-Api-User:user123
func UpdateBuildingDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	buildingdetails := &model.BuildingDetails{}
	if err := readJSON(r, buildingdetails); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := buildingdetails.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	buildingdetails.Prepare()

	if err := buildingdetails.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "building_details", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	buildingdetails, _, err = dao.UpdateBuildingDetails(ctx,
		argID,
		buildingdetails)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, buildingdetails)
}

// DeleteBuildingDetails Delete a single record from building_details table in the rocket_development database
// @Summary Delete a record from building_details
// @Description Delete a single record from building_details table in the rocket_development database
// @Tags BuildingDetails
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BuildingDetails
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /buildingdetails/{argID} [delete]
// http DELETE "https://xinqi.dev:443/buildingdetails/1" X-Api-User:user123
func DeleteBuildingDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "building_details", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBuildingDetails(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
