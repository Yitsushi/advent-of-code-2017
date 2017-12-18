package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* Instruction */
type Instruction struct {
  Command string
  First string
  Second string
}

/* Sound Card */
type SoundCard struct {
  Instructions []*Instruction
  Head int
  LastNote int64
  Memory map[string]int64
}

func (c *SoundCard) Initialize() {
  c.Instructions = make([]*Instruction, 0)
  c.Memory = map[string]int64{}
  c.Head = 0
}

func (c *SoundCard) AddInstruction(ins *Instruction) {
  c.Instructions = append(c.Instructions, ins)
}

func (c *SoundCard) Play() {
  for c.Head < len(c.Instructions) {
    current := c.Instructions[c.Head]
    switch current.Command {
    case "snd":
      c.LastNote = c.ResolveRef(current.First)
    case "set":
      c.MemorySet(current.First, c.ResolveRef(current.Second))
    case "add":
      c.MemorySet(current.First, c.ResolveRef(current.First) + c.ResolveRef(current.Second))
    case "mul":
      c.MemorySet(current.First, c.ResolveRef(current.First) * c.ResolveRef(current.Second))
    case "mod":
      c.MemorySet(current.First, c.ResolveRef(current.First) % c.ResolveRef(current.Second))
    case "rcv":
      if c.ResolveRef(current.First) != 0 {
        fmt.Printf("Recovered: %d\n", c.LastNote)
        return
      }
    case "jgz":
      if c.ResolveRef(current.First) > 0 {
        fmt.Printf("jgz %d %d\n", c.ResolveRef(current.First), c.ResolveRef(current.Second))
        c.Head += int(c.ResolveRef(current.Second))
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

  r := bufio.NewReader(file)
  card := &SoundCard{}
  card.Initialize()
  for s, e := Readln(r); e == nil; s, e = Readln(r) {
    ins := &Instruction{}
    fmt.Sscanf(s, "%s %s %s", &ins.Command, &ins.First, &ins.Second)
    card.AddInstruction(ins)
  }

  card.Play()
}
