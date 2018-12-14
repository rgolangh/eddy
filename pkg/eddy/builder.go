package eddy

import (
	"gopkg.in/ini.v1"
	"io"
)

func Basic() UnitFile {

	return UnitFile{
		Unit: Unit{
			"desc",
			"",
			"",
		},
		Section: struct{}{},
		Install: Install{},
	}
}

// Write will write the given unit file and return a reader,
// unless an error has occurred
func Write(unitfile UnitFile) (*io.PipeReader, error) {
	r, w := io.Pipe()

	iniFile, err := ToIniFile(unitfile)
	if err != nil {
		return nil, err
	}

	_, err = iniFile.WriteTo(w)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func ToIniFile(unitFile UnitFile) (*ini.File, error) {
	file := ini.Empty()
	err := ini.ReflectFrom(file, &unitFile)
	if err != nil {
		return nil, err
	}
	return file, nil
}
