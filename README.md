# go-hpegl-provider
The plugin uses Terraform to configure Virtual Machines on ESXi platform using hpegl provider

## Required Environment variables
* SERVICE_ACCOUNT -> HPE GreenLake vmaas service account (dev/intg/prod).  
* TF_WORKING_DIR -> Directory where the terraform files present.
* HPEGL_TENANT_ID -> The tenant-id to be used.
* HPEGL_USER_ID -> The user id to be used.
* HPEGL_USER_SECRET -> The user secret to be used.
* HPEGL_IAM_SERVICE_URL -> The IAM service URL to be used to generate tokens.

## Usage
* This plugin can be used with following terraform operations.
    * plan
    * apply
    * destroy
* Clone this project.
* To perform `terrform plan`, execute below command    
    ```
    go run helper.go main.go plan
    ```
* To perform `terrform apply`, execute below command    
    ```
    go run helper.go main.go apply
    ```
* To perform `terrform destroy`, execute below command    
    ```
    go run helper.go main.go destroy
    ```