# Configuration reader utility

Purpose of this bruce module is to make handling configuration files simple and seamless.

The `bruce` package enables applications to seamlessly access configuration parameters irrespective of the type of store that the configuration is retrieved from.
`bruce` also automatically refreshes the latest configuration data from the store in a periodic manner enabling dynamic reload of application configurations. This in turn enables applications to be deployed and executed in true build-promotion environments.
The package provides additional security in handling application configurations by hosting it inside a structure in the application processes. This ensures that the configuration data is stored only inside the process memory and not exposed outside of it.

In any running application, there is only a single bruce instance possible. You cannot re'init bruce while the application is running

## Main functionalities

The bruce module provides 2 main functions:

- easy semantics of accessing configuration irrespective of the underlying configuration data store.
- Automatically refresh configuration from configuration data store.

## How to use it

Init() initializes the bruce module, by loading the configuration from the respective data store. It parses in command-line arguments to obtain the following initialization data:

- Provider:
- - local (default) : load the configuration from a local file (name of file obtained from arg `configfile`)
- - aws : load the configuration from AWS Config (connection parameters and credentials are obtained from arg `configfile`)
- - gcp : load the configuration from GCP Secrets Manager (connection parameters and credentials are obtains from arg `configfile`)

- ConfigFile:
- - a local file to load configuration from.
- - for mode local, the application configuration is loaded from this file.
- - for other modes, the cloud provider connection params and credentials are loaded from this file

## Security

Bruce loads application configuration from local or cloud providers into a structure in the process memory space. This way, it ensures that even privileged users of the underlying OS layer are not able to access the application's configuration data.
Since the data is stored in a structure in process memory, the configuration provider (local or cloud) can be disconnected/unmounted after the application bootup with no impact on the application itself.

### Upcoming features

- Implementation of the modes : PROVIDER_GCP, PROVIDER_AWS
- Implementation of a mode to read from ENV params
