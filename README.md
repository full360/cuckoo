# Health

Health is a simple command line tool application that polls a service count from
Consul and post it as a metric in CloudWatch.

## Usage

Use the `-help` flag to display all the available parameters:

    health -help
      -block duration
            Consul blocking query time (default 10m0s)
      -service string
            Name of the Service to check
      -tag string
            Tag name of the Service to check
      -version
            print version and exit

Required parameters are `-service` and `tag`.

If using the default Consul blocking query time (10 minutes):

    health -service voltdb -tag prod

If less blocking time required:

    health -service voltdb -tag prod -block 5m

