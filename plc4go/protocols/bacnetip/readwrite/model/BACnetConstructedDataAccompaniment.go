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

// BACnetConstructedDataAccompaniment is the corresponding interface of BACnetConstructedDataAccompaniment
type BACnetConstructedDataAccompaniment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccompaniment returns Accompaniment (property field)
	GetAccompaniment() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
}

// BACnetConstructedDataAccompanimentExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccompaniment.
// This is useful for switch cases.
type BACnetConstructedDataAccompanimentExactly interface {
	BACnetConstructedDataAccompaniment
	isBACnetConstructedDataAccompaniment() bool
}

// _BACnetConstructedDataAccompaniment is the data-structure of this message
type _BACnetConstructedDataAccompaniment struct {
	*_BACnetConstructedData
	Accompaniment BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccompaniment) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccompaniment) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCOMPANIMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccompaniment) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccompaniment) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccompaniment) GetAccompaniment() BACnetDeviceObjectReference {
	return m.Accompaniment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccompaniment) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetAccompaniment())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccompaniment factory function for _BACnetConstructedDataAccompaniment
func NewBACnetConstructedDataAccompaniment(accompaniment BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccompaniment {
	_result := &_BACnetConstructedDataAccompaniment{
		Accompaniment:          accompaniment,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccompaniment(structType interface{}) BACnetConstructedDataAccompaniment {
	if casted, ok := structType.(BACnetConstructedDataAccompaniment); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccompaniment); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccompaniment) GetTypeName() string {
	return "BACnetConstructedDataAccompaniment"
}

func (m *_BACnetConstructedDataAccompaniment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (accompaniment)
	lengthInBits += m.Accompaniment.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccompaniment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAccompanimentParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccompaniment, error) {
	return BACnetConstructedDataAccompanimentParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAccompanimentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccompaniment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccompaniment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccompaniment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accompaniment)
	if pullErr := readBuffer.PullContext("accompaniment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accompaniment")
	}
	_accompaniment, _accompanimentErr := BACnetDeviceObjectReferenceParseWithBuffer(ctx, readBuffer)
	if _accompanimentErr != nil {
		return nil, errors.Wrap(_accompanimentErr, "Error parsing 'accompaniment' field of BACnetConstructedDataAccompaniment")
	}
	accompaniment := _accompaniment.(BACnetDeviceObjectReference)
	if closeErr := readBuffer.CloseContext("accompaniment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accompaniment")
	}

	// Virtual field
	_actualValue := accompaniment
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccompaniment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccompaniment")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccompaniment{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Accompaniment: accompaniment,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccompaniment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccompaniment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccompaniment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccompaniment")
		}

		// Simple Field (accompaniment)
		if pushErr := writeBuffer.PushContext("accompaniment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accompaniment")
		}
		_accompanimentErr := writeBuffer.WriteSerializable(ctx, m.GetAccompaniment())
		if popErr := writeBuffer.PopContext("accompaniment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accompaniment")
		}
		if _accompanimentErr != nil {
			return errors.Wrap(_accompanimentErr, "Error serializing 'accompaniment' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccompaniment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccompaniment")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccompaniment) isBACnetConstructedDataAccompaniment() bool {
	return true
}

func (m *_BACnetConstructedDataAccompaniment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
