If you want high availability, you must get rid of your single points of failure.

A single point of failure (SPOF) is a service that, if it fails, will stop the entire system from working.

Imagine a web application that relies only on a single authentication service.

If that server crashes, the entire application becomes non-functional.

Good news: you identified a single point of failure. Now, let’s talk about how to solve it.

The answer is Redundancy. If one node fails, the other can take over.

You have 3 main approaches:

Active Components

Instances of your services running all the time and a Load Balancer helping distribute the load.

It helps distribute requests across instances, preparing the system for traffic spikes.

Easier to add or remove instances as needed.

Think of e-commerce sites during a sale. If you have a high load, go for Active Components.

Passive Components

In this case, you still have a second instance but stay idle and ready to take over if the primary component fails.

Only one running instance means a lower operational cost.

Less complexity in management compared to an active-active setup.

But there is a caveat: The passive component may take time to initialize resources. 

And each minute of downtime costs thousands of dollars.

Combining Active and Passive Components

You can have two active and one passive instance offering a middle ground; this was my winning formula.

Fewer active components mean less operational cost compared to a full active setup.

With active and passive, you enjoy quick failover and load distribution.

No system is bulletproof, but designing your systems for Redundancy will:

- Reduce your single points of failure.
- Reduce your downtime.
- Keep your customers happy.
- Help you to stay away from On-Call rotations.


