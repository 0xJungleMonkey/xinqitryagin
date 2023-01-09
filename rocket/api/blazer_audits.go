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

func configBlazerAuditsRouter(router *httprouter.Router) {
	router.GET("/blazeraudits", GetAllBlazerAudits)
	router.POST("/blazeraudits", AddBlazerAudits)
	router.GET("/blazeraudits/:argID", GetBlazerAudits)
	router.PUT("/blazeraudits/:argID", UpdateBlazerAudits)
	router.DELETE("/blazeraudits/:argID", DeleteBlazerAudits)
}

func configGinBlazerAuditsRouter(router gin.IRoutes) {
	router.GET("/blazeraudits", ConverHttprouterToGin(GetAllBlazerAudits))
	router.POST("/blazeraudits", ConverHttprouterToGin(AddBlazerAudits))
	router.GET("/blazeraudits/:argID", ConverHttprouterToGin(GetBlazerAudits))
	router.PUT("/blazeraudits/:argID", ConverHttprouterToGin(UpdateBlazerAudits))
	router.DELETE("/blazeraudits/:argID", ConverHttprouterToGin(DeleteBlazerAudits))
}

// GetAllBlazerAudits is a function to get a slice of record(s) from blazer_audits table in the rocket_development database
// @Summary Get list of BlazerAudits
// @Tags BlazerAudits
// @Description GetAllBlazerAudits is a handler to get a slice of record(s) from blazer_audits table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlazerAudits}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazeraudits [get]
// http "https://xinqi.dev:8080/blazeraudits?page=0&pagesize=20" X-Api-User:user123
func GetAllBlazerAudits(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "blazer_audits", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlazerAudits(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlazerAudits is a function to get a single record from the blazer_audits table in the rocket_development database
// @Summary Get record from table BlazerAudits by  argID
// @Tags BlazerAudits
// @ID argID
// @Description GetBlazerAudits is a function to get a single record from the blazer_audits table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.BlazerAudits
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blazeraudits/{argID} [get]
// http "https://xinqi.dev:8080/blazeraudits/1" X-Api-User:user123
func GetBlazerAudits(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_audits", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlazerAudits(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlazerAudits add to add a single record to blazer_audits table in the rocket_development database
// @Summary Add an record to blazer_audits table
// @Description add to add a single record to blazer_audits table in the rocket_development database
// @Tags BlazerAudits
// @Accept  json
// @Produce  json
// @Param BlazerAudits body model.BlazerAudits true "Add BlazerAudits"
// @Success 200 {object} model.BlazerAudits
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazeraudits [post]
// echo '{"id": 40,"user_id": 17,"query_id": 21,"statement": "KbFwQEcVbDCLofaMNpTbefjrq","data_source": "ujqLDYuYSPMBDDFPTvGXZJUyu","created_at": "2108-01-15T02:15:54.143424326-05:00"}' | http POST "https://xinqi.dev:8080/blazeraudits" X-Api-User:user123
func AddBlazerAudits(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blazeraudits := &model.BlazerAudits{}

	if err := readJSON(r, blazeraudits); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazeraudits.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazeraudits.Prepare()

	if err := blazeraudits.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_audits", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blazeraudits, _, err = dao.AddBlazerAudits(ctx, blazeraudits)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazeraudits)
}

// UpdateBlazerAudits Update a single record from blazer_audits table in the rocket_development database
// @Summary Update an record in table blazer_audits
// @Description Update a single record from blazer_audits table in the rocket_development database
// @Tags BlazerAudits
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  BlazerAudits body model.BlazerAudits true "Update BlazerAudits record"
// @Success 200 {object} model.BlazerAudits
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blazeraudits/{argID} [put]
// echo '{"id": 40,"user_id": 17,"query_id": 21,"statement": "KbFwQEcVbDCLofaMNpTbefjrq","data_source": "ujqLDYuYSPMBDDFPTvGXZJUyu","created_at": "2108-01-15T02:15:54.143424326-05:00"}' | http PUT "https://xinqi.dev:8080/blazeraudits/1"  X-Api-User:user123
func UpdateBlazerAudits(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazeraudits := &model.BlazerAudits{}
	if err := readJSON(r, blazeraudits); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blazeraudits.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blazeraudits.Prepare()

	if err := blazeraudits.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_audits", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blazeraudits, _, err = dao.UpdateBlazerAudits(ctx,
		argID,
		blazeraudits)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blazeraudits)
}

// DeleteBlazerAudits Delete a single record from blazer_audits table in the rocket_development database
// @Summary Delete a record from blazer_audits
// @Description Delete a single record from blazer_audits table in the rocket_development database
// @Tags BlazerAudits
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.BlazerAudits
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blazeraudits/{argID} [delete]
// http DELETE "https://xinqi.dev:8080/blazeraudits/1" X-Api-User:user123
func DeleteBlazerAudits(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blazer_audits", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlazerAudits(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
