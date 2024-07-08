Small Go project to try to simulate PlaceIt game solvers and find an optimal one. Mainly for learning.

Example output:
```
Trials: 1000000
Average for Random: 4.548556
Histogram for Random: [49963 135576 178492 180337 153778 116795 80011 50065 28260 14892 7135 3105 1095 365 97 26 6 1 1 0]
Average for Middle: 6.488814
Histogram for Middle: [0 0 41918 133064 183321 188856 163007 122043 80009 47187 23911 10628 4215 1377 355 89 20 0 0 0]
Average for Bin: 9.858459
Histogram for Bin: [2379 7554 16380 28785 44722 62383 81352 96565 110140 115258 113206 100744 82741 61077 39781 22201 10163 3575 877 117]
Average for Hyper Geom: 9.543664
Histogram for Hyper Geom: [3845 10620 20630 34357 50867 68668 86575 101045 110714 114062 108850 94562 75601 54432 34503 18912 8143 2797 722 95]
Average for Simple Bin: 5.336046
Histogram for Simple Bin: [49000 93186 125980 144387 144772 131481 107580 80686 54660 33872 18892 9256 3991 1564 542 130 19 2 0 0] 
Average for Simple Bin Rand: 8.489348
Histogram for Simple Bin Rand: [2473 12042 29520 52810 76217 98368 114650 123377 121162 109945 91146 68847 46961 28078 14630 6639 2360 662 101 12]

Simulate took 42.9315569s
```

Based on this testing, the optimal strategy for PlaceIt would be binning, where you assign number ranges to each slot that are equally distributed between the 2 limit values, and you place the number your received in the slot with a bin range that includes the pulled number.

Somehow the binning strategy even outperforms the probabilistic strategy where I tried to use the Hypergeometric Distribution to calculate the probability of winning if you place the number in each valid slot, and pick the slot with the highest win probability.

Win rate (20 placed numbers) with optimal strategy is about 0.01%, or 1 win in 10,000 games.

MIT Licensed
