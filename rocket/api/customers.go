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

func configCustomersRouter(router *httprouter.Router) {
	router.GET("/customers", GetAllCustomers)
	router.POST("/customers", AddCustomers)
	router.GET("/customers/:argID", GetCustomers)
	router.PUT("/customers/:argID", UpdateCustomers)
	router.DELETE("/customers/:argID", DeleteCustomers)
}

func configGinCustomersRouter(router gin.IRoutes) {
	router.GET("/customers", ConverHttprouterToGin(GetAllCustomers))
	router.POST("/customers", ConverHttprouterToGin(AddCustomers))
	router.GET("/customers/:argID", ConverHttprouterToGin(GetCustomers))
	router.PUT("/customers/:argID", ConverHttprouterToGin(UpdateCustomers))
	router.DELETE("/customers/:argID", ConverHttprouterToGin(DeleteCustomers))
}

// GetAllCustomers is a function to get a slice of record(s) from customers table in the rocket_development database
// @Summary Get list of Customers
// @Tags Customers
// @Description GetAllCustomers is a handler to get a slice of record(s) from customers table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Customers}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /customers [get]
// http "https://xinqi.dev:8080/customers?page=0&pagesize=20" X-Api-User:user123
func GetAllCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "customers", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllCustomers(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetCustomers is a function to get a single record from the customers table in the rocket_development database
// @Summary Get record from table Customers by  argID
// @Tags Customers
// @ID argID
// @Description GetCustomers is a function to get a single record from the customers table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Customers
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /customers/{argID} [get]
// http "https://xinqi.dev:8080/customers/1" X-Api-User:user123
func GetCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "customers", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetCustomers(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddCustomers add to add a single record to customers table in the rocket_development database
// @Summary Add an record to customers table
// @Description add to add a single record to customers table in the rocket_development database
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param Customers body model.Customers true "Add Customers"
// @Success 200 {object} model.Customers
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /customers [post]
// echo '{"address_id": 57,"user_id": 47,"id": 63,"customer_creation_date": "ECdKQORSTJjnjamVrZSrDeykL","date": "JLRvxXkebNxsvOjOMAoGjyXmA","company_name": "srYriNTciPbtvlmtlhPhyGFMH","company_hq_adress": "xxfLtgwPIsPkLpTbJrbWMoZck","full_name_of_company_contact": "LenoyyCFloMkyhceHBByWuHIV","company_contact_phone": "tjDmQtUNWlcLUwvgKdxQtuiWR","company_contact_e_mail": "JTUNvyAbFPgUyigOclpXtnvaK","company_desc": "ukkBGQUOucJkIWRbbfUSShugl","full_name_service_tech_auth": "WKBUNqFJQkpDQnUfZHGJjecOm","tech_auth_phone_service": "sjNbYheEmKsjVwHEArrYnqDfw","tech_manager_email_service": "oLHbYFBfVRFuctlAZmuSElmTI","created_at": "2108-06-15T03:46:49.633189344-04:00","updated_at": "2278-10-14T13:03:29.305006086-04:00"}' | http POST "https://xinqi.dev:8080/customers" X-Api-User:user123
func AddCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	customers := &model.Customers{}

	if err := readJSON(r, customers); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := customers.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	customers.Prepare()

	if err := customers.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "customers", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	customers, _, err = dao.AddCustomers(ctx, customers)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, customers)
}

// UpdateCustomers Update a single record from customers table in the rocket_development database
// @Summary Update an record in table customers
// @Description Update a single record from customers table in the rocket_development database
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Customers body model.Customers true "Update Customers record"
// @Success 200 {object} model.Customers
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /customers/{argID} [put]
// echo '{"address_id": 57,"user_id": 47,"id": 63,"customer_creation_date": "ECdKQORSTJjnjamVrZSrDeykL","date": "JLRvxXkebNxsvOjOMAoGjyXmA","company_name": "srYriNTciPbtvlmtlhPhyGFMH","company_hq_adress": "xxfLtgwPIsPkLpTbJrbWMoZck","full_name_of_company_contact": "LenoyyCFloMkyhceHBByWuHIV","company_contact_phone": "tjDmQtUNWlcLUwvgKdxQtuiWR","company_contact_e_mail": "JTUNvyAbFPgUyigOclpXtnvaK","company_desc": "ukkBGQUOucJkIWRbbfUSShugl","full_name_service_tech_auth": "WKBUNqFJQkpDQnUfZHGJjecOm","tech_auth_phone_service": "sjNbYheEmKsjVwHEArrYnqDfw","tech_manager_email_service": "oLHbYFBfVRFuctlAZmuSElmTI","created_at": "2108-06-15T03:46:49.633189344-04:00","updated_at": "2278-10-14T13:03:29.305006086-04:00"}' | http PUT "https://xinqi.dev:8080/customers/1"  X-Api-User:user123
func UpdateCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	customers := &model.Customers{}
	if err := readJSON(r, customers); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := customers.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	customers.Prepare()

	if err := customers.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "customers", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	customers, _, err = dao.UpdateCustomers(ctx,
		argID,
		customers)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, customers)
}

// DeleteCustomers Delete a single record from customers table in the rocket_development database
// @Summary Delete a record from customers
// @Description Delete a single record from customers table in the rocket_development database
// @Tags Customers
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Customers
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /customers/{argID} [delete]
// http DELETE "https://xinqi.dev:8080/customers/1" X-Api-User:user123
func DeleteCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "customers", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteCustomers(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
