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

// BACnetAuthenticationFactorEnclosed is the corresponding interface of BACnetAuthenticationFactorEnclosed
type BACnetAuthenticationFactorEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetAuthenticationFactor returns AuthenticationFactor (property field)
	GetAuthenticationFactor() BACnetAuthenticationFactor
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetAuthenticationFactorEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetAuthenticationFactorEnclosed.
// This is useful for switch cases.
type BACnetAuthenticationFactorEnclosedExactly interface {
	BACnetAuthenticationFactorEnclosed
	isBACnetAuthenticationFactorEnclosed() bool
}

// _BACnetAuthenticationFactorEnclosed is the data-structure of this message
type _BACnetAuthenticationFactorEnclosed struct {
	OpeningTag           BACnetOpeningTag
	AuthenticationFactor BACnetAuthenticationFactor
	ClosingTag           BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationFactorEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetAuthenticationFactorEnclosed) GetAuthenticationFactor() BACnetAuthenticationFactor {
	return m.AuthenticationFactor
}

func (m *_BACnetAuthenticationFactorEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAuthenticationFactorEnclosed factory function for _BACnetAuthenticationFactorEnclosed
func NewBACnetAuthenticationFactorEnclosed(openingTag BACnetOpeningTag, authenticationFactor BACnetAuthenticationFactor, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetAuthenticationFactorEnclosed {
	return &_BACnetAuthenticationFactorEnclosed{OpeningTag: openingTag, AuthenticationFactor: authenticationFactor, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationFactorEnclosed(structType interface{}) BACnetAuthenticationFactorEnclosed {
	if casted, ok := structType.(BACnetAuthenticationFactorEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationFactorEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationFactorEnclosed) GetTypeName() string {
	return "BACnetAuthenticationFactorEnclosed"
}

func (m *_BACnetAuthenticationFactorEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (authenticationFactor)
	lengthInBits += m.AuthenticationFactor.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAuthenticationFactorEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationFactorEnclosedParse(theBytes []byte, tagNumber uint8) (BACnetAuthenticationFactorEnclosed, error) {
	return BACnetAuthenticationFactorEnclosedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetAuthenticationFactorEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetAuthenticationFactorEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationFactorEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationFactorEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetAuthenticationFactorEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (authenticationFactor)
	if pullErr := readBuffer.PullContext("authenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationFactor")
	}
	_authenticationFactor, _authenticationFactorErr := BACnetAuthenticationFactorParseWithBuffer(ctx, readBuffer)
	if _authenticationFactorErr != nil {
		return nil, errors.Wrap(_authenticationFactorErr, "Error parsing 'authenticationFactor' field of BACnetAuthenticationFactorEnclosed")
	}
	authenticationFactor := _authenticationFactor.(BACnetAuthenticationFactor)
	if closeErr := readBuffer.CloseContext("authenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationFactor")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetAuthenticationFactorEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationFactorEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationFactorEnclosed")
	}

	// Create the instance
	return &_BACnetAuthenticationFactorEnclosed{
		TagNumber:            tagNumber,
		OpeningTag:           openingTag,
		AuthenticationFactor: authenticationFactor,
		ClosingTag:           closingTag,
	}, nil
}

func (m *_BACnetAuthenticationFactorEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationFactorEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationFactorEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationFactorEnclosed")
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

	// Simple Field (authenticationFactor)
	if pushErr := writeBuffer.PushContext("authenticationFactor"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for authenticationFactor")
	}
	_authenticationFactorErr := writeBuffer.WriteSerializable(ctx, m.GetAuthenticationFactor())
	if popErr := writeBuffer.PopContext("authenticationFactor"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for authenticationFactor")
	}
	if _authenticationFactorErr != nil {
		return errors.Wrap(_authenticationFactorErr, "Error serializing 'authenticationFactor' field")
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

	if popErr := writeBuffer.PopContext("BACnetAuthenticationFactorEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationFactorEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAuthenticationFactorEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetAuthenticationFactorEnclosed) isBACnetAuthenticationFactorEnclosed() bool {
	return true
}

func (m *_BACnetAuthenticationFactorEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
