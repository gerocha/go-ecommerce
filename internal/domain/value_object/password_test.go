package value_object

import (
    "testing"
)


func TestPasswordLength(t *testing.T) {
    pw := "123"
    password := Password{pw}
    if password.IsValid() {
        t.Fatalf("Validação de força de senha não passou")
    }
}
