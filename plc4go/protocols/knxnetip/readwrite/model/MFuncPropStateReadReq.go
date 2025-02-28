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

// MFuncPropStateReadReq is the corresponding interface of MFuncPropStateReadReq
type MFuncPropStateReadReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
}

// MFuncPropStateReadReqExactly can be used when we want exactly this type and not a type which fulfills MFuncPropStateReadReq.
// This is useful for switch cases.
type MFuncPropStateReadReqExactly interface {
	MFuncPropStateReadReq
	isMFuncPropStateReadReq() bool
}

// _MFuncPropStateReadReq is the data-structure of this message
type _MFuncPropStateReadReq struct {
	*_CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MFuncPropStateReadReq) GetMessageCode() uint8 {
	return 0xF9
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MFuncPropStateReadReq) InitializeParent(parent CEMI) {}

func (m *_MFuncPropStateReadReq) GetParent() CEMI {
	return m._CEMI
}

// NewMFuncPropStateReadReq factory function for _MFuncPropStateReadReq
func NewMFuncPropStateReadReq(size uint16) *_MFuncPropStateReadReq {
	_result := &_MFuncPropStateReadReq{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMFuncPropStateReadReq(structType interface{}) MFuncPropStateReadReq {
	if casted, ok := structType.(MFuncPropStateReadReq); ok {
		return casted
	}
	if casted, ok := structType.(*MFuncPropStateReadReq); ok {
		return *casted
	}
	return nil
}

func (m *_MFuncPropStateReadReq) GetTypeName() string {
	return "MFuncPropStateReadReq"
}

func (m *_MFuncPropStateReadReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_MFuncPropStateReadReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MFuncPropStateReadReqParse(theBytes []byte, size uint16) (MFuncPropStateReadReq, error) {
	return MFuncPropStateReadReqParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), size)
}

func MFuncPropStateReadReqParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (MFuncPropStateReadReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MFuncPropStateReadReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MFuncPropStateReadReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MFuncPropStateReadReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MFuncPropStateReadReq")
	}

	// Create a partially initialized instance
	_child := &_MFuncPropStateReadReq{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_MFuncPropStateReadReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MFuncPropStateReadReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropStateReadReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MFuncPropStateReadReq")
		}

		if popErr := writeBuffer.PopContext("MFuncPropStateReadReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MFuncPropStateReadReq")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MFuncPropStateReadReq) isMFuncPropStateReadReq() bool {
	return true
}

func (m *_MFuncPropStateReadReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
