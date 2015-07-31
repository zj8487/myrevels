// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) WBlog(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.WBlog", args).Url
}

func (_ tApp) BlogInfor(
		id string,
		rcnt int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "rcnt", rcnt)
	return revel.MainRouter.Reverse("App.BlogInfor", args).Url
}

func (_ tApp) Message(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Message", args).Url
}

func (_ tApp) History(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.History", args).Url
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


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tWBlog struct {}
var WBlog tWBlog


func (_ tWBlog) Putup(
		blog interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "blog", blog)
	return revel.MainRouter.Reverse("WBlog.Putup", args).Url
}


type tWComment struct {}
var WComment tWComment


func (_ tWComment) Docomment(
		id string,
		rcnt int,
		comment interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "rcnt", rcnt)
	revel.Unbind(args, "comment", comment)
	return revel.MainRouter.Reverse("WComment.Docomment", args).Url
}


type tWMessage struct {}
var WMessage tWMessage


func (_ tWMessage) Putup(
		message interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "message", message)
	return revel.MainRouter.Reverse("WMessage.Putup", args).Url
}


