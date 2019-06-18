package cmd

import (
	"fmt"
	"github.com/rgolangh/eddy/pkg/eddy"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
	"os"
)

var ServiceUnit eddy.ServiceUnit
var SocketUnit eddy.SocketUnit

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a systemd unit file and write the output",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var serviceCmd = &cobra.Command{
	Use: "service",
	Short: "create a unit type service",
	Long: `"create a unit type service",
longer description
`,

	Run: func(cmd *cobra.Command, args []string) {
		ini.PrettyFormat=false
		file := ini.Empty()
		err := ini.ReflectFrom(file, &ServiceUnit)
		if err != nil {
			cmd.PrintErr(err)
		}
		file.WriteTo(os.Stdout)
	},
}

var socketCmd = &cobra.Command{
	Use: "socket",
	Short: "create a unit type socket",
	Long: `"create a unit type socket",
longer description
`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create socket unit called")
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
