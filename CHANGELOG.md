# CHANGE LOGS
###4.3.1
- Support data type blob.
- Add passwordEncrypt option for configuration to encrypt password.
Available values are MD5, LDAP, NOTHING, MD5 is default.

###4.3.0
- Support data type LIST and POINT, remove ARRAY data type.
- AsAttr support LIST data type.
- Support NULL value.
- Implement backup interface.
- Improve error message when failed to create connection pool
- Differentiate timestamp and datetime when parse from string
- Support special character when create/show schema and create property
- 


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