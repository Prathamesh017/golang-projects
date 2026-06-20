Designing A LRU CACHE IN GO

LRU = Least Recently Used
It's a cache with fixed capacity that stores frequently accessed data.
When the cache becomes full and a new item needs to be inserted, it evicts (removes) the item that hasn't been used for the longest time.

This is True LRU Cache , Most of the system don't implement as incrementing the count is very expensive , So they use `approximations`

They use
- FIFO (not LRU)
- Clock Algorithm (Approximate LRU)
- Segmented/2Q caches

Also Important to remember , LRU focus on LRU only cares about the latest access time, not the number of accesses.
so every new element is always put in front, as it is new element  But in actual production is not enough right 

for example:=  Normally, every day these pages are very popular: /home
/profile , for a certain 30 min - /cricket-world-cup , the important cache might get deleted

So In production , we think of
- 1. Recency (LRU)
- 2. Frequency (LFU)
- 3. Expiration (TTL)

Redis always supports TTL, but its eviction policy is configurable. can use different policies like allkeys-lru , allkeys-lfu etc , Also redis use `approximate LRU` not true LRU

Memcached traditionally uses LRU, but it's more sophisticated than textbook LRU.It uses segmented LRU.

Definging data-structure for usage

so as we need to maintain head and tail for most recent and least recent , LinkedList is obv choice , so but getting a value in linkedlist is O(N) ,so we use hashmap

what is usage of hashmap here, 
it is basically to check if a value already exists in our linkedlist , if yes , remove it from the current pos and move it top as it is recently access value

