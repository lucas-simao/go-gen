package templates

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/lucas-simao/go-gen-ca/internal/utils"
	goLog "github.com/lucas-simao/golog"
)

//go:embed main.tmpl
var mainFile string

func GenerateMain(serviceName, projectName string) string {
	tmpl, err := utils.InitTemplate("main", mainFile)
	if err != nil {
		goLog.Error(err)
	}

	b := bytes.Buffer{}

	err = tmpl.Execute(&b, map[string]string{
		"ProjectName": projectName,
		"ServiceName": serviceName,
	})
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return b.String()
}
