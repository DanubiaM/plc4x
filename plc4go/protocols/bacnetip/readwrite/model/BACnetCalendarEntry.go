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

// BACnetCalendarEntry is the corresponding interface of BACnetCalendarEntry
type BACnetCalendarEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetCalendarEntryExactly can be used when we want exactly this type and not a type which fulfills BACnetCalendarEntry.
// This is useful for switch cases.
type BACnetCalendarEntryExactly interface {
	BACnetCalendarEntry
	isBACnetCalendarEntry() bool
}

// _BACnetCalendarEntry is the data-structure of this message
type _BACnetCalendarEntry struct {
	_BACnetCalendarEntryChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetCalendarEntryChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetCalendarEntryParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetCalendarEntry, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetCalendarEntryChild interface {
	utils.Serializable
	InitializeParent(parent BACnetCalendarEntry, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetCalendarEntry

	GetTypeName() string
	BACnetCalendarEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntry) GetPeekedTagHeader() BACnetTagHeader {
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

func (m *_BACnetCalendarEntry) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCalendarEntry factory function for _BACnetCalendarEntry
func NewBACnetCalendarEntry(peekedTagHeader BACnetTagHeader) *_BACnetCalendarEntry {
	return &_BACnetCalendarEntry{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntry(structType interface{}) BACnetCalendarEntry {
	if casted, ok := structType.(BACnetCalendarEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntry) GetTypeName() string {
	return "BACnetCalendarEntry"
}

func (m *_BACnetCalendarEntry) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetCalendarEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCalendarEntryParse(theBytes []byte) (BACnetCalendarEntry, error) {
	return BACnetCalendarEntryParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetCalendarEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCalendarEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCalendarEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntry")
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

	// Validation
	if !(bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{"Validation failed"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetCalendarEntryChildSerializeRequirement interface {
		BACnetCalendarEntry
		InitializeParent(BACnetCalendarEntry, BACnetTagHeader)
		GetParent() BACnetCalendarEntry
	}
	var _childTemp interface{}
	var _child BACnetCalendarEntryChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetCalendarEntryDate
		_childTemp, typeSwitchError = BACnetCalendarEntryDateParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(1): // BACnetCalendarEntryDateRange
		_childTemp, typeSwitchError = BACnetCalendarEntryDateRangeParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(2): // BACnetCalendarEntryWeekNDay
		_childTemp, typeSwitchError = BACnetCalendarEntryWeekNDayParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetCalendarEntry")
	}
	_child = _childTemp.(BACnetCalendarEntryChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntry")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetCalendarEntry) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetCalendarEntry, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetCalendarEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntry")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetCalendarEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCalendarEntry")
	}
	return nil
}

func (m *_BACnetCalendarEntry) isBACnetCalendarEntry() bool {
	return true
}

func (m *_BACnetCalendarEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
