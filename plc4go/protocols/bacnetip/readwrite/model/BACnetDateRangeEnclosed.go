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

// BACnetDateRangeEnclosed is the corresponding interface of BACnetDateRangeEnclosed
type BACnetDateRangeEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetDateRange returns DateRange (property field)
	GetDateRange() BACnetDateRange
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetDateRangeEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetDateRangeEnclosed.
// This is useful for switch cases.
type BACnetDateRangeEnclosedExactly interface {
	BACnetDateRangeEnclosed
	isBACnetDateRangeEnclosed() bool
}

// _BACnetDateRangeEnclosed is the data-structure of this message
type _BACnetDateRangeEnclosed struct {
	OpeningTag BACnetOpeningTag
	DateRange  BACnetDateRange
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDateRangeEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetDateRangeEnclosed) GetDateRange() BACnetDateRange {
	return m.DateRange
}

func (m *_BACnetDateRangeEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDateRangeEnclosed factory function for _BACnetDateRangeEnclosed
func NewBACnetDateRangeEnclosed(openingTag BACnetOpeningTag, dateRange BACnetDateRange, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetDateRangeEnclosed {
	return &_BACnetDateRangeEnclosed{OpeningTag: openingTag, DateRange: dateRange, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetDateRangeEnclosed(structType interface{}) BACnetDateRangeEnclosed {
	if casted, ok := structType.(BACnetDateRangeEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDateRangeEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDateRangeEnclosed) GetTypeName() string {
	return "BACnetDateRangeEnclosed"
}

func (m *_BACnetDateRangeEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (dateRange)
	lengthInBits += m.DateRange.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDateRangeEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDateRangeEnclosedParse(theBytes []byte, tagNumber uint8) (BACnetDateRangeEnclosed, error) {
	return BACnetDateRangeEnclosedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetDateRangeEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetDateRangeEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateRangeEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateRangeEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetDateRangeEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (dateRange)
	if pullErr := readBuffer.PullContext("dateRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateRange")
	}
	_dateRange, _dateRangeErr := BACnetDateRangeParseWithBuffer(ctx, readBuffer)
	if _dateRangeErr != nil {
		return nil, errors.Wrap(_dateRangeErr, "Error parsing 'dateRange' field of BACnetDateRangeEnclosed")
	}
	dateRange := _dateRange.(BACnetDateRange)
	if closeErr := readBuffer.CloseContext("dateRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateRange")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetDateRangeEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetDateRangeEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateRangeEnclosed")
	}

	// Create the instance
	return &_BACnetDateRangeEnclosed{
		TagNumber:  tagNumber,
		OpeningTag: openingTag,
		DateRange:  dateRange,
		ClosingTag: closingTag,
	}, nil
}

func (m *_BACnetDateRangeEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDateRangeEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDateRangeEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateRangeEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (dateRange)
	if pushErr := writeBuffer.PushContext("dateRange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for dateRange")
	}
	_dateRangeErr := writeBuffer.WriteSerializable(ctx, m.GetDateRange())
	if popErr := writeBuffer.PopContext("dateRange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for dateRange")
	}
	if _dateRangeErr != nil {
		return errors.Wrap(_dateRangeErr, "Error serializing 'dateRange' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateRangeEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateRangeEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDateRangeEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetDateRangeEnclosed) isBACnetDateRangeEnclosed() bool {
	return true
}

func (m *_BACnetDateRangeEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
