package d3

const (
	// valid statuses
	in  = "init"
	mIn = "M input"
	uIn = "U input"
	lIn = "L input"
	pIn = "( input"
	f1  = "first factor"
	f2  = "second factor"
	cin = "comma input"
	fin = ") input"
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
