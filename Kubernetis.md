# Kubernetes 
Kubernetes is designed for automate the process of deploying, scaling, and managing containerized applications.  Kubernetes uses YAML files to define resources.

A Pod can contain one or more containers that share the same network, namespace and storage.
Containers in the same Pod share a common IP address and port space. They can communicate with each other using localhost and the same port numbers. Pods can communicate with each other using their IP addresses and defined ports. Pods are often managed by higher-level controllers like Deployments, ReplicaSets, which handle scaling, updates, and ensuring that the desired number of Pod replicas are running.

ReplicaSet is a controller that ensures a specified number of identical Pods are running at any given time.   If a Pod fails or is deleted, the ReplicaSet creates a new Pod to replace it. A ReplicaSet uses a label selector to identify which Pods it should manage. The selector matches Pods based on their labels, and the ReplicaSet ensures that the Pods with these labels are created or deleted as needed.

Services provide a stable endpoint(DNS name and IP address) for accessing a set of Pods. This stable endpoint is crucial because Pods can come and go due to scaling, updates, or failures, but Services provide a consistent way to access these Pods.

Kubernetes provides built-in mechanisms for service discovery, so applications can find and communicate with other services without hardcoding IP addresses.

Types of Services:

ClusterIP : The Service is only accessible within the Kubernetes cluster. It’s used for communication between different services within the cluster.

NodePort : This allows the Service to be accessed from outside the cluster by requesting <NodeIP>:<NodePort>. This type is often used for simple testing and debugging.

LoadBalancer : Creates an external load balancer (e.g., from cloud providers like AWS, GCP, Azure) and assigns a public IP to the Service. This allows the Service to be accessed from outside the cluster using that public IP. It’s often used in production environments.

Headless Service : When you don’t need load balancing or a cluster IP, you can create a headless Service by setting the clusterIP field to None. This allows direct access to the Pods without load balancing, useful for stateful applications.

Services use selectors to determine which Pods belong to the Service. 
selector : Specifies that the Service will route traffic to Pods with the label app: my-app


Deployment is a higher-level abstraction that manages ReplicaSets and Pods. Deployments provide version control and history for your application, making it easy to roll back to previous versions if needed.
Deployments support rolling updates, allowing you to update your application with zero downtime. Kubernetes gradually replaces old Pods with new ones according to the update strategy you specify.
If a new update causes issues, you can roll back to a previous version of your application. Kubernetes maintains a history of revisions, making it easy to revert to a stable state.
Deployments allow you to scale the number of replicas (Pods) up or down easily.


Kubernetes uses a flat network model where every Pod gets its own IP address. Pods communicate directly with each other using their IP addresses, 

Cluster is a set of machines (nodes) that work together to run containerized applications. The cluster consists of a control plane(Master Node) and a set of worker nodes, and it provides the environment needed to deploy, manage, and scale applications.
What are the main components of Kubernetes architecture?
- Master Node: The node that manages the cluster. It includes:

	- API Server: The API server exposes the Kubernetes API which is used by users and other components to interact with the cluster. When we run commands using kubectl or other tools, they communicate with the API Server.  When we deploy a new application, the API Server stores the desired state in etcd and updates the cluster to match this state

	- Controller Manager: Each controller watches the API Server for changes to resources and makes adjustments to ensure that the actual state matches the desired state. For example, if a ReplicaSet specifies that three replicas of a Pod should be running, the controller will create or delete Pods as necessary to meet this requirement.

	- Scheduler: Assigns Pods to Nodes based on resource availability and constraints. if a Pod requires 2 CPU cores and 4 GB of memory, the Scheduler will find a Node with sufficient available resources and assign the Pod to that Node.
	
	- etcd: A distributed key-value store that holds all the configuration data and state of the Kubernetes cluster.  When we make changes to the cluster (e.g., deploying an application or scaling a service), the API Server updates the state in etcd.  This allows Kubernetes to recover from failures by restoring the cluster state from etcd.

- Worker Nodes: worker node refers to a machine (virtual or physical) that runs the applications and workloads in your Kubernetes cluster. On AWS, worker nodes are typically Amazon EC2 instances that are part of your Kubernetes cluster.
Kubernetes Components on Worker Nodes: 
	- Kubelet: An agent that ensures containers are running in Pods.  The Kubelet communicates with the API Server to get the desired state of Pods. It then makes sure that the containers in those Pods are running as expected. If a container fails or crashes, the Kubelet will restart it to maintain the desired state. It also collects and reports metrics about the Node and the Pods running on it.

	- Kube-Proxy: Maintains network rules and load-balances traffic to Pods. Kube-Proxy manages network traffic routing by maintaining iptables or IPVS rules on Nodes. This ensures that network traffic is properly directed to the Pods. For example, if a Service exposes multiple Pods, Kube-Proxy ensures that traffic is distributed evenly across these Pods.

	- Container Runtime: Software that runs and manages containers (e.g., Docker, containerd).

We can scale a cluster by adding or removing worker nodes as needed to handle varying workloads.

AWS provides a managed Kubernetes service called EKS. When you create an EKS cluster, you can specify the EC2 instance types and configurations for your worker nodes.

Elastic Load Balancing (ELB): Kubernetes services that are exposed as LoadBalancer types, can use AWS ELB to provide external access and distribute traffic to the worker nodes.

Amazon EBS: For persistent storage, you can use Amazon Elastic Block Store (EBS) volumes that are attached to worker nodes.

Helm is a package manager for Kubernetes that simplifies the deployment, management, and versioning of applications and services on a Kubernetes cluster.

A Helm chart is a collection of Kubernetes YAML files organized into a directory structure that defines a Kubernetes application. It includes everything needed to run an application, such as deployments, services, ingress configurations, and more.

Helm is installed on your local machine or CI/CD pipeline and communicates with your Kubernetes cluster to deploy and manage applications.

A typical Helm chart directory includes
Chart.yaml: Contains metadata about the chart, such as its name, version, and description.
values.yaml: A file containing default configuration values that can be overridden by the user.
templates/: A directory with Kubernetes manifest templates that Helm will use to generate Kubernetes resources.
charts/: A directory where dependent charts can be stored.
README.md: Documentation about the chart.

Helm can install, upgrade, rollback, and delete releases. Each release can be upgraded or rolled back independently of others.