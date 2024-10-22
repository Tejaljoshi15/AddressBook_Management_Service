TASK-2

*_AddressBook_* *_Management_* *_Service_*


This service is designed to manage user information with an in-memory database. It provides functionalities to add, find, update, and delete users via both gRPC and HTTP protocols.

*Features*

Add New Users: Users can be created with unique usernames along with contact details.
Find Existing Users: Users can be searched by any field (username, address, phone), including partial matches using wildcards.
Update User Information: Modify any field of a user's information.
Delete Users: Users can be removed from the storage.

*Configuration*

Configuration parameters are managed in a configuration file named config.yaml. This file includes:

server_port: Port for running the gRPC server.
http_port: Port for running the HTTP gateway.
log_level: Log level for application logging.