This directory defines resources that must be deployed first in order for the
rest of the Terraform configs to function.

The helper script `deploy.sh` will bootstrap the resources defined.

Currently this only creates the central devops project. After this project has
been created, Terragrunt can bootstrap the state bucket inside the project and
manage all the following resources.
