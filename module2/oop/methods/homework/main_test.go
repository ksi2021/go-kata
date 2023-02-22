package main

import "testing"

type calc struct {
	a, b   float64
	result float64
}

func NewCalc() *calc { // конструктор калькулятора
	return &calc{}
}

func (c *calc) SetA(a float64) *calc {
	c.a = a
	return c
}

func (c *calc) SetB(b float64) *calc {
	c.b = b

	return c
}

func (c *calc) Do(operation func(a, b float64) float64) *calc {
	c.result = operation(c.a, c.b)
	return c
}

func (c calc) Result() float64 {
	return c.result
}

func multiply(a, b float64) float64 { // реализуйте по примеру divide, sum, average
	return a * b
}

func sum(a, b float64) float64 {
	return a + b
}

func average(a, b float64) float64 {
	return (a + b) / 2
}

func divide(a, b float64) float64 {
	return a / b
}

func TestCalc(t *testing.T) {
	calc := NewCalc()

	type args struct {
		a        float64
		b        float64
		function func(a, b float64) float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		//sum
		{
			name: "sum #1 ",
			args: args{a: 5, b: 3, function: sum},
			want: 8,
		},
		{
			name: "sum #2",
			args: args{a: 0, b: 0, function: sum},
			want: 0,
		},
		{
			name: "sum #3",
			args: args{a: 666, b: 333, function: sum},
			want: 999,
		},
		{
			name: "sum #4",
			args: args{a: 5.12, b: 100.67, function: sum},
			want: 105.79,
		},
		{
			name: "sum #5",
			args: args{a: -5, b: 3, function: sum},
			want: -2,
		},
		{
			name: "sum #6",
			args: args{a: 5, b: -3, function: sum},
			want: 2,
		},
		{
			name: "sum #7",
			args: args{a: -5, b: -3, function: sum},
			want: -8,
		},
		{
			name: "sum #8",
			args: args{a: -5677.45, b: 12345.98, function: sum},
			want: 6668.53,
		},
		// divide
		{
			name: "div #1 ",
			args: args{a: 5, b: 3, function: divide},
			want: float64(5) / 3,
		},
		{
			name: "div #2",
			args: args{a: 0, b: 10, function: divide},
			want: float64(0) / 10,
		},
		{
			name: "div #3",
			args: args{a: 666, b: 333, function: divide},
			want: 666 / 333,
		},
		{
			name: "div #4",
			args: args{a: 5.12, b: 100.67, function: divide},
			want: 5.12 / 100.67,
		},
		{
			name: "div #5",
			args: args{a: -5, b: 3, function: divide},
			want: float64(-5) / 3,
		},
		{
			name: "div #6",
			args: args{a: 100, b: 10, function: divide},
			want: 100 / 10,
		},
		{
			name: "div #7",
			args: args{a: -5, b: -3, function: divide},
			want: float64(-5) / -3,
		},
		{
			name: "div #8",
			args: args{a: -5677.45, b: 12345.98, function: divide},
			want: -5677.45 / 12345.98,
		},
		// average
		{
			name: "average #1 ",
			args: args{a: 5, b: 3, function: average},
			want: (float64(5) + 3) / 2,
		},
		{
			name: "average #2",
			args: args{a: 0, b: 0, function: average},
			want: (float64(0) + 0) / 2,
		},
		{
			name: "average #3",
			args: args{a: 666, b: 333, function: average},
			want: (float64(666) + 333) / 2,
		},
		{
			name: "average #4",
			args: args{a: 5.12, b: 100.67, function: average},
			want: (5.12 + 100.67) / 2,
		},
		{
			name: "average #5",
			args: args{a: -5, b: 3, function: average},
			want: (float64(-5) + 3) / 2,
		},
		{
			name: "average #6",
			args: args{a: 5, b: -3, function: average},
			want: (float64(5) + -3) / 2,
		},
		{
			name: "average #7",
			args: args{a: -5, b: -3, function: average},
			want: (float64(-5) + -3) / 2,
		},
		{
			name: "average #8",
			args: args{a: -5677.45, b: 12345.98, function: average},
			want: (-5677.45 + 12345.98) / 2,
		},
		//multiply
		{
			name: "multiply #1 ",
			args: args{a: 5, b: 0, function: multiply},
			want: 5 * 0,
		},
		{
			name: "multiply #2",
			args: args{a: 0, b: 0, function: multiply},
			want: 0 * 0,
		},
		{
			name: "multiply #3",
			args: args{a: 666, b: 333, function: multiply},
			want: 666 * 333,
		},
		{
			name: "multiply #4",
			args: args{a: 5.12, b: 100.67, function: multiply},
			want: 5.12 * 100.67,
		},
		{
			name: "multiply #5",
			args: args{a: -5, b: 3, function: multiply},
			want: -5 * 3,
		},
		{
			name: "multiply #6",
			args: args{a: 5, b: -3, function: multiply},
			want: 5 * -3,
		},
		{
			name: "multiply #7",
			args: args{a: -5, b: -3, function: multiply},
			want: -5 * -3,
		},
		{
			name: "multiply #8",
			args: args{a: -5677.45, b: 12345.98, function: multiply},
			want: -5677.45 * 12345.98,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calc.SetA(tt.args.a).SetB(tt.args.b).Do(tt.args.function).Result(); got != tt.want {
				t.Errorf("Result() = %v, want %v", got, tt.want)
			}
		})
	}
}
