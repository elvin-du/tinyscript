package operand

var _ Operand = &Label{}

type Label struct {
	Label string
	*Offset
}

func NewLabel(label string) *Label {
	return &Label{Label: label, Offset: NewOffset(0)}
}

func (l *Label) String() string {
	return l.Label
}

func (*Label) Typ() OperandType {
	return TYPE_LABEL
}
