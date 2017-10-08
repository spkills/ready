package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
)

var (
	file          = flag.String("file", "route.conf", "Path to route.conf file")
	templatesPath = flag.String("templatesPath", "templates", "Path to route.conf file")
)

// 終了コード
const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeNG
)

type Ready struct {
	outStream, errStream io.Writer
}

type TemplateData struct {
	Name        string
	CapitalName string
	CamelName   string
}

func main() {
	cli := &Ready{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Create(os.Args))
}

func (r *Ready) Create(args []string) int {
	// オプション引数のパース
	var infile, inDir string
	flags := flag.NewFlagSet("awesome-cli", flag.ContinueOnError)
	flags.SetOutput(r.errStream)
	flag.Parse()
	flags.StringVar(&infile, "file", *file, "Path to route.conf file")
	flags.StringVar(&inDir, "templatesPath", *templatesPath, "Path to route.conf file")

	fmt.Println(inDir)
	fp, err := os.Open(infile)
	if err != nil {
		fmt.Println(err)
		return ExitCodeParseFlagError
	}

	scanner := bufio.NewScanner(fp)

	dataList := make([]TemplateData, 0, 10)
	for scanner.Scan() {

		path := scanner.Text()
		separatedPath := strings.Split(path, "/")
		titledPath := strings.Replace(strings.Title(strings.Join(separatedPath, " ")), " ", "", -1)

		filename := ""
		for i, v := range separatedPath {
			if i == 0 {
				filename += v
				continue
			}
			filename += strings.Title(v)
		}

		data := TemplateData{
			Name:        path,
			CapitalName: titledPath,
			CamelName:   filename,
		}

		dataList = append(dataList, data)
		r.createHandler(data, inDir)

	}
	r.createRoutingFile(inDir, dataList)

	return ExitCodeOK
}

func (r *Ready) createRoutingFile(inDir string, dataList []TemplateData) {
	outfile := "controller/routing.go"

	//execute template
	outf, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("cannot createHandler file %q: %s\n", outfile, err)
		panic(err)
	}

	tpl := template.Must(template.ParseFiles(inDir + "/routing.tmpl"))
	err = tpl.Execute(outf, dataList)
	if err != nil {
		panic(err)
	}
}

func (r *Ready) createHandler(data TemplateData, inDir string) {

	outfile := "controller/" + data.CamelName + ".go"
	fmt.Printf("Compiling %q to %q...\n", data.CamelName, outfile)

	if r.fileExists(outfile) {
		fmt.Println("file exists")
		return
	}

	outf, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("cannot createHandler file %q: %s\n", outfile, err)
		panic(err)
	}

	//execute template
	tpl := template.Must(template.ParseFiles(inDir + "/handler.tmpl"))

	err = tpl.Execute(outf, data)
	if err != nil {
		panic(err)
	}
}

func (r *Ready) fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
