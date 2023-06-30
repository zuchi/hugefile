package startup

import (
	"90poe/src/pkg/domain/ports"
	"90poe/src/pkg/domain/use_cases"
	"90poe/src/pkg/infra/parsers/port_parser"
)

type Dependencies struct {
	PortParser ports.PortParser
	PortUC     *use_cases.PortUC
}

func InitDependencies() Dependencies {
	dep := Dependencies{}
	dep.PortParser = port_parser.NewJsonParser()
	dep.PortUC = use_cases.NewPortUC(dep.PortParser)
	return dep
}
