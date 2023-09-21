package validator

import "strconv"

func IsNumber() {

}

func IsDigital(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
