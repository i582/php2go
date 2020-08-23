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

	inputFolder, _ := filepath.Split(inputFile)

	if outputFile == "" {
		dir, file := filepath.Split(inputFile)
		outputFile = dir + strings.TrimSuffix(file, filepath.Ext(file)) + ".go"
	} else {
		dir, _ := filepath.Split(outputFile)
		inputFolder = dir
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

	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("main file not created: %s", err)
	}
	defer f.Close()

	core, err := os.Create(inputFolder + "/core.go")
	if err != nil {
		log.Fatalf("core file not created: %s", err)
	}

	gw := generator.NewGeneratorWalker(f, core, filepath.Base(inputFile))
	rootNode.Walk(&gw)
	gw.Final()

	gw = generator.NewGeneratorWalker(os.Stdout, os.Stdout, filepath.Base(inputFile))
	rootNode.Walk(&gw)
	gw.Final()
}
