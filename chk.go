//

// useful/assert like - but really shouldn't be in production code

package main

import "fmt"

func chk(err error) {
	if err != nil {
		fmt.Printf("FATAL %T %#v\n", err, err)
		panic(err)
	}
}
