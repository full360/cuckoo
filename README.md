# Cuckoo

Cuckoo, like the bird in the clock, is a very simple application that polls
[Consul][consul] for a count of healthy instances of a service and posts it to
[AWS CloudWatch][cloudwatch] as a metric.

## Usage

Use the `-help` flag to display all the available parameters:

    cuckoo -help
      -block duration
            Consul blocking query time (default 10m0s)
      -d    enables debug logging mode
      -metric-name string
            CloudWatch metric data name (default "service_monitoring")
      -metric-namespace string
            CloudWatch metric namespace (default "microservices")
      -service string
            Consul name of the Service to check
      -tag string
            Consul tag of the Service to check
      -version
            print version and exit

The only two required flags are `-service` and `-tag` but others exist to tweak
preferences. For example:

    cuckoo -service voltdb -tag prod

Using the above command will tell Cuckoo to go to Consul and look for all
healthy services named `voltdb` tagged with `prod` and will post that count
number returned from Consul to a CloudWatch metric within the namespace
"microservices" and metric data name of "service_monitoring" because these are
the default values.

If for example we want to reduce the Consul query blocking time from 10 minutes
to 5 we'll need to send the `-block` flag and the argument with the unit of
time. For example:

    cuckoo -service voltdb -tag prod -block 5m

There are two environment variables that can be used to set the AWS region in
which Cuckoo will post metrics too but if no region is set the default one will
be `us-east-1`

- `AWS_DEFAULT_REGION`
- `AWS_REGION`

If Consul is not running in `localhost` use the default Consul address variable
`CONSUL_HTTP_ADDR` to set for a reachable address

[consul]: https://www.consul.io
[cloudwatch]: https://aws.amazon.com/cloudwatch/
