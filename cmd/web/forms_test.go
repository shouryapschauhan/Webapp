package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_Form_Has(t *testing.T) {
	form := NewForm(nil)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has feild when it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = NewForm(postedData)

	has = form.Has("a")
	if !has{
		t.Error("shows form does not have feild when it does")
	}
}

func Test_Form_Required(t *testing.T){
	r := httptest.NewRequest("POST","/whatever",nil)
	form := NewForm(r.PostForm)

	form.Required("a","b","c")

	if form.Valid(){
		t.Error("form shows valid when required feilds are missing")
	}

	postedData := url.Values{}
	postedData.Add("a","a")
	postedData.Add("b","b")
	postedData.Add("c","c")

	r,_ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData

	form = NewForm(r.PostForm)
	form.Required("a","b","c")
	if !form.Valid(){
		t.Error("form shows invalid when required feilds are present")
	}
}

func Test_Form_Check(t *testing.T){
	form := NewForm(nil)

	form.Check(false, "password", "password is required")
	if form.Valid(){
		t.Error("Valid() returns false when it should return true")
	}
}

func Test_Form_ErrorGet(t *testing.T){
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	s := form.Errors.Get("password")

	if len(s) == 0{
		t.Error("should have an error returned from Get, but do not")
	}

	s = form.Errors.Get("whatever")
	if len(s) != 0{
		t.Error("should not have an error but got one")
	}
}
