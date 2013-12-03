// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Search(
		func_name string,
		search_type string,
		call_depth string,
		data_source string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "func_name", func_name)
	revel.Unbind(args, "search_type", search_type)
	revel.Unbind(args, "call_depth", call_depth)
	revel.Unbind(args, "data_source", data_source)
	return revel.MainRouter.Reverse("App.Search", args).Url
}

func (_ tApp) CreateImage(
		func_name string,
		search_type string,
		call_depth uint,
		data_source string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "func_name", func_name)
	revel.Unbind(args, "search_type", search_type)
	revel.Unbind(args, "call_depth", call_depth)
	revel.Unbind(args, "data_source", data_source)
	return revel.MainRouter.Reverse("App.CreateImage", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


