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
package org.apache.plc4x.java.modbus.readwrite;

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

public class ModbusPDUReadCoilsRequest extends ModbusPDU implements Message {

  // Accessors for discriminator values.
  public Boolean getErrorFlag() {
    return (boolean) false;
  }

  public Short getFunctionFlag() {
    return (short) 0x01;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  // Properties.
  protected final int startingAddress;
  protected final int quantity;

  public ModbusPDUReadCoilsRequest(int startingAddress, int quantity) {
    super();
    this.startingAddress = startingAddress;
    this.quantity = quantity;
  }

  public int getStartingAddress() {
    return startingAddress;
  }

  public int getQuantity() {
    return quantity;
  }

  @Override
  protected void serializeModbusPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ModbusPDUReadCoilsRequest");

    // Simple Field (startingAddress)
    writeSimpleField("startingAddress", startingAddress, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (quantity)
    writeSimpleField("quantity", quantity, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("ModbusPDUReadCoilsRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ModbusPDUReadCoilsRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (startingAddress)
    lengthInBits += 16;

    // Simple field (quantity)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static ModbusPDUBuilder staticParseModbusPDUBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("ModbusPDUReadCoilsRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int startingAddress = readSimpleField("startingAddress", readUnsignedInt(readBuffer, 16));

    int quantity = readSimpleField("quantity", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("ModbusPDUReadCoilsRequest");
    // Create the instance
    return new ModbusPDUReadCoilsRequestBuilderImpl(startingAddress, quantity);
  }

  public static class ModbusPDUReadCoilsRequestBuilderImpl implements ModbusPDU.ModbusPDUBuilder {
    private final int startingAddress;
    private final int quantity;

    public ModbusPDUReadCoilsRequestBuilderImpl(int startingAddress, int quantity) {
      this.startingAddress = startingAddress;
      this.quantity = quantity;
    }

    public ModbusPDUReadCoilsRequest build() {
      ModbusPDUReadCoilsRequest modbusPDUReadCoilsRequest =
          new ModbusPDUReadCoilsRequest(startingAddress, quantity);
      return modbusPDUReadCoilsRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModbusPDUReadCoilsRequest)) {
      return false;
    }
    ModbusPDUReadCoilsRequest that = (ModbusPDUReadCoilsRequest) o;
    return (getStartingAddress() == that.getStartingAddress())
        && (getQuantity() == that.getQuantity())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getStartingAddress(), getQuantity());
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
