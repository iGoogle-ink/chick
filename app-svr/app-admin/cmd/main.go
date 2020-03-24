package main

import "chick/app-svr/app-admin/conf"

func main() {
	err := conf.Parse()
	if err != nil {
		panic(err)
	}
}
