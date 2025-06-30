package main

import (
	"fmt"
	"net/http"
)

/*

1) http.Response

type Response struct {
		Status           string // например, "200 OK"
		StatusCode       int    // например, 200
		Proto            string // например, "HTTP/1.0"
		ProtoMajor       int    // например, 1
		ProtoMinor       int    // например, 0
		Header           Header
		Body             io.ReadCloser
		ContentLength    int64
		TransferEncoding []string
		Close            bool
		Uncompressed     bool
		Trailer          Header
		Request          *Request
		TLS              *tls.ConnectionState
	}

	2) http.Request

	type Request struct {
		Method           string
		URL              *url.URL
		Proto            string
		ProtoMajor       int
		ProtoMinor       int
		Header           Header
		Body             io.ReadCloser
		GetBody          func() (io.ReadCloser, error)
		ContentLength    int64
		TransferEncoding []string
		Close            bool
		Host             string
		Form             url.Values
		PostForm         url.Values
		MultipartForm    *multipart.Form
		Trailer          Header
		RemoteAddr       string
		RequestURI       string
		TLS              *tls.ConnectionState
		Cancel           <-chan struct{}
		Response         *Response
	}

	3) http.Transport

	type Transport struct {
		Proxy                  func(*Request) (*url.URL, error)
		DialContext            func(ctx context.Context, network, addr string) (net.Conn, error)
		Dial                   func(network, addr string) (net.Conn, error)
		DialTLSContext         func(ctx context.Context, network, addr string) (net.Conn, error)
		DialTLS                func(network, addr string) (net.Conn, error)
		TLSClientConfig        *tls.Config
		TLSHandshakeTimeout    time.Duration
		DisableKeepAlives      bool
		DisableCompression     bool
		MaxIdleConns           int
		MaxIdleConnsPerHost    int
		MaxConnsPerHost        int
		IdleConnTimeout        time.Duration
		ResponseHeaderTimeout  time.Duration
		ExpectContinueTimeout  time.Duration
		TLSNextProto           map[string]func(authority string, c *tls.Conn) RoundTripper
		ProxyConnectHeader     Header
		GetProxyConnectHeader  func(ctx context.Context, proxyURL *url.URL, target string) (Header, error)
		MaxResponseHeaderBytes int64
		WriteBufferSize        int
		ReadBufferSize         int
		ForceAttemptHTTP2      bool
	}

*/

func main() {
	/*
		мы регистрируем функцию-обработчик для всех запросов, которые приходят на корневой путь (/)
		Функция принимает два аргумента:
		- w http.ResponseWriter: - объект, через который мы отправляем ответ клиенту.
		- r *http.Request: - объект, содержащий информацию о входящем запросе.
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server at port 8080")
	// запускаем сервер на порту 8080
	// Использует стандартные настройки сервера.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Отображаем форму
		fmt.Fprintf(w, `
            <form method="POST">
                <input type="text" name="name" placeholder="Enter your name">
                <button type="submit">Submit</button>
            </form>
        `)
	} else if r.Method == http.MethodPost {
		fmt.Println(r, "this is form")
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "Error parsing form: %v", err)
			return
		}

		name := r.FormValue("name")
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}
