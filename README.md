
#Doubt why  is "temp" allowed but when the same thing is put in the nested structure "ns" it cause error

ns := [][]interface{}{

		{"John", 30},
		{"Alice", 27},
		{"Bob", true},
		{3.14, "pi"},
		{
			[]interface{}{2}, []interface{}{3}, []interface{}{4}, []interface{}{5, 6},
		},
	}
	temp := [][]interface{}{
		{2}, {3}, {4}, {5, 6},
	}
