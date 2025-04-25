# Justify the Hike
I understand that it might seem like a quick switch, but for me,  I’m not switching just for a hike — I’m looking for a role where I can contribute for longer time.

An offer around 30% hike would give me the right stability and motivation to commit myself for the next 2–3 years.
In the past 6 months, I’ve significantly expanded my capabilities.
I believe a 30% hike is a fair reflection of my current capabilities and market worth.

# Justify the Location
My fiancée is currently working in Chennai, near Nungambakkam, so I decided to relocate here. It’s important for me to maintain a good work-life balance, and being in the same city definitely supports that.

# Justify the Fionce thing
Yes, it’s possible that she might move in the future, but she recently joined her current company, so we expect to stay in Chennai for at least the next two years. And even if there is a move, it would most likely be to Bangalore. Since your company also has an office in Bangalore, I would be open to relocating internally if needed.

# What you Learn in Past 6 months
Within my first month, I wrote 150+ unit and functional test cases, which significantly improved test coverage and system stability. Additionally, I had the opportunity to work with gRPC/gNMI, gaining an understanding of how APIs manage data center switch racks. I also gained hands-on experience with Jenkins for CI/CD and expand my knowledge of RabbitMQ. Moreover, I had full ownership of the Inventory Policy and Health microservices, So solve multiple BUGs and added some Features.

Built a centralized /health/summary endpoint aggregating the status of all microservices for dashboards and alerts. Queries all microservices' health endpoints. Aggregates their status (UP/DOWN/DEGRADED). Returns a unified JSON. Forward the JSON to a Prometheus exporter,.
Added alerting for critical service degradation using wPrometheus + Alertmanager integrations.

Added business rule checks to prevent invalid or conflicting switch states.

How did you ensure the quality and reliability of the test cases you wrote?
I made sure the test cases were clearly defined with specific inputs, expected outputs, and edge cases covered. I used code coverage tools to verify that all critical parts of the system were tested. I ensured the tests covered not just the happy path but also potential failure scenarios

gRPC is a general-purpose framework for building efficient remote procedure calls across any network service, suitable for a wide range of applications.

gNMI is a network-specific extension built on top of gRPC for managing network devices, specifically in telecom and data center environments.

gNMI (gRPC Network Management Interface) is specifically designed for managing and configuring network devices, like switches, in a data center environment.

gRPC is primarily used for sending RPC calls, typically for accessing and invoking functions.

gNMI defines three primary operations (Get, Set, and Subscribe) that are focused on retrieving and setting configurations, as well as real-time monitoring of network devices.