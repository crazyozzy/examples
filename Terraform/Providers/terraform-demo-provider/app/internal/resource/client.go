package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var UrlBase *string
var HttpClient *http.Client

type Client struct {
	apiVersion string
	BaseUrl    *url.URL
	httpClient *http.Client
	headers    http.Header
}

func NewClient(h http.Header, hostname, version string) *Client {
	client := &Client{apiVersion: version}

	client.httpClient = &http.Client{
		Timeout: time.Second * 30,
	}

	urlstr := fmt.Sprintf("%s/%s", hostname, version)
	if u, err := url.Parse(urlstr); err != nil {
		panic("Could not init Provider client to Resource microservice")
	} else {
		client.BaseUrl = u
	}

	if h != nil {
		client.headers = h
	}

	return client
}

func (c *Client) doGetDetailsByName(baseURL, resourceName string) (b []byte, err error) {
	path := fmt.Sprintf("http://%s/detail", baseURL)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = fmt.Sprintf("name=%s", resourceName)
	for name, value := range c.headers {
		req.Header.Add(name, strings.Join(value, ""))
	}

	log.Printf("Calling %s\n", req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable to issue get details API call: %s", err.Error())
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("Content-Type of HTTP response is invalid")
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse response from name allocation API: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Name allocation API returned HTTP status code %d", resp.StatusCode)
	}

	log.Printf("Raw output: %v\n", string(b))
	return
}

func (c *Client) doCreateResource(baseURL, resourceType, region string) (b []byte, err error) {
	path := fmt.Sprintf("http://%s/create", baseURL)
	log.Printf("Create resource %s %s \n", resourceType, region)

	reqBody, err := json.Marshal(map[string]string{
		"resource_type": resourceType,
		"region":        region,
	})

	if err != nil {
		return nil, fmt.Errorf("JSON Marchal error: %s", err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	for name, value := range c.headers {
		req.Header.Add(name, strings.Join(value, ""))
	}

	log.Printf("Calling %s\n", req.URL.String())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable to issue name allocation API call: %s", err.Error())
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("Content-Type of HTTP response is invalid")
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse response from name allocation API: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Name allocation API returned HTTP status code %d", resp.StatusCode)
	}

	log.Printf("Raw output: %v\n", string(b))
	return
}
