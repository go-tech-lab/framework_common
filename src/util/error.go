package util

//CheckErr  error
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
