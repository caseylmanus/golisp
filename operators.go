package main

func AddOperator(args ...Symbol) (interface{}, error) {
	var total Number
	for _, i := range args {
		total = total + i.Number()
	}
	return total, nil
}

func SubtractOperator(args ...Symbol) (interface{}, error) {
	var current Number
	for i, n := range args {
		if i == 0 {
			current = n.Number()

		} else {
			current = current - n.Number()
		}
	}
	return current, nil
}

func MultiplyOperator(args ...Symbol) (interface{}, error) {
	var total Number = 1
	for _, i := range args {
		total = total * i.Number()
	}
	return total, nil
}
