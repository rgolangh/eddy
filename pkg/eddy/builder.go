package eddy

import (
	"gopkg.in/ini.v1"
	"io"
)

func Basic() UnitFile {

	return UnitFile{
		Unit: Unit{
			Description: "desc",
		},
		Install: Install{},
	}
}

// Write will write the given unit file and return a reader,
// unless an error has occurred
func Write(unit interface{}, writer io.Writer) error {
	iniFile, err := ToIniFile(unit)
	if err != nil {
		return err
	}

	_, err = iniFile.WriteTo(writer)

	return err
}

func ToIniFile(reader interface{}) (*ini.File, error) {
	ini.PrettyFormat = false
	file := ini.Empty()
	err := ini.ReflectFrom(file, reader)
	if err != nil {
		return nil, err
	}
	return file, nil
}
