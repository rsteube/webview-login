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

				// TODO verbose
				//fmt.Println(Domain)
			}
		}

		// TODO verbose
		//fmt.Println(Match)

		login := &login.WebViewLogin{
			Domain:   Domain,
			LoginUrl: args[0],
			Match:    Match,
			Verbose:  Verbose,
		}
		fmt.Println(login.Login())

		// Do Stuff Here
	},
}

var Verbose bool
var Match string
var Domain string

func init() {
	rootCmd.PersistentFlags().StringVarP(&Match, "domain", "d", "", "cookie domain")
	rootCmd.PersistentFlags().StringVarP(&Match, "match", "m", ".*(_oauth2_proxy)=.*", "cookie regex")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
