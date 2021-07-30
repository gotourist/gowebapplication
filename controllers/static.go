package controllers

import "github.com/komiljon/gowebapplication/views"

func NewStatic() *Static {
	return &Static{
		Home: views.NewView(
			"bootstrap", "static/home"),
		Contact: views.NewView(
			"bootstrap", "static/contact"),
		Faq: views.NewView(
			"bootstrap", "static/faq"),
		Pagenotfound: views.NewView(
			"bootstrap", "static/pagenotfound"),
	}
}

type Static struct {
	Home         *views.View
	Contact      *views.View
	Faq          *views.View
	Pagenotfound *views.View
}
