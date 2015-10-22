package config;

import (
	"strconv"
)

func Atoi(str string)(int){
	result,error := strconv.Atoi(str);
	if error != nil{
		return 0;
	}
	return result;
}

func Itoa(integer int)(string){
	return strconv.Itoa(integer);
}