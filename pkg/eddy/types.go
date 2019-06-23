// eddy types modulerize a systemd unit
package eddy

type UnitFile struct {
	Unit    Unit
	Install Install
}

type ServiceUnit struct {
	Unit    Unit
	Service Service
	Install Install
}

type SocketUnit struct {
	Unit    Unit    `ini:",omitempty"`
	Socket  Socket  `ini:",omitempty"`
	Install Install `ini:",omitempty"`
}

type Unit struct {
	Description           string
	Documentation         string `ini:",omitempty"`
	Requires              string `ini:",omitempty"`
	Requisite             string `ini:",omitempty"`
	Wants                 string `ini:",omitempty"`
	BindsTo               string `ini:",omitempty"`
	PartOf                string `ini:",omitempty"`
	Conflicts             string `ini:",omitempty"`
	Before                string `ini:",omitempty"`
	OnFailure             string `ini:",omitempty"`
	After                 string `ini:",omitempty"`
	StartLimitIntervalSec string `ini:",omitempty"`
	StartLimitBurst       string `ini:",omitempty"`
	StartLimitAction      string `ini:",omitempty"`
}

// Install section, used by systemctl enable/disable tool
type Install struct {
	Before     string   `ini:",omitempty"`
	After      string   `ini:",omitempty"`
	WantedBy   []string `ini:",omitempty"`
	RequiredBy []string `ini:",omitempty"`
	Alias      string   `ini:",omitempty"`
}

type Service struct {
	// Type is one of [simple, exec, forking, oneshot, dbus, notify, idle]
	Type            string `ini:",omitempty"`
	ExecStart       string `ini:",omitempty"`
	ExecStop        string `ini:",omitempty"`
	ExecReload      string `ini:",omitempty"`
	ExecStartPre    string `ini:",omitempty"`
	ExecStartPost   string `ini:",omitempty"`
	PIDFile         string `ini:",omitempty"`
	RemainAfterExit bool   `ini:",omitempty"`
	RestartSec      string `ini:",omitempty"`
	TimeoutStartSec string `ini:",omitempty"`
	TimeoutStopSec  string `ini:",omitempty"`
	TimeoutSec      string `ini:",omitempty"`
	RuntimeMaxSec   string `ini:",omitempty"`
	Restart         string `ini:",omitempty"`
}

type Socket struct {
	ExecStartPre           string `ini:",omitempty"`
	ExecStartPost          string `ini:",omitempty"`
	ExecStopPre           string `ini:",omitempty"`
	ExecStopPost          string `ini:",omitempty"`
	ListenStream           string `ini:",omitempty"`
	ListenDatagram         string `ini:",omitempty"`
	ListenSequentialPacket string `ini:",omitempty"`
	Socket string `ini:",omitempty"`
	Accept bool `ini:",omitempty"`

}

type Exec struct {

}
