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

// BACnetPropertyStatesEnclosed is the corresponding interface of BACnetPropertyStatesEnclosed
type BACnetPropertyStatesEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPropertyState returns PropertyState (property field)
	GetPropertyState() BACnetPropertyStates
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetPropertyStatesEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesEnclosed.
// This is useful for switch cases.
type BACnetPropertyStatesEnclosedExactly interface {
	BACnetPropertyStatesEnclosed
	isBACnetPropertyStatesEnclosed() bool
}

// _BACnetPropertyStatesEnclosed is the data-structure of this message
type _BACnetPropertyStatesEnclosed struct {
	OpeningTag    BACnetOpeningTag
	PropertyState BACnetPropertyStates
	ClosingTag    BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetPropertyStatesEnclosed) GetPropertyState() BACnetPropertyStates {
	return m.PropertyState
}

func (m *_BACnetPropertyStatesEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesEnclosed factory function for _BACnetPropertyStatesEnclosed
func NewBACnetPropertyStatesEnclosed(openingTag BACnetOpeningTag, propertyState BACnetPropertyStates, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetPropertyStatesEnclosed {
	return &_BACnetPropertyStatesEnclosed{OpeningTag: openingTag, PropertyState: propertyState, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesEnclosed(structType interface{}) BACnetPropertyStatesEnclosed {
	if casted, ok := structType.(BACnetPropertyStatesEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesEnclosed) GetTypeName() string {
	return "BACnetPropertyStatesEnclosed"
}

func (m *_BACnetPropertyStatesEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (propertyState)
	lengthInBits += m.PropertyState.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesEnclosedParse(theBytes []byte, tagNumber uint8) (BACnetPropertyStatesEnclosed, error) {
	return BACnetPropertyStatesEnclosedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetPropertyStatesEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetPropertyStatesEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetPropertyStatesEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (propertyState)
	if pullErr := readBuffer.PullContext("propertyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyState")
	}
	_propertyState, _propertyStateErr := BACnetPropertyStatesParseWithBuffer(ctx, readBuffer)
	if _propertyStateErr != nil {
		return nil, errors.Wrap(_propertyStateErr, "Error parsing 'propertyState' field of BACnetPropertyStatesEnclosed")
	}
	propertyState := _propertyState.(BACnetPropertyStates)
	if closeErr := readBuffer.CloseContext("propertyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyState")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetPropertyStatesEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesEnclosed")
	}

	// Create the instance
	return &_BACnetPropertyStatesEnclosed{
		TagNumber:     tagNumber,
		OpeningTag:    openingTag,
		PropertyState: propertyState,
		ClosingTag:    closingTag,
	}, nil
}

func (m *_BACnetPropertyStatesEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPropertyStatesEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesEnclosed")
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

	// Simple Field (propertyState)
	if pushErr := writeBuffer.PushContext("propertyState"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyState")
	}
	_propertyStateErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyState())
	if popErr := writeBuffer.PopContext("propertyState"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyState")
	}
	if _propertyStateErr != nil {
		return errors.Wrap(_propertyStateErr, "Error serializing 'propertyState' field")
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

	if popErr := writeBuffer.PopContext("BACnetPropertyStatesEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPropertyStatesEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetPropertyStatesEnclosed) isBACnetPropertyStatesEnclosed() bool {
	return true
}

func (m *_BACnetPropertyStatesEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
