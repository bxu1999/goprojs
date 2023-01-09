package main 

import ( 
    "os" 
    "text/template" 
) 

func main() { 
    t, err := template.New("Template").Parse("{{if eq . `filler`}}This is filler...{{else}}It's something else...{{end}}\n") 

    if err != nil { 
        panic(err) 
    } 

    t.Execute(os.Stdout, "filler") 
}
