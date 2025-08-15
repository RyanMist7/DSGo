# DSGo

## Directory Structure
```tree
├── core
│   ├── core.go
├── go.mod
└── README.md
```

## Idea
The basic idea is we have nodes that pass messages. We're solving consensus and we need some way to test that. To achieve all of this we'll use a simulator named core. 

**Core** needs to be able to do the following:
* Pass messages between nodes and the client.
* Simulate message failure (loss/duplication/delay/reordering).
* Simulate node failure (crash/recovery). 
* Simulate partitions.

**Nodes** are some implementation of the algorithm. Each node should be able to:
* Recieve and process messages
* Send messages to any node or client (via Core)
* React to timers

We would like to be able to test with some combination of determinstic testing + search tests.

## Design
## Core
Core is the simulation engine. Its function is to model a distributed system through a set of nodes interconnected with a simulated network.

The goals for Core are (1) reproduce distributed system, (2) allow configurable fault injections, and (3) provide a clean way to nodes to send and recieve messages. The natural approach to this is messages (through channels) are routed through Core.

## Node
Each node runs as a go routine and implements the following:
  type Node interface {
      Init(ctx NodeContext)
      HandleMessage(msg Message)
      HandleTimer(timerID string)
  }
Core provides the following context for the node to use:
type NodeContext interface {
    Send(destID string, msg Message)
    SetTimer(timerID string, delay time.Time)
    Log(level string, format string, args ...any)
}
