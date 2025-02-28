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

// BACnetConstructedDataOccupancyCountAdjust is the corresponding interface of BACnetConstructedDataOccupancyCountAdjust
type BACnetConstructedDataOccupancyCountAdjust interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetOccupancyCountAdjust returns OccupancyCountAdjust (property field)
	GetOccupancyCountAdjust() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataOccupancyCountAdjustExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOccupancyCountAdjust.
// This is useful for switch cases.
type BACnetConstructedDataOccupancyCountAdjustExactly interface {
	BACnetConstructedDataOccupancyCountAdjust
	isBACnetConstructedDataOccupancyCountAdjust() bool
}

// _BACnetConstructedDataOccupancyCountAdjust is the data-structure of this message
type _BACnetConstructedDataOccupancyCountAdjust struct {
	*_BACnetConstructedData
	OccupancyCountAdjust BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_COUNT_ADJUST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyCountAdjust) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetOccupancyCountAdjust() BACnetApplicationTagBoolean {
	return m.OccupancyCountAdjust
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetOccupancyCountAdjust())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOccupancyCountAdjust factory function for _BACnetConstructedDataOccupancyCountAdjust
func NewBACnetConstructedDataOccupancyCountAdjust(occupancyCountAdjust BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyCountAdjust {
	_result := &_BACnetConstructedDataOccupancyCountAdjust{
		OccupancyCountAdjust:   occupancyCountAdjust,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyCountAdjust(structType interface{}) BACnetConstructedDataOccupancyCountAdjust {
	if casted, ok := structType.(BACnetConstructedDataOccupancyCountAdjust); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyCountAdjust); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetTypeName() string {
	return "BACnetConstructedDataOccupancyCountAdjust"
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (occupancyCountAdjust)
	lengthInBits += m.OccupancyCountAdjust.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataOccupancyCountAdjustParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOccupancyCountAdjust, error) {
	return BACnetConstructedDataOccupancyCountAdjustParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataOccupancyCountAdjustParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOccupancyCountAdjust, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyCountAdjust"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyCountAdjust")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (occupancyCountAdjust)
	if pullErr := readBuffer.PullContext("occupancyCountAdjust"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for occupancyCountAdjust")
	}
	_occupancyCountAdjust, _occupancyCountAdjustErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _occupancyCountAdjustErr != nil {
		return nil, errors.Wrap(_occupancyCountAdjustErr, "Error parsing 'occupancyCountAdjust' field of BACnetConstructedDataOccupancyCountAdjust")
	}
	occupancyCountAdjust := _occupancyCountAdjust.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("occupancyCountAdjust"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for occupancyCountAdjust")
	}

	// Virtual field
	_actualValue := occupancyCountAdjust
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyCountAdjust"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyCountAdjust")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataOccupancyCountAdjust{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		OccupancyCountAdjust: occupancyCountAdjust,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyCountAdjust"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyCountAdjust")
		}

		// Simple Field (occupancyCountAdjust)
		if pushErr := writeBuffer.PushContext("occupancyCountAdjust"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for occupancyCountAdjust")
		}
		_occupancyCountAdjustErr := writeBuffer.WriteSerializable(ctx, m.GetOccupancyCountAdjust())
		if popErr := writeBuffer.PopContext("occupancyCountAdjust"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for occupancyCountAdjust")
		}
		if _occupancyCountAdjustErr != nil {
			return errors.Wrap(_occupancyCountAdjustErr, "Error serializing 'occupancyCountAdjust' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyCountAdjust"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyCountAdjust")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) isBACnetConstructedDataOccupancyCountAdjust() bool {
	return true
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
