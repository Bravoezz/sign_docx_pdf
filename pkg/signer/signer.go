package signer

import "errors"

type Signer struct {
	strategy SignStrategy
	signOp   SignOp
}

func NewSigner(op *SignOp) *Signer {
	return &Signer{
		strategy: nil,
		signOp:   *op,
	}
}

func (sg *Signer) SetSignStrategy(strategy SignStrategy) {
	sg.strategy = strategy
}

func (sg *Signer) Sign() error {
	if sg.strategy == nil {
		return errors.New("Strategy not set")
	}
	return sg.strategy.SignDocument(sg.signOp)
}
