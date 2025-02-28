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
package org.apache.plc4x.java.profinet.readwrite;

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

public class DceRpc_InterfaceUuid_ParameterInterface extends DceRpc_InterfaceUuid
    implements Message {

  // Accessors for discriminator values.
  public Long getInterfaceType() {
    return (long) 0xDEA00004L;
  }

  public DceRpc_InterfaceUuid_ParameterInterface() {
    super();
  }

  @Override
  protected void serializeDceRpc_InterfaceUuidChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DceRpc_InterfaceUuid_ParameterInterface");

    writeBuffer.popContext("DceRpc_InterfaceUuid_ParameterInterface");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DceRpc_InterfaceUuid_ParameterInterface _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    return lengthInBits;
  }

  public static DceRpc_InterfaceUuidBuilder staticParseDceRpc_InterfaceUuidBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DceRpc_InterfaceUuid_ParameterInterface");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    readBuffer.closeContext("DceRpc_InterfaceUuid_ParameterInterface");
    // Create the instance
    return new DceRpc_InterfaceUuid_ParameterInterfaceBuilderImpl();
  }

  public static class DceRpc_InterfaceUuid_ParameterInterfaceBuilderImpl
      implements DceRpc_InterfaceUuid.DceRpc_InterfaceUuidBuilder {

    public DceRpc_InterfaceUuid_ParameterInterfaceBuilderImpl() {}

    public DceRpc_InterfaceUuid_ParameterInterface build() {
      DceRpc_InterfaceUuid_ParameterInterface dceRpc_InterfaceUuid_ParameterInterface =
          new DceRpc_InterfaceUuid_ParameterInterface();
      return dceRpc_InterfaceUuid_ParameterInterface;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DceRpc_InterfaceUuid_ParameterInterface)) {
      return false;
    }
    DceRpc_InterfaceUuid_ParameterInterface that = (DceRpc_InterfaceUuid_ParameterInterface) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
