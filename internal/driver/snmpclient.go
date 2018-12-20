// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	g "github.com/soniah/gosnmp"
)

// SNMPClient represents the SNMP device is used for getting and setting SNMP device via OID
type SNMPClient struct {
	ip_Addr  string
	base_OID string
}

func NewSNMPClient(addr string) SNMPClient {
	return SNMPClient{
		ip_Addr:  addr,
		base_OID: PATLITE_OID,
	}
}

type DeviceCommand struct {
	operation string
	value     uint8
}

func NewGetDeviceCommand(op string) DeviceCommand {
	return DeviceCommand{
		operation: op,
		// for Gets, value is not used
		value: 0,
	}
}

type snmpCommand struct {
	op_OID    string
	state_OID string
	value     uint8
}

func (c *SNMPClient) GetValues(commands []DeviceCommand) ([]uint8, error) {

	for key, value := range commands {
		
	}
}

func (c *SNMPClient) GetValue(command DeviceCommand) (uint8, error) {

	cmd := new(snmpCommand)
	cmd.op_OID = GET_OID

	switch command.operation {
	case "RED":
		cmd.state_OID = RED_OID
	case "AMBER":
		cmd.state_OID = AMBER_OID
	case "GREEN":
		cmd.state_OID = GREEN_OID
	case "BLUE":
		cmd.state_OID = BLUE_OID
	case "WHITE":
		cmd.state_OID = WHITE_OID
	case "BUZZER":
		cmd.state_OID = BUZZER_OID
	}

	g.Default.Target = c.ip_Addr
	err := g.Default.Connect()
	if err != nil {
		return 0, err
	}
	defer g.Default.Conn.Close()

	oids := []string{c.createOID(cmd)}
	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		return 0, err2
	}

	var y uint64
	// TODO return array or results
	for _, variable := range result.Variables {
		y = g.ToBigInt(variable.Value).Uint64()
	}

	return uint8(y), nil
}

func (c *SNMPClient) createOID(cmd *snmpCommand) string {
	return c.base_OID + cmd.op_OID + cmd.state_OID
}
