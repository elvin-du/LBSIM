package filter

import (
	"net/http"
)


type Filter  interface{
	PreFilter(http.ResponseWriter, *http.Request)bool
	PreFilter(http.ResponseWriter, *http.Request)
	PostFilter(http.ResponseWriter, *http.Request)bool
}

type EmptyFilter struct {}

func (this *EmptyFilter) PreFilter(http.ResponseWriter, *http.Request) bool{return true}
func (this *EmptyFilter) PreErrorHandle(http.ResponseWriter, *http.Request){}
func (this *EmptyFilter) PostFilter(http.ResponseWriter, *http.Request) bool{return true}

