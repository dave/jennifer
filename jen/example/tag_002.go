type foo struct {
	Bar string `valid:"matches(^(\d*\.)?\d+$)"`
}