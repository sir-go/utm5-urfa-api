package main

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
	"strings"

	"urfa-go/internal/urfa"
	"urfa-go/internal/urfa/fn"
)

type RpcReq struct {
	Id      int                    `json:"id"`
	JsonRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
}

type RpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type RpcResp struct {
	Id     int
	Result interface{}
	Error  *RpcError
}

const (
	RpcErrParse          = -32700
	RpcErrInvalidRequest = -32600
	RpcErrMethodNotFound = -32601
	RpcErrInvalidParams  = -32602
	RpcErrInternal       = -32603

	RpcInErrUTMconn = -32005
	RpcInErrUTMcall = -32006
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

	case RpcInErrUTMconn:
		u.Error.Message = "UTM connection error"
		u.Error.Message = "Internal error"
	case RpcInErrUTMcall:
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

type ApiServer struct {
	utmServers CfgUtm
	methods    fn.FMap
}

func NewApiServer(cfg CfgUtm) *ApiServer {
	return &ApiServer{cfg, fn.InitMap()}
}

func (s *ApiServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	answer := &RpcResp{}
	defer writeAnswer(answer, resp)

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

	connParams, billingIsFound := s.utmServers[billingName]
	methodImpl, methodIsFound := s.methods[methodName]

	if !billingIsFound || !methodIsFound {
		answer.Error = &RpcError{Code: RpcErrMethodNotFound}
		return
	}

	username, password, ok := req.BasicAuth()
	if !ok {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	utmConn, err := urfa.Connect(connParams.Addr, username, password, connParams.Cert)
	if err != nil {
		answer.Error = &RpcError{Code: RpcInErrUTMconn, Data: err.Error()}
		return
	}
	defer func() {
		if err := utmConn.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	// log.Debugf("call %s (%v)", methodName, reqObj.Params)
	log.Debugf("%s (%s) calls '%s.%s'", username, req.Header.Get("X-Real-IP"), billingName, methodName)

	defer func() {
		if r := recover(); r != nil {

			answer.Error = &RpcError{Code: RpcInErrUTMcall, Data: r.(error).Error()}

			// todo: debug panic
			log.Error(string(debug.Stack()))
		}
	}()
	answer.Result = methodImpl(utmConn, reqObj.Params)
}

func writeAnswer(answer *RpcResp, respW http.ResponseWriter) {
	answerBuff, err := json.Marshal(answer)
	if err != nil {
		respW.WriteHeader(http.StatusInternalServerError)
	} else {
		// todo: broken pip crashes service
		if _, err = respW.Write(answerBuff); err != nil {
			log.Error(err)
		}
	}
}

func getMethod(reqObj *RpcReq) (bName string, mName string, ok bool) {
	methodParts := strings.Split(reqObj.Method, ".")
	if len(methodParts) < 2 {
		return
	}
	bName, mName = methodParts[0], methodParts[1]
	ok = true
	return
}
