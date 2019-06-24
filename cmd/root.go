package cmd

import (
	"fmt"
	"github.com/rsteube/webview-login/login"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"os/exec"
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
			Clear:    Clear,
			Domain:   Domain,
			Keyring:  Keyring,
			LoginUrl: args[0],
			Match:    Match,
			Verbose:  Verbose,
		}

		if cookie := login.Login(); cookie != "" {
			if Shell {
				startShell(cookie)
			} else {
				fmt.Println(cookie)
			}
		} else {
			os.Exit(1)
		}
	},
}

func startShell(cookie string) {
	shell := exec.Command("bash", "-c", `bash --init-file <(echo "source ~/.bashrc; PS1='[webview-login] '; alias curl='curl -H \"Cookie:`+cookie+`\"'")`)
	shell.Stdout = os.Stdout
	shell.Stdin = os.Stdin
	shell.Stderr = os.Stderr
	shell.Run()
}

var Alias bool
var Clear bool
var Keyring bool
var Verbose bool
var Match string
var Domain string
var Shell bool

func init() {
	// TODO
	rootCmd.PersistentFlags().BoolVarP(&Alias, "alias", "a", false, "TODO set alias for current shell")
	rootCmd.PersistentFlags().BoolVarP(&Shell, "shell", "s", false, "start interactive shell")
	rootCmd.PersistentFlags().StringVarP(&Domain, "domain", "d", "", "cookie domain (default \"{scheme}://{host}\" of login-url)")
	rootCmd.PersistentFlags().BoolVarP(&Clear, "clear", "c", false, "clear domain from keyring")
	rootCmd.PersistentFlags().BoolVarP(&Keyring, "keyring", "k", false, "store/retrieve cookie in/from keyring (webview-login/{domain})")
	rootCmd.PersistentFlags().StringVarP(&Match, "match", "m", ".*(_oauth2_proxy)=.*", "cookie regex")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
