package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

// main funcasfasdfasdfasdfasdfasdfzzz
func main() {
	var filepath = "main.go"
	srcData, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	wokao("dddd")
	ceshi2("dddd")

	fset := token.NewFileSet()
	fparser, err := parser.ParseFile(fset, filepath, srcData, parser.ParseComments)

	for i, k := range fparser.Comments {

		fmt.Println(i, k.Text())
	}

	for _, k := range fparser.Decls { // main funcasfasdfasdfasdfasdfasdfzzz

		if funDecl, ok := k.(*ast.FuncDecl); ok {
			tokenPos := funDecl.Body.Lbrace
			fmt.Println(tokenPos)
			fmt.Println(funDecl.Doc.Text())
			fmt.Println(funDecl.Name.Name, "name.name")
			fmt.Println(funDecl.Recv)
			fmt.Println(funDecl.Type.Pos())
		}
		// main funcasfasdfasdfasdfasdfasdfzzz
	}
}

// woquasdfasdfasdfasdfasdfasdfasdfasdf
func wokao(tt string) {
	fmt.Println(tt)
}
