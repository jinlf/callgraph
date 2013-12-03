// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/robfig/revel"
	_ "github.com/callgraph/app"
	controllers "github.com/callgraph/app/controllers"
	controllers0 "github.com/robfig/revel/modules/static/app/controllers"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					117: []string{ 
						"fnames",
					},
				},
			},
			&revel.MethodType{
				Name: "Search",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "func_name", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "search_type", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "call_depth", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "data_source", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CreateImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "func_name", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "search_type", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "call_depth", Type: reflect.TypeOf((*uint)(nil)) },
					&revel.MethodArg{Name: "data_source", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"github.com/callgraph/app/controllers.App.Search": { 
			121: "func_name",
			124: "search_type",
			127: "call_depth",
			134: "data_source",
		},
	}
	revel.TestSuites = []interface{}{ 
	}

	revel.Run(*port)
}
