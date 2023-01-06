package main

import (
    "os"
    "text/template"
)

func add(a, b int) int {
    return a + b
}

func delimiter(s string) func() string {
    return func() string {
        return s
    }
}

func main() {
    dishesList := []string{"Enciladas con Pollo", "Hot&Spicy Pizza", "Spaghetti Bolognese"}    
    
    tmpl := "Index{{dl}}Dish\n{{range $index, $item:=.}}{{add $index 1}}{{dl}}{{$item}}\n{{end}}\n"    
    funcMap := template.FuncMap{"add": add, "dl": delimiter(",")}    
    t, _ := template.New("Template").Funcs(funcMap).Parse(tmpl)    
    t.Execute(os.Stdout, dishesList)
}
