package main 

import ( 
    "os" 
    "text/template" 
) 

func main() { 
    dishesList := []string{"Enciladas con Pollo", "Hot&Spicy Pizza", "Spaghetti Bolognese"} 

    t, err := template.New("Template").Parse("My favorite dishes are:\n{{range $index, $item:=.}} ({{$index}}) {{$item}}\n{{end}}\n") 
    if err != nil { 
        panic(err) 
    } 
    t.Execute(os.Stdout, dishesList) 
}
