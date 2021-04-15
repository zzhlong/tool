package httpclient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

var client *http.Client

func init() {
	tr := &http.Transport{
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 5,
		MaxConnsPerHost:     10,
		IdleConnTimeout:     time.Second * 3,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
}

//PostFormData form表单请求提交
func PostFormData(url string, headers map[string]string, params []FromDataParams) (*ResponseBody, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for _, p := range params {
		if p.ParamsType == File {
			w, err := writer.CreateFormFile(p.ParamsKey, p.FileName)
			if err != nil {
				return nil, err
			}
			stream := bytes.NewReader(*p.ParamsContent)
			if _, err := io.Copy(w, stream); err != nil {
				return nil, err
			}
		} else if p.ParamsType == Test {
			if err := writer.WriteField(p.ParamsKey, string(*p.ParamsContent)); err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("params ParamsType invalid")
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
		client.CloseIdleConnections()
	}()
	body, err := ioutil.ReadAll(res.Body)
	return &ResponseBody{Body: body, StateCode: res.StatusCode}, nil
}

//Get httpGet请求
func Get(url string, headers map[string]string, params map[string]string) (*ResponseBody, error) {
	u := strings.Builder{}
	u.WriteString(url)
	u.WriteString("?")
	for k, v := range params {
		u.WriteString(k)
		u.WriteString("=")
		u.WriteString(v)
		u.WriteString("&")
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
		client.CloseIdleConnections()
	}()
	body, err := ioutil.ReadAll(res.Body)
	return &ResponseBody{Body: body, StateCode: res.StatusCode}, nil
}
