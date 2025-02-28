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

// SetAttributeSingleRequest is the corresponding interface of SetAttributeSingleRequest
type SetAttributeSingleRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
}

// SetAttributeSingleRequestExactly can be used when we want exactly this type and not a type which fulfills SetAttributeSingleRequest.
// This is useful for switch cases.
type SetAttributeSingleRequestExactly interface {
	SetAttributeSingleRequest
	isSetAttributeSingleRequest() bool
}

// _SetAttributeSingleRequest is the data-structure of this message
type _SetAttributeSingleRequest struct {
	*_CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetAttributeSingleRequest) GetService() uint8 {
	return 0x10
}

func (m *_SetAttributeSingleRequest) GetResponse() bool {
	return bool(false)
}

func (m *_SetAttributeSingleRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetAttributeSingleRequest) InitializeParent(parent CipService) {}

func (m *_SetAttributeSingleRequest) GetParent() CipService {
	return m._CipService
}

// NewSetAttributeSingleRequest factory function for _SetAttributeSingleRequest
func NewSetAttributeSingleRequest(serviceLen uint16) *_SetAttributeSingleRequest {
	_result := &_SetAttributeSingleRequest{
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSetAttributeSingleRequest(structType interface{}) SetAttributeSingleRequest {
	if casted, ok := structType.(SetAttributeSingleRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SetAttributeSingleRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SetAttributeSingleRequest) GetTypeName() string {
	return "SetAttributeSingleRequest"
}

func (m *_SetAttributeSingleRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SetAttributeSingleRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SetAttributeSingleRequestParse(theBytes []byte, connected bool, serviceLen uint16) (SetAttributeSingleRequest, error) {
	return SetAttributeSingleRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func SetAttributeSingleRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (SetAttributeSingleRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetAttributeSingleRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetAttributeSingleRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SetAttributeSingleRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetAttributeSingleRequest")
	}

	// Create a partially initialized instance
	_child := &_SetAttributeSingleRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_SetAttributeSingleRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetAttributeSingleRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetAttributeSingleRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetAttributeSingleRequest")
		}

		if popErr := writeBuffer.PopContext("SetAttributeSingleRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetAttributeSingleRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetAttributeSingleRequest) isSetAttributeSingleRequest() bool {
	return true
}

func (m *_SetAttributeSingleRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
