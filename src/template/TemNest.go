/*模板嵌套*/
package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("header.html", "content.html", "footer.html")
	if err != nil {
		fmt.Printf("parsefile failed:%s\n", err.Error())
	}
	tmpl.ExecuteTemplate(os.Stdout, "content", nil)
}
