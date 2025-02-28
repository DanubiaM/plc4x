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

// ErrorReportingSystemCategory is the corresponding interface of ErrorReportingSystemCategory
type ErrorReportingSystemCategory interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetSystemCategoryClass returns SystemCategoryClass (property field)
	GetSystemCategoryClass() ErrorReportingSystemCategoryClass
	// GetSystemCategoryType returns SystemCategoryType (property field)
	GetSystemCategoryType() ErrorReportingSystemCategoryType
	// GetSystemCategoryVariant returns SystemCategoryVariant (property field)
	GetSystemCategoryVariant() ErrorReportingSystemCategoryVariant
}

// ErrorReportingSystemCategoryExactly can be used when we want exactly this type and not a type which fulfills ErrorReportingSystemCategory.
// This is useful for switch cases.
type ErrorReportingSystemCategoryExactly interface {
	ErrorReportingSystemCategory
	isErrorReportingSystemCategory() bool
}

// _ErrorReportingSystemCategory is the data-structure of this message
type _ErrorReportingSystemCategory struct {
	SystemCategoryClass   ErrorReportingSystemCategoryClass
	SystemCategoryType    ErrorReportingSystemCategoryType
	SystemCategoryVariant ErrorReportingSystemCategoryVariant
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategory) GetSystemCategoryClass() ErrorReportingSystemCategoryClass {
	return m.SystemCategoryClass
}

func (m *_ErrorReportingSystemCategory) GetSystemCategoryType() ErrorReportingSystemCategoryType {
	return m.SystemCategoryType
}

func (m *_ErrorReportingSystemCategory) GetSystemCategoryVariant() ErrorReportingSystemCategoryVariant {
	return m.SystemCategoryVariant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewErrorReportingSystemCategory factory function for _ErrorReportingSystemCategory
func NewErrorReportingSystemCategory(systemCategoryClass ErrorReportingSystemCategoryClass, systemCategoryType ErrorReportingSystemCategoryType, systemCategoryVariant ErrorReportingSystemCategoryVariant) *_ErrorReportingSystemCategory {
	return &_ErrorReportingSystemCategory{SystemCategoryClass: systemCategoryClass, SystemCategoryType: systemCategoryType, SystemCategoryVariant: systemCategoryVariant}
}

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategory(structType interface{}) ErrorReportingSystemCategory {
	if casted, ok := structType.(ErrorReportingSystemCategory); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategory); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategory) GetTypeName() string {
	return "ErrorReportingSystemCategory"
}

func (m *_ErrorReportingSystemCategory) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (systemCategoryClass)
	lengthInBits += 4

	// Simple field (systemCategoryType)
	lengthInBits += m.SystemCategoryType.GetLengthInBits(ctx)

	// Simple field (systemCategoryVariant)
	lengthInBits += 2

	return lengthInBits
}

func (m *_ErrorReportingSystemCategory) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryParse(theBytes []byte) (ErrorReportingSystemCategory, error) {
	return ErrorReportingSystemCategoryParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingSystemCategoryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSystemCategory, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategory"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategory")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (systemCategoryClass)
	if pullErr := readBuffer.PullContext("systemCategoryClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for systemCategoryClass")
	}
	_systemCategoryClass, _systemCategoryClassErr := ErrorReportingSystemCategoryClassParseWithBuffer(ctx, readBuffer)
	if _systemCategoryClassErr != nil {
		return nil, errors.Wrap(_systemCategoryClassErr, "Error parsing 'systemCategoryClass' field of ErrorReportingSystemCategory")
	}
	systemCategoryClass := _systemCategoryClass
	if closeErr := readBuffer.CloseContext("systemCategoryClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for systemCategoryClass")
	}

	// Simple Field (systemCategoryType)
	if pullErr := readBuffer.PullContext("systemCategoryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for systemCategoryType")
	}
	_systemCategoryType, _systemCategoryTypeErr := ErrorReportingSystemCategoryTypeParseWithBuffer(ctx, readBuffer, ErrorReportingSystemCategoryClass(systemCategoryClass))
	if _systemCategoryTypeErr != nil {
		return nil, errors.Wrap(_systemCategoryTypeErr, "Error parsing 'systemCategoryType' field of ErrorReportingSystemCategory")
	}
	systemCategoryType := _systemCategoryType.(ErrorReportingSystemCategoryType)
	if closeErr := readBuffer.CloseContext("systemCategoryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for systemCategoryType")
	}

	// Simple Field (systemCategoryVariant)
	if pullErr := readBuffer.PullContext("systemCategoryVariant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for systemCategoryVariant")
	}
	_systemCategoryVariant, _systemCategoryVariantErr := ErrorReportingSystemCategoryVariantParseWithBuffer(ctx, readBuffer)
	if _systemCategoryVariantErr != nil {
		return nil, errors.Wrap(_systemCategoryVariantErr, "Error parsing 'systemCategoryVariant' field of ErrorReportingSystemCategory")
	}
	systemCategoryVariant := _systemCategoryVariant
	if closeErr := readBuffer.CloseContext("systemCategoryVariant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for systemCategoryVariant")
	}

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategory"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategory")
	}

	// Create the instance
	return &_ErrorReportingSystemCategory{
		SystemCategoryClass:   systemCategoryClass,
		SystemCategoryType:    systemCategoryType,
		SystemCategoryVariant: systemCategoryVariant,
	}, nil
}

func (m *_ErrorReportingSystemCategory) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategory) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategory"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategory")
	}

	// Simple Field (systemCategoryClass)
	if pushErr := writeBuffer.PushContext("systemCategoryClass"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for systemCategoryClass")
	}
	_systemCategoryClassErr := writeBuffer.WriteSerializable(ctx, m.GetSystemCategoryClass())
	if popErr := writeBuffer.PopContext("systemCategoryClass"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for systemCategoryClass")
	}
	if _systemCategoryClassErr != nil {
		return errors.Wrap(_systemCategoryClassErr, "Error serializing 'systemCategoryClass' field")
	}

	// Simple Field (systemCategoryType)
	if pushErr := writeBuffer.PushContext("systemCategoryType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for systemCategoryType")
	}
	_systemCategoryTypeErr := writeBuffer.WriteSerializable(ctx, m.GetSystemCategoryType())
	if popErr := writeBuffer.PopContext("systemCategoryType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for systemCategoryType")
	}
	if _systemCategoryTypeErr != nil {
		return errors.Wrap(_systemCategoryTypeErr, "Error serializing 'systemCategoryType' field")
	}

	// Simple Field (systemCategoryVariant)
	if pushErr := writeBuffer.PushContext("systemCategoryVariant"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for systemCategoryVariant")
	}
	_systemCategoryVariantErr := writeBuffer.WriteSerializable(ctx, m.GetSystemCategoryVariant())
	if popErr := writeBuffer.PopContext("systemCategoryVariant"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for systemCategoryVariant")
	}
	if _systemCategoryVariantErr != nil {
		return errors.Wrap(_systemCategoryVariantErr, "Error serializing 'systemCategoryVariant' field")
	}

	if popErr := writeBuffer.PopContext("ErrorReportingSystemCategory"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategory")
	}
	return nil
}

func (m *_ErrorReportingSystemCategory) isErrorReportingSystemCategory() bool {
	return true
}

func (m *_ErrorReportingSystemCategory) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
