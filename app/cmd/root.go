package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/gSchool/glearn-cli/api/learn"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd is the base for all our commands. It currently just checks for all the
// necessary credentials and prompts the user to set them if they are not there.
var rootCmd = &cobra.Command{
	Use:   "learn [command]",
	Short: "learn is a CLI tool for communicating with Learn",
	Long:  `learn is a CLI tool for communicating with Learn`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires at least 1 argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Unknown command. Try `learn help` for more information")
	},
}

// APIToken is an initialized string used for holding it's flag value
var APIToken string

// UnitsDirectory is a flag for preview command that denotes a location for the units
var UnitsDirectory string

func init() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("Error retrieving your user path information")
		os.Exit(1)
		return
	}

	viper.AddConfigPath(u.HomeDir)
	viper.SetConfigName(".glearn-config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found. Either user's first time using CLI or they deleted it
			configPath := fmt.Sprintf("%s/.glearn-config.yaml", u.HomeDir)
			initialConfig := []byte(`api_token:`)

			// Write a ~/.glearn-config.yaml file with all the needed credential keys to fill in.
			err = ioutil.WriteFile(configPath, initialConfig, 0666)
			if err != nil {
				fmt.Println("Error writing your glearn config file")
				os.Exit(1)
				return
			}
		} else {
			// Config file was found but another error was produced
			fmt.Printf("Error: %s", err)
			os.Exit(1)
			return
		}
	}

	apiToken, ok := viper.Get("api_token").(string)
	if !ok {
		fmt.Println("Please set your api_token in ~/.glearn-config.yaml")
		os.Exit(1)
	}

	client := http.Client{Timeout: 15 * time.Second}
	baseURL := "https://learn-2.galvanize.com"
	alternateURL := os.Getenv("LEARN_BASE_URL")
	if alternateURL != "" {
		baseURL = alternateURL
	}

	learn.API = learn.NewAPI(apiToken, baseURL, &client)

	// Add all the other learn commands defined in cmd/ directory
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(previewCmd)
	rootCmd.AddCommand(publishCmd)

	// Check for flags set by the user and hyrate their corresponding variables.
	setCmd.Flags().StringVarP(&APIToken, "api_token", "", "", "Your Learn api token")
	previewCmd.Flags().StringVarP(&UnitsDirectory, "units", "u", "", "The directory where your units exist")
}

// Execute runs the learn CLI according to the user's command/subcommand/flags
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
