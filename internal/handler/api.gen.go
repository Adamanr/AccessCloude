// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// Defines values for UserRole.
const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleLain    UserRole = "lain"
	UserRoleMember  UserRole = "member"
	UserRoleMentor  UserRole = "mentor"
	UserRoleTeacher UserRole = "teacher"
)

// Defines values for GetUsersParamsRole.
const (
	GetUsersParamsRoleAdmin   GetUsersParamsRole = "admin"
	GetUsersParamsRoleLain    GetUsersParamsRole = "Lain"
	GetUsersParamsRoleMember  GetUsersParamsRole = "member"
	GetUsersParamsRoleMentor  GetUsersParamsRole = "mentor"
	GetUsersParamsRoleTeacher GetUsersParamsRole = "teacher"
)

// InternalServiceError defines model for InternalServiceError.
type InternalServiceError struct {
	Code    *int         `json:"code,omitempty"`
	Message *interface{} `json:"message,omitempty"`
}

// User defines model for User.
type User struct {
	AvatarId    string     `json:"avatar_id"`
	CreatedAt   *string    `json:"created_at,omitempty"`
	Description *string    `json:"description,omitempty"`
	Email       string     `json:"email"`
	Id          *int64     `json:"id,omitempty"`
	Login       string     `json:"login"`
	Password    *string    `json:"password,omitempty"`
	Role        []UserRole `json:"role"`
	UpdatedAt   *string    `json:"updated_at,omitempty"`
}

// UserRole defines model for User.Role.
type UserRole string

// UserSignIn defines model for UserSignIn.
type UserSignIn struct {
	Email    *string `json:"email,omitempty"`
	Login    *string `json:"login,omitempty"`
	Password *string `json:"password,omitempty"`
}

// UserSignUp defines model for UserSignUp.
type UserSignUp struct {
	Email        *string `json:"email,omitempty"`
	Login        *string `json:"login,omitempty"`
	Notification *bool   `json:"notification,omitempty"`
	Password     *string `json:"password,omitempty"`
}

// Users List of user object
type Users = []User

// GetUsersParams defines parameters for GetUsers.
type GetUsersParams struct {
	// Limit Параметр для получения определенное количество курсов
	Limit *int                  `form:"limit,omitempty" json:"limit,omitempty"`
	Role  *[]GetUsersParamsRole `form:"role,omitempty" json:"role,omitempty"`

	// OrderBy Параметр для сортировки курсов по указанному полю и методу сортировки
	OrderBy *string `form:"orderBy,omitempty" json:"orderBy,omitempty"`
}

// GetUsersParamsRole defines parameters for GetUsers.
type GetUsersParamsRole string

// UploadUserAvatarMultipartBody defines parameters for UploadUserAvatar.
type UploadUserAvatarMultipartBody = string

// UploadUserAvatarMultipartRequestBody defines body for UploadUserAvatar for multipart/form-data ContentType.
type UploadUserAvatarMultipartRequestBody = UploadUserAvatarMultipartBody

// UserSignInJSONRequestBody defines body for UserSignIn for application/json ContentType.
type UserSignInJSONRequestBody = UserSignIn

// UserSignUpJSONRequestBody defines body for UserSignUp for application/json ContentType.
type UserSignUpJSONRequestBody = UserSignUp

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// pong
	// (GET /pong)
	Pong(w http.ResponseWriter, r *http.Request)
	// Delete user
	// (DELETE /user/{id})
	DeleteUser(w http.ResponseWriter, r *http.Request, id int)
	// Get user by id
	// (GET /user/{id})
	GetUserById(w http.ResponseWriter, r *http.Request, id int)
	// Update user
	// (PUT /user/{id})
	UpdateUser(w http.ResponseWriter, r *http.Request, id int)
	// List all users
	// (GET /users)
	GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams)
	// Upload avatar
	// (POST /users/avatar)
	UploadUserAvatar(w http.ResponseWriter, r *http.Request)
	// Authorize user
	// (POST /users/sign_in)
	UserSignIn(w http.ResponseWriter, r *http.Request)
	// Create user
	// (POST /users/sign_up)
	UserSignUp(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// pong
// (GET /pong)
func (_ Unimplemented) Pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete user
// (DELETE /user/{id})
func (_ Unimplemented) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get user by id
// (GET /user/{id})
func (_ Unimplemented) GetUserById(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update user
// (PUT /user/{id})
func (_ Unimplemented) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List all users
// (GET /users)
func (_ Unimplemented) GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Upload avatar
// (POST /users/avatar)
func (_ Unimplemented) UploadUserAvatar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Authorize user
// (POST /users/sign_in)
func (_ Unimplemented) UserSignIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create user
// (POST /users/sign_up)
func (_ Unimplemented) UserSignUp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// Pong operation middleware
func (siw *ServerInterfaceWrapper) Pong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Pong(w, r)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUser(w, r, id)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUserById operation middleware
func (siw *ServerInterfaceWrapper) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserById(w, r, id)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUser(w, r, id)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUsers operation middleware
func (siw *ServerInterfaceWrapper) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersParams

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "role" -------------

	err = runtime.BindQueryParameter("form", true, false, "role", r.URL.Query(), &params.Role)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "role", Err: err})
		return
	}

	// ------------- Optional query parameter "orderBy" -------------

	err = runtime.BindQueryParameter("form", true, false, "orderBy", r.URL.Query(), &params.OrderBy)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "orderBy", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsers(w, r, params)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UploadUserAvatar operation middleware
func (siw *ServerInterfaceWrapper) UploadUserAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UploadUserAvatar(w, r)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UserSignIn operation middleware
func (siw *ServerInterfaceWrapper) UserSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UserSignIn(w, r)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UserSignUp operation middleware
func (siw *ServerInterfaceWrapper) UserSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UserSignUp(w, r)
	}))

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/pong", wrapper.Pong)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/user/{id}", wrapper.DeleteUser)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/user/{id}", wrapper.GetUserById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/user/{id}", wrapper.UpdateUser)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users", wrapper.GetUsers)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/users/avatar", wrapper.UploadUserAvatar)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/users/sign_in", wrapper.UserSignIn)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/users/sign_up", wrapper.UserSignUp)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xY3W4Utxd/Fcv8b/7SZGcToKpWilQIiAYhhBTlCq2Qd+ZkxmjGNrYnYRutlA+J3lTl",
	"DVr1DVaIiDTA8gqeN6qOZ/YrO9uEBipBudpZ28c+/p3f+fI+jWSupABhDe3sUxOlkDP/uSksaMGyLdC7",
	"PIK7WkuN40pLBdpy8KsiGQP+wnOWqwxoZyJHDOhd0AS8YEBtX+E0FxYS0HQQ0ByMYcmlxAeDyQay9xQi",
	"i/LbBho0YrvMMv2Ex/P7rq5dv3FzqoaxmosEd4k0MAvxE2bnBdbaq9+vtFdX2qtNUjGYSHNluRTzYvfZ",
	"LiPcEJsC6YGxTcKQM57Niz2VqfjBj7cimTcJnbvQajugO1LnqDaC+t2NRowzmfBzGtoUPHINZyhmzJ7U",
	"l4VOy8xbj1vIPfYgipx2HtOMcUEDyuLc/+aQ9/yBOQhbkQFYlIKm3YZd6wGmNevj/0LFH20gVA6eFVxD",
	"7PXxKIyBrxUPZqjSPU+vgD7PvYkEy3G4QMQGNem2eCI2xSL1/pFdGyx0X6bio8zTCECjv6Dq2+rfV11I",
	"y3d4xBYcxuoCJut7UmbAxCe6q7/YnKPSB9xYIncImpNMTD3h7/807NAOvRZOo2JYh8TQ+8wCPfFweF6F",
	"rDsyajgySjmJZVQg9/31yfr/aUALndEOTa1VphOGiVyJUt7iMrwWemcXO7IKr8KyyBN/XsC2cgg38P6w",
	"+XBjISJR91t56N67kXvl3rth+ZK4127kzspj986duPflkRuWL9wpTnxwI7K1x5IEdIsGNOMRCAMz3L+l",
	"0FnJWqs9p3cnDPf29lrMz7akTsJa1IQPNjfuPty6u7LWardSm2ceN269CeujUKG3/vjywI3ciTtDnciG",
	"LLSBH4seDeguaFPdZrXVbrVxE6lAMMVph15vtVvIBcVs6lEPlRQJfiTg8UJ2e7w3Y9qhj3ASg4JRElXE",
	"FWvt9hhjEF6GKZXVJA2fmoqplf3x6zztFkDfKqIIjNkpMjI53dPTFHnOdJ92qKr0sCwxGJhSYJlNaRcX",
	"hUjKcJ/Hg4pEGVhYvMgdP14HcMU0y8F6qj8+T7yZf0RDwnTMRUK8yBqSHtcgejQYG5rHdDZuVq45vf80",
	"9yxmmkH3iuBe7HqXxTugNz/h2Y1lUIMuy+uWqfEr2/ngM8OBwoeq7iBoZu49sHj/2/3N+JvFvzSL3wNb",
	"5Zpen3isF42uigajb/ua55uXf4k2r2y3zMvHgd4sTVW1w5uLLO9+d8PywA0xo5dH5cE0o7qRe1selz9j",
	"pq9y/Min2RP32p24t34Ya4MT4s5wrTvFteVheeReuRHBMqE8KA+xehhT6FkBuj/lUMZzjpVTA22qufXV",
	"dkM/Mgj2G7er6/HpbgsdxYMlHcW4kZj0Ft1gRhXcd30sc0GjgbpdDmBEpjwoj9ypL11euTN3OgdaVVOV",
	"x+7MDd0bN6zhflcej43zK0ERv68budflceOmS8CXOsZ8sAT+enadFTaVuoNXaiiZP7cjm6/Bk323wLKM",
	"FLU/LnXmsGolfV8lTWNAzySLEZlb1coqBIOxt2XcP3f7vMgsV0zbEDv8lZhZ9pGVaKUOKXRG6uA/H/AH",
	"VzQ/i2OOUyx7NNdGTl4kelwwz9llXjdu1b6GeI+mJWxs1+UkMTwRT6qOeQlLpo8Lf8ePq7lmfUAT8jwR",
	"hAsiYG+cvj4tbf4L6f+Wj7v8p4srgIoPhbqYD9vqM/NhWy3lQ6G+8eEqfNjw78tLyTCYjC2UH3+4D+7U",
	"VxRnvoJzJ1iNlC/cqfuTlIfjWuIX98YXDMPyCKu78qV750uHmZdLQxerm/olol5VP0QMuoO/AgAA///F",
	"9tQiExgAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
