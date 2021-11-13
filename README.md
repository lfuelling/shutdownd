# shutdownd

Service to shut down a system using HTTP requests.

## Usage

Here's a quick example of how you can use this software on Linux.

1. Download or build the `shutdownd` executable and copy it to `/usr/local/bin/shutdownd`
2. Add the shutdownd user: `sudo useradd -r shutdownd`
3. Copy `examples/config.json` to `/etc/shutdownd.json` and edit it to match your setup
4. Copy `examples/shutdownd.service` to `/etc/systemd/system/shutdownd.service`
5. Copy `examples/shutdownd.sudoers` to `/etc/sudoers.d/shutdownd`
6. Start the service: `systemctl daemon-reload && systemctl enable --now shutdownd`

- Make sure the port the server listens on is accessible by the client that triggers the shutdown!
- Make sure that only the service user can read/write the config file to prevent credential leakage!

### Configuration

Below you can find a list of all available configuration options and what they do.

- `authUsername` (string)
    - Username to use for HTTP Basic Auth
- `authPassword` (string)
    - Password to use for HTTP Basic Auth
- `listenAddress` (string)
    - The address string to listen on. Port is required, host/ip is optional.
- `osType` (string)
    - Decides which shutdown command to use 
    - Possible values: `linux`, `bsd`, `windows`
- `useSudo` (boolean)
    - Decides whether to prepend `sudo ` to the shutdown command
    - Ignored when `osType` is `windows`
