package eddy

import (
	"gopkg.in/ini.v1"
	"strings"
)

func Basic() UnitFile {

	return UnitFile{
		Unit: Unit{
			"desc",
			"",
			"",
		},
		Service: Service{},
		Install: Install{},
	}
}

// Write will write the given unit file and return a reader,
// unless an error has occurred
func Write(unitfile UnitFile) (string, error) {
	iniFile, err := ToIniFile(unitfile)
	if err != nil {
		return "", err
	}

	stringBuilder := strings.Builder{}
	_, err = iniFile.WriteTo(&stringBuilder)

	if err != nil {
		return "", err
	}

	return stringBuilder.String(), nil
}

func ToIniFile(unitFile UnitFile) (*ini.File, error) {
	ini.PrettyFormat=false
	file := ini.Empty()
	err := ini.ReflectFrom(file, &unitFile)
	if err != nil {
		return nil, err
	}
	return file, nil
}
