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

// BACnetVTSession is the corresponding interface of BACnetVTSession
type BACnetVTSession interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLocalVtSessionId returns LocalVtSessionId (property field)
	GetLocalVtSessionId() BACnetApplicationTagUnsignedInteger
	// GetRemoveVtSessionId returns RemoveVtSessionId (property field)
	GetRemoveVtSessionId() BACnetApplicationTagUnsignedInteger
	// GetRemoteVtAddress returns RemoteVtAddress (property field)
	GetRemoteVtAddress() BACnetAddress
}

// BACnetVTSessionExactly can be used when we want exactly this type and not a type which fulfills BACnetVTSession.
// This is useful for switch cases.
type BACnetVTSessionExactly interface {
	BACnetVTSession
	isBACnetVTSession() bool
}

// _BACnetVTSession is the data-structure of this message
type _BACnetVTSession struct {
	LocalVtSessionId  BACnetApplicationTagUnsignedInteger
	RemoveVtSessionId BACnetApplicationTagUnsignedInteger
	RemoteVtAddress   BACnetAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetVTSession) GetLocalVtSessionId() BACnetApplicationTagUnsignedInteger {
	return m.LocalVtSessionId
}

func (m *_BACnetVTSession) GetRemoveVtSessionId() BACnetApplicationTagUnsignedInteger {
	return m.RemoveVtSessionId
}

func (m *_BACnetVTSession) GetRemoteVtAddress() BACnetAddress {
	return m.RemoteVtAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetVTSession factory function for _BACnetVTSession
func NewBACnetVTSession(localVtSessionId BACnetApplicationTagUnsignedInteger, removeVtSessionId BACnetApplicationTagUnsignedInteger, remoteVtAddress BACnetAddress) *_BACnetVTSession {
	return &_BACnetVTSession{LocalVtSessionId: localVtSessionId, RemoveVtSessionId: removeVtSessionId, RemoteVtAddress: remoteVtAddress}
}

// Deprecated: use the interface for direct cast
func CastBACnetVTSession(structType interface{}) BACnetVTSession {
	if casted, ok := structType.(BACnetVTSession); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetVTSession); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetVTSession) GetTypeName() string {
	return "BACnetVTSession"
}

func (m *_BACnetVTSession) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (localVtSessionId)
	lengthInBits += m.LocalVtSessionId.GetLengthInBits(ctx)

	// Simple field (removeVtSessionId)
	lengthInBits += m.RemoveVtSessionId.GetLengthInBits(ctx)

	// Simple field (remoteVtAddress)
	lengthInBits += m.RemoteVtAddress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetVTSession) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetVTSessionParse(theBytes []byte) (BACnetVTSession, error) {
	return BACnetVTSessionParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetVTSessionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetVTSession, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetVTSession"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetVTSession")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (localVtSessionId)
	if pullErr := readBuffer.PullContext("localVtSessionId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for localVtSessionId")
	}
	_localVtSessionId, _localVtSessionIdErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _localVtSessionIdErr != nil {
		return nil, errors.Wrap(_localVtSessionIdErr, "Error parsing 'localVtSessionId' field of BACnetVTSession")
	}
	localVtSessionId := _localVtSessionId.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("localVtSessionId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for localVtSessionId")
	}

	// Simple Field (removeVtSessionId)
	if pullErr := readBuffer.PullContext("removeVtSessionId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for removeVtSessionId")
	}
	_removeVtSessionId, _removeVtSessionIdErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _removeVtSessionIdErr != nil {
		return nil, errors.Wrap(_removeVtSessionIdErr, "Error parsing 'removeVtSessionId' field of BACnetVTSession")
	}
	removeVtSessionId := _removeVtSessionId.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("removeVtSessionId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for removeVtSessionId")
	}

	// Simple Field (remoteVtAddress)
	if pullErr := readBuffer.PullContext("remoteVtAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for remoteVtAddress")
	}
	_remoteVtAddress, _remoteVtAddressErr := BACnetAddressParseWithBuffer(ctx, readBuffer)
	if _remoteVtAddressErr != nil {
		return nil, errors.Wrap(_remoteVtAddressErr, "Error parsing 'remoteVtAddress' field of BACnetVTSession")
	}
	remoteVtAddress := _remoteVtAddress.(BACnetAddress)
	if closeErr := readBuffer.CloseContext("remoteVtAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for remoteVtAddress")
	}

	if closeErr := readBuffer.CloseContext("BACnetVTSession"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetVTSession")
	}

	// Create the instance
	return &_BACnetVTSession{
		LocalVtSessionId:  localVtSessionId,
		RemoveVtSessionId: removeVtSessionId,
		RemoteVtAddress:   remoteVtAddress,
	}, nil
}

func (m *_BACnetVTSession) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetVTSession) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetVTSession"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetVTSession")
	}

	// Simple Field (localVtSessionId)
	if pushErr := writeBuffer.PushContext("localVtSessionId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for localVtSessionId")
	}
	_localVtSessionIdErr := writeBuffer.WriteSerializable(ctx, m.GetLocalVtSessionId())
	if popErr := writeBuffer.PopContext("localVtSessionId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for localVtSessionId")
	}
	if _localVtSessionIdErr != nil {
		return errors.Wrap(_localVtSessionIdErr, "Error serializing 'localVtSessionId' field")
	}

	// Simple Field (removeVtSessionId)
	if pushErr := writeBuffer.PushContext("removeVtSessionId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for removeVtSessionId")
	}
	_removeVtSessionIdErr := writeBuffer.WriteSerializable(ctx, m.GetRemoveVtSessionId())
	if popErr := writeBuffer.PopContext("removeVtSessionId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for removeVtSessionId")
	}
	if _removeVtSessionIdErr != nil {
		return errors.Wrap(_removeVtSessionIdErr, "Error serializing 'removeVtSessionId' field")
	}

	// Simple Field (remoteVtAddress)
	if pushErr := writeBuffer.PushContext("remoteVtAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for remoteVtAddress")
	}
	_remoteVtAddressErr := writeBuffer.WriteSerializable(ctx, m.GetRemoteVtAddress())
	if popErr := writeBuffer.PopContext("remoteVtAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for remoteVtAddress")
	}
	if _remoteVtAddressErr != nil {
		return errors.Wrap(_remoteVtAddressErr, "Error serializing 'remoteVtAddress' field")
	}

	if popErr := writeBuffer.PopContext("BACnetVTSession"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetVTSession")
	}
	return nil
}

func (m *_BACnetVTSession) isBACnetVTSession() bool {
	return true
}

func (m *_BACnetVTSession) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
