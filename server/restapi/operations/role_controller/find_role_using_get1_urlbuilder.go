// Code generated by go-swagger; DO NOT EDIT.

package role_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// FindRoleUsingGET1URL generates an URL for the find role using g e t 1 operation
type FindRoleUsingGET1URL struct {
	RoleID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindRoleUsingGET1URL) WithBasePath(bp string) *FindRoleUsingGET1URL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindRoleUsingGET1URL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FindRoleUsingGET1URL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/role/_roles/{roleId}"

	roleID := o.RoleID
	if roleID != "" {
		_path = strings.Replace(_path, "{roleId}", roleID, -1)
	} else {
		return nil, errors.New("RoleID is required on FindRoleUsingGET1URL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FindRoleUsingGET1URL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FindRoleUsingGET1URL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FindRoleUsingGET1URL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FindRoleUsingGET1URL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FindRoleUsingGET1URL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *FindRoleUsingGET1URL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}