package Model

type Problem struct {
	detail string
}

func (p *Problem) Error() string {
	return p.detail
}

func HaveError(description string) error {
	return &Problem{detail: description}
}
