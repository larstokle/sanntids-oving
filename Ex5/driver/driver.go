package driver

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: ${SRCDIR}/simelev.a /usr/lib/x86_64-linux-gnu/libphobos2.a -lpthread -lcomedi -lm
#include "io.h"
#include "channels.h"
*/
import "C"
import(
	"time"
)

const (
	N_FLOORS = 4

	BTN_UP = 0
	BTN_DOWN = 1
	BTN_CMD = 2

	DIR_DOWN = -1
	DIR_STOP = 0
	DIR_UP = 1
	)



func Set_button_light(button int, floor int, value bool) {
	channel := C.int(encode_light(button, floor))
	if (value) {
		C.io_set_bit(channel)
	} else {
		C.io_clear_bit(channel)
	}
}

func encode_light(button int, floor int) int {
	
	channel := C.LIGHT_COMMAND1
	if (button == BTN_CMD) {
		channel = channel - floor
	} else if (button == BTN_UP && floor == 0) {
		channel = C.LIGHT_UP1
	} else if (button == BTN_DOWN && floor == 3) {
		channel = C.LIGHT_DOWN4
	} else {
		channel = C.LIGHT_UP2
		channel = channel - button - 2 * (floor - 1)
	}
	return channel
}

func Init() int {
	return int(C.io_init(ET_comedi))
}

func RunTopFloor() {
	if GetFloorSignal() != 3 {
		C.io_clear_bit(C.MOTORDIR)
		//time.Sleep(time.Second * 1)
		C.io_write_analog(C.MOTOR, 2800)
		for C.io_read_bit(C.SENSOR_FLOOR4) == 0 {
			SetFloorIndicator(GetFloorSignal())
			time.Sleep(time.Millisecond*200)
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
			time.Sleep(time.Millisecond*200)
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
