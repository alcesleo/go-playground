# Word Ladders

Finds [word ladders](https://en.wikipedia.org/wiki/Word_ladder)
between words. Play around with the main function to try different
ones, it is using the `/usr/share/dict/words` dictionary present on
Macs which is easy but not the greatest dictionary...

```bash
$ go run ladders.go
13 steps between chaos and order
chaos -> chais -> chair -> cheir -> cheer -> sheer -> smeer -> emeer -> emmer -> ammer -> armer -> ormer -> order

21 steps between right and wrong
right -> dight -> digit -> dimit -> demit -> remit -> refit -> befit -> besit -> beset -> beret -> buret -> burst -> burse -> birse -> biose -> brose -> prose -> prone -> prong -> wrong
```
