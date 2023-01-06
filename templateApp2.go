package main

import (
    "os"
    "text/template"
)

func main() {
    t, _ := template.New("Template").Parse("{{$var:=2150}}Internal variable has the value [{{$var}}]\n")    
    t.Execute(os.Stdout, nil)
}
