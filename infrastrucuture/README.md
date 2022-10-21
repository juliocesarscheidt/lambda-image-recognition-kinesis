## Deploy IaC

In order to deploy, run the following commands:

```bash
export ENV="development"
# aws settings
export AWS_ACCESS_KEY_ID="$AWS_ACCESS_KEY_ID"
export AWS_SECRET_ACCESS_KEY="$AWS_SECRET_ACCESS_KEY"
export AWS_DEFAULT_REGION="$AWS_DEFAULT_REGION"
# S3 bucket for terraform backend
export AWS_BACKEND_BUCKET="backend-bucket-$(cat /dev/urandom | tr -dc 'a-z0-9' | fold -w 12 | head -n1)"
export AWS_BACKEND_ACCESS_KEY_ID="$AWS_BACKEND_ACCESS_KEY_ID"
export AWS_BACKEND_SECRET_ACCESS_KEY="$AWS_BACKEND_SECRET_ACCESS_KEY"
export AWS_BACKEND_REGION="$AWS_BACKEND_REGION"

cd terraform/

# run vet and compile on application code, also it prepares the executables
make compile

# creates the backend bucket on S3 (if doesn't exist, this could take a few seconds), initializes the terraform, create the workspaces, validate and do the plan
make init

# apply the previous plan
make apply
```
