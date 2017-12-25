package main

import (
  "fmt"
)


type State func(byte)(byte, int, string)

func CreateState(wOnZero byte, mOnZero int, nOnZero string, wOnOne byte, mOnOne int, nOnOne string) State {
  return func(current byte) (byte, int, string) {
    if (current == 0) { return wOnZero, mOnZero, nOnZero }
    return wOnOne, mOnOne, nOnOne
  }
}

type TuringMachine struct {
  Memory []byte
  Pointer int
  States map[string]State
  NextState string
}

func (tm *TuringMachine) AddState(name string, state State) {
  tm.States[name] = state
}

func NewTuringMachine() *TuringMachine {
  machine := &TuringMachine{}
  machine.Memory = make([]byte, 1)
  machine.Pointer = 0
  machine.States = map[string]State{}
  machine.NextState = "A"

  machine.AddState("A", CreateState(1, 1, "B", 0, -1, "E"))
  machine.AddState("B", CreateState(1, -1, "C", 0, 1, "A"))
  machine.AddState("C", CreateState(1, -1, "D", 0, 1, "C"))
  machine.AddState("D", CreateState(1, -1, "E", 0, -1, "F"))
  machine.AddState("E", CreateState(1, -1, "A", 1, -1, "C"))
  machine.AddState("F", CreateState(1, -1, "E", 1, 1, "A"))

  return machine
}

func (tm *TuringMachine) Step() {
  moveHead := 0
  tm.Memory[tm.Pointer], moveHead, tm.NextState = tm.States[tm.NextState](tm.Memory[tm.Pointer])

  tm.Pointer += moveHead

  if (tm.Pointer < 0) {
    tm.Memory = append([]byte{0}, tm.Memory...)
    tm.Pointer = 0
  }

  if (tm.Pointer >= len(tm.Memory)) {
    tm.Memory = append(tm.Memory, 0)
  }
}

func (tm *TuringMachine) Checksum() int {
  counter := 0
  for _, c := range tm.Memory {
    if c == 1 {
      counter++
    }
  }

  return counter
}

func main() {
  machine := NewTuringMachine()
  for i := 0; i < 12208951; i++ {
    machine.Step()
  }

  fmt.Printf("Checksum: %d\n", machine.Checksum())
}
