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

// BACnetConstructedDataLimitEnable is the corresponding interface of BACnetConstructedDataLimitEnable
type BACnetConstructedDataLimitEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLimitEnable returns LimitEnable (property field)
	GetLimitEnable() BACnetLimitEnableTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLimitEnableTagged
}

// BACnetConstructedDataLimitEnableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLimitEnable.
// This is useful for switch cases.
type BACnetConstructedDataLimitEnableExactly interface {
	BACnetConstructedDataLimitEnable
	isBACnetConstructedDataLimitEnable() bool
}

// _BACnetConstructedDataLimitEnable is the data-structure of this message
type _BACnetConstructedDataLimitEnable struct {
	*_BACnetConstructedData
	LimitEnable BACnetLimitEnableTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLimitEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIMIT_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLimitEnable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLimitEnable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetLimitEnable() BACnetLimitEnableTagged {
	return m.LimitEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetActualValue() BACnetLimitEnableTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLimitEnableTagged(m.GetLimitEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLimitEnable factory function for _BACnetConstructedDataLimitEnable
func NewBACnetConstructedDataLimitEnable(limitEnable BACnetLimitEnableTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLimitEnable {
	_result := &_BACnetConstructedDataLimitEnable{
		LimitEnable:            limitEnable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLimitEnable(structType interface{}) BACnetConstructedDataLimitEnable {
	if casted, ok := structType.(BACnetConstructedDataLimitEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLimitEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLimitEnable) GetTypeName() string {
	return "BACnetConstructedDataLimitEnable"
}

func (m *_BACnetConstructedDataLimitEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (limitEnable)
	lengthInBits += m.LimitEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLimitEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLimitEnableParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLimitEnable, error) {
	return BACnetConstructedDataLimitEnableParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLimitEnableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLimitEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLimitEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLimitEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (limitEnable)
	if pullErr := readBuffer.PullContext("limitEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for limitEnable")
	}
	_limitEnable, _limitEnableErr := BACnetLimitEnableTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _limitEnableErr != nil {
		return nil, errors.Wrap(_limitEnableErr, "Error parsing 'limitEnable' field of BACnetConstructedDataLimitEnable")
	}
	limitEnable := _limitEnable.(BACnetLimitEnableTagged)
	if closeErr := readBuffer.CloseContext("limitEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for limitEnable")
	}

	// Virtual field
	_actualValue := limitEnable
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLimitEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLimitEnable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLimitEnable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LimitEnable: limitEnable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLimitEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLimitEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLimitEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLimitEnable")
		}

		// Simple Field (limitEnable)
		if pushErr := writeBuffer.PushContext("limitEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for limitEnable")
		}
		_limitEnableErr := writeBuffer.WriteSerializable(ctx, m.GetLimitEnable())
		if popErr := writeBuffer.PopContext("limitEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for limitEnable")
		}
		if _limitEnableErr != nil {
			return errors.Wrap(_limitEnableErr, "Error serializing 'limitEnable' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLimitEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLimitEnable")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLimitEnable) isBACnetConstructedDataLimitEnable() bool {
	return true
}

func (m *_BACnetConstructedDataLimitEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
