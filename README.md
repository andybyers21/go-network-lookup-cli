# Network Lookup CLI

Query things like the name-servers of a website, the CNAME's, IP addresses and
so on in the command line.

Using [urfave/cli](https://github.com/urfave/cli)

## Commands

| Command  | Will do                                                |
|----------|--------------------------------------------------------|
| `ns`     | will retrieve the name servers                         |
| `cname`  | will lookup the CNAME for a given host                 |
| `mx`     | will lookup the mail exchange records for a given host |
| `ip`     | will lookup the IP addresses for a given host.         |
