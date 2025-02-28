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

// AdsDeleteDeviceNotificationResponse is the corresponding interface of AdsDeleteDeviceNotificationResponse
type AdsDeleteDeviceNotificationResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
}

// AdsDeleteDeviceNotificationResponseExactly can be used when we want exactly this type and not a type which fulfills AdsDeleteDeviceNotificationResponse.
// This is useful for switch cases.
type AdsDeleteDeviceNotificationResponseExactly interface {
	AdsDeleteDeviceNotificationResponse
	isAdsDeleteDeviceNotificationResponse() bool
}

// _AdsDeleteDeviceNotificationResponse is the data-structure of this message
type _AdsDeleteDeviceNotificationResponse struct {
	*_AmsPacket
	Result ReturnCode
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) GetCommandId() CommandId {
	return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
}

func (m *_AdsDeleteDeviceNotificationResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) {
	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsDeleteDeviceNotificationResponse) GetParent() AmsPacket {
	return m._AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDeleteDeviceNotificationResponse) GetResult() ReturnCode {
	return m.Result
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDeleteDeviceNotificationResponse factory function for _AdsDeleteDeviceNotificationResponse
func NewAdsDeleteDeviceNotificationResponse(result ReturnCode, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsDeleteDeviceNotificationResponse {
	_result := &_AdsDeleteDeviceNotificationResponse{
		Result:     result,
		_AmsPacket: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDeleteDeviceNotificationResponse(structType interface{}) AdsDeleteDeviceNotificationResponse {
	if casted, ok := structType.(AdsDeleteDeviceNotificationResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDeleteDeviceNotificationResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDeleteDeviceNotificationResponse) GetTypeName() string {
	return "AdsDeleteDeviceNotificationResponse"
}

func (m *_AdsDeleteDeviceNotificationResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsDeleteDeviceNotificationResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDeleteDeviceNotificationResponseParse(theBytes []byte) (AdsDeleteDeviceNotificationResponse, error) {
	return AdsDeleteDeviceNotificationResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AdsDeleteDeviceNotificationResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDeleteDeviceNotificationResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsDeleteDeviceNotificationResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDeleteDeviceNotificationResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for result")
	}
	_result, _resultErr := ReturnCodeParseWithBuffer(ctx, readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field of AdsDeleteDeviceNotificationResponse")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for result")
	}

	if closeErr := readBuffer.CloseContext("AdsDeleteDeviceNotificationResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDeleteDeviceNotificationResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsDeleteDeviceNotificationResponse{
		_AmsPacket: &_AmsPacket{},
		Result:     result,
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsDeleteDeviceNotificationResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDeleteDeviceNotificationResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeleteDeviceNotificationResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDeleteDeviceNotificationResponse")
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

		if popErr := writeBuffer.PopContext("AdsDeleteDeviceNotificationResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDeleteDeviceNotificationResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDeleteDeviceNotificationResponse) isAdsDeleteDeviceNotificationResponse() bool {
	return true
}

func (m *_AdsDeleteDeviceNotificationResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
