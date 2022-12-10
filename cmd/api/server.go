package main

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/pkg/errors"
	zlog "github.com/rs/zerolog/log"

	"urfa-go/internal/urfa"
	"urfa-go/internal/urfa/fn"
)

type (
	// RpcReq - standard json-rpc request
	RpcReq struct {
		Id      int                    `json:"id"`
		JsonRPC string                 `json:"jsonrpc"`
		Method  string                 `json:"method"`
		Params  map[string]interface{} `json:"params"`
	}

	// RpcError - standard json-rpc error struct
	RpcError struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}

	// RpcResp - standard json-rpc response
	RpcResp struct {
		Id     int
		Result interface{}
		Error  *RpcError
	}
)

const (
	RpcErrParse          = -32700
	RpcErrInvalidRequest = -32600
	RpcErrMethodNotFound = -32601
	RpcErrInvalidParams  = -32602
	RpcErrInternal       = -32603

	RpcInErrUtmConn = -32005
	RpcInErrUrmCall = -32006
)

func (u *RpcResp) MarshalJSON() ([]byte, error) {
	if u.Error == nil {
		return json.Marshal(&struct {
			JsonRPC string      `json:"jsonrpc"`
			Id      int         `json:"id"`
			Result  interface{} `json:"result"`
		}{
			JsonRPC: "2.0",
			Id:      u.Id,
			Result:  u.Result,
		})
	}

	switch u.Error.Code {
	case RpcErrParse:
		u.Error.Message = "Parse error"
	case RpcErrInvalidRequest:
		u.Error.Message = "Invalid request"
	case RpcErrMethodNotFound:
		u.Error.Message = "Method not found"
	case RpcErrInvalidParams:
		u.Error.Message = "Invalid params"
	case RpcErrInternal:
		u.Error.Message = "Internal error"

	case RpcInErrUtmConn:
		u.Error.Message = "UTM connection error"
		u.Error.Message = "Internal error"
	case RpcInErrUrmCall:
		u.Error.Message = "UTM call function error"
	}

	return json.Marshal(&struct {
		JsonRPC string    `json:"jsonrpc"`
		Id      int       `json:"id"`
		Error   *RpcError `json:"error"`
	}{
		JsonRPC: "2.0",
		Id:      u.Id,
		Error:   u.Error,
	})
}

type Server struct {
	billings billingsMap
	methods  fn.FMap
}

func NewServer(bMap billingsMap) *Server {
	return &Server{bMap, fn.InitMap()}
}

func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	answer := &RpcResp{}
	defer writeAnswer(answer, resp) // send answer to the client in any case

	// parse the request

	decoder := json.NewDecoder(req.Body)
	reqObj := &RpcReq{}
	if err := decoder.Decode(reqObj); err != nil {
		answer.Error = &RpcError{Code: RpcErrParse, Data: err.Error()}
		return
	}

	answer.Id = reqObj.Id

	billingName, methodName, ok := getMethod(reqObj)
	if !ok {
		answer.Error = &RpcError{Code: RpcErrMethodNotFound}
		return
	}

	// check if we have that billing endpoint and the RPC method

	connParams, billingIsFound := s.billings[billingName]
	methodImpl, methodIsFound := s.methods[methodName]

	if !billingIsFound || !methodIsFound {
		answer.Error = &RpcError{Code: RpcErrMethodNotFound}
		return
	}

	// check auth

	username, password, ok := req.BasicAuth()
	if !ok {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	// connect to certain billing server

	utmConn, err := urfa.Connect(connParams.Addr, username, password, connParams.Cert)
	if err != nil {
		answer.Error = &RpcError{Code: RpcInErrUtmConn, Data: err.Error()}
		return
	}

	// plan to disconnect and make an answer

	defer func() {
		if err := utmConn.Disconnect(); err != nil {
			zlog.Err(err).Str("billing", billingName).Msg("disconnection error")
		}
		if r := recover(); r != nil {
			answer.Error = &RpcError{Code: RpcInErrUrmCall, Data: r.(error).Error()}
			zlog.Err(errors.New(string(debug.Stack())))
		}
	}()

	zlog.Debug().
		Str("username", username).
		Str("ip", req.Header.Get("X-Real-IP")).
		Str("billing", billingName).
		Str("method", methodName).
		Msg("got request")

	// call the method

	answer.Result = methodImpl(utmConn, reqObj.Params)
}

// writeAnswer - send a response to the client in any case
func writeAnswer(answer *RpcResp, respW http.ResponseWriter) {
	answerBuff, err := json.Marshal(answer)
	if err != nil {
		respW.WriteHeader(http.StatusInternalServerError)
	} else {
		if _, err = respW.Write(answerBuff); err != nil {
			zlog.Err(err).Msg("response write error")
		}
	}
}

// getMethod "bill.method" -> "bill", "method", ok if both is ok
func getMethod(reqObj *RpcReq) (bName string, mName string, ok bool) {
	methodParts := strings.Split(reqObj.Method, ".")
	if len(methodParts) < 2 {
		return
	}
	bName, mName = methodParts[0], methodParts[1]
	if len(bName) > 0 && len(mName) > 0 {
		return bName, mName, true
	}
	return "", "", false
}
