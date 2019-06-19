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

	Run: func(cmd *cobra.Command, args []string) {
		iniFile, err := eddy.ToIniFile(&ServiceUnit)
		if err != nil {
			cmd.PrintErr(err)
		}
		iniFile.WriteTo(os.Stdout)
	},
}

var socketCmd = &cobra.Command{
	Use:   "socket",
	Short: "create a unit type socket",
	Long: `"create a unit type socket",
longer description
`,

	Run: func(cmd *cobra.Command, args []string) {
		err := eddy.Write(&SocketUnit, os.Stdout)
		if err != nil {
			cmd.PrintErr(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(serviceCmd)
	createCmd.AddCommand(socketCmd)

	serviceCmd.Flags().StringVar(&ServiceUnit.Unit.Description, "description", "", "a description of the service")
	serviceCmd.Flags().StringVar(&ServiceUnit.Service.ExecStart, "exec-start", "", "exec start of the service")
	serviceCmd.Flags().StringVar(&ServiceUnit.Service.PIDFile, "pid-file", "", "the pid file of the service")
	serviceCmd.Flags().StringVar(&ServiceUnit.Service.Type, "service-type", "", "the type of the service e.g forking")
	serviceCmd.Flags().StringArrayVar(&ServiceUnit.Install.WantedBy, "install-wanted-by", []string{}, "a WantedBy specification, e.g multi-user.target")
	serviceCmd.Flags().StringArrayVar(&ServiceUnit.Install.RequiredBy, "install-required-by", []string{}, "a WantedBy specification, e.g multi-user.target")

	socketCmd.Flags().StringVar(&SocketUnit.Unit.Description, "description", "", "a description of the socket")
}
