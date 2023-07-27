#### Metrics Logger
A metrics logger is a service that is responsible for logging metrics to a database. The metrics logger will be installed at the client and will connect to the logging service via a REST API. The metrics logger will be responsible for logging the following metrics:

### Functional Requirements

- Receive metrics from the client
- Agnostic to the type of metrics being logged
- Store metrics in a database
- Provide a REST API for the client to connect to

### Non-Functional Requirements

- Near real-time logging of metrics
  - Metrics should be logged within 1 second of being received
  - Low latency
- High availability
  - Metrics should be logged even if the logging service is not entirely functional
  - High uptime
- High throughput
  - The logging service should be able to handle a large number of metrics without slowing down

### Assumptions

- Each metric payload will be less than 1KB
- Each client will send metrics at a rate of 1 metric per second
- The number of clients will be about 1000
- Users will be able to view metrics in a dashboard
- Users will be able to filter metrics by time, client, and metric type

### Constraints

