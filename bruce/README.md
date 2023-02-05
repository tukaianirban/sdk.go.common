# Configuration reader utility

Purpose of this bruce module is to make handling configuration files easier.
It is meant to be used in applications that read their configuration either from a configuration file, from Google Cloud Platform, AWS or similar sources.

## Main functionalities

The bruce module provides 2 main functions:

- a lightweight means of making any json-structured configuration available at runtime in a convenient ( dot[.] separated ) format.
- Automatically reload configuration from a configfile if it changes.

## How to use it

Init(<mode>, <optional param>) initializes the bruce module. The module can run in either of the following modes.
Depending on the mode, it might be required to pass in optional parameter(s).

### Modes of operation

The bruce module can execute in either of the following modes:

- MODE_LOCAL : In this mode, bruce attempts to load the application's configuration from a json-formatted file available locally to the application.
  Init() must (MANDATORY) be called with the absolute or relative path of the configuration file passed in.

- MODE_GCP : In this mode, bruce attempts to load the application's configuration from a remote GCP account.
  A local file ( with json-formatted data ) must be provided with the configuration to read from the GCP account.
  Init() must (MANADATORY) be called with the absolute or relative path of this local file.
  Yet to be implemented

-MODE_AWS : In this mode, bruce attempts to load the application's configuration from a remote AWS account.
A local file ( with json-formatted data ) must be provided with the configuration to read from the GCP account.
Init() must (MANADATORY) be called with the absolute or relative path of this local file.
Yet to be implemented

### Upcoming features

- Implementation of the modes : MODE_GCP, MODE_AWS
- Implementation of a mode to read from ENV params
