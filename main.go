package main

import (
	"fmt"
	"os"
)

var (
	valid_terraform_operations []string
	required_env_variables     []string
	workingdir                 string
)

func main() {
	// Check the Environment variables
	required_env_variables = []string{"SERVICE_ACCOUNT", "TF_WORKING_DIR"}
	for _, env := range required_env_variables {
		_, present := os.LookupEnv(env)
		if !present {
			fmt.Printf("[ERROR] : %v environment variable is required.\n", env)
			os.Exit(1)
		}
	}

	workingdir = os.Getenv("TF_WORKING_DIR")

	valid_terraform_operations = []string{"plan", "apply", "destroy"}

	if len(os.Args) != 2 {
		fmt.Println("[ERROR] : Terraform operation should be specified. Choose one from ", valid_terraform_operations)
		os.Exit(1)
	}

	operation := os.Args[1]
	valid_op := contains(valid_terraform_operations, operation)
	if !valid_op {
		fmt.Printf("[ERROR] : %v is not a valid terraform operation. Choose one from %v\n", operation, valid_terraform_operations)
		os.Exit(1)
	}

	// Terraform init
	fmt.Println("[INFO] : Creating Terraform binary")
	tf := GetInstance(workingdir)
	fmt.Printf("[INFO] : Terraform instance initialised at : %+v\n", tf.WorkingDir())

	switch operation {
	case "plan":
		// Terraform plan
		fmt.Println("[INFO] : Running Terraform Plan")
		hasPlanChanges, err := planWithDefaults(tf)
		if err != nil {
			panic(err)
		}

		if !hasPlanChanges {
			fmt.Printf("[WARN] : Plan has no changes to be provisioned in working directory %v\n", workingdir)
			os.Exit(1)
		}
		fmt.Println("[INFO] : Terraform Plan is completed")

	case "apply":
		// Terraform apply
		fmt.Println("[INFO] : Running Terraform apply")
		err := applyWithDefaults(tf)
		if err != nil {
			panic(err)
		}
		fmt.Println("[INFO] : Applying Terraform resources completed")

	case "destroy":
		// Terraform destroy
		fmt.Println("[INFO] : Running Terraform destroy")
		err := destroyWithDefaults(tf)
		if err != nil {
			panic(err)
		}
		fmt.Println("[INFO] : Destroying Terraform resources completed")
	}
}
