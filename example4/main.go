package example4

import (
	"fmt"
	"golang.org/x/image/math/f64"
)

// igop run -p /path/to/igop_pluginx_x.so example4

func main() {
	fmt.Printf("--from plugin--\n")
	fmt.Printf("mat4 %v\n", f64.Mat4{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
}
