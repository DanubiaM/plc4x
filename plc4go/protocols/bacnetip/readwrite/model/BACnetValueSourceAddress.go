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

// BACnetValueSourceAddress is the corresponding interface of BACnetValueSourceAddress
type BACnetValueSourceAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetValueSource
	// GetAddress returns Address (property field)
	GetAddress() BACnetAddressEnclosed
}

// BACnetValueSourceAddressExactly can be used when we want exactly this type and not a type which fulfills BACnetValueSourceAddress.
// This is useful for switch cases.
type BACnetValueSourceAddressExactly interface {
	BACnetValueSourceAddress
	isBACnetValueSourceAddress() bool
}

// _BACnetValueSourceAddress is the data-structure of this message
type _BACnetValueSourceAddress struct {
	*_BACnetValueSource
	Address BACnetAddressEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetValueSourceAddress) InitializeParent(parent BACnetValueSource, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetValueSourceAddress) GetParent() BACnetValueSource {
	return m._BACnetValueSource
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetValueSourceAddress) GetAddress() BACnetAddressEnclosed {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetValueSourceAddress factory function for _BACnetValueSourceAddress
func NewBACnetValueSourceAddress(address BACnetAddressEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetValueSourceAddress {
	_result := &_BACnetValueSourceAddress{
		Address:            address,
		_BACnetValueSource: NewBACnetValueSource(peekedTagHeader),
	}
	_result._BACnetValueSource._BACnetValueSourceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetValueSourceAddress(structType interface{}) BACnetValueSourceAddress {
	if casted, ok := structType.(BACnetValueSourceAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetValueSourceAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetValueSourceAddress) GetTypeName() string {
	return "BACnetValueSourceAddress"
}

func (m *_BACnetValueSourceAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += m.Address.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetValueSourceAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetValueSourceAddressParse(theBytes []byte) (BACnetValueSourceAddress, error) {
	return BACnetValueSourceAddressParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetValueSourceAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetValueSourceAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetValueSourceAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetValueSourceAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (address)
	if pullErr := readBuffer.PullContext("address"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for address")
	}
	_address, _addressErr := BACnetAddressEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of BACnetValueSourceAddress")
	}
	address := _address.(BACnetAddressEnclosed)
	if closeErr := readBuffer.CloseContext("address"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for address")
	}

	if closeErr := readBuffer.CloseContext("BACnetValueSourceAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetValueSourceAddress")
	}

	// Create a partially initialized instance
	_child := &_BACnetValueSourceAddress{
		_BACnetValueSource: &_BACnetValueSource{},
		Address:            address,
	}
	_child._BACnetValueSource._BACnetValueSourceChildRequirements = _child
	return _child, nil
}

func (m *_BACnetValueSourceAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetValueSourceAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetValueSourceAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetValueSourceAddress")
		}

		// Simple Field (address)
		if pushErr := writeBuffer.PushContext("address"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for address")
		}
		_addressErr := writeBuffer.WriteSerializable(ctx, m.GetAddress())
		if popErr := writeBuffer.PopContext("address"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for address")
		}
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("BACnetValueSourceAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetValueSourceAddress")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetValueSourceAddress) isBACnetValueSourceAddress() bool {
	return true
}

func (m *_BACnetValueSourceAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
