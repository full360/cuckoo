# Cuckoo

Cuckoo, like the bird in the clock is a very simple command line tool that polls
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

Required only required flags are `-service` and `-tag` but others exist to tweak
preferences. For example:

    cuckoo -service voltdb -tag prod

Using that command will tell Cuckoo to go to Consul and look for all the healthy
services named `voltdb` and tagged with `prod`. If we want to reduce the Consul
query blocking time we'll need to send the `-block` flag and send the argument
with the unit of time. For example for 5 minutes blocking time:

    cuckoo -service voltdb -tag prod -block 5m

There are two environment variables that can be used to set the AWS region in
which Cuckoo will post metrics too. The default region is `us-east-1`

- `AWS_DEFAULT_REGION`
- `AWS_REGION`

If Consul is not running in localhost and Cuckoo needs to connect to a different
one, use the default consul address variable `CONSUL_HTTP_ADDR`

[consul]: https://www.consul.io
[cloudwatch]: https://aws.amazon.com/cloudwatch/
