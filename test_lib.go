package main

/*
#ifdef __linux__
	#include <Python.h>
#elif __APPLE__
	#include <Python/Python.h>
#endif

PyObject *test_lib_error;
PyObject * Concat(PyObject *, PyObject *);

int Parse_Args(PyObject * args, char * str, int * index) {
    return PyArg_ParseTuple(args, "ls", index, str);
}


static PyMethodDef test_lib_methods[] = {
    {"concat", Concat, METH_VARARGS, "Test string concat"},
    {NULL, NULL, 0, NULL}
};

PyMODINIT_FUNC
inittest_lib(void) {
     PyObject *m;

    m = Py_InitModule("test_lib", test_lib_methods);
    if (m == NULL)
        return;

    test_lib_error = PyErr_NewException("test_lib.error", NULL, NULL);
    Py_INCREF(test_lib_error);
    PyModule_AddObject(m, "error", test_lib_error);
}
*/
import "C"

func setException(text string) {
	C.PyErr_SetString(C.test_lib_error, C.CString(text))
	return
}

func main() {}
