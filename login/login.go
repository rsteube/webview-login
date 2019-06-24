package login

import (
	"fmt"
	"github.com/zserge/webview"
	"os"
	"regexp"
	"time"
)

type WebViewLogin struct {
	CheckCookie bool
	Domain      string
	LoginUrl    string
	Match       string
	Verbose     bool
}

func (w *WebViewLogin) Login() string {
	webView := webview.New(webview.Settings{
		Title:                  "Login",
		URL:                    w.LoginUrl,
		Width:                  1000,
		Height:                 800,
		Resizable:              true,
		Debug:                  true,
		ExternalInvokeCallback: nil,
	})

	if !w.Verbose {
		disableJsLog(webView)
	} else {
		fmt.Fprintf(os.Stderr, "%+v\n", w)
	}

	r, _ := regexp.Compile(w.Match)

	go w.timer()
	result := ""
	for result == "" && webView.Loop(true) {
		if w.CheckCookie {
			webView.GetCookie(w.Domain, func(cookie string) {
				if r.MatchString(cookie) {
					result = cookie
				}
			})
			w.CheckCookie = false
		}
	}

	webView.Exit()
	webView.Terminate()

	// TODO use err
	return result
}

func (w *WebViewLogin) timer() {
	for {
		time.Sleep(100 * time.Millisecond)
		w.CheckCookie = true
	}
}

func disableJsLog(webView webview.WebView) {
	webView.Dispatch(func() {
		webView.Eval("console = function() {}")
	})
}
