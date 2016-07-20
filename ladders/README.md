# Word Ladders

Finds [word ladders](https://en.wikipedia.org/wiki/Word_ladder)
between words. Play around with the main function to try different
ones, it is using the `/usr/share/dict/words` dictionary present on
Macs which is easy but not the greatest dictionary...

```bash
$ go run ladders.go
4 steps between cold and warm
cold -> cord -> word -> ward -> warm

12 steps between chaos and order
chaos -> chais -> chair -> cheir -> cheer -> sheer -> smeer -> emeer -> emmer -> ammer -> armer -> ormer -> order

20 steps between right and wrong
right -> dight -> digit -> dimit -> demit -> remit -> remix -> remex -> resex -> resee -> besee -> belee -> belve -> beeve -> breve -> brede -> bride -> brine -> bring -> wring -> wrong
```
