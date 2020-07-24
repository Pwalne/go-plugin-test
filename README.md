# go-plugin-test
Testing golang native plugins with a shared package on different enviornments/GOPATHS.

## Experiment
The experiment is pretty straight forward. The binary of `plugin-test` was built using go1.13.3 linux/amd64.
One was built and committed just in the chance another environment is not available, requiring only the plugin to be built in another environment.


## Results
There are a few things required in order to get native plugins to function properly.

* Only available on linux/macOS
* Must be using the same version of golang
* Seems you need to use `-trimpath` to help with standardizing GOPATH/modules.
* Shared libraries have the following restrictions
    * You must use the same module version.
    * Use `replace` to reference the module location if module's are not available.
    * Shared libraries should be placed in their own module / repo to reduce incompatibilities