package d3

const (
	// valid statuses
	in  = Status("init")
	mIn = Status("M input")
	uIn = Status("U input")
	lIn = Status("L input")
	pIn = Status("( input")
	f1  = Status("first factor")
	f2  = Status("second factor")
	cin = Status("comma input")
	fin = Status(") input")

	dIn  = Status("D input")
	oIn  = Status("O input")
	nIn  = Status("N input")
	apIn = Status("' input")
	tIn  = Status("T input")
	pIn2 = Status("( input no par expected")
)

type Status string
type Next func(int32, *Node) Status

var statusMap = map[Status]Next{
	in:  Next(fIn),
	mIn: Next(fMin),
	uIn: Next(fUin),
	lIn: Next(fLin),
	pIn: Next(fPin),
	f1:  Next(fF1),
	f2:  Next(fF2),
	cin: Next(fCin),
	fin: Next(fFin),

	dIn:  Next(fDin),
	oIn:  Next(fOin),
	nIn:  Next(fNin),
	apIn: Next(fApin),
	tIn:  Next(fTin),
	pIn2: Next(fP2in),
}

func fP2in(c int32, node *Node) Status {
	if c == cpar {
		node.op = node.op + ")"
		return fin
	}
	resetNode(node)
	return in
}

func fTin(c int32, node *Node) Status {
	if c == opar {
		node.op = "don't("
		return pIn2
	}
	resetNode(node)
	return in
}

func fNin(c int32, node *Node) Status {
	if c == ap {
		node.op = "don'"
		return apIn
	}
	resetNode(node)
	return in
}

func fDin(c int32, node *Node) Status {
	if c == o {
		node.op = "do"
		return oIn
	}
	resetNode(node)
	return in
}

func fOin(c int32, node *Node) Status {
	if c == n {
		node.op = "don"
		return nIn
	}
	if c == opar {
		node.op = "do("
		return pIn2
	}
	resetNode(node)
	return in
}

func fApin(c int32, node *Node) Status {
	if c == t {
		node.op = "don't"
		return tIn
	}
	resetNode(node)
	return in
}

func fFin(c int32, n *Node) Status {
	panic("unexpected call")
}

func fF2(c int32, n *Node) Status {
	if c == cpar {
		return fin
	}
	if c >= zero && c <= nine {
		n.factor[1] = n.factor[1]*10 + int(c-zero)
		return f2
	}
	resetNode(n)
	return in
}

func fF1(c int32, n *Node) Status {
	if c == comm {
		return cin
	}
	if c >= zero && c <= nine {
		n.factor[0] = n.factor[0]*10 + int(c-zero)
		return f1
	}
	resetNode(n)
	return in
}

func resetNode(n *Node) {
	// we reset here
	n.factor = make([]int, 0)
	n.op = ""
}

func fPin(c int32, n *Node) Status {
	if c >= zero && c <= nine {
		n.factor = append(n.factor, int(c-zero))
		return f1
	}
	resetNode(n)
	return in
}

func fLin(c int32, n *Node) Status {
	if c == opar {
		return pIn
	}
	resetNode(n)
	return in
}

func fUin(c int32, n *Node) Status {
	if c == l {
		n.op = "mul"
		return lIn
	}
	resetNode(n)
	return in
}

func fCin(c int32, n *Node) Status {
	if c >= zero && c <= nine {
		n.factor = append(n.factor, int(c-zero))
		return f2
	}
	resetNode(n)
	return in
}

func fIn(c int32, n *Node) Status {
	if c == m {
		n.op = "m"
		return mIn
	}
	if c == d {
		n.op = "d"
		return dIn
	}
	return in
}

func fMin(c int32, n *Node) Status {
	if c == u {
		n.op = "mu"
		return uIn
	}
	resetNode(n)
	return in
}
