package main 

import ( 
    "os" 
    "text/template" 
) 

type Book struct { 
    Title     string 
    Publisher string 
    Year      int 
} 

func main() { 
    t1 := template.New("Template") 
    t1, _ = t1.Parse("External variable has the value [{{.}}]\n") 
    t1.Execute(os.Stdout, "Amazing")     

    b := Book{"The CSound Book", "MIT Press", 2002} 
    t1.Execute(os.Stdout, b)     
    
    t2 := template.New("Template") 
    t2, _ = t2.Parse("External variable Book has the values [Title: {{.Title}}, Publisher: {{.Publisher}}, Year: {{.Year}}]\n") 
    t2.Execute(os.Stdout, b) 
}
