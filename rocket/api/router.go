package api

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unsafe"

	"rocket/dao"
	"rocket/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

var (
	_             = time.Second // import time.Second for unknown usage in api
	crudEndpoints map[string]*CrudAPI
)

// CrudAPI describes requests available for tables in the database
type CrudAPI struct {
	Name            string           `json:"name"`
	CreateURL       string           `json:"create_url"`
	RetrieveOneURL  string           `json:"retrieve_one_url"`
	RetrieveManyURL string           `json:"retrieve_many_url"`
	UpdateURL       string           `json:"update_url"`
	DeleteURL       string           `json:"delete_url"`
	FetchDDLURL     string           `json:"fetch_ddl_url"`
	TableInfo       *model.TableInfo `json:"table_info"`
}

// PagedResults results for pages GetAll results.
type PagedResults struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"page_size"`
	Data         interface{} `json:"data"`
	TotalRecords int         `json:"total_records"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// ConfigRouter configure http.Handler router
func ConfigRouter() http.Handler {
	router := httprouter.New()
	configActiveAdminCommentsRouter(router)
	configActiveStorageAttachmentsRouter(router)
	configActiveStorageBlobsRouter(router)
	configAddressesRouter(router)
	configAdminUsersRouter(router)
	configArInternalMetadataRouter(router)
	configBatteriesRouter(router)
	configBlazerAuditsRouter(router)
	configBlazerChecksRouter(router)
	configBlazerDashboardQueriesRouter(router)
	configBlazerDashboardsRouter(router)
	configBlazerQueriesRouter(router)
	configBuildingDetailsRouter(router)
	configBuildingsRouter(router)
	configColumnsRouter(router)
	configCustomersRouter(router)
	configElevatorsRouter(router)
	configEmployeesRouter(router)
	configInterventionsRouter(router)
	configLeadsRouter(router)
	configMapsRouter(router)
	configQuotesRouter(router)
	configSchemaMigrationsRouter(router)
	configUsers_Router(router)

	router.GET("/ddl/:argID", GetDdl)
	router.GET("/ddl", GetDdlEndpoints)
	return router
}

// ConfigGinRouter configure gin router
func ConfigGinRouter(router gin.IRoutes) {
	configGinActiveAdminCommentsRouter(router)
	configGinActiveStorageAttachmentsRouter(router)
	configGinActiveStorageBlobsRouter(router)
	configGinAddressesRouter(router)
	configGinAdminUsersRouter(router)
	configGinArInternalMetadataRouter(router)
	configGinBatteriesRouter(router)
	configGinBlazerAuditsRouter(router)
	configGinBlazerChecksRouter(router)
	configGinBlazerDashboardQueriesRouter(router)
	configGinBlazerDashboardsRouter(router)
	configGinBlazerQueriesRouter(router)
	configGinBuildingDetailsRouter(router)
	configGinBuildingsRouter(router)
	configGinColumnsRouter(router)
	configGinCustomersRouter(router)
	configGinElevatorsRouter(router)
	configGinEmployeesRouter(router)
	configGinInterventionsRouter(router)
	configGinLeadsRouter(router)
	configGinMapsRouter(router)
	configGinQuotesRouter(router)
	configGinSchemaMigrationsRouter(router)
	configGinUsers_Router(router)

	router.GET("/ddl/:argID", ConverHttprouterToGin(GetDdl))
	router.GET("/ddl", ConverHttprouterToGin(GetDdlEndpoints))
	return
}

// ConverHttprouterToGin wrap httprouter.Handle to gin.HandlerFunc
func ConverHttprouterToGin(f httprouter.Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params httprouter.Params
		_len := len(c.Params)
		if _len == 0 {
			params = nil
		} else {
			params = ((*[1 << 10]httprouter.Param)(unsafe.Pointer(&c.Params[0])))[:_len]
		}

		f(c.Writer, c.Request, params)
	}
}

func initializeContext(r *http.Request) (ctx context.Context) {
	if ContextInitializer != nil {
		ctx = ContextInitializer(r)
	} else {
		ctx = r.Context()
	}
	return ctx
}

func ValidateRequest(ctx context.Context, r *http.Request, table string, action model.Action) error {
	if RequestValidator != nil {
		return RequestValidator(ctx, r, table, action)
	}

	return nil
}

type RequestValidatorFunc func(ctx context.Context, r *http.Request, table string, action model.Action) error

var RequestValidator RequestValidatorFunc

type ContextInitializerFunc func(r *http.Request) (ctx context.Context)

var ContextInitializer ContextInitializerFunc

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func writeRowsAffected(w http.ResponseWriter, rowsAffected int64) {
	data, _ := json.Marshal(rowsAffected)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}

func returnError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	status := 0
	switch err {
	case dao.ErrNotFound:
		status = http.StatusBadRequest
	case dao.ErrUnableToMarshalJSON:
		status = http.StatusBadRequest
	case dao.ErrUpdateFailed:
		status = http.StatusBadRequest
	case dao.ErrInsertFailed:
		status = http.StatusBadRequest
	case dao.ErrDeleteFailed:
		status = http.StatusBadRequest
	case dao.ErrBadParams:
		status = http.StatusBadRequest
	default:
		status = http.StatusBadRequest
	}
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	SendJSON(w, r, er.Code, er)
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

func parseUint8(ps httprouter.Params, key string) (uint8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return uint8(id), err
	}
	return uint8(id), err
}
func parseUint16(ps httprouter.Params, key string) (uint16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return uint16(id), err
	}
	return uint16(id), err
}
func parseUint32(ps httprouter.Params, key string) (uint32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return uint32(id), err
	}
	return uint32(id), err
}
func parseUint64(ps httprouter.Params, key string) (uint64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return uint64(id), err
	}
	return uint64(id), err
}
func parseInt(ps httprouter.Params, key string) (int, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(id), err
}
func parseInt8(ps httprouter.Params, key string) (int8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return -1, err
	}
	return int8(id), err
}
func parseInt16(ps httprouter.Params, key string) (int16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return -1, err
	}
	return int16(id), err
}
func parseInt32(ps httprouter.Params, key string) (int32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(id), err
}
func parseInt64(ps httprouter.Params, key string) (int64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 54)
	if err != nil {
		return -1, err
	}
	return id, err
}
func parseString(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}
func parseUUID(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}

// GetDdl is a function to get table info for a table in the rocket_development database
// @Summary Get table info for a table in the rocket_development database by argID
// @Tags TableInfo
// @ID argID
// @Description GetDdl is a function to get table info for a table in the rocket_development database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} api.CrudAPI
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /ddl/{argID} [get]
// http "https://xinqi.dev:443/ddl/xyz" X-Api-User:user123
func GetDdl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID := ps.ByName("argID")

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, ok := crudEndpoints[argID]
	if !ok {
		returnError(ctx, w, r, fmt.Errorf("unable to find table: %s", argID))
		return
	}

	writeJSON(ctx, w, record)
}

// GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the rocket_development database
// @Summary Gets a list of ddl endpoints available for tables in the rocket_development database
// @Tags TableInfo
// @Description GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the rocket_development database
// @Accept  json
// @Produce  json
// @Success 200 {object} api.CrudAPI
// @Router /ddl [get]
// http "https://xinqi.dev:443/ddl" X-Api-User:user123
func GetDdlEndpoints(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, crudEndpoints)
}

func init() {
	crudEndpoints = make(map[string]*CrudAPI)

	var tmp *CrudAPI

	tmp = &CrudAPI{
		Name:            "active_admin_comments",
		CreateURL:       "/activeadmincomments",
		RetrieveOneURL:  "/activeadmincomments",
		RetrieveManyURL: "/activeadmincomments",
		UpdateURL:       "/activeadmincomments",
		DeleteURL:       "/activeadmincomments",
		FetchDDLURL:     "/ddl/active_admin_comments",
	}

	tmp.TableInfo, _ = model.GetTableInfo("active_admin_comments")
	crudEndpoints["active_admin_comments"] = tmp

	tmp = &CrudAPI{
		Name:            "active_storage_attachments",
		CreateURL:       "/activestorageattachments",
		RetrieveOneURL:  "/activestorageattachments",
		RetrieveManyURL: "/activestorageattachments",
		UpdateURL:       "/activestorageattachments",
		DeleteURL:       "/activestorageattachments",
		FetchDDLURL:     "/ddl/active_storage_attachments",
	}

	tmp.TableInfo, _ = model.GetTableInfo("active_storage_attachments")
	crudEndpoints["active_storage_attachments"] = tmp

	tmp = &CrudAPI{
		Name:            "active_storage_blobs",
		CreateURL:       "/activestorageblobs",
		RetrieveOneURL:  "/activestorageblobs",
		RetrieveManyURL: "/activestorageblobs",
		UpdateURL:       "/activestorageblobs",
		DeleteURL:       "/activestorageblobs",
		FetchDDLURL:     "/ddl/active_storage_blobs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("active_storage_blobs")
	crudEndpoints["active_storage_blobs"] = tmp

	tmp = &CrudAPI{
		Name:            "addresses",
		CreateURL:       "/addresses",
		RetrieveOneURL:  "/addresses",
		RetrieveManyURL: "/addresses",
		UpdateURL:       "/addresses",
		DeleteURL:       "/addresses",
		FetchDDLURL:     "/ddl/addresses",
	}

	tmp.TableInfo, _ = model.GetTableInfo("addresses")
	crudEndpoints["addresses"] = tmp

	tmp = &CrudAPI{
		Name:            "admin_users",
		CreateURL:       "/adminusers",
		RetrieveOneURL:  "/adminusers",
		RetrieveManyURL: "/adminusers",
		UpdateURL:       "/adminusers",
		DeleteURL:       "/adminusers",
		FetchDDLURL:     "/ddl/admin_users",
	}

	tmp.TableInfo, _ = model.GetTableInfo("admin_users")
	crudEndpoints["admin_users"] = tmp

	tmp = &CrudAPI{
		Name:            "ar_internal_metadata",
		CreateURL:       "/arinternalmetadata",
		RetrieveOneURL:  "/arinternalmetadata",
		RetrieveManyURL: "/arinternalmetadata",
		UpdateURL:       "/arinternalmetadata",
		DeleteURL:       "/arinternalmetadata",
		FetchDDLURL:     "/ddl/ar_internal_metadata",
	}

	tmp.TableInfo, _ = model.GetTableInfo("ar_internal_metadata")
	crudEndpoints["ar_internal_metadata"] = tmp

	tmp = &CrudAPI{
		Name:            "batteries",
		CreateURL:       "/batteries",
		RetrieveOneURL:  "/batteries",
		RetrieveManyURL: "/batteries",
		UpdateURL:       "/batteries",
		DeleteURL:       "/batteries",
		FetchDDLURL:     "/ddl/batteries",
	}

	tmp.TableInfo, _ = model.GetTableInfo("batteries")
	crudEndpoints["batteries"] = tmp

	tmp = &CrudAPI{
		Name:            "blazer_audits",
		CreateURL:       "/blazeraudits",
		RetrieveOneURL:  "/blazeraudits",
		RetrieveManyURL: "/blazeraudits",
		UpdateURL:       "/blazeraudits",
		DeleteURL:       "/blazeraudits",
		FetchDDLURL:     "/ddl/blazer_audits",
	}

	tmp.TableInfo, _ = model.GetTableInfo("blazer_audits")
	crudEndpoints["blazer_audits"] = tmp

	tmp = &CrudAPI{
		Name:            "blazer_checks",
		CreateURL:       "/blazerchecks",
		RetrieveOneURL:  "/blazerchecks",
		RetrieveManyURL: "/blazerchecks",
		UpdateURL:       "/blazerchecks",
		DeleteURL:       "/blazerchecks",
		FetchDDLURL:     "/ddl/blazer_checks",
	}

	tmp.TableInfo, _ = model.GetTableInfo("blazer_checks")
	crudEndpoints["blazer_checks"] = tmp

	tmp = &CrudAPI{
		Name:            "blazer_dashboard_queries",
		CreateURL:       "/blazerdashboardqueries",
		RetrieveOneURL:  "/blazerdashboardqueries",
		RetrieveManyURL: "/blazerdashboardqueries",
		UpdateURL:       "/blazerdashboardqueries",
		DeleteURL:       "/blazerdashboardqueries",
		FetchDDLURL:     "/ddl/blazer_dashboard_queries",
	}

	tmp.TableInfo, _ = model.GetTableInfo("blazer_dashboard_queries")
	crudEndpoints["blazer_dashboard_queries"] = tmp

	tmp = &CrudAPI{
		Name:            "blazer_dashboards",
		CreateURL:       "/blazerdashboards",
		RetrieveOneURL:  "/blazerdashboards",
		RetrieveManyURL: "/blazerdashboards",
		UpdateURL:       "/blazerdashboards",
		DeleteURL:       "/blazerdashboards",
		FetchDDLURL:     "/ddl/blazer_dashboards",
	}

	tmp.TableInfo, _ = model.GetTableInfo("blazer_dashboards")
	crudEndpoints["blazer_dashboards"] = tmp

	tmp = &CrudAPI{
		Name:            "blazer_queries",
		CreateURL:       "/blazerqueries",
		RetrieveOneURL:  "/blazerqueries",
		RetrieveManyURL: "/blazerqueries",
		UpdateURL:       "/blazerqueries",
		DeleteURL:       "/blazerqueries",
		FetchDDLURL:     "/ddl/blazer_queries",
	}

	tmp.TableInfo, _ = model.GetTableInfo("blazer_queries")
	crudEndpoints["blazer_queries"] = tmp

	tmp = &CrudAPI{
		Name:            "building_details",
		CreateURL:       "/buildingdetails",
		RetrieveOneURL:  "/buildingdetails",
		RetrieveManyURL: "/buildingdetails",
		UpdateURL:       "/buildingdetails",
		DeleteURL:       "/buildingdetails",
		FetchDDLURL:     "/ddl/building_details",
	}

	tmp.TableInfo, _ = model.GetTableInfo("building_details")
	crudEndpoints["building_details"] = tmp

	tmp = &CrudAPI{
		Name:            "buildings",
		CreateURL:       "/buildings",
		RetrieveOneURL:  "/buildings",
		RetrieveManyURL: "/buildings",
		UpdateURL:       "/buildings",
		DeleteURL:       "/buildings",
		FetchDDLURL:     "/ddl/buildings",
	}

	tmp.TableInfo, _ = model.GetTableInfo("buildings")
	crudEndpoints["buildings"] = tmp

	tmp = &CrudAPI{
		Name:            "columns",
		CreateURL:       "/columns",
		RetrieveOneURL:  "/columns",
		RetrieveManyURL: "/columns",
		UpdateURL:       "/columns",
		DeleteURL:       "/columns",
		FetchDDLURL:     "/ddl/columns",
	}

	tmp.TableInfo, _ = model.GetTableInfo("columns")
	crudEndpoints["columns"] = tmp

	tmp = &CrudAPI{
		Name:            "customers",
		CreateURL:       "/customers",
		RetrieveOneURL:  "/customers",
		RetrieveManyURL: "/customers",
		UpdateURL:       "/customers",
		DeleteURL:       "/customers",
		FetchDDLURL:     "/ddl/customers",
	}

	tmp.TableInfo, _ = model.GetTableInfo("customers")
	crudEndpoints["customers"] = tmp

	tmp = &CrudAPI{
		Name:            "elevators",
		CreateURL:       "/elevators",
		RetrieveOneURL:  "/elevators",
		RetrieveManyURL: "/elevators",
		UpdateURL:       "/elevators",
		DeleteURL:       "/elevators",
		FetchDDLURL:     "/ddl/elevators",
	}

	tmp.TableInfo, _ = model.GetTableInfo("elevators")
	crudEndpoints["elevators"] = tmp

	tmp = &CrudAPI{
		Name:            "employees",
		CreateURL:       "/employees",
		RetrieveOneURL:  "/employees",
		RetrieveManyURL: "/employees",
		UpdateURL:       "/employees",
		DeleteURL:       "/employees",
		FetchDDLURL:     "/ddl/employees",
	}

	tmp.TableInfo, _ = model.GetTableInfo("employees")
	crudEndpoints["employees"] = tmp

	tmp = &CrudAPI{
		Name:            "interventions",
		CreateURL:       "/interventions",
		RetrieveOneURL:  "/interventions",
		RetrieveManyURL: "/interventions",
		UpdateURL:       "/interventions",
		DeleteURL:       "/interventions",
		FetchDDLURL:     "/ddl/interventions",
	}

	tmp.TableInfo, _ = model.GetTableInfo("interventions")
	crudEndpoints["interventions"] = tmp

	tmp = &CrudAPI{
		Name:            "leads",
		CreateURL:       "/leads",
		RetrieveOneURL:  "/leads",
		RetrieveManyURL: "/leads",
		UpdateURL:       "/leads",
		DeleteURL:       "/leads",
		FetchDDLURL:     "/ddl/leads",
	}

	tmp.TableInfo, _ = model.GetTableInfo("leads")
	crudEndpoints["leads"] = tmp

	tmp = &CrudAPI{
		Name:            "maps",
		CreateURL:       "/maps",
		RetrieveOneURL:  "/maps",
		RetrieveManyURL: "/maps",
		UpdateURL:       "/maps",
		DeleteURL:       "/maps",
		FetchDDLURL:     "/ddl/maps",
	}

	tmp.TableInfo, _ = model.GetTableInfo("maps")
	crudEndpoints["maps"] = tmp

	tmp = &CrudAPI{
		Name:            "quotes",
		CreateURL:       "/quotes",
		RetrieveOneURL:  "/quotes",
		RetrieveManyURL: "/quotes",
		UpdateURL:       "/quotes",
		DeleteURL:       "/quotes",
		FetchDDLURL:     "/ddl/quotes",
	}

	tmp.TableInfo, _ = model.GetTableInfo("quotes")
	crudEndpoints["quotes"] = tmp

	tmp = &CrudAPI{
		Name:            "schema_migrations",
		CreateURL:       "/schemamigrations",
		RetrieveOneURL:  "/schemamigrations",
		RetrieveManyURL: "/schemamigrations",
		UpdateURL:       "/schemamigrations",
		DeleteURL:       "/schemamigrations",
		FetchDDLURL:     "/ddl/schema_migrations",
	}

	tmp.TableInfo, _ = model.GetTableInfo("schema_migrations")
	crudEndpoints["schema_migrations"] = tmp

	tmp = &CrudAPI{
		Name:            "users",
		CreateURL:       "/users_",
		RetrieveOneURL:  "/users_",
		RetrieveManyURL: "/users_",
		UpdateURL:       "/users_",
		DeleteURL:       "/users_",
		FetchDDLURL:     "/ddl/users",
	}

	tmp.TableInfo, _ = model.GetTableInfo("users")
	crudEndpoints["users"] = tmp

}
