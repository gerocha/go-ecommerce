package value_object

type Password struct {
    Value string
}

func (password Password) IsValid() bool {
    return len(password.Value) > 2
}
