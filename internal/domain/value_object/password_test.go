package value_object

import (
    "testing"
)


func TestPasswordLength(t *testing.T) {
    pw := "123"
    password := Password{pw}
    is_valid, _ := password.IsValid()
    if !is_valid {
        t.Fatalf("Validação de força de senha não passou")
    }
}
