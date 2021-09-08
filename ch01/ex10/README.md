The latencies of the two trials are significantly different, so the results seem to be cached. The reponse bodies are exactly same.

```console
$ go build main.go 
$ ./main out1.txt && ./main out2.txt
2.20s  305541 https://hippocampus-garden.com/best_papers_2020/
2.198516s elapsed
1.58s  305541 https://hippocampus-garden.com/best_papers_2020/
1.577605s elapsed
$ diff out1.txt out2.txt 
```