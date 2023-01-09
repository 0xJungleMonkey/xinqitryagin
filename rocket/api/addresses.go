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

func configAddressesRouter(router *httprouter.Router) {
	router.GET("/addresses", GetAllAddresses)
	router.POST("/addresses", AddAddresses)
	router.GET("/addresses/:argID", GetAddresses)
	router.PUT("/addresses/:argID", UpdateAddresses)
	router.DELETE("/addresses/:argID", DeleteAddresses)
}

func configGinAddressesRouter(router gin.IRoutes) {
	router.GET("/addresses", ConverHttprouterToGin(GetAllAddresses))
	router.POST("/addresses", ConverHttprouterToGin(AddAddresses))
	router.GET("/addresses/:argID", ConverHttprouterToGin(GetAddresses))
	router.PUT("/addresses/:argID", ConverHttprouterToGin(UpdateAddresses))
	router.DELETE("/addresses/:argID", ConverHttprouterToGin(DeleteAddresses))
}

// GetAllAddresses is a function to get a slice of record(s) from addresses table in the rocket_development database
// @Summary Get list of Addresses
// @Tags Addresses
// @Description GetAllAddresses is a handler to get a slice of record(s) from addresses table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Addresses}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /addresses [get]
// http "https://xinqi.dev:8080/addresses?page=0&pagesize=20" X-Api-User:user123
func GetAllAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "addresses", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAddresses(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetAddresses is a function to get a single record from the addresses table in the rocket_development database
// @Summary Get record from table Addresses by  argID
// @Tags Addresses
// @ID argID
// @Description GetAddresses is a function to get a single record from the addresses table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Addresses
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /addresses/{argID} [get]
// http "https://xinqi.dev:8080/addresses/1" X-Api-User:user123
func GetAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "addresses", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAddresses(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddAddresses add to add a single record to addresses table in the rocket_development database
// @Summary Add an record to addresses table
// @Description add to add a single record to addresses table in the rocket_development database
// @Tags Addresses
// @Accept  json
// @Produce  json
// @Param Addresses body model.Addresses true "Add Addresses"
// @Success 200 {object} model.Addresses
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /addresses [post]
// echo '{"id": 22,"address_type": "lQeQwkUejSrsnTBCVlNIpUOYp","status": "cPyqErFmGxpDQeWHZcgpyCruO","entity": "hlrmhtySIrXuJIsHkWBZdNgJG","number_and_street": "NuOqLcvfIpyajbWgGUWKGxMvR","suite_or_apartment": "TiLyHJTWWWuPXZNKgbPdnhGvI","city": "wBWXAvqMopyfaqoWNEPuUGhkR","postal_code": "TbvHpsJbOuJeliqaQNmMDgFRJ","country": "PIZJQViroOyJVRjDshJYtOgmV","notes": "bNChjHABfLmLvARXQQkpUhbhb","created_at": "2180-07-28T09:10:25.217150908-04:00","updated_at": "2178-07-04T01:29:04.490206665-04:00","latitude": 0.1873465,"longitude": 0.011193146}' | http POST "https://xinqi.dev:8080/addresses" X-Api-User:user123
func AddAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	addresses := &model.Addresses{}

	if err := readJSON(r, addresses); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := addresses.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	addresses.Prepare()

	if err := addresses.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "addresses", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	addresses, _, err = dao.AddAddresses(ctx, addresses)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, addresses)
}

// UpdateAddresses Update a single record from addresses table in the rocket_development database
// @Summary Update an record in table addresses
// @Description Update a single record from addresses table in the rocket_development database
// @Tags Addresses
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Addresses body model.Addresses true "Update Addresses record"
// @Success 200 {object} model.Addresses
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /addresses/{argID} [put]
// echo '{"id": 22,"address_type": "lQeQwkUejSrsnTBCVlNIpUOYp","status": "cPyqErFmGxpDQeWHZcgpyCruO","entity": "hlrmhtySIrXuJIsHkWBZdNgJG","number_and_street": "NuOqLcvfIpyajbWgGUWKGxMvR","suite_or_apartment": "TiLyHJTWWWuPXZNKgbPdnhGvI","city": "wBWXAvqMopyfaqoWNEPuUGhkR","postal_code": "TbvHpsJbOuJeliqaQNmMDgFRJ","country": "PIZJQViroOyJVRjDshJYtOgmV","notes": "bNChjHABfLmLvARXQQkpUhbhb","created_at": "2180-07-28T09:10:25.217150908-04:00","updated_at": "2178-07-04T01:29:04.490206665-04:00","latitude": 0.1873465,"longitude": 0.011193146}' | http PUT "https://xinqi.dev:8080/addresses/1"  X-Api-User:user123
func UpdateAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	addresses := &model.Addresses{}
	if err := readJSON(r, addresses); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := addresses.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	addresses.Prepare()

	if err := addresses.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "addresses", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	addresses, _, err = dao.UpdateAddresses(ctx,
		argID,
		addresses)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, addresses)
}

// DeleteAddresses Delete a single record from addresses table in the rocket_development database
// @Summary Delete a record from addresses
// @Description Delete a single record from addresses table in the rocket_development database
// @Tags Addresses
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Addresses
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /addresses/{argID} [delete]
// http DELETE "https://xinqi.dev:8080/addresses/1" X-Api-User:user123
func DeleteAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "addresses", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAddresses(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
