package main 

import ( 
    "os" 
    "text/template" 
) 

func add(a, b int) int { 
    return a + b 
} 

func main() { 
    dishesList := []string{"Enciladas con Pollo", "Hot&Spicy Pizza", "Spaghetti Bolognese"}     
    
    t, err := template.New("Template").Funcs(template.FuncMap{"add": add}).Parse("My favorite dishes are:\n{{range $index, $item:=.}}({{add $index 1}}) {{$item}}\n{{end}}\n")     
    if err != nil { 
        panic(err) 
    }     
    
    t.Execute(os.Stdout, dishesList) 
}
