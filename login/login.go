package login

import (
	"github.com/zserge/webview"
	"regexp"
	"time"
)

type WebViewLogin struct {
	CheckCookie bool
	Domain      string
	LoginUrl    string
	Match       string
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

	r, _ := regexp.Compile(w.Match)

	go w.timer()
	result := ""
	for result == "" && webView.Loop(true) {
		if w.CheckCookie {
			webView.GetCookie(w.Domain, func(cookie string) {
				if r.MatchString(cookie) {
					result = cookie
					webView.Terminate()
				}
			})
			w.CheckCookie = false
		}
	}

	// TODO use err
	return result
}

func (w *WebViewLogin) timer() {
	for {
		time.Sleep(100 * time.Millisecond)
		w.CheckCookie = true
	}
}
