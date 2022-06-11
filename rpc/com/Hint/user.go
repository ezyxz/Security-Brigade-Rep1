// Code generated by Thrift Compiler (0.16.0). DO NOT EDIT.

package Hint

import (
	"bytes"
	"context"
	"fmt"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// Attributes:
//  - ResultCode
//  - Hints
type HintResult_ struct {
	ResultCode int32    `thrift:"resultCode,1" db:"resultCode" json:"resultCode"`
	Hints      []string `thrift:"hints,2" db:"hints" json:"hints"`
}

func NewHintResult_() *HintResult_ {
	return &HintResult_{}
}

func (p *HintResult_) GetResultCode() int32 {
	return p.ResultCode
}

func (p *HintResult_) GetHints() []string {
	return p.Hints
}
func (p *HintResult_) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *HintResult_) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ResultCode = v
	}
	return nil
}

func (p *HintResult_) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Hints = tSlice
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.Hints = append(p.Hints, _elem0)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *HintResult_) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "HintResult"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *HintResult_) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "resultCode", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:resultCode: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.ResultCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.resultCode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:resultCode: ", p), err)
	}
	return err
}

func (p *HintResult_) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "hints", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:hints: ", p), err)
	}
	if err := oprot.WriteListBegin(ctx, thrift.STRING, len(p.Hints)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Hints {
		if err := oprot.WriteString(ctx, string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:hints: ", p), err)
	}
	return err
}

func (p *HintResult_) Equals(other *HintResult_) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.ResultCode != other.ResultCode {
		return false
	}
	if len(p.Hints) != len(other.Hints) {
		return false
	}
	for i, _tgt := range p.Hints {
		_src1 := other.Hints[i]
		if _tgt != _src1 {
			return false
		}
	}
	return true
}

func (p *HintResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HintResult_(%+v)", *p)
}

type GetHintService interface {
	TestConnection(ctx context.Context) (_r int32, _err error)
	// Parameters:
	//  - Word
	GetString(ctx context.Context, word string) (_r *HintResult_, _err error)
}

type GetHintServiceClient struct {
	c    thrift.TClient
	meta thrift.ResponseMeta
}

func NewGetHintServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *GetHintServiceClient {
	return &GetHintServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewGetHintServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *GetHintServiceClient {
	return &GetHintServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewGetHintServiceClient(c thrift.TClient) *GetHintServiceClient {
	return &GetHintServiceClient{
		c: c,
	}
}

func (p *GetHintServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *GetHintServiceClient) LastResponseMeta_() thrift.ResponseMeta {
	return p.meta
}

func (p *GetHintServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	p.meta = meta
}

func (p *GetHintServiceClient) TestConnection(ctx context.Context) (_r int32, _err error) {
	var _args2 GetHintServiceTestConnectionArgs
	var _result4 GetHintServiceTestConnectionResult
	var _meta3 thrift.ResponseMeta
	_meta3, _err = p.Client_().Call(ctx, "testConnection", &_args2, &_result4)
	p.SetLastResponseMeta_(_meta3)
	if _err != nil {
		return
	}
	return _result4.GetSuccess(), nil
}

// Parameters:
//  - Word
func (p *GetHintServiceClient) GetString(ctx context.Context, word string) (_r *HintResult_, _err error) {
	var _args5 GetHintServiceGetStringArgs
	_args5.Word = word
	var _result7 GetHintServiceGetStringResult
	var _meta6 thrift.ResponseMeta
	_meta6, _err = p.Client_().Call(ctx, "getString", &_args5, &_result7)
	p.SetLastResponseMeta_(_meta6)
	if _err != nil {
		return
	}
	if _ret8 := _result7.GetSuccess(); _ret8 != nil {
		return _ret8, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "getString failed: unknown result")
}

type GetHintServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      GetHintService
}

func (p *GetHintServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *GetHintServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *GetHintServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewGetHintServiceProcessor(handler GetHintService) *GetHintServiceProcessor {

	self9 := &GetHintServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self9.processorMap["testConnection"] = &getHintServiceProcessorTestConnection{handler: handler}
	self9.processorMap["getString"] = &getHintServiceProcessorGetString{handler: handler}
	return self9
}

func (p *GetHintServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
	if err2 != nil {
		return false, thrift.WrapTException(err2)
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(ctx, thrift.STRUCT)
	iprot.ReadMessageEnd(ctx)
	x10 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
	x10.Write(ctx, oprot)
	oprot.WriteMessageEnd(ctx)
	oprot.Flush(ctx)
	return false, x10

}

type getHintServiceProcessorTestConnection struct {
	handler GetHintService
}

func (p *getHintServiceProcessorTestConnection) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetHintServiceTestConnectionArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "testConnection", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := GetHintServiceTestConnectionResult{}
	var retval int32
	if retval, err2 = p.handler.TestConnection(ctx); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing testConnection: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "testConnection", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = &retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "testConnection", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

type getHintServiceProcessorGetString struct {
	handler GetHintService
}

func (p *getHintServiceProcessorGetString) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetHintServiceGetStringArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "getString", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := GetHintServiceGetStringResult{}
	var retval *HintResult_
	if retval, err2 = p.handler.GetString(ctx, args.Word); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getString: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "getString", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "getString", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type GetHintServiceTestConnectionArgs struct {
}

func NewGetHintServiceTestConnectionArgs() *GetHintServiceTestConnectionArgs {
	return &GetHintServiceTestConnectionArgs{}
}

func (p *GetHintServiceTestConnectionArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(ctx, fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetHintServiceTestConnectionArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "testConnection_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetHintServiceTestConnectionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetHintServiceTestConnectionArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GetHintServiceTestConnectionResult struct {
	Success *int32 `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGetHintServiceTestConnectionResult() *GetHintServiceTestConnectionResult {
	return &GetHintServiceTestConnectionResult{}
}

var GetHintServiceTestConnectionResult_Success_DEFAULT int32

func (p *GetHintServiceTestConnectionResult) GetSuccess() int32 {
	if !p.IsSetSuccess() {
		return GetHintServiceTestConnectionResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *GetHintServiceTestConnectionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetHintServiceTestConnectionResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetHintServiceTestConnectionResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *GetHintServiceTestConnectionResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "testConnection_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetHintServiceTestConnectionResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.I32, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteI32(ctx, int32(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *GetHintServiceTestConnectionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetHintServiceTestConnectionResult(%+v)", *p)
}

// Attributes:
//  - Word
type GetHintServiceGetStringArgs struct {
	Word string `thrift:"word,1" db:"word" json:"word"`
}

func NewGetHintServiceGetStringArgs() *GetHintServiceGetStringArgs {
	return &GetHintServiceGetStringArgs{}
}

func (p *GetHintServiceGetStringArgs) GetWord() string {
	return p.Word
}
func (p *GetHintServiceGetStringArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetHintServiceGetStringArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Word = v
	}
	return nil
}

func (p *GetHintServiceGetStringArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getString_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetHintServiceGetStringArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "word", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:word: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Word)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.word (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:word: ", p), err)
	}
	return err
}

func (p *GetHintServiceGetStringArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetHintServiceGetStringArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GetHintServiceGetStringResult struct {
	Success *HintResult_ `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGetHintServiceGetStringResult() *GetHintServiceGetStringResult {
	return &GetHintServiceGetStringResult{}
}

var GetHintServiceGetStringResult_Success_DEFAULT *HintResult_

func (p *GetHintServiceGetStringResult) GetSuccess() *HintResult_ {
	if !p.IsSetSuccess() {
		return GetHintServiceGetStringResult_Success_DEFAULT
	}
	return p.Success
}
func (p *GetHintServiceGetStringResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetHintServiceGetStringResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetHintServiceGetStringResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &HintResult_{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *GetHintServiceGetStringResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getString_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetHintServiceGetStringResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *GetHintServiceGetStringResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetHintServiceGetStringResult(%+v)", *p)
}