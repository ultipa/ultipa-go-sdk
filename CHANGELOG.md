# CHANGE LOGS

## Version 4.2.1

- fixed bug: clear task should send to global graphset

## Version 4.2.0

- add InstallExta Methods
- add UninstallExta Methods

## Version 4.0.10

- add insertRequestConfig
- add insertResponse
- update insertBatch* Methods

## Version 4.0.5

- Support more layouts when parsing string to UltipaTime
- Support timestamp when deserializing from bytes to GoLang Type and return to an interface
- Attr add PropertyType
- Add more comments to request configurations and connection configurations



## Version 4.0.4 - release

- Add Model Manager
  - manage model(graph)
    - manage schemas for a model
- Add load db config from yaml
- Add asFirstNode、asFirstEdge Methods for DataItem
- Bug Fix
  - Fix Context Memory leak