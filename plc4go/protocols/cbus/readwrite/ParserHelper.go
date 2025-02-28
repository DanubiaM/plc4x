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

package readwrite

import (
	"context"

	"github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type CbusParserHelper struct {
}

func (m CbusParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (interface{}, error) {
	switch typeName {
	case "HVACStatusFlags":
		return model.HVACStatusFlagsParseWithBuffer(context.Background(), io)
	case "ParameterValue":
		parameterType, _ := model.ParameterTypeByName(arguments[0])
		numBytes, err := utils.StrToUint8(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.ParameterValueParseWithBuffer(context.Background(), io, parameterType, numBytes)
	case "ReplyOrConfirmation":
		var cBusOptions model.CBusOptions
		var requestContext model.RequestContext
		return model.ReplyOrConfirmationParseWithBuffer(context.Background(), io, cBusOptions, requestContext)
	case "CBusOptions":
		return model.CBusOptionsParseWithBuffer(context.Background(), io)
	case "TemperatureBroadcastData":
		return model.TemperatureBroadcastDataParseWithBuffer(context.Background(), io)
	case "PanicStatus":
		return model.PanicStatusParseWithBuffer(context.Background(), io)
	case "IdentifyReplyCommandUnitSummary":
		return model.IdentifyReplyCommandUnitSummaryParseWithBuffer(context.Background(), io)
	case "InterfaceOptions1PowerUpSettings":
		return model.InterfaceOptions1PowerUpSettingsParseWithBuffer(context.Background(), io)
	case "MonitoredSAL":
		var cBusOptions model.CBusOptions
		return model.MonitoredSALParseWithBuffer(context.Background(), io, cBusOptions)
	case "ReplyNetwork":
		return model.ReplyNetworkParseWithBuffer(context.Background(), io)
	case "SerialNumber":
		return model.SerialNumberParseWithBuffer(context.Background(), io)
	case "CBusPointToMultiPointCommand":
		var cBusOptions model.CBusOptions
		return model.CBusPointToMultiPointCommandParseWithBuffer(context.Background(), io, cBusOptions)
	case "StatusRequest":
		return model.StatusRequestParseWithBuffer(context.Background(), io)
	case "InterfaceOptions3":
		return model.InterfaceOptions3ParseWithBuffer(context.Background(), io)
	case "InterfaceOptions1":
		return model.InterfaceOptions1ParseWithBuffer(context.Background(), io)
	case "InterfaceOptions2":
		return model.InterfaceOptions2ParseWithBuffer(context.Background(), io)
	case "HVACModeAndFlags":
		return model.HVACModeAndFlagsParseWithBuffer(context.Background(), io)
	case "LightingData":
		return model.LightingDataParseWithBuffer(context.Background(), io)
	case "SALData":
		applicationId, _ := model.ApplicationIdByName(arguments[0])
		return model.SALDataParseWithBuffer(context.Background(), io, applicationId)
	case "CBusCommand":
		var cBusOptions model.CBusOptions
		return model.CBusCommandParseWithBuffer(context.Background(), io, cBusOptions)
	case "HVACHumidity":
		return model.HVACHumidityParseWithBuffer(context.Background(), io)
	case "HVACHumidityModeAndFlags":
		return model.HVACHumidityModeAndFlagsParseWithBuffer(context.Background(), io)
	case "CBusConstants":
		return model.CBusConstantsParseWithBuffer(context.Background(), io)
	case "SerialInterfaceAddress":
		return model.SerialInterfaceAddressParseWithBuffer(context.Background(), io)
	case "MeasurementData":
		return model.MeasurementDataParseWithBuffer(context.Background(), io)
	case "HVACZoneList":
		return model.HVACZoneListParseWithBuffer(context.Background(), io)
	case "MediaTransportControlData":
		return model.MediaTransportControlDataParseWithBuffer(context.Background(), io)
	case "StatusByte":
		return model.StatusByteParseWithBuffer(context.Background(), io)
	case "TriggerControlLabelOptions":
		return model.TriggerControlLabelOptionsParseWithBuffer(context.Background(), io)
	case "HVACAuxiliaryLevel":
		return model.HVACAuxiliaryLevelParseWithBuffer(context.Background(), io)
	case "ErrorReportingData":
		return model.ErrorReportingDataParseWithBuffer(context.Background(), io)
	case "UnitAddress":
		return model.UnitAddressParseWithBuffer(context.Background(), io)
	case "SecurityArmCode":
		return model.SecurityArmCodeParseWithBuffer(context.Background(), io)
	case "MeteringData":
		return model.MeteringDataParseWithBuffer(context.Background(), io)
	case "EnableControlData":
		return model.EnableControlDataParseWithBuffer(context.Background(), io)
	case "ApplicationAddress2":
		return model.ApplicationAddress2ParseWithBuffer(context.Background(), io)
	case "ApplicationAddress1":
		return model.ApplicationAddress1ParseWithBuffer(context.Background(), io)
	case "RequestContext":
		return model.RequestContextParseWithBuffer(context.Background(), io)
	case "TriggerControlData":
		return model.TriggerControlDataParseWithBuffer(context.Background(), io)
	case "HVACStartTime":
		return model.HVACStartTimeParseWithBuffer(context.Background(), io)
	case "HVACTemperature":
		return model.HVACTemperatureParseWithBuffer(context.Background(), io)
	case "RequestTermination":
		return model.RequestTerminationParseWithBuffer(context.Background(), io)
	case "CBusMessage":
		isResponse, err := utils.StrToBool(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		var requestContext model.RequestContext
		var cBusOptions model.CBusOptions
		return model.CBusMessageParseWithBuffer(context.Background(), io, isResponse, requestContext, cBusOptions)
	case "ErrorReportingSystemCategory":
		return model.ErrorReportingSystemCategoryParseWithBuffer(context.Background(), io)
	case "PowerUp":
		return model.PowerUpParseWithBuffer(context.Background(), io)
	case "Reply":
		var cBusOptions model.CBusOptions
		var requestContext model.RequestContext
		return model.ReplyParseWithBuffer(context.Background(), io, cBusOptions, requestContext)
	case "TelephonyData":
		return model.TelephonyDataParseWithBuffer(context.Background(), io)
	case "HVACHumidityStatusFlags":
		return model.HVACHumidityStatusFlagsParseWithBuffer(context.Background(), io)
	case "ParameterChange":
		return model.ParameterChangeParseWithBuffer(context.Background(), io)
	case "ErrorReportingSystemCategoryType":
		errorReportingSystemCategoryClass, _ := model.ErrorReportingSystemCategoryClassByName(arguments[0])
		return model.ErrorReportingSystemCategoryTypeParseWithBuffer(context.Background(), io, errorReportingSystemCategoryClass)
	case "Confirmation":
		return model.ConfirmationParseWithBuffer(context.Background(), io)
	case "SecurityData":
		return model.SecurityDataParseWithBuffer(context.Background(), io)
	case "NetworkProtocolControlInformation":
		return model.NetworkProtocolControlInformationParseWithBuffer(context.Background(), io)
	case "CBusHeader":
		return model.CBusHeaderParseWithBuffer(context.Background(), io)
	case "Request":
		var cBusOptions model.CBusOptions
		return model.RequestParseWithBuffer(context.Background(), io, cBusOptions)
	case "Alpha":
		return model.AlphaParseWithBuffer(context.Background(), io)
	case "CALData":
		var requestContext model.RequestContext
		return model.CALDataParseWithBuffer(context.Background(), io, requestContext)
	case "Checksum":
		return model.ChecksumParseWithBuffer(context.Background(), io)
	case "CALReply":
		var cBusOptions model.CBusOptions
		var requestContext model.RequestContext
		return model.CALReplyParseWithBuffer(context.Background(), io, cBusOptions, requestContext)
	case "CustomManufacturer":
		numBytes, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.CustomManufacturerParseWithBuffer(context.Background(), io, numBytes)
	case "AccessControlData":
		return model.AccessControlDataParseWithBuffer(context.Background(), io)
	case "ClockAndTimekeepingData":
		return model.ClockAndTimekeepingDataParseWithBuffer(context.Background(), io)
	case "NetworkRoute":
		return model.NetworkRouteParseWithBuffer(context.Background(), io)
	case "ResponseTermination":
		return model.ResponseTerminationParseWithBuffer(context.Background(), io)
	case "LevelInformation":
		return model.LevelInformationParseWithBuffer(context.Background(), io)
	case "TamperStatus":
		return model.TamperStatusParseWithBuffer(context.Background(), io)
	case "IdentifyReplyCommand":
		attribute, _ := model.AttributeByName(arguments[0])
		numBytes, err := utils.StrToUint8(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.IdentifyReplyCommandParseWithBuffer(context.Background(), io, attribute, numBytes)
	case "HVACRawLevels":
		return model.HVACRawLevelsParseWithBuffer(context.Background(), io)
	case "ZoneStatus":
		return model.ZoneStatusParseWithBuffer(context.Background(), io)
	case "BridgeAddress":
		return model.BridgeAddressParseWithBuffer(context.Background(), io)
	case "LightingLabelOptions":
		return model.LightingLabelOptionsParseWithBuffer(context.Background(), io)
	case "CustomTypes":
		numBytes, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.CustomTypesParseWithBuffer(context.Background(), io, numBytes)
	case "EncodedReply":
		var cBusOptions model.CBusOptions
		var requestContext model.RequestContext
		return model.EncodedReplyParseWithBuffer(context.Background(), io, cBusOptions, requestContext)
	case "CBusPointToPointToMultiPointCommand":
		var cBusOptions model.CBusOptions
		return model.CBusPointToPointToMultiPointCommandParseWithBuffer(context.Background(), io, cBusOptions)
	case "CBusPointToPointCommand":
		var cBusOptions model.CBusOptions
		return model.CBusPointToPointCommandParseWithBuffer(context.Background(), io, cBusOptions)
	case "AirConditioningData":
		return model.AirConditioningDataParseWithBuffer(context.Background(), io)
	case "LogicAssignment":
		return model.LogicAssignmentParseWithBuffer(context.Background(), io)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
