package controllers

import "log"

func checkErr(err error, ctx string) {
	if err != nil {
		ErrMain = err
		log.Println(red(ctx))
	}
}
