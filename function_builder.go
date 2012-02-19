package main

type FunctionBuilder struct {
	Function *FunctionInfo
	Args     []FunctionBuilderArg
	Rets     []FunctionBuilderArg
	OrigArgs []*ArgInfo
}

type FunctionBuilderArg struct {
	Index    int
	ArgInfo  *ArgInfo
	TypeInfo *TypeInfo
}

func IntSliceContains(haystack []int, needle int) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}

func NewFunctionBuilder(fi *FunctionInfo) *FunctionBuilder {
	fb := new(FunctionBuilder)
	fb.Function = fi


	// prepare an array of ArgInfos
	for i, n := 0, fi.NumArg(); i < n; i++ {
		arg := fi.Arg(i)
		fb.OrigArgs = append(fb.OrigArgs, arg)
	}

	// build skip list
	var skiplist []int
	for _, arg := range fb.OrigArgs {
		ti := arg.Type()

		len := ti.ArrayLength()
		if len != -1 {
			skiplist = append(skiplist, len)
		}

		clo := arg.Closure()
		if clo != -1 {
			skiplist = append(skiplist, clo)
		}

		des := arg.Destroy()
		if des != -1 {
			skiplist = append(skiplist, des)
		}
	}

	// then walk over arguments
	for i, ai := range fb.OrigArgs {
		if IntSliceContains(skiplist, i) {
			continue
		}

		ti := ai.Type()

		switch ai.Direction() {
		case DIRECTION_IN:
			fb.Args = append(fb.Args, FunctionBuilderArg{i, ai, ti})
		case DIRECTION_INOUT:
			fb.Args = append(fb.Args, FunctionBuilderArg{i, ai, ti})
			fb.Rets = append(fb.Rets, FunctionBuilderArg{i, ai, ti})
		case DIRECTION_OUT:
			fb.Rets = append(fb.Rets, FunctionBuilderArg{i, ai, ti})
		}
	}

	// add return value if it exists to 'rets'
	if ret := fi.ReturnType(); ret != nil && ret.Tag() != TYPE_TAG_VOID {
		fb.Rets = append(fb.Rets, FunctionBuilderArg{-1, nil, ret})
	}

	// add GError special argument (if any)
	if fi.Flags()&FUNCTION_THROWS != 0 {
		fb.Rets = append(fb.Rets, FunctionBuilderArg{-2, nil, nil})
	}

	return fb
}

func (fb *FunctionBuilder) HasReturnValue() bool {
	return (len(fb.Rets) > 0 && fb.Rets[len(fb.Rets)-1].Index == -1) ||
		(len(fb.Rets) > 1 && fb.Rets[len(fb.Rets)-2].Index == -1)
}

func (fb *FunctionBuilder) HasClosureArgument() (int, int, ScopeType) {
	for _, arg := range fb.Args {
		userdata := arg.ArgInfo.Closure()
		if userdata == -1 {
			continue
		}

		if arg.TypeInfo.Tag() != TYPE_TAG_INTERFACE {
			continue
		}

		if arg.TypeInfo.Interface().Type() != INFO_TYPE_CALLBACK {
			continue
		}

		destroy := arg.ArgInfo.Destroy()
		scope := arg.ArgInfo.Scope()
		return userdata, destroy, scope
	}
	return -1, -1, SCOPE_TYPE_INVALID
}
