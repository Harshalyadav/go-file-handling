package main

import ( "testing"
"example.com/learn/corePackages/testing/process"

)

func TestCheckEven(t *testing.T){
	i:=10
	expected :="YES"
	res := process.CheckEven(i)
	if expected != res{
		t.Errorf("Expected:%v,got:%v", expected,res)
	}
}