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

func configBuildingsRouter(router *httprouter.Router) {
	router.GET("/buildings", GetAllBuildings)
	router.POST("/buildings", AddBuildings)
	router.GET("/buildings/:argID", GetBuildings)
	router.PUT("/buildings/:argID", UpdateBuildings)
	router.DELETE("/buildings/:argID", DeleteBuildings)
}

func configGinBuildingsRouter(router gin.IRoutes) {
	router.GET("/buildings", ConverHttprouterToGin(GetAllBuildings))
	router.POST("/buildings", ConverHttprouterToGin(AddBuildings))
	router.GET("/buildings/:argID", ConverHttprouterToGin(GetBuildings))
	router.PUT("/buildings/:argID", ConverHttprouterToGin(UpdateBuildings))
	router.DELETE("/buildings/:argID", ConverHttprouterToGin(DeleteBuildings))
}

// GetAllBuildings is a function to get a slice of record(s) from buildings table in the rocket_development database
// @Summary Get list of Buildings
// @Tags Buildings
// @Description GetAllBuildings is a handler to get a slice of record(s) from buildings table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Buildings}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildings [get]
// http "http://localhost:8080/buildings?page=0&pagesize=20" X-Api-User:user123
func GetAllBuildings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "buildings", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBuildings(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBuildings is a function to get a single record from the buildings table in the rocket_development database
// @Summary Get record from table Buildings by  argID
// @Tags Buildings
// @ID argID
// @Description GetBuildings is a function to get a single record from the buildings table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Buildings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /buildings/{argID} [get]
// http "http://localhost:8080/buildings/1" X-Api-User:user123
func GetBuildings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "buildings", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBuildings(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBuildings add to add a single record to buildings table in the rocket_development database
// @Summary Add an record to buildings table
// @Description add to add a single record to buildings table in the rocket_development database
// @Tags Buildings
// @Accept  json
// @Produce  json
// @Param Buildings body model.Buildings true "Add Buildings"
// @Success 200 {object} model.Buildings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildings [post]
// echo '{"customer_id": 84,"address_id": 69,"id": 35,"full_name_of_building_admin": "xQOHZKXqDuOUNlBPfVwgXfinM","email_of_admin_of_building": "eRbPrKXbxxNBBZfmBWwrirwmI","phone_num_of_building_admin": 45,"full_name_of_tech_contact_for_building": "bFoVEClufqWeoaiuxNQsRYmEK","tech_contact_email_for_building": "xhXvaBxNZyUyiWqlCEBmeYgRx","tech_contact_phone_for_building": 73,"created_at": "2168-01-02T23:41:54.743188052-05:00","updated_at": "2207-12-05T10:47:51.931699991-05:00"}' | http POST "http://localhost:8080/buildings" X-Api-User:user123
func AddBuildings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	buildings := &model.Buildings{}

	if err := readJSON(r, buildings); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := buildings.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	buildings.Prepare()

	if err := buildings.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "buildings", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	buildings, _, err = dao.AddBuildings(ctx, buildings)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, buildings)
}

// UpdateBuildings Update a single record from buildings table in the rocket_development database
// @Summary Update an record in table buildings
// @Description Update a single record from buildings table in the rocket_development database
// @Tags Buildings
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Buildings body model.Buildings true "Update Buildings record"
// @Success 200 {object} model.Buildings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /buildings/{argID} [put]
// echo '{"customer_id": 84,"address_id": 69,"id": 35,"full_name_of_building_admin": "xQOHZKXqDuOUNlBPfVwgXfinM","email_of_admin_of_building": "eRbPrKXbxxNBBZfmBWwrirwmI","phone_num_of_building_admin": 45,"full_name_of_tech_contact_for_building": "bFoVEClufqWeoaiuxNQsRYmEK","tech_contact_email_for_building": "xhXvaBxNZyUyiWqlCEBmeYgRx","tech_contact_phone_for_building": 73,"created_at": "2168-01-02T23:41:54.743188052-05:00","updated_at": "2207-12-05T10:47:51.931699991-05:00"}' | http PUT "http://localhost:8080/buildings/1"  X-Api-User:user123
func UpdateBuildings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	buildings := &model.Buildings{}
	if err := readJSON(r, buildings); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := buildings.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	buildings.Prepare()

	if err := buildings.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "buildings", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	buildings, _, err = dao.UpdateBuildings(ctx,
		argID,
		buildings)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, buildings)
}

// DeleteBuildings Delete a single record from buildings table in the rocket_development database
// @Summary Delete a record from buildings
// @Description Delete a single record from buildings table in the rocket_development database
// @Tags Buildings
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Buildings
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /buildings/{argID} [delete]
// http DELETE "http://localhost:8080/buildings/1" X-Api-User:user123
func DeleteBuildings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "buildings", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBuildings(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
