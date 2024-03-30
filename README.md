Small Go project to try to simulate PlaceIt game solvers and find an optimal one. Mainly for learning.

Example output:
```
Trials: 1000000
Average for Random: 4.552462
Histogram for Random: [49910 134814 178714 180114 153684 116871 80893 49686 28840 14863 6987 2948 1134 401 104 29 6 2 0 0]
Average for Middle: 6.492687
Histogram for Middle: [0 0 41475 133424 183596 188115 162443 122246 80688 47075 23962 10877 4251 1333 403 98 11 3 0 0]
Average for Bin: 9.900919
Histogram for Bin: [1736 6689 15518 27924 43793 61630 80947 97903 109907 115775 113211 102075 83639 61813 40334 22398 10106 3587 912 103]
Average for Hyper Geom: 9.547571
Histogram for Hyper Geom: [3778 10377 20764 34331 50681 68762 86180 101962 111062 113300 108146 94946 75913 54353 34674 18668 8417 2854 736 96]

Simulate took 31.2525743s
```

Based on this testing, the optimal strategy for PlaceIt would be binning, where you assign number ranges to each slot that are equally distributed between the 2 limit values, and you place the number your received in the slot with a bin range that includes the pulled number.

Somehow the binning strategy even outperforms the probabilistic strategy where I tried to use the Hypergeometric Distribution to calculate the probability of winning if you place the number in each valid slot, and pick the slot with the highest win probability.

MIT Licensed
