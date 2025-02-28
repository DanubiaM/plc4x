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

// Constant values.
const MultipleServiceRequest_REQUESTPATHSIZE uint8 = 0x02
const MultipleServiceRequest_REQUESTPATH uint32 = 0x01240220

// MultipleServiceRequest is the corresponding interface of MultipleServiceRequest
type MultipleServiceRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetData returns Data (property field)
	GetData() Services
}

// MultipleServiceRequestExactly can be used when we want exactly this type and not a type which fulfills MultipleServiceRequest.
// This is useful for switch cases.
type MultipleServiceRequestExactly interface {
	MultipleServiceRequest
	isMultipleServiceRequest() bool
}

// _MultipleServiceRequest is the data-structure of this message
type _MultipleServiceRequest struct {
	*_CipService
	Data Services
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MultipleServiceRequest) GetService() uint8 {
	return 0x0A
}

func (m *_MultipleServiceRequest) GetResponse() bool {
	return bool(false)
}

func (m *_MultipleServiceRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MultipleServiceRequest) InitializeParent(parent CipService) {}

func (m *_MultipleServiceRequest) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MultipleServiceRequest) GetData() Services {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_MultipleServiceRequest) GetRequestPathSize() uint8 {
	return MultipleServiceRequest_REQUESTPATHSIZE
}

func (m *_MultipleServiceRequest) GetRequestPath() uint32 {
	return MultipleServiceRequest_REQUESTPATH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMultipleServiceRequest factory function for _MultipleServiceRequest
func NewMultipleServiceRequest(data Services, serviceLen uint16) *_MultipleServiceRequest {
	_result := &_MultipleServiceRequest{
		Data:        data,
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMultipleServiceRequest(structType interface{}) MultipleServiceRequest {
	if casted, ok := structType.(MultipleServiceRequest); ok {
		return casted
	}
	if casted, ok := structType.(*MultipleServiceRequest); ok {
		return *casted
	}
	return nil
}

func (m *_MultipleServiceRequest) GetTypeName() string {
	return "MultipleServiceRequest"
}

func (m *_MultipleServiceRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (requestPathSize)
	lengthInBits += 8

	// Const Field (requestPath)
	lengthInBits += 32

	// Simple field (data)
	lengthInBits += m.Data.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MultipleServiceRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MultipleServiceRequestParse(theBytes []byte, connected bool, serviceLen uint16) (MultipleServiceRequest, error) {
	return MultipleServiceRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func MultipleServiceRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (MultipleServiceRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MultipleServiceRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MultipleServiceRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (requestPathSize)
	requestPathSize, _requestPathSizeErr := readBuffer.ReadUint8("requestPathSize", 8)
	if _requestPathSizeErr != nil {
		return nil, errors.Wrap(_requestPathSizeErr, "Error parsing 'requestPathSize' field of MultipleServiceRequest")
	}
	if requestPathSize != MultipleServiceRequest_REQUESTPATHSIZE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", MultipleServiceRequest_REQUESTPATHSIZE) + " but got " + fmt.Sprintf("%d", requestPathSize))
	}

	// Const Field (requestPath)
	requestPath, _requestPathErr := readBuffer.ReadUint32("requestPath", 32)
	if _requestPathErr != nil {
		return nil, errors.Wrap(_requestPathErr, "Error parsing 'requestPath' field of MultipleServiceRequest")
	}
	if requestPath != MultipleServiceRequest_REQUESTPATH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", MultipleServiceRequest_REQUESTPATH) + " but got " + fmt.Sprintf("%d", requestPath))
	}

	// Simple Field (data)
	if pullErr := readBuffer.PullContext("data"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	_data, _dataErr := ServicesParseWithBuffer(ctx, readBuffer, uint16(uint16(serviceLen)-uint16(uint16(6))))
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field of MultipleServiceRequest")
	}
	data := _data.(Services)
	if closeErr := readBuffer.CloseContext("data"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	if closeErr := readBuffer.CloseContext("MultipleServiceRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MultipleServiceRequest")
	}

	// Create a partially initialized instance
	_child := &_MultipleServiceRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		Data: data,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_MultipleServiceRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MultipleServiceRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MultipleServiceRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MultipleServiceRequest")
		}

		// Const Field (requestPathSize)
		_requestPathSizeErr := writeBuffer.WriteUint8("requestPathSize", 8, 0x02)
		if _requestPathSizeErr != nil {
			return errors.Wrap(_requestPathSizeErr, "Error serializing 'requestPathSize' field")
		}

		// Const Field (requestPath)
		_requestPathErr := writeBuffer.WriteUint32("requestPath", 32, 0x01240220)
		if _requestPathErr != nil {
			return errors.Wrap(_requestPathErr, "Error serializing 'requestPath' field")
		}

		// Simple Field (data)
		if pushErr := writeBuffer.PushContext("data"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for data")
		}
		_dataErr := writeBuffer.WriteSerializable(ctx, m.GetData())
		if popErr := writeBuffer.PopContext("data"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for data")
		}
		if _dataErr != nil {
			return errors.Wrap(_dataErr, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("MultipleServiceRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MultipleServiceRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MultipleServiceRequest) isMultipleServiceRequest() bool {
	return true
}

func (m *_MultipleServiceRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
