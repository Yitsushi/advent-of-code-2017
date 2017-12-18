package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "time"
)

const (
  Running = 1
  Waiting = 2
  Stopped = 3
)

/* Instruction */
type Instruction struct {
  Command string
  First string
  Second string
}

/* Sound Card */
type SoundCard struct {
  ID int64
  Instructions []*Instruction
  Head int
  State int64
  Memory map[string]int64
  NumberOfSend int64
  inCh chan int64
  outCh chan int64
}

func (c *SoundCard) Initialize(progId int64, inCh chan int64, outCh chan int64) {
  c.ID = progId
  c.Instructions = make([]*Instruction, 0)
  c.inCh = inCh
  c.outCh = outCh
  c.Memory = map[string]int64{}
  c.Memory["p"] = progId
  c.Head = 0
  c.NumberOfSend = 0
}

func (c *SoundCard) AddInstruction(ins *Instruction) {
  c.Instructions = append(c.Instructions, ins)
}

func (c *SoundCard) Play() {
  c.State = Running
  for c.Head < len(c.Instructions) {
    current := c.Instructions[c.Head]
    switch current.Command {
    case "snd":
      c.NumberOfSend++
      c.outCh <- c.ResolveRef(current.First)
    case "set":
      c.MemorySet(current.First, c.ResolveRef(current.Second))
    case "add":
      c.MemorySet(current.First, c.ResolveRef(current.First) + c.ResolveRef(current.Second))
    case "mul":
      c.MemorySet(current.First, c.ResolveRef(current.First) * c.ResolveRef(current.Second))
    case "mod":
      c.MemorySet(current.First, c.ResolveRef(current.First) % c.ResolveRef(current.Second))
    case "rcv":
      select {
      case val := <-c.inCh:
        c.State = Running
        c.MemorySet(current.First, val)
      default:
        c.State = Waiting
        continue
      }
    case "jgz":
      if c.ResolveRef(current.First) > 0 {
        c.Head += int(c.ResolveRef(current.Second))
        c.State = Running
        continue
      }
    }
    c.Head++
  }
}

func (c *SoundCard) ResolveRef(ref string) int64 {
  v, e := strconv.ParseInt(ref, 10, 64)
  if e == nil {
    return v
  }

  return c.MemoryGet(ref)
}

func (c *SoundCard) MemoryGet(key string) int64 {
  if _, ok := c.Memory[key]; !ok {
    c.Memory[key] = 0
  }
  v, _ := c.Memory[key]
  return v
}

func (c *SoundCard) MemorySet(key string, val int64) {
  c.Memory[key] = val
}

/* Utilities */
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

/* Main */
func main() {
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

  cardOne := &SoundCard{}
  cardTwo := &SoundCard{}

  cardBusOne := make(chan int64, 1024)
  cardBusTwo := make(chan int64, 1024)
  cardOne.Initialize(0, cardBusOne, cardBusTwo)
  cardTwo.Initialize(1, cardBusTwo, cardBusOne)

  r := bufio.NewReader(file)
  for s, e := Readln(r); e == nil; s, e = Readln(r) {
    ins := &Instruction{}
    fmt.Sscanf(s, "%s %s %s", &ins.Command, &ins.First, &ins.Second)
    cardOne.AddInstruction(ins)
    cardTwo.AddInstruction(ins)
  }

  go cardOne.Play()
  go cardTwo.Play()

  for {
    if (cardOne.State == Waiting && cardTwo.State == Waiting) {
      fmt.Printf("Card#%d sent %d packages.\n", cardOne.ID, cardOne.NumberOfSend)
      fmt.Printf("Card#%d sent %d packages.\n", cardTwo.ID, cardTwo.NumberOfSend)
      break
    }

    time.Sleep(time.Millisecond * 10)
  }
}
