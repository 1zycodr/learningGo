package main

import (
	"fmt"
	"os"
)

type Custom struct {
	data string
}

func (c *Custom) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(c.data)

	return err
}

func (c *Custom) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	c.data = string(data)
	return nil
}

func main() {
	data := Custom{}
	err := data.Load("test")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data.data)
}
