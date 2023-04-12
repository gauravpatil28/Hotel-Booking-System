package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/GG", nil)

	resp := New(r.PostForm)

	isValid := resp.Valid()
	if !isValid {
		t.Error("Got invalid whereas it should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/gg", nil)

	resp := New(r.PostForm)

	resp.Required("a", "b", "c")
	if resp.Valid() {
		t.Error("It should have given it as False")
	}

	postdata := url.Values{}
	postdata.Add("a", "a")
	postdata.Add("b", "a")
	postdata.Add("c", "a")

	r = httptest.NewRequest("POST", "/gg", nil)

	r.PostForm = postdata
	resp = New(r.PostForm)
	resp.Required("a", "b", "c")
	if !resp.Valid() {
		t.Error("It should not show this Error")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/gg", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")

	if has {
		t.Error("Shows Field when it is absent")
	}

	posteddata := url.Values{}
	posteddata.Add("a", "a")
	form = New(posteddata)

	has = form.Has("a")
	if !has {
		t.Error("Shows form does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/gg", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)

	if form.Valid() {
		t.Error("Shows Field when it is absent")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("Should have a error but it did not get an Error")
	}

	posteddata := url.Values{}
	posteddata.Add("some_field", "some value")
	form = New(posteddata)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("Shows form does not have field when it should")
	}

	posteddata = url.Values{}
	posteddata.Add("another_field", "abc123")
	form = New(posteddata)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("Shows minlength of 1 when not meant to be")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("Should not have a error but it got Error")
	}
}

func TestForm_IsEmail(t *testing.T) {
	posteddata := url.Values{}

	form := New(posteddata)

	form.IsEmail("x")

	if form.Valid() {
		t.Error("SHowing valid when it should not")
	}

	posteddata = url.Values{}
	posteddata.Add("email", "abc@gamas.com")
	form = New(posteddata)
	form.IsEmail("email")

	if !form.Valid() {
		t.Error("It's Showing Error when it shouldn't have shown")
	}

	posteddata = url.Values{}
	posteddata.Add("email", "x")
	form = New(posteddata)
	form.IsEmail("a")

	if form.Valid() {
		t.Error("It's Showing Error when it shouldn't have shown")
	}
}
