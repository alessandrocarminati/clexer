package main

import "C"

var ProcessorSTR []string;
var FuncName []string;


//export GoAddFuncName
func GoAddFuncName(dataPtr *C.char) {
        FuncName=append(FuncName,  C.GoString(dataPtr));
}

//export GoAddPreProcessorSTR
func GoAddPreProcessorSTR(dataPtr *C.char) {
	ProcessorSTR=append(ProcessorSTR,  C.GoString(dataPtr));
}


