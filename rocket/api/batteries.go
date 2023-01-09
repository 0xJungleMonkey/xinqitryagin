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

func configBatteriesRouter(router *httprouter.Router) {
	router.GET("/batteries", GetAllBatteries)
	router.POST("/batteries", AddBatteries)
	router.GET("/batteries/:argID", GetBatteries)
	router.PUT("/batteries/:argID", UpdateBatteries)
	router.DELETE("/batteries/:argID", DeleteBatteries)
}

func configGinBatteriesRouter(router gin.IRoutes) {
	router.GET("/batteries", ConverHttprouterToGin(GetAllBatteries))
	router.POST("/batteries", ConverHttprouterToGin(AddBatteries))
	router.GET("/batteries/:argID", ConverHttprouterToGin(GetBatteries))
	router.PUT("/batteries/:argID", ConverHttprouterToGin(UpdateBatteries))
	router.DELETE("/batteries/:argID", ConverHttprouterToGin(DeleteBatteries))
}

// GetAllBatteries is a function to get a slice of record(s) from batteries table in the rocket_development database
// @Summary Get list of Batteries
// @Tags Batteries
// @Description GetAllBatteries is a handler to get a slice of record(s) from batteries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Batteries}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /batteries [get]
// http "http://localhost:8080/batteries?page=0&pagesize=20" X-Api-User:user123
func GetAllBatteries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "batteries", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBatteries(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBatteries is a function to get a single record from the batteries table in the rocket_development database
// @Summary Get record from table Batteries by  argID
// @Tags Batteries
// @ID argID
// @Description GetBatteries is a function to get a single record from the batteries table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Batteries
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /batteries/{argID} [get]
// http "http://localhost:8080/batteries/1" X-Api-User:user123
func GetBatteries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "batteries", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBatteries(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBatteries add to add a single record to batteries table in the rocket_development database
// @Summary Add an record to batteries table
// @Description add to add a single record to batteries table in the rocket_development database
// @Tags Batteries
// @Accept  json
// @Produce  json
// @Param Batteries body model.Batteries true "Add Batteries"
// @Success 200 {object} model.Batteries
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /batteries [post]
// echo '{"employee_id": 5,"building_id": 55,"id": 96,"type": "TghWZfDcVEQwpQqeOuPWuiswo","status": "gNsqnBwMLFwtUJaNeLmFGRiWg","commission_date": "2055-07-20T14:52:13.733091611-04:00","last_inspection_date": "2139-11-30T08:11:43.479907759-05:00","operations_cert": "bfAGtoMPNKdJDHIFGCNQeUMcY","information": "uTaCLmHosbHaBMKRHlWDxZemk","notes": "wprKksnbraZXMqthbWiUFJFiM","created_at": "2228-11-12T16:57:16.615003887-05:00","updated_at": "2188-03-30T18:33:29.152735497-04:00"}' | http POST "http://localhost:8080/batteries" X-Api-User:user123
func AddBatteries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	batteries := &model.Batteries{}

	if err := readJSON(r, batteries); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := batteries.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	batteries.Prepare()

	if err := batteries.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "batteries", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	batteries, _, err = dao.AddBatteries(ctx, batteries)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, batteries)
}

// UpdateBatteries Update a single record from batteries table in the rocket_development database
// @Summary Update an record in table batteries
// @Description Update a single record from batteries table in the rocket_development database
// @Tags Batteries
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Batteries body model.Batteries true "Update Batteries record"
// @Success 200 {object} model.Batteries
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /batteries/{argID} [put]
// echo '{"employee_id": 5,"building_id": 55,"id": 96,"type": "TghWZfDcVEQwpQqeOuPWuiswo","status": "gNsqnBwMLFwtUJaNeLmFGRiWg","commission_date": "2055-07-20T14:52:13.733091611-04:00","last_inspection_date": "2139-11-30T08:11:43.479907759-05:00","operations_cert": "bfAGtoMPNKdJDHIFGCNQeUMcY","information": "uTaCLmHosbHaBMKRHlWDxZemk","notes": "wprKksnbraZXMqthbWiUFJFiM","created_at": "2228-11-12T16:57:16.615003887-05:00","updated_at": "2188-03-30T18:33:29.152735497-04:00"}' | http PUT "http://localhost:8080/batteries/1"  X-Api-User:user123
func UpdateBatteries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	batteries := &model.Batteries{}
	if err := readJSON(r, batteries); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := batteries.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	batteries.Prepare()

	if err := batteries.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "batteries", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	batteries, _, err = dao.UpdateBatteries(ctx,
		argID,
		batteries)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, batteries)
}

// DeleteBatteries Delete a single record from batteries table in the rocket_development database
// @Summary Delete a record from batteries
// @Description Delete a single record from batteries table in the rocket_development database
// @Tags Batteries
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Batteries
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /batteries/{argID} [delete]
// http DELETE "http://localhost:8080/batteries/1" X-Api-User:user123
func DeleteBatteries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "batteries", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBatteries(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
