package cmd

import (
	"github.com/rgolangh/eddy/pkg/eddy"
	"github.com/spf13/cobra"
	"os"
)

var ServiceUnit eddy.ServiceUnit
var SocketUnit eddy.SocketUnit

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a systemd unit file and write the output",
	Long: `create a systemd unit file and write the output. For example:

eddy create service --description "my sleeping daemon" --exec-start "sleep 5m" --install-required-by "multi-user.target"
`,
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "create a unit type service",
	Long: "create a unit type service",

	RunE: func(cmd *cobra.Command, args []string) error {
		return eddy.Write(&ServiceUnit, os.Stdout)

	},
}

var socketCmd = &cobra.Command{
	Use:   "socket",
	Short: "create a unit type socket",
	Long: "create a unit type socket",

	RunE: func(cmd *cobra.Command, args []string) error {
		return eddy.Write(&SocketUnit, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(serviceCmd)
	createCmd.AddCommand(socketCmd)

	serviceCmd.Flags().StringVar(&ServiceUnit.Unit.Description, "description", "", "a description of the service")
	serviceCmd.Flags().StringVar(&ServiceUnit.Unit.Documentation, "documentation", "", "a space-separated list of URIs referencing documentation for this unit or its configuration. description of the service")
	serviceCmd.Flags().StringVar(&ServiceUnit.Unit.Requires, "requires", "", "configures requirement dependencies on other unites")
	serviceCmd.Flags().StringVar(&ServiceUnit.Unit.Requisite, "requisite", "", "similar to requires, units listed here will not be started")
	serviceCmd.Flags().StringVar(&ServiceUnit.Service.ExecStart, "exec-start", "", "exec start of the service")
	serviceCmd.Flags().StringVar(&ServiceUnit.Service.PIDFile, "pid-file", "", "the pid file of the service")
	serviceCmd.Flags().StringVar(&ServiceUnit.Service.Type, "service-type", "", "the type of the service e.g forking")
	serviceCmd.Flags().StringArrayVar(&ServiceUnit.Install.WantedBy, "install-wanted-by", []string{}, "a WantedBy specification, e.g multi-user.target")
	serviceCmd.Flags().StringArrayVar(&ServiceUnit.Install.RequiredBy, "install-required-by", []string{}, "a WantedBy specification, e.g multi-user.target")

	socketCmd.Flags().StringVar(&SocketUnit.Unit.Description, "description", "", "a description of the socket")
	socketCmd.Flags().StringVar(&SocketUnit.Unit.Documentation, "documentation", "", "a space-separated list of URIs referencing documentation for this unit or its configuration. description of the service")
	socketCmd.Flags().StringVar(&SocketUnit.Unit.Requires, "requires", "", "configures requirement dependencies on other unites")
	socketCmd.Flags().StringVar(&SocketUnit.Unit.Requisite, "requisite", "", "similar to requires, units listed here will not be started")
	socketCmd.Flags().StringVar(&SocketUnit.Socket.ExecStartPre, "exec-start-pre", "", "pre exec start of the service")
	socketCmd.Flags().StringVar(&SocketUnit.Socket.ExecStopPre, "exec-stop-post", "", "pre exec stop of the service")
	socketCmd.Flags().StringVar(&SocketUnit.Socket.ListenDatagram, "socket-listen-datagram", "", "")
	socketCmd.Flags().StringVar(&SocketUnit.Socket.ListenSequentialPacket, "socket-listen-sequential-packet", "", "")
	socketCmd.Flags().StringVar(&SocketUnit.Socket.ListenStream, "socket-listen-stream", "", "")
	socketCmd.Flags().StringArrayVar(&SocketUnit.Install.WantedBy, "install-wanted-by", []string{}, "a WantedBy specification, e.g multi-user.target")
	socketCmd.Flags().StringArrayVar(&SocketUnit.Install.RequiredBy, "install-required-by", []string{}, "a WantedBy specification, e.g multi-user.target")

}
