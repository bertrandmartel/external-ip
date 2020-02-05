package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const apifyURL = "https://api.ipify.org?format=json"
const googleDNSURL = "https://dns.google.com/resolve?type=A"

const defaultPort = 3000
const defaultHostname = "example.com"

type Result struct {
	OutboundIP string `json:"outbound_ip"`
	InboundIP  string `json:"inbound_ip"`
	TicToc     string `json:"tic_toc"`
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type ApifyResponse struct {
	IP string `json:"ip"`
}

type GoogleDNSResponse struct {
	Answer []GoogleDNSResult `json:"Answer"`
}

type GoogleDNSResult struct {
	Data string `json:"data"`
}

func main() {

	port := getEnv("PORT", strconv.Itoa(defaultPort))
	host := getEnv("HOSTNAME", defaultHostname)

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		start := time.Now()
		apifyResponse := new(ApifyResponse)
		err := fetch(httpClient, apifyResponse, apifyURL)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, SendError("server_error", "an error occurred while fetching outbound ip"))
		}
		log.Println(apifyResponse)

		googleDNSResponse := new(GoogleDNSResponse)
		err = fetch(httpClient, googleDNSResponse, fmt.Sprintf("%v&name=%v", googleDNSURL, host))
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, SendError("server_error", "an error occurred while fetching inbound ip"))
		}

		log.Println(googleDNSResponse)
		inboundIP := ""
		for i := 0; i < len(googleDNSResponse.Answer); i++ {
			inboundIP += googleDNSResponse.Answer[i].Data + ","
		}
		inboundIP = strings.TrimSuffix(inboundIP, ",")
		elapsed := time.Since(start).String()
		return c.JSON(http.StatusOK, &Result{
			OutboundIP: apifyResponse.IP,
			InboundIP:  inboundIP,
			TicToc:     elapsed,
		})
	})
	e.Logger.Fatal(e.Start(":" + port))
}

func fetch(httpClient *http.Client, target interface{}, url string) error {
	if httpClient == nil {
		return errors.New("no http client specified")
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")

	r, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode == 404 {
		return errors.New("not found")
	}
	if r.StatusCode != 200 {
		return errors.New("received incorrect status : " + strconv.Itoa(r.StatusCode))
	}
	return json.NewDecoder(r.Body).Decode(target)
}

func SendError(errorMessage string, errorDescription string) *ErrorResponse {
	return &ErrorResponse{
		Error:            errorMessage,
		ErrorDescription: errorDescription,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
