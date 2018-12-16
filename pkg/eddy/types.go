package eddy

type UnitFile struct {
	Unit    Unit
	Service Service
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

type Service struct {
	ExecStart string
	PidFile   string
	Type      string
}
