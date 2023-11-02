package member

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	str := "1,2,3,4,5    aad b  ,sdaf,  ,6"
	str = strings.ReplaceAll(str, " ", "")
	arr := strings.Split(str, ",")
	intArr := make([]int, 0)
	for _, s := range arr {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		intArr = append(intArr, i)
	}
	expectArr := []int{1, 2, 3, 4, 6}
	for i, e := range expectArr {
		assert.Equal(t, e, intArr[i])
	}
	fmt.Println(arrayToString(intArr, ","))
}

func Test_arrayToString(t *testing.T) {

}
