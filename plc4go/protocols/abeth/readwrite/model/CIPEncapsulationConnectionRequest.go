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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CIPEncapsulationConnectionRequest is the corresponding interface of CIPEncapsulationConnectionRequest
type CIPEncapsulationConnectionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CIPEncapsulationPacket
}

// CIPEncapsulationConnectionRequestExactly can be used when we want exactly this type and not a type which fulfills CIPEncapsulationConnectionRequest.
// This is useful for switch cases.
type CIPEncapsulationConnectionRequestExactly interface {
	CIPEncapsulationConnectionRequest
	isCIPEncapsulationConnectionRequest() bool
}

// _CIPEncapsulationConnectionRequest is the data-structure of this message
type _CIPEncapsulationConnectionRequest struct {
	*_CIPEncapsulationPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CIPEncapsulationConnectionRequest) GetCommandType() uint16 {
	return 0x0101
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CIPEncapsulationConnectionRequest) InitializeParent(parent CIPEncapsulationPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_CIPEncapsulationConnectionRequest) GetParent() CIPEncapsulationPacket {
	return m._CIPEncapsulationPacket
}

// NewCIPEncapsulationConnectionRequest factory function for _CIPEncapsulationConnectionRequest
func NewCIPEncapsulationConnectionRequest(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *_CIPEncapsulationConnectionRequest {
	_result := &_CIPEncapsulationConnectionRequest{
		_CIPEncapsulationPacket: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
	}
	_result._CIPEncapsulationPacket._CIPEncapsulationPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCIPEncapsulationConnectionRequest(structType interface{}) CIPEncapsulationConnectionRequest {
	if casted, ok := structType.(CIPEncapsulationConnectionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CIPEncapsulationConnectionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CIPEncapsulationConnectionRequest) GetTypeName() string {
	return "CIPEncapsulationConnectionRequest"
}

func (m *_CIPEncapsulationConnectionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_CIPEncapsulationConnectionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPEncapsulationConnectionRequestParse(theBytes []byte) (CIPEncapsulationConnectionRequest, error) {
	return CIPEncapsulationConnectionRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func CIPEncapsulationConnectionRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CIPEncapsulationConnectionRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationConnectionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationConnectionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CIPEncapsulationConnectionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationConnectionRequest")
	}

	// Create a partially initialized instance
	_child := &_CIPEncapsulationConnectionRequest{
		_CIPEncapsulationPacket: &_CIPEncapsulationPacket{},
	}
	_child._CIPEncapsulationPacket._CIPEncapsulationPacketChildRequirements = _child
	return _child, nil
}

func (m *_CIPEncapsulationConnectionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CIPEncapsulationConnectionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationConnectionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationConnectionRequest")
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationConnectionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CIPEncapsulationConnectionRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CIPEncapsulationConnectionRequest) isCIPEncapsulationConnectionRequest() bool {
	return true
}

func (m *_CIPEncapsulationConnectionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
