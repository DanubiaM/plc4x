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

// BACnetConstructedDataDatabaseRevision is the corresponding interface of BACnetConstructedDataDatabaseRevision
type BACnetConstructedDataDatabaseRevision interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDatabaseRevision returns DatabaseRevision (property field)
	GetDatabaseRevision() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataDatabaseRevisionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDatabaseRevision.
// This is useful for switch cases.
type BACnetConstructedDataDatabaseRevisionExactly interface {
	BACnetConstructedDataDatabaseRevision
	isBACnetConstructedDataDatabaseRevision() bool
}

// _BACnetConstructedDataDatabaseRevision is the data-structure of this message
type _BACnetConstructedDataDatabaseRevision struct {
	*_BACnetConstructedData
	DatabaseRevision BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDatabaseRevision) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDatabaseRevision) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DATABASE_REVISION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDatabaseRevision) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDatabaseRevision) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDatabaseRevision) GetDatabaseRevision() BACnetApplicationTagUnsignedInteger {
	return m.DatabaseRevision
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDatabaseRevision) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetDatabaseRevision())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDatabaseRevision factory function for _BACnetConstructedDataDatabaseRevision
func NewBACnetConstructedDataDatabaseRevision(databaseRevision BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDatabaseRevision {
	_result := &_BACnetConstructedDataDatabaseRevision{
		DatabaseRevision:       databaseRevision,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDatabaseRevision(structType interface{}) BACnetConstructedDataDatabaseRevision {
	if casted, ok := structType.(BACnetConstructedDataDatabaseRevision); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDatabaseRevision); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDatabaseRevision) GetTypeName() string {
	return "BACnetConstructedDataDatabaseRevision"
}

func (m *_BACnetConstructedDataDatabaseRevision) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (databaseRevision)
	lengthInBits += m.DatabaseRevision.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDatabaseRevision) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataDatabaseRevisionParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDatabaseRevision, error) {
	return BACnetConstructedDataDatabaseRevisionParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataDatabaseRevisionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDatabaseRevision, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDatabaseRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDatabaseRevision")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (databaseRevision)
	if pullErr := readBuffer.PullContext("databaseRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for databaseRevision")
	}
	_databaseRevision, _databaseRevisionErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _databaseRevisionErr != nil {
		return nil, errors.Wrap(_databaseRevisionErr, "Error parsing 'databaseRevision' field of BACnetConstructedDataDatabaseRevision")
	}
	databaseRevision := _databaseRevision.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("databaseRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for databaseRevision")
	}

	// Virtual field
	_actualValue := databaseRevision
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDatabaseRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDatabaseRevision")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDatabaseRevision{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DatabaseRevision: databaseRevision,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDatabaseRevision) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDatabaseRevision) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDatabaseRevision"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDatabaseRevision")
		}

		// Simple Field (databaseRevision)
		if pushErr := writeBuffer.PushContext("databaseRevision"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for databaseRevision")
		}
		_databaseRevisionErr := writeBuffer.WriteSerializable(ctx, m.GetDatabaseRevision())
		if popErr := writeBuffer.PopContext("databaseRevision"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for databaseRevision")
		}
		if _databaseRevisionErr != nil {
			return errors.Wrap(_databaseRevisionErr, "Error serializing 'databaseRevision' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDatabaseRevision"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDatabaseRevision")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDatabaseRevision) isBACnetConstructedDataDatabaseRevision() bool {
	return true
}

func (m *_BACnetConstructedDataDatabaseRevision) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
