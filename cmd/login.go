package cmd

import (
	opaal "davidallendj/opaal/internal"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Start the login flow",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			err := opaal.Login(&config)
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			} else if config.RunOnce {
				break
			}
		}
	},
}

func init() {
	loginCmd.Flags().StringVar(&config.Client.Id, "client.id", config.Client.Id, "set the client ID")
	loginCmd.Flags().StringVar(&config.Client.Secret, "client.secret", config.Client.Secret, "set the client secret")
	loginCmd.Flags().StringSliceVar(&config.Client.RedirectUris, "redirect-uri", config.Client.RedirectUris, "set the redirect URI")
	loginCmd.Flags().StringVar(&config.ResponseType, "response-type", config.ResponseType, "set the response-type")
	loginCmd.Flags().StringSliceVar(&config.Scope, "scope", config.Scope, "set the scopes")
	loginCmd.Flags().StringVar(&config.State, "state", config.State, "set the state")
	loginCmd.Flags().StringVar(&config.Server.Host, "host", config.Server.Host, "set the listening host")
	loginCmd.Flags().IntVar(&config.Server.Port, "port", config.Server.Port, "set the listening port")
	loginCmd.Flags().BoolVar(&config.OpenBrowser, "open-browser", config.OpenBrowser, "automatically open link in browser")
	loginCmd.Flags().BoolVar(&config.DecodeIdToken, "decode-id-token", config.DecodeIdToken, "decode and print ID token from identity provider")
	loginCmd.Flags().BoolVar(&config.DecodeAccessToken, "decore-access-token", config.DecodeAccessToken, "decode and print access token from authorization server")
	loginCmd.Flags().BoolVar(&config.RunOnce, "once", config.RunOnce, "set whether to run login once and exit")
	rootCmd.AddCommand(loginCmd)
}
