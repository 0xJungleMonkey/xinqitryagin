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

func configSchemaMigrationsRouter(router *httprouter.Router) {
	router.GET("/schemamigrations", GetAllSchemaMigrations)
	router.POST("/schemamigrations", AddSchemaMigrations)
	router.GET("/schemamigrations/:argVersion", GetSchemaMigrations)
	router.PUT("/schemamigrations/:argVersion", UpdateSchemaMigrations)
	router.DELETE("/schemamigrations/:argVersion", DeleteSchemaMigrations)
}

func configGinSchemaMigrationsRouter(router gin.IRoutes) {
	router.GET("/schemamigrations", ConverHttprouterToGin(GetAllSchemaMigrations))
	router.POST("/schemamigrations", ConverHttprouterToGin(AddSchemaMigrations))
	router.GET("/schemamigrations/:argVersion", ConverHttprouterToGin(GetSchemaMigrations))
	router.PUT("/schemamigrations/:argVersion", ConverHttprouterToGin(UpdateSchemaMigrations))
	router.DELETE("/schemamigrations/:argVersion", ConverHttprouterToGin(DeleteSchemaMigrations))
}

// GetAllSchemaMigrations is a function to get a slice of record(s) from schema_migrations table in the rocket_development database
// @Summary Get list of SchemaMigrations
// @Tags SchemaMigrations
// @Description GetAllSchemaMigrations is a handler to get a slice of record(s) from schema_migrations table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.SchemaMigrations}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /schemamigrations [get]
// http "http://localhost:8080/schemamigrations?page=0&pagesize=20" X-Api-User:user123
func GetAllSchemaMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "schema_migrations", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllSchemaMigrations(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetSchemaMigrations is a function to get a single record from the schema_migrations table in the rocket_development database
// @Summary Get record from table SchemaMigrations by  argVersion
// @Tags SchemaMigrations
// @ID argVersion
// @Description GetSchemaMigrations is a function to get a single record from the schema_migrations table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argVersion path string true "version"
// @Success 200 {object} model.SchemaMigrations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /schemamigrations/{argVersion} [get]
// http "http://localhost:8080/schemamigrations/hello world" X-Api-User:user123
func GetSchemaMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argVersion, err := parseString(ps, "argVersion")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "schema_migrations", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetSchemaMigrations(ctx, argVersion)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddSchemaMigrations add to add a single record to schema_migrations table in the rocket_development database
// @Summary Add an record to schema_migrations table
// @Description add to add a single record to schema_migrations table in the rocket_development database
// @Tags SchemaMigrations
// @Accept  json
// @Produce  json
// @Param SchemaMigrations body model.SchemaMigrations true "Add SchemaMigrations"
// @Success 200 {object} model.SchemaMigrations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /schemamigrations [post]
// echo '{"version": "lkkVOmEKCiJaPYkqeBkhEMAlB"}' | http POST "http://localhost:8080/schemamigrations" X-Api-User:user123
func AddSchemaMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	schemamigrations := &model.SchemaMigrations{}

	if err := readJSON(r, schemamigrations); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := schemamigrations.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	schemamigrations.Prepare()

	if err := schemamigrations.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "schema_migrations", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	schemamigrations, _, err = dao.AddSchemaMigrations(ctx, schemamigrations)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, schemamigrations)
}

// UpdateSchemaMigrations Update a single record from schema_migrations table in the rocket_development database
// @Summary Update an record in table schema_migrations
// @Description Update a single record from schema_migrations table in the rocket_development database
// @Tags SchemaMigrations
// @Accept  json
// @Produce  json
// @Param  argVersion path string true "version"
// @Param  SchemaMigrations body model.SchemaMigrations true "Update SchemaMigrations record"
// @Success 200 {object} model.SchemaMigrations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /schemamigrations/{argVersion} [put]
// echo '{"version": "lkkVOmEKCiJaPYkqeBkhEMAlB"}' | http PUT "http://localhost:8080/schemamigrations/hello world"  X-Api-User:user123
func UpdateSchemaMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argVersion, err := parseString(ps, "argVersion")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	schemamigrations := &model.SchemaMigrations{}
	if err := readJSON(r, schemamigrations); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := schemamigrations.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	schemamigrations.Prepare()

	if err := schemamigrations.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "schema_migrations", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	schemamigrations, _, err = dao.UpdateSchemaMigrations(ctx,
		argVersion,
		schemamigrations)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, schemamigrations)
}

// DeleteSchemaMigrations Delete a single record from schema_migrations table in the rocket_development database
// @Summary Delete a record from schema_migrations
// @Description Delete a single record from schema_migrations table in the rocket_development database
// @Tags SchemaMigrations
// @Accept  json
// @Produce  json
// @Param  argVersion path string true "version"
// @Success 204 {object} model.SchemaMigrations
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schemamigrations/{argVersion} [delete]
// http DELETE "http://localhost:8080/schemamigrations/hello world" X-Api-User:user123
func DeleteSchemaMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argVersion, err := parseString(ps, "argVersion")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "schema_migrations", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteSchemaMigrations(ctx, argVersion)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
