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

// BACnetHostAddress is the corresponding interface of BACnetHostAddress
type BACnetHostAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetHostAddressExactly can be used when we want exactly this type and not a type which fulfills BACnetHostAddress.
// This is useful for switch cases.
type BACnetHostAddressExactly interface {
	BACnetHostAddress
	isBACnetHostAddress() bool
}

// _BACnetHostAddress is the data-structure of this message
type _BACnetHostAddress struct {
	_BACnetHostAddressChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetHostAddressChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetHostAddressParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetHostAddress, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetHostAddressChild interface {
	utils.Serializable
	InitializeParent(parent BACnetHostAddress, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetHostAddress

	GetTypeName() string
	BACnetHostAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostAddress) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetHostAddress) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostAddress factory function for _BACnetHostAddress
func NewBACnetHostAddress(peekedTagHeader BACnetTagHeader) *_BACnetHostAddress {
	return &_BACnetHostAddress{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetHostAddress(structType interface{}) BACnetHostAddress {
	if casted, ok := structType.(BACnetHostAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostAddress) GetTypeName() string {
	return "BACnetHostAddress"
}

func (m *_BACnetHostAddress) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetHostAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetHostAddressParse(theBytes []byte) (BACnetHostAddress, error) {
	return BACnetHostAddressParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetHostAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetHostAddressChildSerializeRequirement interface {
		BACnetHostAddress
		InitializeParent(BACnetHostAddress, BACnetTagHeader)
		GetParent() BACnetHostAddress
	}
	var _childTemp interface{}
	var _child BACnetHostAddressChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetHostAddressNull
		_childTemp, typeSwitchError = BACnetHostAddressNullParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(1): // BACnetHostAddressIpAddress
		_childTemp, typeSwitchError = BACnetHostAddressIpAddressParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(2): // BACnetHostAddressName
		_childTemp, typeSwitchError = BACnetHostAddressNameParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetHostAddress")
	}
	_child = _childTemp.(BACnetHostAddressChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetHostAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddress")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetHostAddress) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetHostAddress, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetHostAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetHostAddress")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetHostAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetHostAddress")
	}
	return nil
}

func (m *_BACnetHostAddress) isBACnetHostAddress() bool {
	return true
}

func (m *_BACnetHostAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
