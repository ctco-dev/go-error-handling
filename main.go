package main

import (
	"errors"
	"fmt"
	"log"
	pkg_errors "github.com/pkg/errors"
)

func main() {
	// presentSimpleErrors()
	presentPErrors()
}


func presentPErrors(){
	var e error
	e = funcWithPError()

	log.Println(e)
	log.Println("----------------")
	fmt.Printf("%+v\n", e)

	e=funcWithWrappedError()
	log.Println("----------------")
	fmt.Printf("%+v\n", e)

	e=funcWithChainOfPErrors()
	log.Println("----------------")
	fmt.Printf("%+v\n", e)
}

func presentSimpleErrors(){
	var e error
	e = funcWithError()
	log.Println(e)

	log.Println("----------------")
	e = funcWithRecoverablePanic()
	log.Println(e)

	log.Println("----------------")
	funcWithPanic()
}

func funcWithError() error {
	return errors.New("failure occured")
}

func funcWithPanic() {
	panic("I forgot to switch off an iron")
}

func funcWithRecoverablePanic() (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()

	panic("I lost my car keys")
}

func funcWithPError() error {
	//stack info is stored in error
	return pkg_errors.New("my cool error")
}

func funcWithWrappedError() error {
	e:= funcWithError()
	return pkg_errors.Wrapf(e, "Error processing func")
}

func funcWithChainOfPErrors() error{
	e:= funcWithPError()
	return pkg_errors.Wrapf(e, "Error processing func")
}

