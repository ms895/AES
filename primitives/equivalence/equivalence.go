package equivalence

import (
	"github.com/OpenWhiteBox/AES/primitives/encoding"
	"github.com/OpenWhiteBox/AES/primitives/matrix"
)

// A (linear) equivalence is a pair of matrices, (A, B) such that f(A(x)) = B(g(x)) for all x.
type Equivalence struct {
	A, B matrix.Matrix
}

// LinearEquivalence finds linear equivalences between f and g. cap is the maximum number of equivalences to return.
func LinearEquivalence(f, g encoding.Byte, cap int) []Equivalence {
	return search(f, g, NewMatrix(), NewMatrix(), cap)
}

// search contains the search logic of our dynamic programming algorithm.
// f and g are the functions we're finding equivalences for. A and B are the parasites. As we find equivalences, we send
// them back on the channel res.
func search(f, g encoding.Byte, A, B *Matrix, cap int) (res []Equivalence) {
	x := A.NovelInput()

	// 1. Take a guess for A(x).
	// 2. Check if its possible for any matrix B to satisfy an equivalence relation with what we know about A.
	for _, guess := range Universe {
		if A.IsInSpan(guess) { // Our guess for A(x) can't result in A being singular.
			continue
		}

		AT, BT := A.Dup(), B.Dup()
		AT.Assert(x, guess)

		consistent := learn(f, g, AT, BT)

		// Our guess for A(x) ...
		if !consistent { // ... isn't consistent with any equivalence relation.
			continue
		} else if AT.FullyDefined() { // ... uniquely specified an equivalence relation.
			res = append(res, Equivalence{
				A: AT.Matrix(),
				B: BT.Matrix(),
			})
		} else { // ... has neither led to a contradiction nor a full definition.
			res = append(res, search(f, g, AT, BT, cap-len(res))...)
		}

		if len(res) >= cap {
			return
		}
	}

	return
}

// learn finds the logical implications of what we've specified about A and B.
// f and g are the functions we're finding equivalences for. A and B are the parasites.
// learn returns whether or not A and B are consistent with any possible equivalence. A and B are mutated to contain the
// new information.
func learn(f, g encoding.Byte, A, B *Matrix) (consistent bool) {
	defer func() {
		if r := recover(); r != nil {
			consistent = false
		}
	}()

	consistent = true
	learning := true

	// We have to loop because of the "Needlework Effect." A gives info on B, which may in turn give more info on A, ....
	for learning {
		learning = false

		// For every vector x in the domain of A, we have that f(Ax) = B * g(x) -- the span of A gives us input-output
		// behavior that B *has* to satsify, and for every A(x) = y that we guess correctly, we get twice as many
		// input-output pairs. This means that we end up with a complete definition of B (allowing us to derive A) much
		// faster than if we were to try to guess all of A first (and then derive B).
		for x, y := range A.Space {
			xT := matrix.Row{g.Encode(x)}
			yT := matrix.Row{f.Encode(y)}

			learning = learning || B.Assert(xT, yT)
		}

		// Lets say that we know two input-output pairs of B: f(Ax) = B * g(x) and f(Ay) = B * g(y). Because B is linear,
		// we get the following statement about B for free: f(Ax) + f(Ay) = B * (g(x) + g(y)). But this means there has to
		// be some z such that f(Ax) + f(Ay) = f(Az) and g(x) + g(y) = g(z). Find z and Az by solving the two previous
		// equations.
		for x, y := range B.Space {
			z := matrix.Row{g.Decode(x)}
			Az := matrix.Row{f.Decode(y)}

			learning = learning || A.Assert(z, Az)
		}
	}

	return
}