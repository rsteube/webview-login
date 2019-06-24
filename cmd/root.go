package cmd

import (
	"fmt"
	"github.com/rsteube/webview-login/login"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
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

		if Clear {
			if err := keyring.Delete("webview-login", Domain); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}

		if Keyring {
			if p, err := keyring.Get("webview-login", Domain); err == nil {
				fmt.Println(p)
				os.Exit(0)
			}
		}

		login := &login.WebViewLogin{
			Domain:   Domain,
			LoginUrl: args[0],
			Match:    Match,
			Verbose:  Verbose,
		}

		if cookie := login.Login(); cookie != "" {
			if Keyring {
				if err := keyring.Set("webview-login", Domain, cookie); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
			fmt.Println(cookie)

			shell := exec.Command("bash", "-c", `bash --init-file <(echo "source ~/.bashrc; PS1='[webview-login] '; alias curl='curl -H \"Cookie:`+cookie+`\"'")`)
			shell.Stdout = os.Stdout
			shell.Stdin = os.Stdin
			shell.Stderr = os.Stderr
			shell.Run()
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	},
}

var Alias bool
var Clear bool
var Keyring bool
var Verbose bool
var Match string
var Domain string

func init() {
	// TODO
	rootCmd.PersistentFlags().BoolVarP(&Alias, "alias", "a", false, "TODO set alias for current shell")
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
