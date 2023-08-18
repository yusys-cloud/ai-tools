// Author: yangzq80@gmail.com
// Date: 2023/4/27
package time

import (
	"fmt"
	"testing"
)

func TestGetCurrentDate(t *testing.T) {
	fmt.Println(GetCurrentDate())
	fmt.Println(GetCurrentDateTime())
}
