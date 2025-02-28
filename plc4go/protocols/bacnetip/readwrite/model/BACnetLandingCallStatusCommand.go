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

// BACnetLandingCallStatusCommand is the corresponding interface of BACnetLandingCallStatusCommand
type BACnetLandingCallStatusCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetLandingCallStatusCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetLandingCallStatusCommand.
// This is useful for switch cases.
type BACnetLandingCallStatusCommandExactly interface {
	BACnetLandingCallStatusCommand
	isBACnetLandingCallStatusCommand() bool
}

// _BACnetLandingCallStatusCommand is the data-structure of this message
type _BACnetLandingCallStatusCommand struct {
	_BACnetLandingCallStatusCommandChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetLandingCallStatusCommandChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetLandingCallStatusCommandParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetLandingCallStatusCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetLandingCallStatusCommandChild interface {
	utils.Serializable
	InitializeParent(parent BACnetLandingCallStatusCommand, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetLandingCallStatusCommand

	GetTypeName() string
	BACnetLandingCallStatusCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLandingCallStatusCommand) GetPeekedTagHeader() BACnetTagHeader {
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

func (m *_BACnetLandingCallStatusCommand) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLandingCallStatusCommand factory function for _BACnetLandingCallStatusCommand
func NewBACnetLandingCallStatusCommand(peekedTagHeader BACnetTagHeader) *_BACnetLandingCallStatusCommand {
	return &_BACnetLandingCallStatusCommand{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetLandingCallStatusCommand(structType interface{}) BACnetLandingCallStatusCommand {
	if casted, ok := structType.(BACnetLandingCallStatusCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLandingCallStatusCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLandingCallStatusCommand) GetTypeName() string {
	return "BACnetLandingCallStatusCommand"
}

func (m *_BACnetLandingCallStatusCommand) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetLandingCallStatusCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLandingCallStatusCommandParse(theBytes []byte) (BACnetLandingCallStatusCommand, error) {
	return BACnetLandingCallStatusCommandParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetLandingCallStatusCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLandingCallStatusCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingCallStatusCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLandingCallStatusCommand")
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
	type BACnetLandingCallStatusCommandChildSerializeRequirement interface {
		BACnetLandingCallStatusCommand
		InitializeParent(BACnetLandingCallStatusCommand, BACnetTagHeader)
		GetParent() BACnetLandingCallStatusCommand
	}
	var _childTemp interface{}
	var _child BACnetLandingCallStatusCommandChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(1): // BACnetLandingCallStatusCommandDirection
		_childTemp, typeSwitchError = BACnetLandingCallStatusCommandDirectionParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(2): // BACnetLandingCallStatusCommandDestination
		_childTemp, typeSwitchError = BACnetLandingCallStatusCommandDestinationParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetLandingCallStatusCommand")
	}
	_child = _childTemp.(BACnetLandingCallStatusCommandChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetLandingCallStatusCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLandingCallStatusCommand")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetLandingCallStatusCommand) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetLandingCallStatusCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLandingCallStatusCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLandingCallStatusCommand")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetLandingCallStatusCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLandingCallStatusCommand")
	}
	return nil
}

func (m *_BACnetLandingCallStatusCommand) isBACnetLandingCallStatusCommand() bool {
	return true
}

func (m *_BACnetLandingCallStatusCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
