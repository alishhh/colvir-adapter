package mappers

import (
	"errors"

	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/config"
	"gitlab.apps.bcc.kz/digital-banking-platform/packages/go/dbp-adapters-framework/logger"
)

type IObjectMapper interface {
	Map(src map[string]interface{}) ([]byte, error)
	IsAppliable(command string) bool
}

type MapperManager struct {
	conf    config.IConfigStorage
	log     *logger.ContextLogManager
	Mappers []IObjectMapper
}

func NewMapperManager(conf config.IConfigStorage, log *logger.ContextLogManager) *MapperManager {
	return &MapperManager{
		conf:    conf,
		log:     log,
		Mappers: []IObjectMapper{},
	}
}

func (m *MapperManager) Map(command string, src map[string]interface{}) ([]byte, error) {
	if command == "" {
		return nil, errors.New("command should not be empty")
	}
	if len(src) == 0 {
		return nil, errors.New("src should not be an empty or nil map")
	}
	found := false
	i := 0
	for i = range m.Mappers {
		if m.Mappers[i].IsAppliable(command) {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("there is no appliable mapper for this command")
	}

	return m.Mappers[i].Map(src)
}

func (m *MapperManager) Register(mappers ...IObjectMapper) {
	for i := range mappers {
		if mappers[i] == nil {
			continue
		}
		m.Mappers = append(m.Mappers, mappers[i])
	}

}
