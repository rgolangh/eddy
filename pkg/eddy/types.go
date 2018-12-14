package eddy

type UnitFile struct {
	Unit    Unit
	Section Section
	Install Install
}

type Unit struct {
	Description string
	Requires    string
	After       string
}

type Install struct {
	WantedBy   string
	RequiredBy string
}

type Section struct {
}
