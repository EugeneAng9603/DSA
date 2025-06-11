Scenario	            Approach	                                                    Time Complexity
All in [0,100]	    Counting array only	                                        O(1 insert), O(100) median
99% in [0,100]	    Counting + small buffer	                               O(1 insert), O(n log n) median (only for outliers)
Arbitrary numbers	Two heaps	                                                O(log n) insert, O(1) median