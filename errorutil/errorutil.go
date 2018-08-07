package errorutil

import "log"
//@Deprecated
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}