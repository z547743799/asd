package main

import "fmt"

type Run interface {
	run()
}

type Runner struct {
	name string
}

func (r Runner) run() {
	fmt.Println(r.name, "，开始跑步。。。慢跑")
}

type Pig struct {
	name string
}

func (p Pig) run() {
	fmt.Println(p.name, ",撅着屁股跑。。。")
}

// 定义一个函数：
func testRun(r Run) { // 将接口作为参数，可以传入的任意是实现类的对象。
	r.run()
}
func main() {
	/*
		多态：事物的多种形态
				靠接口实现：定义接口类型，创建任意的实现类对象。

						定义接口类型的对象，但是实际上可以创建该接口的任意实现那类的对象。
								接口类型的对象，不能访问实现类中的字段的。
	*/
	var r1 Run // 对象的声明
	r1 = Runner{"跑步者"}
	r1.run()

	r1 = Pig{"佩奇"}
	r1.run()

	//r1.name //语法错误，接口类型对象，不能访问实现类的字段

	testRun(Pig{"二师兄"})
}
