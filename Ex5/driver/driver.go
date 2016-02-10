package driver

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
#include "channels.h"
*/
import "C"

func Init() C.int {
	return C.io_init()
}

func RunTopFloor() {
	if GetFloorSignal() != 3 {
		C.io_clear_bit(C.MOTORDIR)
		//time.Sleep(time.Second * 1)
		C.io_write_analog(C.MOTOR, 2800)
		for C.io_read_bit(C.SENSOR_FLOOR4) == 0 {
			SetFloorIndicator(GetFloorSignal())
		}
		SetFloorIndicator(GetFloorSignal())
		C.io_write_analog(C.MOTOR, 0)
	}
}

func RunBottomFloor() {
	if GetFloorSignal() != 0 {
		C.io_set_bit(C.MOTORDIR)
		//time.Sleep(time.Second * 1)
		C.io_write_analog(C.MOTOR, 2800)
		for C.io_read_bit(C.SENSOR_FLOOR1) == 0 {
			SetFloorIndicator(GetFloorSignal())
		}
		SetFloorIndicator(GetFloorSignal())
		C.io_write_analog(C.MOTOR, 0)
	}
}

func RunUp() {
	C.io_clear_bit(C.MOTORDIR)
	//time.Sleep(time.Second * 1)
	C.io_write_analog(C.MOTOR, 2800)
}

func RunDown() {
	C.io_set_bit(C.MOTORDIR)
	//time.Sleep(time.Second * 1)
	C.io_write_analog(C.MOTOR, 2800)
}

func RunStop() {
	C.io_write_analog(C.MOTOR, 0)
}

func SetFloorIndicator(floor int) bool {
	if floor < 0 || floor > 3 {
		return false
	}

	if (floor & 0x02) != 0 {
		C.io_set_bit(C.LIGHT_FLOOR_IND1)
	} else {
		C.io_clear_bit(C.LIGHT_FLOOR_IND1)
	}

	if (floor & 0x01) != 0 {
		C.io_set_bit(C.LIGHT_FLOOR_IND2)
	} else {
		C.io_clear_bit(C.LIGHT_FLOOR_IND2)
	}

	return true
}

func GetFloorSignal() int {
	if C.io_read_bit(C.SENSOR_FLOOR1) != 0 {
		return 0
	} else if C.io_read_bit(C.SENSOR_FLOOR2) != 0 {
		return 1
	} else if C.io_read_bit(C.SENSOR_FLOOR3) != 0 {
		return 2
	} else if C.io_read_bit(C.SENSOR_FLOOR4) != 0 {
		return 3
	} else {
		return -1
	}
}
