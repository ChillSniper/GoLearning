package ErrorLearning

import (
	"errors"
)

func GetPositiveSelfAdd(num int) (int, error) {
	if num <= 0 {
		return -1, errors.New("num must be positive")
	}
	return num + 1, nil
}
