package driver

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
#include "channels.h"
*/
import "C"
import (
	"fmt"
	"time"
)

func Init() C.int {
	return C.io_init()
}

func RunFloorUp() {
	if C.io_read_bit(C.SENSOR_FLOOR4) == 0 {
		C.io_set_bit(C.MOTORDIR)
		time.Sleep(time.Second * 1)
		C.io_write_analog(C.MOTOR, 2800)
		for C.io_read_bit(C.SENSOR_FLOOR4) == 0 {
			fmt.Println(C.io_read_bit(C.MOTORDIR))

		}
		C.io_write_analog(C.MOTOR, 0)
	}
}

func RunFloorDown() {
	if C.io_read_bit(C.SENSOR_FLOOR1) == 0 {
		C.io_clear_bit(C.MOTORDIR)
		time.Sleep(time.Second * 1)
		C.io_write_analog(C.MOTOR, 2800)
		for C.io_read_bit(C.SENSOR_FLOOR1) == 0 {
			fmt.Println(C.io_read_bit(C.MOTORDIR))
		}
		C.io_write_analog(C.MOTOR, 0)
	}
}

func RunUpDown() {
	C.io_clear_bit(C.MOTORDIR)
	C.io_write_analog(C.MOTOR, 2800)

	C.io_write_analog(C.MOTOR, 0)

	C.io_clear_bit(C.MOTORDIR)
	C.io_write_analog(C.MOTOR, 2800)
}
