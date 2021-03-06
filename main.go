//main comment
package main

import (
	_ "haipi/memory"
	"haipi/session"
	_ "time"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

// entrance
func main() {
	//----------basicTest-----------
	// addTest()
	// typeFuncTest()
	// ArrayTest()
	// getAverageTest()
	// booksTest()
	// sliceTest()
	// rangeTest()
	// mapTest()
	// chanelTest()
	// stringTest()
	// errorTest()
	// constTest()
	// sliceAndArray()
	// mapTest1()
	// switchTest()
	// sumAndProductTest()
	// paramsTest(1, 2, 3, 4)

	// parseJSONTest()
	// parseJSONTest2()
	// simpleJSONTest()
	// generateJSONTest()
	// jsonTagTest()
	// cmdParamsTest()
	// cmdParasByFlagTest()
	// isIP("1.2.2.1")
	// resolveContentFromWeb()
	// regexpFindTest()
	// regexExpandTest()
	// goSvr()
	// irisSvr()
	// mysqlConnTest()
	// parseXML()
	// generateXML()
	// irisSvr()
	configIris()
	// db.MysqlConn()
}
