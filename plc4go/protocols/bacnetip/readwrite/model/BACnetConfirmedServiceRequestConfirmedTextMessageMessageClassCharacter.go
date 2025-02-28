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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetCharacterValue returns CharacterValue (property field)
	GetCharacterValue() BACnetContextTagCharacterString
}

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterExactly interface {
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
	isBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter() bool
}

// _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter struct {
	*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	CharacterValue BACnetContextTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) InitializeParent(parent BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetParent() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	return m._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetCharacterValue() BACnetContextTagCharacterString {
	return m.CharacterValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter factory function for _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter(characterValue BACnetContextTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	_result := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter{
		CharacterValue: characterValue,
		_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass: NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter(structType interface{}) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (characterValue)
	lengthInBits += m.CharacterValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterParse(theBytes []byte, tagNumber uint8) (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter, error) {
	return BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (characterValue)
	if pullErr := readBuffer.PullContext("characterValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for characterValue")
	}
	_characterValue, _characterValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _characterValueErr != nil {
		return nil, errors.Wrap(_characterValueErr, "Error parsing 'characterValue' field of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
	}
	characterValue := _characterValue.(BACnetContextTagCharacterString)
	if closeErr := readBuffer.CloseContext("characterValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for characterValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter{
		_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass: &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass{
			TagNumber: tagNumber,
		},
		CharacterValue: characterValue,
	}
	_child._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
		}

		// Simple Field (characterValue)
		if pushErr := writeBuffer.PushContext("characterValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for characterValue")
		}
		_characterValueErr := writeBuffer.WriteSerializable(ctx, m.GetCharacterValue())
		if popErr := writeBuffer.PopContext("characterValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for characterValue")
		}
		if _characterValueErr != nil {
			return errors.Wrap(_characterValueErr, "Error serializing 'characterValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) isBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassCharacter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
