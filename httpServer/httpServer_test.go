package httpServer

import "testing"

func Test_StartServer_NegativePort_ErrorReturned(test *testing.T) {
	err := StartServer(-1)
	if err == nil {
		test.Errorf("Error not returned")
	}
}

func Test_StartServer_LargePort_ErrorReturned(test *testing.T) {
	err := StartServer(65536)
	if err == nil {
		test.Errorf("Error not returned")
	}
}
