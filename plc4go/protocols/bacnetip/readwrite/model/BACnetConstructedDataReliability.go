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

// BACnetConstructedDataReliability is the corresponding interface of BACnetConstructedDataReliability
type BACnetConstructedDataReliability interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetReliability returns Reliability (property field)
	GetReliability() BACnetReliabilityTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetReliabilityTagged
}

// BACnetConstructedDataReliabilityExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataReliability.
// This is useful for switch cases.
type BACnetConstructedDataReliabilityExactly interface {
	BACnetConstructedDataReliability
	isBACnetConstructedDataReliability() bool
}

// _BACnetConstructedDataReliability is the data-structure of this message
type _BACnetConstructedDataReliability struct {
	*_BACnetConstructedData
	Reliability BACnetReliabilityTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataReliability) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataReliability) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELIABILITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataReliability) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataReliability) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataReliability) GetReliability() BACnetReliabilityTagged {
	return m.Reliability
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataReliability) GetActualValue() BACnetReliabilityTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetReliabilityTagged(m.GetReliability())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataReliability factory function for _BACnetConstructedDataReliability
func NewBACnetConstructedDataReliability(reliability BACnetReliabilityTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataReliability {
	_result := &_BACnetConstructedDataReliability{
		Reliability:            reliability,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataReliability(structType interface{}) BACnetConstructedDataReliability {
	if casted, ok := structType.(BACnetConstructedDataReliability); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReliability); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataReliability) GetTypeName() string {
	return "BACnetConstructedDataReliability"
}

func (m *_BACnetConstructedDataReliability) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (reliability)
	lengthInBits += m.Reliability.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataReliability) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataReliabilityParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataReliability, error) {
	return BACnetConstructedDataReliabilityParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataReliabilityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataReliability, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReliability"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataReliability")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reliability)
	if pullErr := readBuffer.PullContext("reliability"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reliability")
	}
	_reliability, _reliabilityErr := BACnetReliabilityTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _reliabilityErr != nil {
		return nil, errors.Wrap(_reliabilityErr, "Error parsing 'reliability' field of BACnetConstructedDataReliability")
	}
	reliability := _reliability.(BACnetReliabilityTagged)
	if closeErr := readBuffer.CloseContext("reliability"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reliability")
	}

	// Virtual field
	_actualValue := reliability
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReliability"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataReliability")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataReliability{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Reliability: reliability,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataReliability) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataReliability) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReliability"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataReliability")
		}

		// Simple Field (reliability)
		if pushErr := writeBuffer.PushContext("reliability"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reliability")
		}
		_reliabilityErr := writeBuffer.WriteSerializable(ctx, m.GetReliability())
		if popErr := writeBuffer.PopContext("reliability"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reliability")
		}
		if _reliabilityErr != nil {
			return errors.Wrap(_reliabilityErr, "Error serializing 'reliability' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReliability"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataReliability")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataReliability) isBACnetConstructedDataReliability() bool {
	return true
}

func (m *_BACnetConstructedDataReliability) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
