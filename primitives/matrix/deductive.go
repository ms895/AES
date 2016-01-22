package matrix

import (
	"crypto/rand"
)

// DeductiveMatrix is a generalization of IncrementalMatrix that allows the incremental deduction of matrices.
//
// Like incremental matrices, its primary use-case is in cryptanalyses and search algorithms, where we can do some work
// to obtain an (input, output) pair that we believe defines a matrix. We don't want to do more work than necessary and
// we also can't just obtain n (input, output) pairs because of linear dependence, etc.
type DeductiveMatrix struct {
	input, output IncrementalMatrix
}

// NewDeductiveMatrix returns a new n-by-n deductive matrix.
func NewDeductiveMatrix(n int) DeductiveMatrix {
	return DeductiveMatrix{
		input:  NewIncrementalMatrix(n),
		output: NewIncrementalMatrix(n),
	}
}

// Assert represents an assertion that A(in) = out. The function will panic if this is inconsistent with previous
// assertions. It it's not, it returns whether or not the assertion contained new information about A.
func (dm *DeductiveMatrix) Assert(in, out Row) (learned bool) {
	inReduced, inInverse := dm.input.reduce(in)
	outReduced, outInverse := dm.output.reduce(out)

	if inReduced.IsZero() || outReduced.IsZero() {
		real := Matrix{inInverse}.Compose(dm.output.Matrix())[0]

		if !real.Equals(out) {
			panic("Asserted input, output pair is inconsistent with previous assertions!")
		}
		return false
	}

	dm.input.addRows(in, inReduced, inInverse)
	dm.output.addRows(out, outReduced, outInverse)
	return true
}

// FullyDefined returns true if the assertions made give a fully defined matrix.
func (dm *DeductiveMatrix) FullyDefined() bool {
	return dm.input.FullyDefined() && dm.output.FullyDefined()
}

// novel returns a random novel input to a given matrix.
func (dm *DeductiveMatrix) novel(m Matrix) Row {
	null := Matrix(m.NullSpace())
	if len(null) == 0 {
		return nil
	}

	// The space of unspanned elements can be represented by a non-zero element of the nullspace plus any element of the
	// rowspace.
	x := GenerateRandomNonZeroRow(rand.Reader, len(null))
	a := null.Transpose().Mul(x)

	y := GenerateRandomRow(rand.Reader, len(m))
	b := m.Transpose().Mul(y)

	return a.Add(b)
}

// NovelInput returns a random x not in the domain of A.
func (dm *DeductiveMatrix) NovelInput() Row {
	return dm.novel(dm.input.Matrix())
}

// NovelOutput returns a random y not in the span of A.
func (dm *DeductiveMatrix) NovelOutput() Row {
	return dm.novel(dm.output.Matrix())
}

// IsInDomain returns whether or not x is in the known span of A.
func (dm *DeductiveMatrix) IsInDomain(x Row) bool {
	reduced, _ := dm.input.reduce(x)
	return !reduced.IsZero()
}

// IsInSpan returns whether or not y is in the known span of A.
func (dm *DeductiveMatrix) IsInSpan(y Row) bool {
	reduced, _ := dm.output.reduce(y)
	return !reduced.IsZero()
}

// Matrix returns the deduced matrix.
func (dm *DeductiveMatrix) Matrix() Matrix {
	if !dm.FullyDefined() {
		return nil
	}
	return dm.input.Inverse().Compose(dm.output.Matrix()).Transpose()
}

// Inverse returns the deduced matrix's inverse.
func (dm *DeductiveMatrix) Inverse() Matrix {
	if !dm.FullyDefined() {
		return nil
	}
	return dm.output.Inverse().Compose(dm.input.Matrix()).Transpose()
}

// Dup returns a duplicate of dm.
func (dm *DeductiveMatrix) Dup() DeductiveMatrix {
	return DeductiveMatrix{
		input:  dm.input.Dup(),
		output: dm.output.Dup(),
	}
}