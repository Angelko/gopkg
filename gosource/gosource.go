package gosource

import (
	"bytes"
	"errors"
)

// GoSource describe the content to a go source file
type GoSource struct {
	buf      bytes.Buffer
	name     string
	describe string
	comments []string
	imports  []string
	consts   map[string]Type
	globals  map[string]Type
	funcs    []*Function
}

// New return a GoSource type initialized
// describe will be comment add on package
// do not set go syntax, just the text which you want
func New(packageName, describe string) *GoSource {
	return &GoSource{
		name:     packageName,
		describe: describe,
		consts:   make(map[string]Type),
		globals:  make(map[string]Type),
	}
}

// SetComments which you want add on your package
// each entry on only one line
func (g *GoSource) SetComments(comments ...string) {
	for _, v := range comments {
		g.comments = append(g.comments, v)
	}
}

// SetImports which you want add on your package
// each entry to a complete import
func (g *GoSource) SetImports(imports ...string) {
	for _, v := range imports {
		g.imports = append(g.imports, v)
	}
}

// AddConst in the GoSource
// name describe the constant name
// nameType describe the constant type
// value describe the constant value
func (g *GoSource) AddConst(name, nameType string, value interface{}) error {
	if value == nil {
		return errors.New("value not define to the constant definition")
	}
	g.consts[name] = Type{
		name:  nameType,
		value: value,
	}
	return nil
}

// AddGlobal in the GoSource where
// name describe the global name
// nameType describe the global type
// value describe the global value
func (g *GoSource) AddGlobal(name, nameType string, value interface{}) error {
	if value == nil {
		return errors.New("value not define to the global definition")
	}
	g.globals[name] = Type{
		name:  nameType,
		value: value,
	}
	return nil
}

// AddFunction in the GoSource
func (g *GoSource) AddFunction(f *Function) error {
	if f == nil {
		return errors.New("function not define on the function definition")
	}
	g.funcs = append(g.funcs, f)
	return nil
}
