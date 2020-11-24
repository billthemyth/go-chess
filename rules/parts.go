package rules


var King  = Part{
	IsDead: false,
	justOneStep : true,
	steps: [][][]int{
		{
			{1},{0},
		},
		{
			{0},{1},
		},
		{
			{-1},{0},
		},
		{
			{0},{-1},
		},

	},
}

var Queen  = Part{
	IsDead: false,
	justOneStep: false,
	steps: [][][]int{
		{
			{1},{0},
		},
		{
			{0},{1},
		},
		{
			{-1},{0},
		},
		{
			{0},{-1},
		},
		{
			{1},{1},
		},
		{
			{-1},{1},
		},
		{
			{1},{-1},
		},
		{
			{-1},{-1},
		},

	},
}

var Towers  = Part{
	IsDead: false,
	justOneStep: false,
	steps: [][][]int{
		{
			{1},{0},
		},
		{
			{0},{1},
		},
		{
			{-1},{0},
		},
		{
			{0},{-1},
		},
	},

}

var Bishop  = Part{
	IsDead: false,
	justOneStep: false,
	steps: [][][]int{
		{
			{1},{1},
		},
		{
			{-1},{1},
		},
		{
			{1},{-1},
		},
		{
			{-1},{-1},
		},
	},
}


var Horse  = Part{
	IsDead: false,
	justOneStep: false,
	steps: [][][]int{
		{
			{3},{1},
		},
		{
			{-3},{1},
		},
		{
			{3},{-1},
		},
		{
			{-3},{-1},
		},
	},
}


var Pawn = Part{
	IsDead: false,
	justOneStep: false,
	steps: [][][]int{
		{
			{1},{0},
		},
		{
			{-1},{0},
		},
	},
}


type Part struct {

	IsDead bool
	justOneStep bool
	steps [][][]int

}