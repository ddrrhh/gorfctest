// Package gorfc provides SAP NetWeawer RFC SDK client bindings for GO
package gorfc

import (
	"fmt"
	"log"
)

//################################################################################
//# ERRORS                                                             	 	     #
//################################################################################

// RfcError is returned by SAP NWRFC SDK
type RfcError struct {
	Description string
	ErrorInfo   rfcSDKError
}

func (err RfcError) Error() string {
	return fmt.Sprintf("NWRFC SDK error: %s | %s", err.Description, err.ErrorInfo)
}

// GoRfcError is returned by gorfc
type GoRfcError struct {
	Description string
	GoError     error
}

func (err GoRfcError) Error() string {
	if err.GoError != nil {
		return fmt.Sprintf("GORFC error: %s | %s", err.Description, err.GoError.Error())
	}
	return fmt.Sprintf("GORFC error: %s", err.Description)
}

func goRfcError(description string, goerror error) *GoRfcError {
	return &GoRfcError{description, goerror}
}

// ################################################################################
// # FILL FUNCTIONS                                                            	 #
// ################################################################################
// # Fill functions take Go values and return C values
type rfcSDKError struct {
	Message       string
	Code          string
	Key           string
	AbapMsgClass  string
	AbapMsgType   string
	AbapMsgNumber string
	AbapMsgV1     string
	AbapMsgV2     string
	AbapMsgV3     string
	AbapMsgV4     string
}

func (err rfcSDKError) String() string {
	return fmt.Sprintf("rfcSDKError[%v, %v, %v, %v, %v, %v, %v, %v, %v, %v]", err.Message, err.Code, err.Key, err.AbapMsgClass, err.AbapMsgType, err.AbapMsgNumber, err.AbapMsgV1, err.AbapMsgV2, err.AbapMsgV3, err.AbapMsgV4)
}

// ConnectionAttributes returned by getConnectionInfo() method
type ConnectionAttributes map[string]string

// FieldDescription type
type FieldDescription struct {
	Name      string
	FieldType string
	NucLength uint
	NucOffset uint
	UcLength  uint
	UcOffset  uint
	Decimals  uint
	TypeDesc  TypeDescription
}

// TypeDescription type
type TypeDescription struct {
	Name      string
	NucLength uint
	UcLength  uint
	Fields    []FieldDescription
}

// ParameterDescription type
type ParameterDescription struct {
	Name          string
	ParameterType string
	Direction     string
	NucLength     uint
	UcLength      uint
	Decimals      uint
	DefaultValue  string
	ParameterText string
	Optional      bool
	TypeDesc      TypeDescription
	// ExtendedDescription interface{} //This field can be used by the application programmer (i.e. you) to store arbitrary extra information.
}

func (paramDesc ParameterDescription) String() string {
	return fmt.Sprintf("paramDesc(name= %v, paramType= %v, dir= %v, nucLen= %v, ucLen= %v, dec= %v, defValue= %v, paramText= %v, optional= %v, typeDesc= %v)",
		paramDesc.Name, paramDesc.ParameterType, paramDesc.Direction, paramDesc.NucLength, paramDesc.UcLength, paramDesc.Decimals, paramDesc.DefaultValue, paramDesc.ParameterText, paramDesc.Optional, paramDesc.TypeDesc)
}

// FunctionDescription type
type FunctionDescription struct {
	Name       string
	Parameters []ParameterDescription
}

func (funcDesc FunctionDescription) String() (result string) {
	result = fmt.Sprintf("FunctionDescription:\n Name: %v\n Parameters:\n", funcDesc.Name)
	for i := 0; i < len(funcDesc.Parameters); i++ {
		result += fmt.Sprintf("    %v\n", funcDesc.Parameters[i])
	}
	return
}

//################################################################################
//# NW RFC LIB FUNCTIONALITY                                                     #
//################################################################################

// GetNWRFCLibVersion returnd the major version, minor version and patchlevel of the SAP NetWeaver RFC library used.
func GetNWRFCLibVersion() (major, minor, patchlevel uint) {
	/*
		var cmaj, cmin, cpatch C.uint
		C.RfcGetVersion(&cmaj, &cmin, &cpatch)
		major = uint(cmaj)
		minor = uint(cmin)
		patchlevel = uint(cpatch)
	*/
	major = 0
	minor = 0
	patchlevel = 0
	return
}

//################################################################################
//# CONNECTION                                                                   #
//################################################################################

// Connection Parameters
type ConnectionParameters map[string]string

// Client Connection
type Connection struct {
	//handle             C.RFC_CONNECTION_HANDLE
	rstrip             bool
	returnImportParams bool
	alive              bool
	//paramCount         C.uint
	//connParams         []C.RFC_CONNECTION_PARAMETER
	connectionParams ConnectionParameters
}

/*
func connectionFinalizer(conn *Connection) {
	for _, connParam := range conn.connParams {
		C.free(unsafe.Pointer(connParam.name))
		C.free(unsafe.Pointer(connParam.value))
	}
}
*/

// ConnectionFromParams creates a new connection with the given connection parameters and tries to open it.
// Returns the connection if successfull, otherwise nil.
func ConnectionFromParams(connectionParams ConnectionParameters) (conn *Connection, err error) {
	log.Println("Calling: ConnectionFromParams with:", connectionParams)
	err = nil
	return
	/*
	   conn = new(Connection)

	   conn.handle = nil
	   conn.rstrip = true
	   conn.returnImportParams = false
	   conn.alive = false

	   runtime.SetFinalizer(conn, connectionFinalizer)
	   conn.paramCount = C.uint(len(connectionParams))
	   conn.connectionParams = connectionParams
	   conn.connParams = make([]C.RFC_CONNECTION_PARAMETER, conn.paramCount, conn.paramCount)
	   i := 0

	   	for name, value := range conn.connectionParams {
	   		conn.connParams[i].name, err = fillString(name)
	   		conn.connParams[i].value, err = fillString(value)
	   		i++
	   	}

	   	if err != nil {
	   		return nil, err
	   	}

	   err = conn.Open()

	   	if err != nil {
	   		return nil, err
	   	}

	   return
	*/
}

// ConnectionFromDest creates a new connection with just the dest system id.
func ConnectionFromDest(dest string) (conn *Connection, err error) {
	log.Println("Calling: ConnectionFromDest with:", dest)
	err = nil
	return
	/*
		return ConnectionFromParams(ConnectionParameters{"dest": dest})
	*/
}

// RStrip sets rstrip of the given connection to the passed parameter and returns the connection
// right strips strings returned from RFC call (default is true)
func (conn *Connection) RStrip(rstrip bool) *Connection {
	log.Println("Calling: RStrip with:", rstrip)
	return nil
	/*
		conn.rstrip = rstrip
		return conn
	*/
}

// ReturnImportParams sets returnImportParams of the given connection to the passed parameter and returns the connection
func (conn *Connection) ReturnImportParams(returnImportParams bool) *Connection {
	log.Println("Calling: ReturnImportParams with:", returnImportParams)
	return nil
	/*
		conn.returnImportParams = returnImportParams
		return conn
	*/
}

// Alive returns true if the connection is open else returns false.
func (conn *Connection) Alive() bool {
	log.Println("Calling: Alive")
	return true
	/*
	   return conn.alive
	*/
}

// Close closes the connection and sets alive to false.
func (conn *Connection) Close() (err error) {
	log.Println("Calling: Close")
	err = nil
	return
	/*
	   var errorInfo C.RFC_ERROR_INFO

	   	if conn.alive {
	   		conn.alive = false
	   		rc := C.RfcCloseConnection(conn.handle, &errorInfo)
	   		if rc != C.RFC_OK {
	   			return rfcError(errorInfo, "Connection could not be closed")
	   		}
	   	}

	   return
	*/
}

// Open opens the connection and sets alive to true.
func (conn *Connection) Open() (err error) {
	log.Println("Calling: Open")
	err = nil
	return
	/*
		var errorInfo C.RFC_ERROR_INFO
		conn.handle = C.RfcOpenConnection(&conn.connParams[0], conn.paramCount, &errorInfo)
		if errorInfo.code != C.RFC_OK {
			return rfcError(errorInfo, "Connection could not be opened")
		}
		conn.alive = true
		return
	*/
}

// Reopen closes and opens the connection.
func (conn *Connection) Reopen() (err error) {
	log.Println("Calling: Reopen")
	err = nil
	return
	/*
		err = conn.Close()
		if err != nil {
			return
		}
		err = conn.Open()
		return
	*/
}

// Ping pings the server which the client is connected to and does nothing with the error if one occurs.
func (conn *Connection) Ping() (err error) {
	log.Println("Calling: Ping")
	err = nil
	return
	/*
		var errorInfo C.RFC_ERROR_INFO
		if !conn.alive {
			err = conn.Open()
			if err != nil {
				return
			}
		}
		rc := C.RfcPing(conn.handle, &errorInfo)
		if rc != C.RFC_OK {
			return rfcError(errorInfo, "Server could not be pinged")
		}
		return
	*/
}

// GetConnectionAttributes returns the wrapped connection attributes of the connection.
func (conn *Connection) GetConnectionAttributes() (connAttr ConnectionAttributes, err error) {
	log.Println("Called: GetConnectionAttributes")
	err = nil
	return
	/*
		var errorInfo C.RFC_ERROR_INFO
		var attributes C.RFC_ATTRIBUTES

		rc := C.RfcGetConnectionAttributes(conn.handle, &attributes, &errorInfo)
		if rc != C.RFC_OK || errorInfo.code != C.RFC_OK {
			return nil, rfcError(errorInfo, "Could not get connection attributes")
		}
		return wrapConnectionAttributes(attributes, conn.rstrip)
	*/
}

// GetFunctionDescription returns the wrapped function description of the given function.
func (conn *Connection) GetFunctionDescription(goFuncName string) (goFuncDesc FunctionDescription, err error) {
	log.Println("Calling: Getfunctiondescription for", goFuncName)
	err = nil
	return
	/*
		var errorInfo C.RFC_ERROR_INFO

		funcName, err := fillString(goFuncName)
		defer C.free(unsafe.Pointer(funcName))
		if err != nil {
			return
		}

		if !conn.alive {
			err = conn.Open()
			if err != nil {
				return
			}
		}

		funcDesc := C.RfcGetFunctionDesc(conn.handle, funcName, &errorInfo)
		if funcDesc == nil {
			return goFuncDesc, rfcError(errorInfo, "Could not get function description for \"%v\"", goFuncName)
		}

		return wrapFunctionDescription(funcDesc)
	*/
}

// Call calls the given function with the given parameters and wraps the results returned.
func (conn *Connection) Call(goFuncName string, params interface{}) (result map[string]interface{}, err error) {
	log.Println("Calling: Call", goFuncName, "with params:", params)
	err = nil
	return
	/*
		if !conn.alive {
			return nil, goRfcError("Call() method requires an open connection", nil)
		}

		var errorInfo C.RFC_ERROR_INFO

		funcName, err := fillString(goFuncName)
		defer C.free(unsafe.Pointer(funcName))
		if err != nil {
			return
		}

		if !conn.alive {
			err = conn.Open()
			if err != nil {
				return
			}
		}

		funcDesc := C.RfcGetFunctionDesc(conn.handle, funcName, &errorInfo)
		if funcDesc == nil {
			return result, rfcError(errorInfo, "Could not get function description for \"%v\"", funcName)
		}

		funcCont := C.RfcCreateFunction(funcDesc, &errorInfo)
		if funcCont == nil {
			return result, rfcError(errorInfo, "Could not create function")
		}

		defer C.RfcDestroyFunction(funcCont, nil)

		paramsValue := reflect.ValueOf(params)
		if paramsValue.Type().Kind() == reflect.Map {
			keys := paramsValue.MapKeys()
			if len(keys) > 0 {
				if keys[0].Kind() == reflect.String {
					for _, nameValue := range keys {
						fieldName := nameValue.String()
						fieldValue := paramsValue.MapIndex(nameValue).Interface()

						err = fillFunctionParameter(funcDesc, funcCont, fieldName, fieldValue)
						if err != nil {
							return
						}
					}
				} else {
					return result, rfcError(errorInfo, "Could not fill parameters passed as map with non-string keys")
				}
			}
		} else if paramsValue.Type().Kind() == reflect.Struct {
			for i := 0; i < paramsValue.NumField(); i++ {
				fieldName := paramsValue.Type().Field(i).Name
				fieldValue := paramsValue.Field(i).Interface()

				err = fillFunctionParameter(funcDesc, funcCont, fieldName, fieldValue)
				if err != nil {
					return
				}
			}
		} else {
			return result, rfcError(errorInfo, "Parameters can only be passed as types map[string]interface{} or go-structures")
		}

		rc := C.RfcInvoke(conn.handle, funcCont, &errorInfo)

		if rc != C.RFC_OK {
			return result, rfcError(errorInfo, "Could not invoke function \"%v\"", goFuncName)
		}

		if conn.returnImportParams {
			return wrapResult(funcDesc, funcCont, (C.RFC_DIRECTION)(0), conn.rstrip)
		}
		return wrapResult(funcDesc, funcCont, C.RFC_IMPORT, conn.rstrip)
	*/
}
