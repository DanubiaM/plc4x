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

// ApduDataExtKeyResponse is the corresponding interface of ApduDataExtKeyResponse
type ApduDataExtKeyResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtKeyResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtKeyResponse.
// This is useful for switch cases.
type ApduDataExtKeyResponseExactly interface {
	ApduDataExtKeyResponse
	isApduDataExtKeyResponse() bool
}

// _ApduDataExtKeyResponse is the data-structure of this message
type _ApduDataExtKeyResponse struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtKeyResponse) GetExtApciType() uint8 {
	return 0x14
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtKeyResponse) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtKeyResponse) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtKeyResponse factory function for _ApduDataExtKeyResponse
func NewApduDataExtKeyResponse(length uint8) *_ApduDataExtKeyResponse {
	_result := &_ApduDataExtKeyResponse{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtKeyResponse(structType interface{}) ApduDataExtKeyResponse {
	if casted, ok := structType.(ApduDataExtKeyResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtKeyResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtKeyResponse) GetTypeName() string {
	return "ApduDataExtKeyResponse"
}

func (m *_ApduDataExtKeyResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtKeyResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtKeyResponseParse(theBytes []byte, length uint8) (ApduDataExtKeyResponse, error) {
	return ApduDataExtKeyResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtKeyResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtKeyResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtKeyResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtKeyResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtKeyResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtKeyResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtKeyResponse{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtKeyResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtKeyResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtKeyResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtKeyResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtKeyResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtKeyResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtKeyResponse) isApduDataExtKeyResponse() bool {
	return true
}

func (m *_ApduDataExtKeyResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
