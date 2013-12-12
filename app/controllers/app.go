package controllers

import (
  "github.com/robfig/revel"
  "strconv"
  "fmt"
  "net/http"
  "os/exec"
  "os"
  "bufio"
  "strings"
  "bytes"
  "sort"
)

type arrow struct {
  from string
  to string
  attr string
}

type arrowSlice []arrow
func (a arrowSlice) Len() int {
  return len(a)
}
func (a arrowSlice) Less(i, j int) bool {
  if a[i].from != a[j].from {
    return a[i].from < a[j].from
  } else {
    return a[i].to < a[j].to
  }
}
func (a arrowSlice) Swap(i, j int) {
  a[i], a[j] = a[j], a[i]
}

var allCallers map[string]arrowSlice = make(map[string]arrowSlice)
var allCallees map[string]arrowSlice = make(map[string]arrowSlice)

func init() {
  f, err := os.Open("public/data")
  defer f.Close()
  if err != nil {
    revel.ERROR.Println("Can not read public/data: ", err)
  }

  fnames, err := f.Readdirnames(0)
  if err != nil {
    revel.ERROR.Println("Error on reading public/data: ", err)
  }

  for _, v := range(fnames) {
    loadData(v)
  }
}

func loadData(fname string) {
  var callers arrowSlice
  var callees arrowSlice
  f, err := os.Open("public/data/" + fname)
  defer f.Close()
  if err != nil {
    revel.ERROR.Println(err)
    return
  }

  br := bufio.NewReader(f)
  var allLine []byte
  for {
    line, isPrefix, err := br.ReadLine()
    if err != nil {
      break
    }
    allLine = append(allLine, line...)
    if !isPrefix {
      var str = string(allLine)
      parseLine(str, &callers, &callees)
      allLine = []byte{}
    }
  }
  sort.Sort(callers)
  sort.Sort(callees)
  allCallers[fname] = callers
  allCallees[fname] = callees
}

func parseLine(str string, callers *arrowSlice, callees *arrowSlice) {
  if !strings.Contains(str, "->") {
    return
  }
  items := strings.Split(str, "\"")
  if len(items) != 5 {
    revel.ERROR.Println("Can not parse '", str, "'")
  }

  first := items[1]
  second := items[3]

  //caller
  *callers = append(*callers, arrow{second, first, str})
  //callee
  *callees = append(*callees, arrow{first, second, str})
}


type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
  fnames := make([]string, len(allCallers))
  i := 0
  for key, _ := range(allCallers) {
    fnames[i] = key
    i=i+1
  }
	return c.Render(fnames)
}

func (c App) Search(func_name, search_type, call_depth, direction, data_source string) revel.Result {
  c.Validation.Required(func_name).Message("Function name is required")
  c.Flash.Out["func_name"] = func_name

  c.Validation.Required(search_type).Message("Search type is required")
  c.Flash.Out["search_type"] = search_type

  c.Validation.Required(call_depth).Message("Call depth is required")
  callDepth, err :=strconv.ParseUint(call_depth, 0, 0)
  if err != nil {
    c.Validation.Error("Call depth invalid")
  }
  c.Flash.Out["call_depth"] = call_depth

  c.Validation.Required(data_source).Message("Data source is requried")
  c.Flash.Out["data_source"] = data_source

  if c.Validation.HasErrors() {
    // Store the validation errors in the flash context and redirect.
    c.Validation.Keep()
    c.FlashParams()

    delete(c.Flash.Out, "graph")
    return c.Redirect(App.Index)
  }

  c.Flash.Out["direction"] = direction

  c.Flash.Out["graph"] = fmt.Sprintf("/App/CreateImage?func_name=%s&search_type=%s&call_depth=%d&direction=%s&data_source=%s", func_name, search_type, callDepth, direction, data_source)

	return c.Redirect(App.Index)
}

type dynamicImage struct {
  funcName string
  searchType string
  callDepth uint
  direction string
  dataSource string
}

func NewDynamicImage(func_name, search_type string, call_depth uint, direction, data_source string) dynamicImage {
  return dynamicImage{func_name, search_type, call_depth, direction, data_source}
}

func (r dynamicImage) calcCalls(source *arrowSlice, func_name string, call_depth uint, buffer *bytes.Buffer, visitedCalls sort.StringSlice) {
  for _, v := range(*source) {
    if v.from < func_name {
      continue
    } else if v.from > func_name {
      break
    }
    buffer.WriteString(v.attr)
    buffer.WriteString("\n")

    idx := visitedCalls.Search(v.to)
    if idx < visitedCalls.Len() && visitedCalls[idx] == v.to {
    } else if r.callDepth == 0 || call_depth + 1 < r.callDepth {
      visitedCalls = append(visitedCalls, v.to)
      visitedCalls.Sort()
      r.calcCalls(source, v.to, call_depth + 1, buffer, visitedCalls)
    }
  }
}

func (r dynamicImage) Apply(req *revel.Request, resp *revel.Response) {
  buffer := bytes.NewBufferString("digraph callgraph{\n")
  if r.searchType == "caller" {
    var visitedCallers = sort.StringSlice {r.funcName}
    callers, ok := allCallers[r.dataSource]
    if ok {
      r.calcCalls(&callers, r.funcName, 0, buffer, visitedCallers)
    }
  } else {
    var visitedCallees = sort.StringSlice {r.funcName}
    callees, ok := allCallees[r.dataSource]
    if ok {
      r.calcCalls(&callees, r.funcName, 0, buffer, visitedCallees)
    }
  }
  buffer.WriteString("}")

  /* set direction */
  param0 :="-Grankdir=" + r.direction
  cmdDot := exec.Command("dot", param0, "-Tpng")
  cmdDot.Stdin = bytes.NewReader(buffer.Bytes())


  output, err := cmdDot.Output()
  if err != nil {
    resp.WriteHeader(http.StatusOK, "text/html")
    resp.Out.Write([]byte("No image: " + err.Error()))
    return
  }

  resp.WriteHeader(http.StatusOK, "image/png")
  resp.Out.Write(output)
}

func (c App) CreateImage(func_name, search_type string, call_depth uint, direction, data_source string) revel.Result {
	return NewDynamicImage(func_name, search_type, call_depth, direction, data_source)
}


