## apictl delete

Delete an API/APIProduct/Application in an environment

### Synopsis

Delete an API available in the environment specified by flag (--environment, -e)
Delete an API Product available in the environment specified by flag (--environment, -e)
Delete an Application of a specific user in the environment specified by flag (--environment, -e)

```
apictl delete [flags]
```

### Examples

```
apictl delete api -n TwitterAPI -v 1.0.0 -r admin -e dev
apictl delete api-product -n TwitterAPI -r admin -e dev 
apictl delete app -n TestApplication -o admin -e dev
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
  -k, --insecure   Allow connections to SSL endpoints without certs
      --verbose    Enable verbose mode
```

### SEE ALSO

* [apictl](apictl.md)	 - CLI for Importing and Exporting APIs and Applications and Managing WSO2 Micro Integrator
* [apictl delete api](apictl_delete_api.md)	 - Delete API
* [apictl delete api-product](apictl_delete_api-product.md)	 - Delete API Product
* [apictl delete app](apictl_delete_app.md)	 - Delete App
* [apictl delete policy](apictl_delete_policy.md)	 - Delete a Policy

