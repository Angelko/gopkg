package gosource

import (
	"fmt"
	"go/format"
)

func (g *GoSource) writeN(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format+"\n", args...)
}

func (g *GoSource) write(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *GoSource) printHeader() {
	if g.describe != "" {
		g.writeN("// package %s %s", g.name, g.describe)
	}
	for _, v := range g.comments {
		g.writeN("// %s", v)
	}
	g.writeN("package %s", g.name)
	g.writeN("")
}

func (g *GoSource) printImport() {
	g.writeN("import (")
	for _, v := range g.imports {
		g.writeN("\t\"%s\"", v)
	}
	g.writeN(")")
	g.writeN("")
}

func (g *GoSource) printConst() {
	g.writeN("const (")
	for name, t := range g.consts {
		if t.name == "string" {
			g.writeN("\t%s %s = \"%v\"", name, t.name, t.value)
		} else {
			g.writeN("\t%s %s = %v", name, t.name, t.value)
		}
	}
	g.writeN(")")
	g.writeN("")
}

func (g *GoSource) printGlobal() {
	for name, t := range g.globals {
		if t.name == "string" {
			g.writeN("var %s %s = \"%v\"", name, t.name, t.value)
		} else {
			g.writeN("var %s %s = %v", name, t.name, t.value)
		}
	}
	g.writeN("")
}

func (g *GoSource) printFunction() {
	for _, f := range g.funcs {
		g.write("func %s(", f.name)
		first := true
		for _, p := range f.args {
			if !first {
				g.write(", ")
			}
			g.write("%s %s", p.Name, p.Type)
			first = false
		}
		if len(f.rets) > 0 {
			g.write(") (")
			first = true
			for _, p := range f.rets {
				if !first {
					g.write(", ")
				}
				g.write("%s %s", p.Name, p.Type)
				first = false
			}
		}
		g.writeN(") {")
		g.write(f.content)
		g.writeN("}")
		g.writeN("")
	}
}

// Bytes return the content src file describe on the GoSource
// If GoSource don't describe a go file, err is return and byte will be empty
func (g *GoSource) Bytes() ([]byte, error) {
	var content []byte
	var err error

	g.printHeader()
	if len(g.imports) > 0 {
		g.printImport()
	}
	if len(g.consts) > 0 {
		g.printConst()
	}
	if len(g.globals) > 0 {
		g.printGlobal()
	}
	g.printFunction()

	if content, err = format.Source(g.buf.Bytes()); err != nil {
		fmt.Println(string(g.buf.Bytes()))
		return nil, err
	}
	return content, nil
}
