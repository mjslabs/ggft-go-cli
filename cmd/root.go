 package cmd

 import (
	 "fmt"
	 "os"

	 "github.com/spf13/cobra"
 )

 var version = "undefined"

 // rootCmd represents the base command when called without any subcommands
 var rootCmd = &cobra.Command{
	 Use:   "{{ .ProjectName }}",
	 Short: "{{ .ProjectNameFullTitle }}",
	 Long: `
 `,

	 Version: version,
 }

 func init() {
	 cobra.OnInitialize(func() {
		 // Add Viper init here
	 })
 }

 // Execute adds all child commands to the root command and sets flags appropriately.
 // This is called by main.main(). It only needs to happen once to the rootCmd.
 func Execute() {
	 if err := rootCmd.Execute(); err != nil {
		 fmt.Println(err)
		 os.Exit(1)
	 }
 }
