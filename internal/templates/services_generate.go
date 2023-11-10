package templates

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/lucas-simao/go-gen-ca/internal/utils"
	goLog "github.com/lucas-simao/golog"
)

//go:embed services.tmpl
var servicesFile string

func GenerateServices(serviceName, projectName string) string {
	tmpl, err := utils.InitTemplate("usecases", servicesFile)
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
