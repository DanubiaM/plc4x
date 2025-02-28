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

// BACnetConstructedDataMaintenanceRequired is the corresponding interface of BACnetConstructedDataMaintenanceRequired
type BACnetConstructedDataMaintenanceRequired interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaintenanceRequired returns MaintenanceRequired (property field)
	GetMaintenanceRequired() BACnetMaintenanceTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetMaintenanceTagged
}

// BACnetConstructedDataMaintenanceRequiredExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMaintenanceRequired.
// This is useful for switch cases.
type BACnetConstructedDataMaintenanceRequiredExactly interface {
	BACnetConstructedDataMaintenanceRequired
	isBACnetConstructedDataMaintenanceRequired() bool
}

// _BACnetConstructedDataMaintenanceRequired is the data-structure of this message
type _BACnetConstructedDataMaintenanceRequired struct {
	*_BACnetConstructedData
	MaintenanceRequired BACnetMaintenanceTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAINTENANCE_REQUIRED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetMaintenanceRequired() BACnetMaintenanceTagged {
	return m.MaintenanceRequired
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetActualValue() BACnetMaintenanceTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetMaintenanceTagged(m.GetMaintenanceRequired())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaintenanceRequired factory function for _BACnetConstructedDataMaintenanceRequired
func NewBACnetConstructedDataMaintenanceRequired(maintenanceRequired BACnetMaintenanceTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaintenanceRequired {
	_result := &_BACnetConstructedDataMaintenanceRequired{
		MaintenanceRequired:    maintenanceRequired,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaintenanceRequired(structType interface{}) BACnetConstructedDataMaintenanceRequired {
	if casted, ok := structType.(BACnetConstructedDataMaintenanceRequired); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaintenanceRequired); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetTypeName() string {
	return "BACnetConstructedDataMaintenanceRequired"
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (maintenanceRequired)
	lengthInBits += m.MaintenanceRequired.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataMaintenanceRequiredParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaintenanceRequired, error) {
	return BACnetConstructedDataMaintenanceRequiredParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataMaintenanceRequiredParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaintenanceRequired, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaintenanceRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaintenanceRequired")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maintenanceRequired)
	if pullErr := readBuffer.PullContext("maintenanceRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maintenanceRequired")
	}
	_maintenanceRequired, _maintenanceRequiredErr := BACnetMaintenanceTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _maintenanceRequiredErr != nil {
		return nil, errors.Wrap(_maintenanceRequiredErr, "Error parsing 'maintenanceRequired' field of BACnetConstructedDataMaintenanceRequired")
	}
	maintenanceRequired := _maintenanceRequired.(BACnetMaintenanceTagged)
	if closeErr := readBuffer.CloseContext("maintenanceRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maintenanceRequired")
	}

	// Virtual field
	_actualValue := maintenanceRequired
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaintenanceRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaintenanceRequired")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMaintenanceRequired{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaintenanceRequired: maintenanceRequired,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMaintenanceRequired) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaintenanceRequired) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaintenanceRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaintenanceRequired")
		}

		// Simple Field (maintenanceRequired)
		if pushErr := writeBuffer.PushContext("maintenanceRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maintenanceRequired")
		}
		_maintenanceRequiredErr := writeBuffer.WriteSerializable(ctx, m.GetMaintenanceRequired())
		if popErr := writeBuffer.PopContext("maintenanceRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maintenanceRequired")
		}
		if _maintenanceRequiredErr != nil {
			return errors.Wrap(_maintenanceRequiredErr, "Error serializing 'maintenanceRequired' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaintenanceRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaintenanceRequired")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaintenanceRequired) isBACnetConstructedDataMaintenanceRequired() bool {
	return true
}

func (m *_BACnetConstructedDataMaintenanceRequired) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
