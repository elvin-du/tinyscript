package operand

var _ Operand = &Label{}

type Label struct {
	Label string
}

func NewLabel(label string) *Label {
	return &Label{Label: label}
}

func (l *Label) String() string {
	return l.Label
}

func (*Label) Typ() OperandType {
	return TYPE_LABEL
}
