# Kademlia DHT quick-faq
Copyright (c) 2023 Johanness A. Nilsson. All Rights Reserved.

Kademlia is a distributed hash table (DHT) for decentralized peer-to-peer networks. 
---

## Routing in Kademlia has some unique aspects:

1. **Node IDs and XOR Metric**: Nodes are identified using a fixed-length binary string. The distance between nodes is calculated using the XOR metric.
2. **KBuckets**: Nodes maintain lists of other nodes in "KBuckets", which are split based on distance ranges from the node. Each KBucket holds nodes of a certain distance range.
3. **Iterative Routing**: Nodes find each other via iterative routing, where a node queries others to find closer nodes to the target, refining the search with each step.
4. **Lookup Algorithm**: To find a node or value, Kademlia employs a lookup algorithm that queries the KBuckets to find the closest nodes, and iteratively narrows down the search.
5. **Routing Table Refresh**: To maintain updated network info, nodes refresh KBuckets periodically, ensuring the routing table remains accurate over time.

This design makes routing in Kademlia efficient and resilient, allowing for quick lookups and robustness in dynamic network conditions.

## XOR Metric (exclusive or) 

XOR is used as a metric to define distances in the network:

1. **Node ID**: As mentioned before, each node has a unique identifier, and the XOR of two node IDs is used to calculate the "distance" between them.
2. **Distance Metric**: The XOR value of two node IDs is treated as the distance between them. This metric has properties that convey the triangle inequality. From the disciplines of Metric Space Theory we know that; *Triangle Inequality* states that: *d(A, B) + d(B, C) ≥ d(A, C).* 
3. **Routing**: When a node needs to find another node or a value, it starts by querying nodes in its own buckets that are closest to the target ID, as per the XOR metric. The queried nodes return their known closest nodes to the target, and the querying node then iteratively queries these nodes, getting closer to the target with each step.
4. **Efficient Lookups**: This method allows for efficient lookups as the number of steps required to find a node or value grows logarithmically with the size of the network, making routing scalable.
  
XOR routing is fundamental to Kademlia’s design, enabling the network to function efficiently and effectively.

## KBuckets
Data structures used for routing and network organization. Here's a concise breakdown:

1. **Organization**: KBuckets help organize known nodes based on the XOR distance metric from the local node. Each KBucket stores nodes at specific distance ranges.
2. **Fixed Size**: Each KBucket has a fixed size, often 20 nodes. When a KBucket is full and a new node is discovered, if the node is closer than others, the oldest unresponsive node is evicted.
3. **Splitting**: KBuckets can split when they become full, especially if the new node is closer, to maintain finer granularity of distance ranges.
4. **Updating**: KBuckets are updated with new node information over time, which helps in maintaining an accurate view of the network.
5. **Lookup**: During a lookup, KBuckets are queried to find closer nodes to a desired node or key, facilitating efficient routing.

The implementation of KBuckets plays a central role in the KDHT's ability to ensure effective and efficient routing within the network.

## Iterative routing 

The process where a node actively seeks information through a series of iterative queries to other nodes, getting closer to the desired data with each step.

1. **Lookup Start**: Begins with the node looking for either a value or another node’s info. It starts by querying nodes in its KBucket(s) that are closest to the target identifier.
2. **Refined Queries**: With each query, the node receives info about other nodes that are closer to the target. It then queries these closer nodes in the next iteration.
3. **Convergence**: Through successive iterations, the querying node converges towards the nodes that hold the desired data or are closest to the target identifier.
4. **Parallelism**: Multiple queries are optionally sent out in parallel to expedite the process and enhance the robustness of the lookup.  Parallelism can be dynamically controlled considering a node's strength or position. A node could adjust its parallelism based on performance metrics like query response times, or feedback from peers regarding network congestion. For instance, a well-positioned node might increase parallelism for faster lookups, while a peripheral node might reduce parallelism to ease it's own localized   Masload. 

This iterative process makes routing in Kademlia highly efficient and economic. Kademlia will allow the network to function effectively even in dynamic conditions with nodes constantly joining and leaving. All of Kademlia's features add up to being naturally suited for large, globally distributed,  multi-party real-time gaming services due to its efficient routing, low latency, and self-organizing nature, and serverless architecture. The XOR-based distance metric and KBucket system enable quick lookup times and efficient network organization, crucial for low-latency real-time interactions in a gaming environment. Iterative routing with controlled parallelism facilitates rapid location of game nodes or assets, while the decentralized architecture enhances scalability, allowing the network to handle a constantly growing number of players with minimal performance degradation. Furthermore, Kademlia's ability to maintain network integrity amidst dynamic node entrances and exits ensures a robust and resilient infrastructure, critical for providing a seamless gaming experience across a globally dispersed player base.
