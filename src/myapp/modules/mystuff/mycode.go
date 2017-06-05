package mystuff

import (
	"fmt"

	gm "github.com/deprepo"
)

func Stuff() {
	fmt.Println("from inside modules/mystuff")
	gm.Stuff("from mystuff")
}
