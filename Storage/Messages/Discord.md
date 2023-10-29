How Discord Stores Trillions of Messages

In 2015, Discord started using MongoDB to store messages. By November 2015, there were already 100 million messages that did not fit in RAM, causing unpredictable latency.

In 2017, Discord moved to storing billions of messages in a Cassandra database. By 2022, the number of stored messages had ballooned to trillions across 177 nodes.

But Cassandra was plagued by several serious issues.

- Hot partitions occurred when a small number of high traffic channels overwhelmed nodes, cascading latency across the cluster.
- Garbage collection pauses created latency spikes.
- Compactions fell behind, forcing expensive reads to query multiple SSTables.
- Maintenance like node repairs interrupted service.

Discord migrated to ScyllaDB for the following benefits:

- Written in C++ instead of Java, eliminating disruptive garbage collection pauses
- Shard-per-core model provides stronger workload isolation to prevent hot partitions from cascading latency across nodes.
- Reverse query performance optimized to meet Discord's needs
- They reduced nodes to 72 while increasing disk space per node to 9TB.

To further protect ScyllaDB, Discord:

- Built intermediary data services in Rust that limit concurrent traffic spikes
- Data services sit between the API and database, coalescing requests
- Query the database just once even if multiple users request the same data
- Rust provided fast, safe concurrency ideal for this workload

The results? Tail latencies down from 40-125ms to a steady 15ms. Database uptime improved from weekend-long outages to smooth sailing. The system easily handled World Cup traffic spikes, processing events like goals and penalties without breaking a sweat. Discord continues to scale, now reliably storing trillions of messages with ScyllaDB.


