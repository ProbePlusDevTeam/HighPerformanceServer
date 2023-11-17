Horizontal scalability is crucial for systems that handle high loads.

Here are a few characteristics of a horizontally scalable API:

- Stateless sessions
- Load balancing
- Fault tolerance
- Caching

All right, that all sounds nice.

But what does it mean?

𝗦𝘁𝗮𝘁𝗲𝗹𝗲𝘀𝘀 𝘀𝗲𝘀𝘀𝗶𝗼𝗻𝘀

A horizontally scalable API should be stateless. It doesn't store any session data on the server side.
Instead, the client maintains the state of the application.

This is by design using REST APIs.

𝗟𝗼𝗮𝗱 𝗯𝗮𝗹𝗮𝗻𝗰𝗶𝗻𝗴

Load balancing distributes requests across different API instances. This reduces the overall load on a single instance, which helps to improve API performance. 

Also in some ways this increases High Availability (HA)

𝗙𝗮𝘂𝗹𝘁 𝘁𝗼𝗹𝗲𝗿𝗮𝗻𝗰𝗲

Fault tolerance means the API can handle failures gracefully. This includes:
- Handling errors
- Retrying failed requests
- Maintaining high availability

𝗖𝗮𝗰𝗵𝗶𝗻𝗴

Use caching to improve performance and reduce the database load. The important thing is to use a distributed cache (like Redis).
Otherwise, caching won't help when scaling out APIs.

