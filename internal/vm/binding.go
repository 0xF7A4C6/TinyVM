package vm

import (
	"fmt"
	"io"
	"net/http"
)

func _BIND_ADD(v *Vm) int {
	if v.stack.len() != 2 {
		return 0
	}

	a := v.stack.pop()
	b := v.stack.pop()

	return int(a) + int(b)
}

func _BIND_HTTP_REQ(v *Vm) int {
	// load ptr of url
	url_ptr := v.stack.pop()
	url_len := v.stack.pop()
	body_ptr := v.stack.pop()

	// load url from ram
	url := v.getString(int(url_ptr), int(url_len))
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	
	// set response body in ram
	v.setRam(body_ptr, body)

	return 0
}
