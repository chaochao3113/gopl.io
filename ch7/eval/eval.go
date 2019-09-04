package eval

// Expr: 算术表达式
type Expr interface {
	// Eval 返回表达式在env上下文下的值
	Eval(env Env) float64
}

// Var 表示一个变量, 比如 x
type Var string 

// literal 是一个数字变量, 比如3.141
type literal float64

// unary 表示一元操作符表达式, 比如 -x
type unary struct {
	op rune  // '+', '-'中的一个
	x Expr
}

// binary 表示二元操作符表达式, 比如 x+y
type binary struct {
	op rune // '+', '-', '*', '/'中的一个
	x, y Expr
} 

// call 表示函数调用表达式, 比如 sin(x)
type call struct {
	fn string // one of "pow", "sin", "sqrt"中的一个
	args []Expr
}

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}