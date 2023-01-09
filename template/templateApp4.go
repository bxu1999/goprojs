package main 

import ( 
    "os" 
    "text/template" 
) 

func main() { 
    computerList := []string{"Arduino", "Raspberri Pi", "NVidia Jetson Nano"} 

    t, err := template.New("Template").Parse("My favorite computers are:\n{{range .}}{{.}}\n{{end}}\n") 
    if err != nil { 
        panic(err) 
    } 

    t.Execute(os.Stdout, computerList) 
}
