/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// runUrlCmd represents the runUrl command
var runUrlCmd = &cobra.Command{
	Use:   "run-url",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		urlValue, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalln("💥 Error with url flags", err)
		}

		_, err = url.ParseRequestURI(urlValue)
		if err != nil {
			log.Fatalln("💥 Url coould not be parsed", err)
		}

		runCompose(urlValue)
	},
}

func init() {
	rootCmd.AddCommand(runUrlCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runUrlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runUrlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	runUrlCmd.Flags().StringP("url", "u", "", "The url of the docker-compose file")
	runUrlCmd.MarkPersistentFlagRequired("url")
}

func runCompose(url string) {
	r, err := http.Get(url)
	if err != nil {
		log.Fatalln("💥 Error retrieving the url", err)
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("💥 Error reading the content from url", err)
	}

	location, _ := os.Executable()
	err = os.Mkdir(path.Join(location, url), 0755)
	if err != nil {
		log.Fatalln("💥 Error reading the content2 from url", err)
	}
}
