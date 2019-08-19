package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rakyll/statik/fs"
	_ "github.com/t10471/converter/statik"
)

var (
	typeName = flag.String("type", "", "must be set")
	useCase  = flag.String("useCase", "", "converter or decoder, if undefined both")
)

func Usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of converter:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tconverter [flags] -type T [directory]\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tconverter [flags] -type T files... # Must be a encode package\n")
	_, _ = fmt.Fprintf(os.Stderr, "For more information, see:\n")
	_, _ = fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

type Writer struct {
	*template.Template
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("converter: ")
	flag.Usage = Usage
	flag.Parse()
	if len(*typeName) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	var dir string
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		dir = filepath.Dir(args[0])
	}

	pkg, err := newPackage(args)
	if err != nil {
		log.Fatal(err)
	}

	w := newWriter()
	ss, err := pkg.makeStructureSource(*typeName)
	if err != nil {
		log.Fatal(err)
	}
	switch *useCase {
	case "encoder":
		ss.UseCase = Encoder.String()
	case "decoder":
		ss.UseCase = Decoder.String()
	case "":
		ss.UseCase = Both.String()
	default:
		flag.Usage()
		os.Exit(2)
	}

	w.outputToFile(dir, ss)
}

func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

func newWriter() *Writer {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	f, err := statikFS.Open("/converter.go.tpl")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fm := template.FuncMap{
		"concatWithPrefix": func(prefix string, ss []string) string {
			ret := ""
			for i, s := range ss {
				ret = prefix + s
				if len(ss)-i != i {
					ret += ","
				}
			}
			return ret
		},
	}
	t, err := template.New("converter").Funcs(fm).Parse(string(b))
	if err != nil {
		log.Fatal(err)
	}
	return &Writer{t}
}

func (w *Writer) outputToFile(dir string, ss *structureSource) {
	baseName := fmt.Sprintf("%s_converter.go", ss.TypeName)
	outputName := filepath.Join(dir, strings.ToLower(baseName))
	src := new(bytes.Buffer)
	err := w.Execute(src, ss)
	if err != nil {
		log.Fatalf("template Execute: %s", err)
	}
	err = ioutil.WriteFile(outputName, src.Bytes(), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
	err = exec.Command("goimports", "-w", outputName).Run()
	if err != nil {
		log.Fatalf("goimports error: %s", err)
	}

}
