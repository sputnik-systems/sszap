package sszap

import (
	"bytes"
	"io"
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	HTTPRequestKey = "http_request"
)

type HttpRequest struct {
	RequestMethod                  string
	RequestURL                     string
	RequestSize                    string
	Status                         int
	ResponseSize                   string
	UserAgent                      string
	RemoteIP                       string
	ServerIP                       string
	Referer                        string
	Latency                        string
	CacheLookup                    bool
	CacheHit                       bool
	CacheValidatedWithOriginServer bool
	CacheFillBytes                 string
	Protocol                       string
}

func (req *HttpRequest) Clone() *HttpRequest {
	return &HttpRequest{
		RequestMethod:                  req.RequestMethod,
		RequestURL:                     req.RequestURL,
		RequestSize:                    req.RequestSize,
		Status:                         req.Status,
		ResponseSize:                   req.ResponseSize,
		UserAgent:                      req.UserAgent,
		RemoteIP:                       req.RemoteIP,
		ServerIP:                       req.ServerIP,
		Referer:                        req.Referer,
		Latency:                        req.Latency,
		CacheLookup:                    req.CacheLookup,
		CacheHit:                       req.CacheHit,
		CacheValidatedWithOriginServer: req.CacheValidatedWithOriginServer,
		CacheFillBytes:                 req.CacheFillBytes,
		Protocol:                       req.Protocol,
	}
}

func (req *HttpRequest) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("request_method", req.RequestMethod)
	enc.AddString("request_url", req.RequestURL)
	enc.AddString("request_size", req.RequestSize)
	enc.AddInt("status", req.Status)
	enc.AddString("response_size", req.ResponseSize)
	enc.AddString("user_agent", req.UserAgent)
	enc.AddString("remote_ip", req.RemoteIP)
	enc.AddString("server_ip", req.ServerIP)
	enc.AddString("referer", req.Referer)
	enc.AddString("latency", req.Latency)
	enc.AddBool("cache_lookup", req.CacheLookup)
	enc.AddBool("cache_hit", req.CacheHit)
	enc.AddBool("cache_validated_with_origin_server", req.CacheValidatedWithOriginServer)
	enc.AddString("cache_fill_bytes", req.CacheFillBytes)
	enc.AddString("protocol", req.Protocol)

	return nil
}

func NewHttpRequest(req *http.Request, statusCode int, respSize int) *HttpRequest {
	if req == nil {
		req = &http.Request{}
	}

	r := &HttpRequest{
		RequestMethod: req.Method,
		Status:        statusCode,
		UserAgent:     req.UserAgent(),
		RemoteIP:      req.RemoteAddr,
		Referer:       req.Referer(),
		Protocol:      req.Proto,
		ResponseSize:  strconv.Itoa(respSize),
	}

	if req.URL != nil {
		r.RequestURL = req.URL.String()
	}

	buf := new(bytes.Buffer)
	if req.Body != nil {
		n, _ := io.Copy(buf, req.Body)
		r.RequestSize = strconv.FormatInt(n, 10)
	}
	buf.Reset()

	return r
}

func HttpRequestField(req *HttpRequest) zap.Field {
	return zap.Object(HTTPRequestKey, req)
}
