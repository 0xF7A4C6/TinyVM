package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	_ "tinyvm/internal/vm"
)

var (
	baseRegister = 10
	byteCode     []byte
)

func ParseFunc(node ast.Node) bool {
	funcType := node.(*ast.FuncDecl)
	log.Println("func ->", funcType, funcType.Recv, funcType.Type.Func)

	return true
}

func ParseBinary(node ast.Node) bool {
	funcType := node.(*ast.BinaryExpr)
	log.Println("binary ->", funcType, funcType.Op, funcType.X, funcType.Y)

	switch funcType.Op {
	case token.ADD:
		fmt.Println(funcType.X.(*ast.Ident))
		/*byteCode = append(byteCode,
			// LOAD_NUM, reg, val
			vm.LOAD_NUM, byte(baseRegister), funcType.X.(*ast.BasicLit).Value[0]-'0',
			// LOAD_NUM, reg, val
			vm.LOAD_NUM, byte(baseRegister+1), funcType.Y.(*ast.BasicLit).Value[0]-'0',
			// ADD, reg, left, right
			vm.ADD, byte(baseRegister+2), byte(baseRegister), byte(baseRegister+1),
		)*/
		baseRegister += 10

	case token.SUB:
		/*byteCode = append(byteCode,
			// LOAD_NUM, reg, val
			vm.LOAD_NUM, byte(baseRegister), funcType.X.(*ast.BasicLit).Value[0]-'0',
			// LOAD_NUM, reg, val
			vm.LOAD_NUM, byte(baseRegister+1), funcType.Y.(*ast.BasicLit).Value[0]-'0',
			// ADD, reg, left, right
			vm.ADD, byte(baseRegister+2), byte(baseRegister), byte(baseRegister+1),
		)*/
		baseRegister += 10
	}

	return true
}

func ParseNode(node ast.Node) bool {
	switch n := node.(type) {
	case *ast.FuncDecl:
		ParseFunc(n)
	case *ast.BinaryExpr:
		ParseBinary(n)
	}

	return true
}

func Parse(input string) error {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, input, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(node, ParseNode)

	return nil
}
