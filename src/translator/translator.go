package translator

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/i582/php2go/src/generator"
	"github.com/i582/php2go/src/meta"
	"github.com/i582/php2go/src/php/php7"
	"github.com/i582/php2go/src/root"
)

type Translator struct {
}

func NewTranslator() Translator {
	return Translator{}
}

func (t Translator) Run() {
	var inputFile string
	flag.StringVar(&inputFile, "i", "", "input file")

	var outputFile string
	flag.StringVar(&outputFile, "o", "", "output file")

	flag.Parse()

	if outputFile == "" {
		dir, file := filepath.Split(inputFile)
		outputFile = dir + strings.TrimSuffix(file, filepath.Ext(file)) + ".go"
	}

	src, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	parser := php7.NewParser(src, "7.4")
	parser.Parse()

	for _, e := range parser.GetErrors() {
		fmt.Println(e)
	}

	rw := root.RootWalker{}
	rootNode := parser.GetRootNode()
	rootNode.Walk(&rw)

	for _, v := range meta.AllVariables.Vars {
		fmt.Println(v)
	}

	for _, f := range meta.AllFunctions.Functions {
		fmt.Println(f)

		for _, v := range f.Variables.Vars {
			fmt.Println(v)
		}
	}

	f, err := os.Create(outputFile)
	if err != nil {
		panic("file not created")
	}
	defer f.Close()

	gw := generator.NewGeneratorWalker(f, filepath.Base(inputFile))
	rootNode.Walk(&gw)
	gw.Final()

	gw = generator.NewGeneratorWalker(os.Stdout, filepath.Base(inputFile))
	rootNode.Walk(&gw)
	gw.Final()
}
