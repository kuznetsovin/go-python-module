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
	"errors"
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
	var err error = nil

	if result, err = sum(C.GoString(string), int(c_idx)); err != nil {
		setException(err.Error())
	}

	go func() {
		for range time.NewTicker(time.Second).C {
			log.Println("Вызов из горутины")
		}
	}()

	return C.PyString_FromString(C.CString(result))
}

func sum(s string, i int) (string, error) {
	result := ""
	var err error = nil

	if i != 0 {
		result = s + strconv.Itoa(i)
	} else {
		err = errors.New("Индекс не может быть 0")
	}

	return result, err
}
