package client

import (
    "testing"
)

func TestCreateClientInstance(t *testing.T) {
    _, error := NewClient(1, "bla", "Vini Jr","12")
    if error != nil{
        t.Fatalf(error.Error())
    }
}
