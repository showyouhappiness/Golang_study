package template

import (
	"fmt"
	"log"
	"os"
	"testing"
	"text/template"
)

type Person struct {
	Name                string
	nonExportedAgeField string
}

func TestField(t *testing.T) {
	tem := template.New("hello")
	tem, _ = tem.Parse("hello {{.Name}}!")
	p := Person{Name: "Mary", nonExportedAgeField: "31"}
	if err := tem.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}

func TestValidation(t *testing.T) {
	tOk := template.New("ok")
	tErr := template.New("error_template")
	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	//a valid template, so no panic with Must:
	template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name }}"))
	fmt.Println("The first one parsed OK.")
	fmt.Println("The next one ought to fail.")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}
