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

// BACnetConstructedDataLightingOutputTrackingValue is the corresponding interface of BACnetConstructedDataLightingOutputTrackingValue
type BACnetConstructedDataLightingOutputTrackingValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTrackingValue returns TrackingValue (property field)
	GetTrackingValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataLightingOutputTrackingValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLightingOutputTrackingValue.
// This is useful for switch cases.
type BACnetConstructedDataLightingOutputTrackingValueExactly interface {
	BACnetConstructedDataLightingOutputTrackingValue
	isBACnetConstructedDataLightingOutputTrackingValue() bool
}

// _BACnetConstructedDataLightingOutputTrackingValue is the data-structure of this message
type _BACnetConstructedDataLightingOutputTrackingValue struct {
	*_BACnetConstructedData
	TrackingValue BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLightingOutputTrackingValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIGHTING_OUTPUT
}

func (m *_BACnetConstructedDataLightingOutputTrackingValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRACKING_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLightingOutputTrackingValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLightingOutputTrackingValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLightingOutputTrackingValue) GetTrackingValue() BACnetApplicationTagReal {
	return m.TrackingValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLightingOutputTrackingValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetTrackingValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLightingOutputTrackingValue factory function for _BACnetConstructedDataLightingOutputTrackingValue
func NewBACnetConstructedDataLightingOutputTrackingValue(trackingValue BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLightingOutputTrackingValue {
	_result := &_BACnetConstructedDataLightingOutputTrackingValue{
		TrackingValue:          trackingValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLightingOutputTrackingValue(structType interface{}) BACnetConstructedDataLightingOutputTrackingValue {
	if casted, ok := structType.(BACnetConstructedDataLightingOutputTrackingValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingOutputTrackingValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLightingOutputTrackingValue) GetTypeName() string {
	return "BACnetConstructedDataLightingOutputTrackingValue"
}

func (m *_BACnetConstructedDataLightingOutputTrackingValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (trackingValue)
	lengthInBits += m.TrackingValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLightingOutputTrackingValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLightingOutputTrackingValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLightingOutputTrackingValue, error) {
	return BACnetConstructedDataLightingOutputTrackingValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLightingOutputTrackingValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLightingOutputTrackingValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingOutputTrackingValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingOutputTrackingValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (trackingValue)
	if pullErr := readBuffer.PullContext("trackingValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for trackingValue")
	}
	_trackingValue, _trackingValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _trackingValueErr != nil {
		return nil, errors.Wrap(_trackingValueErr, "Error parsing 'trackingValue' field of BACnetConstructedDataLightingOutputTrackingValue")
	}
	trackingValue := _trackingValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("trackingValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for trackingValue")
	}

	// Virtual field
	_actualValue := trackingValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingOutputTrackingValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingOutputTrackingValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLightingOutputTrackingValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		TrackingValue: trackingValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLightingOutputTrackingValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLightingOutputTrackingValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingOutputTrackingValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingOutputTrackingValue")
		}

		// Simple Field (trackingValue)
		if pushErr := writeBuffer.PushContext("trackingValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for trackingValue")
		}
		_trackingValueErr := writeBuffer.WriteSerializable(ctx, m.GetTrackingValue())
		if popErr := writeBuffer.PopContext("trackingValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for trackingValue")
		}
		if _trackingValueErr != nil {
			return errors.Wrap(_trackingValueErr, "Error serializing 'trackingValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingOutputTrackingValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingOutputTrackingValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLightingOutputTrackingValue) isBACnetConstructedDataLightingOutputTrackingValue() bool {
	return true
}

func (m *_BACnetConstructedDataLightingOutputTrackingValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
