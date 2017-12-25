package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type State func(byte)(byte, int, string)

type StateDescriptor struct {
  Write byte
  Move int
  Next string
}

func NewState(zero *StateDescriptor, one *StateDescriptor) State {
  return func(current byte) (byte, int, string) {
    if (current == 0) { return zero.Write, zero.Move, zero.Next }
    return one.Write, one.Move, one.Next
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

func Readln(r *bufio.Reader) (string, error) {
  var (
    isPrefix bool  = true
    err      error = nil
    line, ln []byte
  )
  for isPrefix && err == nil {
    line, isPrefix, err = r.ReadLine()
    ln = append(ln, line...)
  }
  return string(ln), err
}

func LoadInstructions(machine *TuringMachine) (stepCount int) {
  args := os.Args[1:]

  if len(args) < 1 {
    fmt.Printf("First argument must be a file.")
    os.Exit(1)
  }

  file, err := os.Open(args[0])
  if err != nil {
    fmt.Printf("%s\n", err)
    os.Exit(1)
  }

  defer file.Close()

  var currentBranch int
  var activeStateName string
  states := make([]*StateDescriptor, 2, 2)
  states[0] = &StateDescriptor{}
  states[1] = &StateDescriptor{}

  r := bufio.NewReader(file)
  for s, e := Readln(r); e == nil; s, e = Readln(r) {
    if strings.HasPrefix(s, "Begin") {
      fmt.Sscanf(s, "Begin in state %1s.", &machine.NextState)
      continue
    }

    if strings.HasPrefix(s, "Perform") {
      fmt.Sscanf(s, "Perform a diagnostic checksum after %d steps.", &stepCount)
      continue
    }

    if s == "" && activeStateName != "" {
      machine.AddState(
        activeStateName,
        NewState(states[0], states[1]),
      )
      activeStateName = ""
      states[0] = &StateDescriptor{}
      states[1] = &StateDescriptor{}
    }

    if strings.HasPrefix(s, "In state") {
      fmt.Sscanf(s, "In state %1s:", &activeStateName)
    } else if strings.HasPrefix(s, "  If the current") {
      fmt.Sscanf(s, "  If the current value is %d:", &currentBranch)
    } else if strings.HasPrefix(s, "    - Write the") {
      fmt.Sscanf(s, "    - Write the value %d.", states[currentBranch].Write)
    } else if strings.HasPrefix(s, "    - Move one slot to") {
      var direction string
      fmt.Sscanf(s, "    - Move one slot to the %s", &direction)
      if (direction == "left.") {
        states[currentBranch].Move = -1
      } else {
        states[currentBranch].Move = 1
      }
    } else if strings.HasPrefix(s, "    - Continue with state") {
      var n string
      fmt.Sscanf(s, "   - Continue with state %1s.", &n)
      states[currentBranch].Next = n
    }
  }

  if activeStateName != "" {
      machine.AddState(
        activeStateName,
        NewState(states[0], states[1]),
      )
  }

  return stepCount
}

func main() {
  machine := NewTuringMachine()
  stepCount := LoadInstructions(machine)
  fmt.Println(stepCount)
  for i := 0; i < stepCount; i++ {
    fmt.Printf("%.2f%%\r", float64(i)/float64(stepCount) * 100)
    machine.Step()
  }
  fmt.Printf("Checksum: %d\n", machine.Checksum())
}
