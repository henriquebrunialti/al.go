//Package cmd holds all the CLI commands such as "al.go sorting mergesort" or "al.go misc moving-rectangle"
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "al.go",
	Short: "CLI Algorithms and Data Structures Visualizer",
	Long:  `This is a CLI Algorithms and Data Structures Visualizer`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("This is a CLI Algorithms and Data structures Visualizer. Use help or -h to checkout supported algorithms")
		fmt.Printf("\n\nUsage Example:\nal.go [Algorithm Type] [Algorithm Name]\n")

		fmt.Printf("\nTYPES\n\n")

		fmt.Printf("al.go misc                 Fun Miscellaneous Animations\n")
		fmt.Printf("al.go sorting              Sorting Algorithms\n")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.al.go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".al.go" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".al.go")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
