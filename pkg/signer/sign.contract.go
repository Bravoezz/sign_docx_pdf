package signer

type SignOp struct {
	InputPath     string
	OutputPath    string
	SignaturePath string
	SearchText    string
}

type SignStrategy interface {
	SignDocument(op SignOp) error
}
