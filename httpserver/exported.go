package httpserver

import (
	// stdlib
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	// local
	"sources.dev.pztrn.name/gonews/gonews/configuration"

	// other
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	Srv *echo.Echo
)

func Initialize() {
	log.Println("Initializing HTTP server...")

	Srv = echo.New()
	Srv.Use(middleware.Recover())
	Srv.Use(middleware.Logger())
	Srv.DisableHTTP2 = true
	Srv.HideBanner = true
	Srv.HidePort = true
	Srv.Binder = echo.Binder(&StrictJSONBinder{})

	Srv.GET("/_internal/waitForOnline", waitForHTTPServerToBeUpHandler)
}

// Shutdown stops HTTP server. Returns true on success and false on failure.
func Shutdown() {
	log.Println("Shutting down HTTP server...")

	err := Srv.Shutdown(context.Background())
	if err != nil {
		log.Fatalln("Failed to stop HTTP server:", err.Error())
	}

	log.Println("HTTP server shutted down")
}

// Start starts HTTP server and checks that server is ready to process
// requests. Returns true on success and false on failure.
func Start() {
	log.Println("Starting HTTP server on " + configuration.Cfg.HTTP.Listen + "...")

	go func() {
		err := Srv.Start(configuration.Cfg.HTTP.Listen)
		if !strings.Contains(err.Error(), "Server closed") {
			log.Fatalln("HTTP server critical error occurred:", err.Error())
		}
	}()

	// Check that HTTP server was started.
	httpc := &http.Client{Timeout: time.Second * 1}
	checks := 0

	for {
		checks++

		if checks >= configuration.Cfg.HTTP.WaitForSeconds {
			log.Fatalln("HTTP server isn't up after", checks, "seconds")
		}

		time.Sleep(time.Second * 1)

		resp, err := httpc.Get("http://" + configuration.Cfg.HTTP.Listen + "/_internal/waitForOnline")
		if err != nil {
			log.Println("HTTP error occurred, HTTP server isn't ready, waiting (error was: '" + err.Error() + "')")
			continue
		}

		response, err1 := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err1 != nil {
			log.Println("Failed to read response body, HTTP server isn't ready, waiting (error was: '" + err1.Error() + "')")
			continue
		}

		log.Println("HTTP response with status '" + resp.Status + "' received")

		if resp.StatusCode == http.StatusOK {
			if len(response) == 0 {
				log.Println("Response is empty, HTTP server isn't ready, waiting...")
				continue
			}

			log.Printf("Got response with status code %d and body: %+v\n", resp.StatusCode, string(response))

			if len(response) == 17 {
				break
			}
		}
	}
	log.Println("HTTP server is ready to process requests")
}

func waitForHTTPServerToBeUpHandler(ec echo.Context) error {
	response := map[string]string{
		"error": "None",
	}

	return ec.JSON(200, response)
}
