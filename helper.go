package main

import (
	"context"
	"sync"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

var (
	_instance *tfexec.Terraform = nil
	once      sync.Once
)

// terraform plan
func planWithDefaults(tf *tfexec.Terraform) (bool, error) {
	hasChanges, err := tf.Plan(context.Background(),
		tfexec.Parallelism(99),
		tfexec.Lock(false),
		tfexec.Refresh(true))
	return hasChanges, err
}

// terraform apply
func applyWithDefaults(tf *tfexec.Terraform) error {
	err := tf.Apply(context.TODO(),
		tfexec.Parallelism(99),
		tfexec.Lock(false),
		tfexec.Refresh(true))
	return err
}

// terraform destroy
func destroyWithDefaults(tf *tfexec.Terraform) error {
	err := tf.Destroy(context.TODO(),
		tfexec.Parallelism(99),
		tfexec.Lock(false),
		tfexec.Refresh(true))
	return err
}

// Make terraform runtime environment
func GetInstance(withWorkdir string) *tfexec.Terraform {
	once.Do(func() {
		installer := &releases.ExactVersion{
			Product: product.Terraform,
			Version: version.Must(version.NewVersion("1.0.10")),
		}

		execPath, err := installer.Install(context.Background())
		if err != nil {
			panic(err)
		}

		tf, err := tfexec.NewTerraform(withWorkdir, execPath)
		if err != nil {
			panic(err)
		}

		// Initialize if not done already
		err = tf.Init(context.Background(), tfexec.Upgrade(true))
		if err != nil {
			panic(err)
		}

		_instance = tf
	})
	return _instance
}

func contains(str []string, searchstr string) bool {
	for _, st := range str {
		if st == searchstr {
			return true
		}
	}
	return false
}
