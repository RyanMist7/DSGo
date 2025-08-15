# DSGo

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