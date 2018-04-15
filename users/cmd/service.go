package cmd

import (
	"github.com/spf13/cobra"
	"go-micro-demo/users/starter"
)

var UserServiceCMD = &cobra.Command{
	Use: "userservice" ,
	Short: "manager users services" ,
}

var StartCMD = &cobra.Command{
	Use: "start",
	Short: "start the user service",
	Run: func(cmd *cobra.Command, args []string) {
		service := starter.NewService()
		err := service.Start()
		if err != nil {
			service.Stop()
			panic(err)
		}
	},
}

func init()  {
	UserServiceCMD.AddCommand(StartCMD)
}

func Execute() error {
	return UserServiceCMD.Execute()
}
