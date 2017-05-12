package main

/*
#cgo pkg-config: python-2.7 --cflags --libs
#ifdef __linux__
	#include <Python.h>
#elif __APPLE__
	#include <Python/Python.h>
#endif
int Parse_Args(PyObject * args, char * str, int * index);
*/
import "C"
import (
	"log"
	"strconv"
	"time"
)

//export Concat
func Concat(self *C.PyObject, args *C.PyObject) *C.PyObject {
	var string = new(C.char)
	var c_idx C.int

	if C.Parse_Args(args, string, &c_idx) == 0 {
		setException("Ошибка парсинга аргументов")
		return nil
	}

	result := ""
	idx := int(c_idx)
	if idx != 0 {
		result = C.GoString(string) + strconv.Itoa(idx)
	} else {
		setException("Индекс не может быть 0")
	}

	go func() {
		for range time.NewTicker(time.Second).C {
			log.Println("Вызов из горутины")
		}
	}()

	return C.PyString_FromString(C.CString(result))
}
