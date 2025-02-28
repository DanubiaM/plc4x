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

// BACnetTagHeader is the corresponding interface of BACnetTagHeader
type BACnetTagHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTagNumber returns TagNumber (property field)
	GetTagNumber() uint8
	// GetTagClass returns TagClass (property field)
	GetTagClass() TagClass
	// GetLengthValueType returns LengthValueType (property field)
	GetLengthValueType() uint8
	// GetExtTagNumber returns ExtTagNumber (property field)
	GetExtTagNumber() *uint8
	// GetExtLength returns ExtLength (property field)
	GetExtLength() *uint8
	// GetExtExtLength returns ExtExtLength (property field)
	GetExtExtLength() *uint16
	// GetExtExtExtLength returns ExtExtExtLength (property field)
	GetExtExtExtLength() *uint32
	// GetActualTagNumber returns ActualTagNumber (virtual field)
	GetActualTagNumber() uint8
	// GetIsBoolean returns IsBoolean (virtual field)
	GetIsBoolean() bool
	// GetIsConstructed returns IsConstructed (virtual field)
	GetIsConstructed() bool
	// GetIsPrimitiveAndNotBoolean returns IsPrimitiveAndNotBoolean (virtual field)
	GetIsPrimitiveAndNotBoolean() bool
	// GetActualLength returns ActualLength (virtual field)
	GetActualLength() uint32
}

// BACnetTagHeaderExactly can be used when we want exactly this type and not a type which fulfills BACnetTagHeader.
// This is useful for switch cases.
type BACnetTagHeaderExactly interface {
	BACnetTagHeader
	isBACnetTagHeader() bool
}

// _BACnetTagHeader is the data-structure of this message
type _BACnetTagHeader struct {
	TagNumber       uint8
	TagClass        TagClass
	LengthValueType uint8
	ExtTagNumber    *uint8
	ExtLength       *uint8
	ExtExtLength    *uint16
	ExtExtExtLength *uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagHeader) GetTagNumber() uint8 {
	return m.TagNumber
}

func (m *_BACnetTagHeader) GetTagClass() TagClass {
	return m.TagClass
}

func (m *_BACnetTagHeader) GetLengthValueType() uint8 {
	return m.LengthValueType
}

func (m *_BACnetTagHeader) GetExtTagNumber() *uint8 {
	return m.ExtTagNumber
}

func (m *_BACnetTagHeader) GetExtLength() *uint8 {
	return m.ExtLength
}

func (m *_BACnetTagHeader) GetExtExtLength() *uint16 {
	return m.ExtExtLength
}

func (m *_BACnetTagHeader) GetExtExtExtLength() *uint32 {
	return m.ExtExtExtLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagHeader) GetActualTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	extTagNumber := m.ExtTagNumber
	_ = extTagNumber
	extLength := m.ExtLength
	_ = extLength
	extExtLength := m.ExtExtLength
	_ = extExtLength
	extExtExtLength := m.ExtExtExtLength
	_ = extExtExtLength
	return uint8(utils.InlineIf(bool((m.GetTagNumber()) < (15)), func() interface{} { return uint8(m.GetTagNumber()) }, func() interface{} { return uint8((*m.GetExtTagNumber())) }).(uint8))
}

func (m *_BACnetTagHeader) GetIsBoolean() bool {
	ctx := context.Background()
	_ = ctx
	extTagNumber := m.ExtTagNumber
	_ = extTagNumber
	extLength := m.ExtLength
	_ = extLength
	extExtLength := m.ExtExtLength
	_ = extExtLength
	extExtExtLength := m.ExtExtExtLength
	_ = extExtExtLength
	return bool(bool(bool((m.GetTagNumber()) == (1))) && bool(bool((m.GetTagClass()) == (TagClass_APPLICATION_TAGS))))
}

func (m *_BACnetTagHeader) GetIsConstructed() bool {
	ctx := context.Background()
	_ = ctx
	extTagNumber := m.ExtTagNumber
	_ = extTagNumber
	extLength := m.ExtLength
	_ = extLength
	extExtLength := m.ExtExtLength
	_ = extExtLength
	extExtExtLength := m.ExtExtExtLength
	_ = extExtExtLength
	return bool(bool(bool((m.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) && bool(bool((m.GetLengthValueType()) == (6))))
}

func (m *_BACnetTagHeader) GetIsPrimitiveAndNotBoolean() bool {
	ctx := context.Background()
	_ = ctx
	extTagNumber := m.ExtTagNumber
	_ = extTagNumber
	extLength := m.ExtLength
	_ = extLength
	extExtLength := m.ExtExtLength
	_ = extExtLength
	extExtExtLength := m.ExtExtExtLength
	_ = extExtExtLength
	return bool(bool(!(m.GetIsConstructed())) && bool(!(m.GetIsBoolean())))
}

func (m *_BACnetTagHeader) GetActualLength() uint32 {
	ctx := context.Background()
	_ = ctx
	extTagNumber := m.ExtTagNumber
	_ = extTagNumber
	extLength := m.ExtLength
	_ = extLength
	extExtLength := m.ExtExtLength
	_ = extExtLength
	extExtExtLength := m.ExtExtExtLength
	_ = extExtExtLength
	return uint32(utils.InlineIf(bool(bool((m.GetLengthValueType()) == (5))) && bool(bool((*m.GetExtLength()) == (255))), func() interface{} { return uint32((*m.GetExtExtExtLength())) }, func() interface{} {
		return uint32((utils.InlineIf(bool(bool((m.GetLengthValueType()) == (5))) && bool(bool((*m.GetExtLength()) == (254))), func() interface{} { return uint32((*m.GetExtExtLength())) }, func() interface{} {
			return uint32((utils.InlineIf(bool((m.GetLengthValueType()) == (5)), func() interface{} { return uint32((*m.GetExtLength())) }, func() interface{} { return uint32(m.GetLengthValueType()) }).(uint32)))
		}).(uint32)))
	}).(uint32))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagHeader factory function for _BACnetTagHeader
func NewBACnetTagHeader(tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *_BACnetTagHeader {
	return &_BACnetTagHeader{TagNumber: tagNumber, TagClass: tagClass, LengthValueType: lengthValueType, ExtTagNumber: extTagNumber, ExtLength: extLength, ExtExtLength: extExtLength, ExtExtExtLength: extExtExtLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagHeader(structType interface{}) BACnetTagHeader {
	if casted, ok := structType.(BACnetTagHeader); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagHeader); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagHeader) GetTypeName() string {
	return "BACnetTagHeader"
}

func (m *_BACnetTagHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (tagNumber)
	lengthInBits += 4

	// Simple field (tagClass)
	lengthInBits += 1

	// Simple field (lengthValueType)
	lengthInBits += 3

	// Optional Field (extTagNumber)
	if m.ExtTagNumber != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (extLength)
	if m.ExtLength != nil {
		lengthInBits += 8
	}

	// Optional Field (extExtLength)
	if m.ExtExtLength != nil {
		lengthInBits += 16
	}

	// Optional Field (extExtExtLength)
	if m.ExtExtExtLength != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTagHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagHeaderParse(theBytes []byte) (BACnetTagHeader, error) {
	return BACnetTagHeaderParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetTagHeaderParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (tagNumber)
	_tagNumber, _tagNumberErr := readBuffer.ReadUint8("tagNumber", 4)
	if _tagNumberErr != nil {
		return nil, errors.Wrap(_tagNumberErr, "Error parsing 'tagNumber' field of BACnetTagHeader")
	}
	tagNumber := _tagNumber

	// Simple Field (tagClass)
	if pullErr := readBuffer.PullContext("tagClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for tagClass")
	}
	_tagClass, _tagClassErr := TagClassParseWithBuffer(ctx, readBuffer)
	if _tagClassErr != nil {
		return nil, errors.Wrap(_tagClassErr, "Error parsing 'tagClass' field of BACnetTagHeader")
	}
	tagClass := _tagClass
	if closeErr := readBuffer.CloseContext("tagClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for tagClass")
	}

	// Simple Field (lengthValueType)
	_lengthValueType, _lengthValueTypeErr := readBuffer.ReadUint8("lengthValueType", 3)
	if _lengthValueTypeErr != nil {
		return nil, errors.Wrap(_lengthValueTypeErr, "Error parsing 'lengthValueType' field of BACnetTagHeader")
	}
	lengthValueType := _lengthValueType

	// Optional Field (extTagNumber) (Can be skipped, if a given expression evaluates to false)
	var extTagNumber *uint8 = nil
	if bool((tagNumber) == (15)) {
		_val, _err := readBuffer.ReadUint8("extTagNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extTagNumber' field of BACnetTagHeader")
		}
		extTagNumber = &_val
	}

	// Virtual field
	_actualTagNumber := utils.InlineIf(bool((tagNumber) < (15)), func() interface{} { return uint8(tagNumber) }, func() interface{} { return uint8((*extTagNumber)) }).(uint8)
	actualTagNumber := uint8(_actualTagNumber)
	_ = actualTagNumber

	// Virtual field
	_isBoolean := bool(bool((tagNumber) == (1))) && bool(bool((tagClass) == (TagClass_APPLICATION_TAGS)))
	isBoolean := bool(_isBoolean)
	_ = isBoolean

	// Virtual field
	_isConstructed := bool(bool((tagClass) == (TagClass_CONTEXT_SPECIFIC_TAGS))) && bool(bool((lengthValueType) == (6)))
	isConstructed := bool(_isConstructed)
	_ = isConstructed

	// Virtual field
	_isPrimitiveAndNotBoolean := bool(!(isConstructed)) && bool(!(isBoolean))
	isPrimitiveAndNotBoolean := bool(_isPrimitiveAndNotBoolean)
	_ = isPrimitiveAndNotBoolean

	// Optional Field (extLength) (Can be skipped, if a given expression evaluates to false)
	var extLength *uint8 = nil
	if bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5))) {
		_val, _err := readBuffer.ReadUint8("extLength", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extLength' field of BACnetTagHeader")
		}
		extLength = &_val
	}

	// Optional Field (extExtLength) (Can be skipped, if a given expression evaluates to false)
	var extExtLength *uint16 = nil
	if bool(bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5)))) && bool(bool((*extLength) == (254))) {
		_val, _err := readBuffer.ReadUint16("extExtLength", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extExtLength' field of BACnetTagHeader")
		}
		extExtLength = &_val
	}

	// Optional Field (extExtExtLength) (Can be skipped, if a given expression evaluates to false)
	var extExtExtLength *uint32 = nil
	if bool(bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5)))) && bool(bool((*extLength) == (255))) {
		_val, _err := readBuffer.ReadUint32("extExtExtLength", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extExtExtLength' field of BACnetTagHeader")
		}
		extExtExtLength = &_val
	}

	// Virtual field
	_actualLength := utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (255))), func() interface{} { return uint32((*extExtExtLength)) }, func() interface{} {
		return uint32((utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (254))), func() interface{} { return uint32((*extExtLength)) }, func() interface{} {
			return uint32((utils.InlineIf(bool((lengthValueType) == (5)), func() interface{} { return uint32((*extLength)) }, func() interface{} { return uint32(lengthValueType) }).(uint32)))
		}).(uint32)))
	}).(uint32)
	actualLength := uint32(_actualLength)
	_ = actualLength

	if closeErr := readBuffer.CloseContext("BACnetTagHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagHeader")
	}

	// Create the instance
	return &_BACnetTagHeader{
		TagNumber:       tagNumber,
		TagClass:        tagClass,
		LengthValueType: lengthValueType,
		ExtTagNumber:    extTagNumber,
		ExtLength:       extLength,
		ExtExtLength:    extExtLength,
		ExtExtExtLength: extExtExtLength,
	}, nil
}

func (m *_BACnetTagHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTagHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagHeader")
	}

	// Simple Field (tagNumber)
	tagNumber := uint8(m.GetTagNumber())
	_tagNumberErr := writeBuffer.WriteUint8("tagNumber", 4, (tagNumber))
	if _tagNumberErr != nil {
		return errors.Wrap(_tagNumberErr, "Error serializing 'tagNumber' field")
	}

	// Simple Field (tagClass)
	if pushErr := writeBuffer.PushContext("tagClass"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for tagClass")
	}
	_tagClassErr := writeBuffer.WriteSerializable(ctx, m.GetTagClass())
	if popErr := writeBuffer.PopContext("tagClass"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for tagClass")
	}
	if _tagClassErr != nil {
		return errors.Wrap(_tagClassErr, "Error serializing 'tagClass' field")
	}

	// Simple Field (lengthValueType)
	lengthValueType := uint8(m.GetLengthValueType())
	_lengthValueTypeErr := writeBuffer.WriteUint8("lengthValueType", 3, (lengthValueType))
	if _lengthValueTypeErr != nil {
		return errors.Wrap(_lengthValueTypeErr, "Error serializing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if the value is null)
	var extTagNumber *uint8 = nil
	if m.GetExtTagNumber() != nil {
		extTagNumber = m.GetExtTagNumber()
		_extTagNumberErr := writeBuffer.WriteUint8("extTagNumber", 8, *(extTagNumber))
		if _extTagNumberErr != nil {
			return errors.Wrap(_extTagNumberErr, "Error serializing 'extTagNumber' field")
		}
	}
	// Virtual field
	if _actualTagNumberErr := writeBuffer.WriteVirtual(ctx, "actualTagNumber", m.GetActualTagNumber()); _actualTagNumberErr != nil {
		return errors.Wrap(_actualTagNumberErr, "Error serializing 'actualTagNumber' field")
	}
	// Virtual field
	if _isBooleanErr := writeBuffer.WriteVirtual(ctx, "isBoolean", m.GetIsBoolean()); _isBooleanErr != nil {
		return errors.Wrap(_isBooleanErr, "Error serializing 'isBoolean' field")
	}
	// Virtual field
	if _isConstructedErr := writeBuffer.WriteVirtual(ctx, "isConstructed", m.GetIsConstructed()); _isConstructedErr != nil {
		return errors.Wrap(_isConstructedErr, "Error serializing 'isConstructed' field")
	}
	// Virtual field
	if _isPrimitiveAndNotBooleanErr := writeBuffer.WriteVirtual(ctx, "isPrimitiveAndNotBoolean", m.GetIsPrimitiveAndNotBoolean()); _isPrimitiveAndNotBooleanErr != nil {
		return errors.Wrap(_isPrimitiveAndNotBooleanErr, "Error serializing 'isPrimitiveAndNotBoolean' field")
	}

	// Optional Field (extLength) (Can be skipped, if the value is null)
	var extLength *uint8 = nil
	if m.GetExtLength() != nil {
		extLength = m.GetExtLength()
		_extLengthErr := writeBuffer.WriteUint8("extLength", 8, *(extLength))
		if _extLengthErr != nil {
			return errors.Wrap(_extLengthErr, "Error serializing 'extLength' field")
		}
	}

	// Optional Field (extExtLength) (Can be skipped, if the value is null)
	var extExtLength *uint16 = nil
	if m.GetExtExtLength() != nil {
		extExtLength = m.GetExtExtLength()
		_extExtLengthErr := writeBuffer.WriteUint16("extExtLength", 16, *(extExtLength))
		if _extExtLengthErr != nil {
			return errors.Wrap(_extExtLengthErr, "Error serializing 'extExtLength' field")
		}
	}

	// Optional Field (extExtExtLength) (Can be skipped, if the value is null)
	var extExtExtLength *uint32 = nil
	if m.GetExtExtExtLength() != nil {
		extExtExtLength = m.GetExtExtExtLength()
		_extExtExtLengthErr := writeBuffer.WriteUint32("extExtExtLength", 32, *(extExtExtLength))
		if _extExtExtLengthErr != nil {
			return errors.Wrap(_extExtExtLengthErr, "Error serializing 'extExtExtLength' field")
		}
	}
	// Virtual field
	if _actualLengthErr := writeBuffer.WriteVirtual(ctx, "actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagHeader")
	}
	return nil
}

func (m *_BACnetTagHeader) isBACnetTagHeader() bool {
	return true
}

func (m *_BACnetTagHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
