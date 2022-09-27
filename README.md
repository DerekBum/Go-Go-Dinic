# Go-Go-Dinic
This is [Dinic's algorithm](https://en.wikipedia.org/wiki/Dinic%27s_algorithm) for computing the maximum flow in a flow network implementation using Golang.
## Input format
The first must line contain two space-separated integers *n* and *m* (1 ≤ *n* ≤ 1000) – the number of vertices and the number of edges correspondingly.  
Then *m* lines follow – the dedescription of *m* edges. The *i*-th line contains three integers $a_i$, $b_i$, $c_i$, indicating an oriented edge from vertice $a_i$ to $b_i$ with capacity $c_i$.  
The last line must contain two space-separated integers *s* and *t* – the numbers of source and sink correspondingly.
## Output format
One integer – the maximum flow between *s* and *t*.
