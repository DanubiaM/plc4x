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

// BACnetConstructedDataHigherDeck is the corresponding interface of BACnetConstructedDataHigherDeck
type BACnetConstructedDataHigherDeck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetHigherDeck returns HigherDeck (property field)
	GetHigherDeck() BACnetApplicationTagObjectIdentifier
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagObjectIdentifier
}

// BACnetConstructedDataHigherDeckExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataHigherDeck.
// This is useful for switch cases.
type BACnetConstructedDataHigherDeckExactly interface {
	BACnetConstructedDataHigherDeck
	isBACnetConstructedDataHigherDeck() bool
}

// _BACnetConstructedDataHigherDeck is the data-structure of this message
type _BACnetConstructedDataHigherDeck struct {
	*_BACnetConstructedData
	HigherDeck BACnetApplicationTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataHigherDeck) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataHigherDeck) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_HIGHER_DECK
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataHigherDeck) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataHigherDeck) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataHigherDeck) GetHigherDeck() BACnetApplicationTagObjectIdentifier {
	return m.HigherDeck
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataHigherDeck) GetActualValue() BACnetApplicationTagObjectIdentifier {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagObjectIdentifier(m.GetHigherDeck())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataHigherDeck factory function for _BACnetConstructedDataHigherDeck
func NewBACnetConstructedDataHigherDeck(higherDeck BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataHigherDeck {
	_result := &_BACnetConstructedDataHigherDeck{
		HigherDeck:             higherDeck,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataHigherDeck(structType interface{}) BACnetConstructedDataHigherDeck {
	if casted, ok := structType.(BACnetConstructedDataHigherDeck); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataHigherDeck); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataHigherDeck) GetTypeName() string {
	return "BACnetConstructedDataHigherDeck"
}

func (m *_BACnetConstructedDataHigherDeck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (higherDeck)
	lengthInBits += m.HigherDeck.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataHigherDeck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataHigherDeckParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataHigherDeck, error) {
	return BACnetConstructedDataHigherDeckParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataHigherDeckParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataHigherDeck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataHigherDeck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataHigherDeck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (higherDeck)
	if pullErr := readBuffer.PullContext("higherDeck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for higherDeck")
	}
	_higherDeck, _higherDeckErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _higherDeckErr != nil {
		return nil, errors.Wrap(_higherDeckErr, "Error parsing 'higherDeck' field of BACnetConstructedDataHigherDeck")
	}
	higherDeck := _higherDeck.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("higherDeck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for higherDeck")
	}

	// Virtual field
	_actualValue := higherDeck
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataHigherDeck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataHigherDeck")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataHigherDeck{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		HigherDeck: higherDeck,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataHigherDeck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataHigherDeck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataHigherDeck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataHigherDeck")
		}

		// Simple Field (higherDeck)
		if pushErr := writeBuffer.PushContext("higherDeck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for higherDeck")
		}
		_higherDeckErr := writeBuffer.WriteSerializable(ctx, m.GetHigherDeck())
		if popErr := writeBuffer.PopContext("higherDeck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for higherDeck")
		}
		if _higherDeckErr != nil {
			return errors.Wrap(_higherDeckErr, "Error serializing 'higherDeck' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataHigherDeck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataHigherDeck")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataHigherDeck) isBACnetConstructedDataHigherDeck() bool {
	return true
}

func (m *_BACnetConstructedDataHigherDeck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
