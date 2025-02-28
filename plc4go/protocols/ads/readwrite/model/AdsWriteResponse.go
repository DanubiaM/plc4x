/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsWriteResponse is the corresponding interface of AdsWriteResponse
type AdsWriteResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
}

// AdsWriteResponseExactly can be used when we want exactly this type and not a type which fulfills AdsWriteResponse.
// This is useful for switch cases.
type AdsWriteResponseExactly interface {
	AdsWriteResponse
	isAdsWriteResponse() bool
}

// _AdsWriteResponse is the data-structure of this message
type _AdsWriteResponse struct {
	*_AmsPacket
	Result ReturnCode
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsWriteResponse) GetCommandId() CommandId {
	return CommandId_ADS_WRITE
}

func (m *_AdsWriteResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsWriteResponse) InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) {
	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsWriteResponse) GetParent() AmsPacket {
	return m._AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsWriteResponse) GetResult() ReturnCode {
	return m.Result
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsWriteResponse factory function for _AdsWriteResponse
func NewAdsWriteResponse(result ReturnCode, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsWriteResponse {
	_result := &_AdsWriteResponse{
		Result:     result,
		_AmsPacket: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsWriteResponse(structType interface{}) AdsWriteResponse {
	if casted, ok := structType.(AdsWriteResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsWriteResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsWriteResponse) GetTypeName() string {
	return "AdsWriteResponse"
}

func (m *_AdsWriteResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsWriteResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsWriteResponseParse(theBytes []byte) (AdsWriteResponse, error) {
	return AdsWriteResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AdsWriteResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsWriteResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsWriteResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsWriteResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for result")
	}
	_result, _resultErr := ReturnCodeParseWithBuffer(ctx, readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field of AdsWriteResponse")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for result")
	}

	if closeErr := readBuffer.CloseContext("AdsWriteResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsWriteResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsWriteResponse{
		_AmsPacket: &_AmsPacket{},
		Result:     result,
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsWriteResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsWriteResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsWriteResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsWriteResponse")
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for result")
		}
		_resultErr := writeBuffer.WriteSerializable(ctx, m.GetResult())
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for result")
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		if popErr := writeBuffer.PopContext("AdsWriteResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsWriteResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsWriteResponse) isAdsWriteResponse() bool {
	return true
}

func (m *_AdsWriteResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
