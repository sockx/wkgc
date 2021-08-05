package checkerr

/*
	Check Error
*/
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
