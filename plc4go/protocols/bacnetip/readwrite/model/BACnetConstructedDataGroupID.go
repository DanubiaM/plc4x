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

// BACnetConstructedDataGroupID is the corresponding interface of BACnetConstructedDataGroupID
type BACnetConstructedDataGroupID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetGroupId returns GroupId (property field)
	GetGroupId() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataGroupIDExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataGroupID.
// This is useful for switch cases.
type BACnetConstructedDataGroupIDExactly interface {
	BACnetConstructedDataGroupID
	isBACnetConstructedDataGroupID() bool
}

// _BACnetConstructedDataGroupID is the data-structure of this message
type _BACnetConstructedDataGroupID struct {
	*_BACnetConstructedData
	GroupId BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGroupID) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataGroupID) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GROUP_ID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGroupID) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataGroupID) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGroupID) GetGroupId() BACnetApplicationTagUnsignedInteger {
	return m.GroupId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGroupID) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetGroupId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataGroupID factory function for _BACnetConstructedDataGroupID
func NewBACnetConstructedDataGroupID(groupId BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGroupID {
	_result := &_BACnetConstructedDataGroupID{
		GroupId:                groupId,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGroupID(structType interface{}) BACnetConstructedDataGroupID {
	if casted, ok := structType.(BACnetConstructedDataGroupID); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupID); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGroupID) GetTypeName() string {
	return "BACnetConstructedDataGroupID"
}

func (m *_BACnetConstructedDataGroupID) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (groupId)
	lengthInBits += m.GroupId.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataGroupID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataGroupIDParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataGroupID, error) {
	return BACnetConstructedDataGroupIDParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataGroupIDParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataGroupID, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupID"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGroupID")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (groupId)
	if pullErr := readBuffer.PullContext("groupId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for groupId")
	}
	_groupId, _groupIdErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _groupIdErr != nil {
		return nil, errors.Wrap(_groupIdErr, "Error parsing 'groupId' field of BACnetConstructedDataGroupID")
	}
	groupId := _groupId.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("groupId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for groupId")
	}

	// Virtual field
	_actualValue := groupId
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupID"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGroupID")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataGroupID{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		GroupId: groupId,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataGroupID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGroupID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupID"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGroupID")
		}

		// Simple Field (groupId)
		if pushErr := writeBuffer.PushContext("groupId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for groupId")
		}
		_groupIdErr := writeBuffer.WriteSerializable(ctx, m.GetGroupId())
		if popErr := writeBuffer.PopContext("groupId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for groupId")
		}
		if _groupIdErr != nil {
			return errors.Wrap(_groupIdErr, "Error serializing 'groupId' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupID"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGroupID")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGroupID) isBACnetConstructedDataGroupID() bool {
	return true
}

func (m *_BACnetConstructedDataGroupID) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
