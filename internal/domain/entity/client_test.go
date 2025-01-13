package client

import (
    "testing"
)

func TestCreateClientInstance(t *testing.T) {
    _, error := NewClientEntity(1, "bla", "Vini Jr","1234")
    if error != nil{
        t.Fatalf(error.Error())
    }
}
