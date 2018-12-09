package eddy

type UnitFile struct {
	Unit 	Unit
	Section Section
	Install Install
}

type Section interface {

}

type Unit struct {
	Description string
	Requires 	string
	After 		string
}

type Install struct {
	WantedBy	string
	RequiredBy	string
}

func Basic() UnitFile  {

	return UnitFile{
		Unit: Unit{
			"desc",
			"requires",
			"after",
		},
		Section: nil,
		Install: Install{},
	}
}