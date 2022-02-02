# nem-morpheus-plugin
The plugin uses Terraform to configure Virtual Machines on EXSi platform using hpegl provider

## Required Environment variables
SERVICE_ACCOUNT -> vmaas service account (dev/intg/prod).  
TF_WORKING_DIR -> Directory where the terraform files present.

## Usage
* This plugin can be used with following terraform operations.
    * plan
    * apply
    * destroy
* Clone this project and run as shown below
    ```
    go run helper.go main.go plan
    ```