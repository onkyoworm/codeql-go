// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for github.com/emicklei/go-restful/v3, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: github.com/emicklei/go-restful/v3 (exports: Request; functions: )

// Package gorestfulstub is a stub of github.com/emicklei/go-restful/v3, generated by depstubber.
package gorestfulstub

import (
	http "net/http"
)

type Request struct {
	Request *http.Request
}

func (_ Request) Attribute(_ string) interface{} {
	return nil
}

func (_ Request) SelectedRoutePath() string {
	return ""
}

func (_ *Request) BodyParameter(_ string) (string, error) {
	return "", nil
}

func (_ *Request) HeaderParameter(_ string) string {
	return ""
}

func (_ *Request) PathParameter(_ string) string {
	return ""
}

func (_ *Request) PathParameters() map[string]string {
	return nil
}

func (_ *Request) QueryParameter(_ string) string {
	return ""
}

func (_ *Request) QueryParameters(_ string) []string {
	return nil
}

func (_ *Request) ReadEntity(_ interface{}) error {
	return nil
}

func (_ *Request) SetAttribute(_ string, _ interface{}) {}
