package cmd

import (
	"fmt"
	"github.com/rsteube/webview-login/login"
	"github.com/spf13/cobra"
	"net/url"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "webview-login [login-url]",
	Short: "TODO",
	Long: `TODO
                TODO
                TODO`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if Domain == "" {
			if u, err := url.Parse(args[0]); err != nil {
				panic(err)
			} else {
				// TODO support with/without port
				//host, _, _ := net.SplitHostPort(u.Host)
				//Domain = u.Scheme + "://" + host
				Domain = u.Scheme + "://" + u.Host
			}
		}

		login := &login.WebViewLogin{
			Domain:   Domain,
			LoginUrl: args[0],
			Match:    Match,
			Verbose:  Verbose,
		}

		if cookie := login.Login(); cookie != "" {
			fmt.Println(cookie)
		} else {
			os.Exit(1)
		}
	},
}

var Verbose bool
var Match string
var Domain string

func init() {
	rootCmd.PersistentFlags().StringVarP(&Domain, "domain", "d", "", "cookie domain (default \"{scheme}://{host}\" of LoginUrl)")
	rootCmd.PersistentFlags().StringVarP(&Match, "match", "m", ".*(_oauth2_proxy)=.*", "cookie regex")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
