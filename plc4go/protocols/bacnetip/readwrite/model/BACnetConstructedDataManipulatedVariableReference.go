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

// BACnetConstructedDataManipulatedVariableReference is the corresponding interface of BACnetConstructedDataManipulatedVariableReference
type BACnetConstructedDataManipulatedVariableReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetManipulatedVariableReference returns ManipulatedVariableReference (property field)
	GetManipulatedVariableReference() BACnetObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectPropertyReference
}

// BACnetConstructedDataManipulatedVariableReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataManipulatedVariableReference.
// This is useful for switch cases.
type BACnetConstructedDataManipulatedVariableReferenceExactly interface {
	BACnetConstructedDataManipulatedVariableReference
	isBACnetConstructedDataManipulatedVariableReference() bool
}

// _BACnetConstructedDataManipulatedVariableReference is the data-structure of this message
type _BACnetConstructedDataManipulatedVariableReference struct {
	*_BACnetConstructedData
	ManipulatedVariableReference BACnetObjectPropertyReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataManipulatedVariableReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataManipulatedVariableReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MANIPULATED_VARIABLE_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataManipulatedVariableReference) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataManipulatedVariableReference) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataManipulatedVariableReference) GetManipulatedVariableReference() BACnetObjectPropertyReference {
	return m.ManipulatedVariableReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataManipulatedVariableReference) GetActualValue() BACnetObjectPropertyReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetObjectPropertyReference(m.GetManipulatedVariableReference())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataManipulatedVariableReference factory function for _BACnetConstructedDataManipulatedVariableReference
func NewBACnetConstructedDataManipulatedVariableReference(manipulatedVariableReference BACnetObjectPropertyReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataManipulatedVariableReference {
	_result := &_BACnetConstructedDataManipulatedVariableReference{
		ManipulatedVariableReference: manipulatedVariableReference,
		_BACnetConstructedData:       NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataManipulatedVariableReference(structType interface{}) BACnetConstructedDataManipulatedVariableReference {
	if casted, ok := structType.(BACnetConstructedDataManipulatedVariableReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataManipulatedVariableReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataManipulatedVariableReference) GetTypeName() string {
	return "BACnetConstructedDataManipulatedVariableReference"
}

func (m *_BACnetConstructedDataManipulatedVariableReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (manipulatedVariableReference)
	lengthInBits += m.ManipulatedVariableReference.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataManipulatedVariableReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataManipulatedVariableReferenceParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataManipulatedVariableReference, error) {
	return BACnetConstructedDataManipulatedVariableReferenceParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataManipulatedVariableReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataManipulatedVariableReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataManipulatedVariableReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataManipulatedVariableReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (manipulatedVariableReference)
	if pullErr := readBuffer.PullContext("manipulatedVariableReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for manipulatedVariableReference")
	}
	_manipulatedVariableReference, _manipulatedVariableReferenceErr := BACnetObjectPropertyReferenceParseWithBuffer(ctx, readBuffer)
	if _manipulatedVariableReferenceErr != nil {
		return nil, errors.Wrap(_manipulatedVariableReferenceErr, "Error parsing 'manipulatedVariableReference' field of BACnetConstructedDataManipulatedVariableReference")
	}
	manipulatedVariableReference := _manipulatedVariableReference.(BACnetObjectPropertyReference)
	if closeErr := readBuffer.CloseContext("manipulatedVariableReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for manipulatedVariableReference")
	}

	// Virtual field
	_actualValue := manipulatedVariableReference
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataManipulatedVariableReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataManipulatedVariableReference")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataManipulatedVariableReference{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ManipulatedVariableReference: manipulatedVariableReference,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataManipulatedVariableReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataManipulatedVariableReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataManipulatedVariableReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataManipulatedVariableReference")
		}

		// Simple Field (manipulatedVariableReference)
		if pushErr := writeBuffer.PushContext("manipulatedVariableReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for manipulatedVariableReference")
		}
		_manipulatedVariableReferenceErr := writeBuffer.WriteSerializable(ctx, m.GetManipulatedVariableReference())
		if popErr := writeBuffer.PopContext("manipulatedVariableReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for manipulatedVariableReference")
		}
		if _manipulatedVariableReferenceErr != nil {
			return errors.Wrap(_manipulatedVariableReferenceErr, "Error serializing 'manipulatedVariableReference' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataManipulatedVariableReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataManipulatedVariableReference")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataManipulatedVariableReference) isBACnetConstructedDataManipulatedVariableReference() bool {
	return true
}

func (m *_BACnetConstructedDataManipulatedVariableReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
