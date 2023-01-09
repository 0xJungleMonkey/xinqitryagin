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

func configQuotesRouter(router *httprouter.Router) {
	router.GET("/quotes", GetAllQuotes)
	router.POST("/quotes", AddQuotes)
	router.GET("/quotes/:argID", GetQuotes)
	router.PUT("/quotes/:argID", UpdateQuotes)
	router.DELETE("/quotes/:argID", DeleteQuotes)
}

func configGinQuotesRouter(router gin.IRoutes) {
	router.GET("/quotes", ConverHttprouterToGin(GetAllQuotes))
	router.POST("/quotes", ConverHttprouterToGin(AddQuotes))
	router.GET("/quotes/:argID", ConverHttprouterToGin(GetQuotes))
	router.PUT("/quotes/:argID", ConverHttprouterToGin(UpdateQuotes))
	router.DELETE("/quotes/:argID", ConverHttprouterToGin(DeleteQuotes))
}

// GetAllQuotes is a function to get a slice of record(s) from quotes table in the rocket_development database
// @Summary Get list of Quotes
// @Tags Quotes
// @Description GetAllQuotes is a handler to get a slice of record(s) from quotes table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Quotes}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /quotes [get]
// http "https://xinqi.dev:8080/quotes?page=0&pagesize=20" X-Api-User:user123
func GetAllQuotes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "quotes", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllQuotes(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetQuotes is a function to get a single record from the quotes table in the rocket_development database
// @Summary Get record from table Quotes by  argID
// @Tags Quotes
// @ID argID
// @Description GetQuotes is a function to get a single record from the quotes table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Quotes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /quotes/{argID} [get]
// http "https://xinqi.dev:8080/quotes/1" X-Api-User:user123
func GetQuotes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "quotes", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetQuotes(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddQuotes add to add a single record to quotes table in the rocket_development database
// @Summary Add an record to quotes table
// @Description add to add a single record to quotes table in the rocket_development database
// @Tags Quotes
// @Accept  json
// @Produce  json
// @Param Quotes body model.Quotes true "Add Quotes"
// @Success 200 {object} model.Quotes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /quotes [post]
// echo '{"id": 70,"building_type": "wjLmFbsXZSYYxlDUHjncHvdLg","service_quality": "bUjriwnVTiPSgZrJDaRKdpveB","number_of_apartments": "ZOqIlVFppxuekerppZyKaKKRe","number_of_floors": "rudeSmVXwAvPACKOqEhhvVCJE","number_of_businesses": "UwokCfoRChGSYMlZfsPdYTQKa","number_of_basements": "KHQDIdNqmwRxLGcccCkfHBkGy","number_of_parking": "NgpAydvmanheyloRwgKoICHIo","number_of_cages": "lhIJseMaEHFSxoxFUZPhLKhVW","number_of_occupants": "tRUJEaaaXKDWUcmGjpmhWOkRj","number_of_hours": "dMwGmcchfrkjfdZjncfKfKUme","number_of_elevators_needed": "tJBqcTLFlmDISESnNEKldeEYu","price_per_unit": "bLWfTkcyaJAXQfqcYWNWQPxvG","elevator_price": "OvcQachhslgLtynQAWwuHpAvc","installation_fee": "ejQHLvXiYLpbHkUsRyvqBgMSr","final_price": "FlFJlnxCnPTfgAnIuTgBEAYqg","created_at": "2125-10-02T03:47:51.649898761-04:00","updated_at": "2158-11-27T22:08:01.381669191-05:00","name": "MOkmMRhOmYCYKxfafGkGnWcrr","company_name": "jqHFPkWSUZtKnxZfNSeXdPOfB","email": "qEJlMHCdqXGaZfuHDARrGoiyv","phone": "MXReGpAwoAWeGXbNSvpSVjxNf","department": "rnAJcEqayUFPOrEhcbXoiaBws","project_name": "YAJoQEbjjBbeCoPSMlavjGgaB","project_description": "EiwmESTxbDAEAglFBIQSfjNGc"}' | http POST "https://xinqi.dev:8080/quotes" X-Api-User:user123
func AddQuotes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	quotes := &model.Quotes{}

	if err := readJSON(r, quotes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := quotes.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	quotes.Prepare()

	if err := quotes.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "quotes", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	quotes, _, err = dao.AddQuotes(ctx, quotes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, quotes)
}

// UpdateQuotes Update a single record from quotes table in the rocket_development database
// @Summary Update an record in table quotes
// @Description Update a single record from quotes table in the rocket_development database
// @Tags Quotes
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Quotes body model.Quotes true "Update Quotes record"
// @Success 200 {object} model.Quotes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /quotes/{argID} [put]
// echo '{"id": 70,"building_type": "wjLmFbsXZSYYxlDUHjncHvdLg","service_quality": "bUjriwnVTiPSgZrJDaRKdpveB","number_of_apartments": "ZOqIlVFppxuekerppZyKaKKRe","number_of_floors": "rudeSmVXwAvPACKOqEhhvVCJE","number_of_businesses": "UwokCfoRChGSYMlZfsPdYTQKa","number_of_basements": "KHQDIdNqmwRxLGcccCkfHBkGy","number_of_parking": "NgpAydvmanheyloRwgKoICHIo","number_of_cages": "lhIJseMaEHFSxoxFUZPhLKhVW","number_of_occupants": "tRUJEaaaXKDWUcmGjpmhWOkRj","number_of_hours": "dMwGmcchfrkjfdZjncfKfKUme","number_of_elevators_needed": "tJBqcTLFlmDISESnNEKldeEYu","price_per_unit": "bLWfTkcyaJAXQfqcYWNWQPxvG","elevator_price": "OvcQachhslgLtynQAWwuHpAvc","installation_fee": "ejQHLvXiYLpbHkUsRyvqBgMSr","final_price": "FlFJlnxCnPTfgAnIuTgBEAYqg","created_at": "2125-10-02T03:47:51.649898761-04:00","updated_at": "2158-11-27T22:08:01.381669191-05:00","name": "MOkmMRhOmYCYKxfafGkGnWcrr","company_name": "jqHFPkWSUZtKnxZfNSeXdPOfB","email": "qEJlMHCdqXGaZfuHDARrGoiyv","phone": "MXReGpAwoAWeGXbNSvpSVjxNf","department": "rnAJcEqayUFPOrEhcbXoiaBws","project_name": "YAJoQEbjjBbeCoPSMlavjGgaB","project_description": "EiwmESTxbDAEAglFBIQSfjNGc"}' | http PUT "https://xinqi.dev:8080/quotes/1"  X-Api-User:user123
func UpdateQuotes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	quotes := &model.Quotes{}
	if err := readJSON(r, quotes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := quotes.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	quotes.Prepare()

	if err := quotes.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "quotes", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	quotes, _, err = dao.UpdateQuotes(ctx,
		argID,
		quotes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, quotes)
}

// DeleteQuotes Delete a single record from quotes table in the rocket_development database
// @Summary Delete a record from quotes
// @Description Delete a single record from quotes table in the rocket_development database
// @Tags Quotes
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Quotes
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /quotes/{argID} [delete]
// http DELETE "https://xinqi.dev:8080/quotes/1" X-Api-User:user123
func DeleteQuotes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "quotes", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteQuotes(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
