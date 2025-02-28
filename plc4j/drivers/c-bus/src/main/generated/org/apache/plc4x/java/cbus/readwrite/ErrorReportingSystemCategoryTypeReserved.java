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
package org.apache.plc4x.java.cbus.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ErrorReportingSystemCategoryTypeReserved extends ErrorReportingSystemCategoryType
    implements Message {

  // Accessors for discriminator values.
  public ErrorReportingSystemCategoryClass getErrorReportingSystemCategoryClass() {
    return null;
  }

  // Properties.
  protected final byte reservedValue;

  public ErrorReportingSystemCategoryTypeReserved(byte reservedValue) {
    super();
    this.reservedValue = reservedValue;
  }

  public byte getReservedValue() {
    return reservedValue;
  }

  @Override
  protected void serializeErrorReportingSystemCategoryTypeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ErrorReportingSystemCategoryTypeReserved");

    // Simple Field (reservedValue)
    writeSimpleField("reservedValue", reservedValue, writeUnsignedByte(writeBuffer, 4));

    writeBuffer.popContext("ErrorReportingSystemCategoryTypeReserved");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ErrorReportingSystemCategoryTypeReserved _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (reservedValue)
    lengthInBits += 4;

    return lengthInBits;
  }

  public static ErrorReportingSystemCategoryTypeBuilder
      staticParseErrorReportingSystemCategoryTypeBuilder(
          ReadBuffer readBuffer,
          ErrorReportingSystemCategoryClass errorReportingSystemCategoryClass)
          throws ParseException {
    readBuffer.pullContext("ErrorReportingSystemCategoryTypeReserved");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte reservedValue = readSimpleField("reservedValue", readUnsignedByte(readBuffer, 4));

    readBuffer.closeContext("ErrorReportingSystemCategoryTypeReserved");
    // Create the instance
    return new ErrorReportingSystemCategoryTypeReservedBuilderImpl(reservedValue);
  }

  public static class ErrorReportingSystemCategoryTypeReservedBuilderImpl
      implements ErrorReportingSystemCategoryType.ErrorReportingSystemCategoryTypeBuilder {
    private final byte reservedValue;

    public ErrorReportingSystemCategoryTypeReservedBuilderImpl(byte reservedValue) {
      this.reservedValue = reservedValue;
    }

    public ErrorReportingSystemCategoryTypeReserved build() {
      ErrorReportingSystemCategoryTypeReserved errorReportingSystemCategoryTypeReserved =
          new ErrorReportingSystemCategoryTypeReserved(reservedValue);
      return errorReportingSystemCategoryTypeReserved;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ErrorReportingSystemCategoryTypeReserved)) {
      return false;
    }
    ErrorReportingSystemCategoryTypeReserved that = (ErrorReportingSystemCategoryTypeReserved) o;
    return (getReservedValue() == that.getReservedValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getReservedValue());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
