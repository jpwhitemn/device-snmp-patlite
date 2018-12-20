// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 Dell Technologies
//
// SPDX-License-Identifier: Apache-2.0

package driver

const (
	PATLITE_OID = "1.3.6.1.4.1.20440.4.1.5.1.2.1"

	SET_OID   = ".2"
	TIMER_OID = ".3"
	GET_OID   = ".4"

	RED_OID   = ".1"
	AMBER_OID = ".2"
	GREEN_OID = ".3"
	//blue and white are not available on all patlites
	BLUE_OID   = ".4"
	WHITE_OID  = ".5"
	BUZZER_OID = ".6"

	// PATLITE light values
	//LIGHT_OFF   = 1
	//LIGHT_ON    = 2
	//LIGHT_BLINK = 3
	//LIGHT_FLASH = 5

	//PATLITE buzzer values
	//BUZZ_ON       = 5
	//BUZZ_OFF      = 1
	//BUZZ_PATTERN1 = 2
	//BUZZ_PATTERN2 = 3
	//BUZZ_PATTERN3 = 4
)
