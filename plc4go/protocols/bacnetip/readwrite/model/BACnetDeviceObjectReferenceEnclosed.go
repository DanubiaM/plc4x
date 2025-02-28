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

// BACnetDeviceObjectReferenceEnclosed is the corresponding interface of BACnetDeviceObjectReferenceEnclosed
type BACnetDeviceObjectReferenceEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetObjectReference returns ObjectReference (property field)
	GetObjectReference() BACnetDeviceObjectReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetDeviceObjectReferenceEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetDeviceObjectReferenceEnclosed.
// This is useful for switch cases.
type BACnetDeviceObjectReferenceEnclosedExactly interface {
	BACnetDeviceObjectReferenceEnclosed
	isBACnetDeviceObjectReferenceEnclosed() bool
}

// _BACnetDeviceObjectReferenceEnclosed is the data-structure of this message
type _BACnetDeviceObjectReferenceEnclosed struct {
	OpeningTag      BACnetOpeningTag
	ObjectReference BACnetDeviceObjectReference
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDeviceObjectReferenceEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetDeviceObjectReferenceEnclosed) GetObjectReference() BACnetDeviceObjectReference {
	return m.ObjectReference
}

func (m *_BACnetDeviceObjectReferenceEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDeviceObjectReferenceEnclosed factory function for _BACnetDeviceObjectReferenceEnclosed
func NewBACnetDeviceObjectReferenceEnclosed(openingTag BACnetOpeningTag, objectReference BACnetDeviceObjectReference, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetDeviceObjectReferenceEnclosed {
	return &_BACnetDeviceObjectReferenceEnclosed{OpeningTag: openingTag, ObjectReference: objectReference, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetDeviceObjectReferenceEnclosed(structType interface{}) BACnetDeviceObjectReferenceEnclosed {
	if casted, ok := structType.(BACnetDeviceObjectReferenceEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDeviceObjectReferenceEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDeviceObjectReferenceEnclosed) GetTypeName() string {
	return "BACnetDeviceObjectReferenceEnclosed"
}

func (m *_BACnetDeviceObjectReferenceEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (objectReference)
	lengthInBits += m.ObjectReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDeviceObjectReferenceEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDeviceObjectReferenceEnclosedParse(theBytes []byte, tagNumber uint8) (BACnetDeviceObjectReferenceEnclosed, error) {
	return BACnetDeviceObjectReferenceEnclosedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetDeviceObjectReferenceEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetDeviceObjectReferenceEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDeviceObjectReferenceEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDeviceObjectReferenceEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetDeviceObjectReferenceEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (objectReference)
	if pullErr := readBuffer.PullContext("objectReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectReference")
	}
	_objectReference, _objectReferenceErr := BACnetDeviceObjectReferenceParseWithBuffer(ctx, readBuffer)
	if _objectReferenceErr != nil {
		return nil, errors.Wrap(_objectReferenceErr, "Error parsing 'objectReference' field of BACnetDeviceObjectReferenceEnclosed")
	}
	objectReference := _objectReference.(BACnetDeviceObjectReference)
	if closeErr := readBuffer.CloseContext("objectReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectReference")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetDeviceObjectReferenceEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetDeviceObjectReferenceEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDeviceObjectReferenceEnclosed")
	}

	// Create the instance
	return &_BACnetDeviceObjectReferenceEnclosed{
		TagNumber:       tagNumber,
		OpeningTag:      openingTag,
		ObjectReference: objectReference,
		ClosingTag:      closingTag,
	}, nil
}

func (m *_BACnetDeviceObjectReferenceEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDeviceObjectReferenceEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDeviceObjectReferenceEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDeviceObjectReferenceEnclosed")
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

	// Simple Field (objectReference)
	if pushErr := writeBuffer.PushContext("objectReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectReference")
	}
	_objectReferenceErr := writeBuffer.WriteSerializable(ctx, m.GetObjectReference())
	if popErr := writeBuffer.PopContext("objectReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for objectReference")
	}
	if _objectReferenceErr != nil {
		return errors.Wrap(_objectReferenceErr, "Error serializing 'objectReference' field")
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

	if popErr := writeBuffer.PopContext("BACnetDeviceObjectReferenceEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDeviceObjectReferenceEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDeviceObjectReferenceEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetDeviceObjectReferenceEnclosed) isBACnetDeviceObjectReferenceEnclosed() bool {
	return true
}

func (m *_BACnetDeviceObjectReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
