module multiple-modules-A

go 1.23

require github.com/sheldon/submodule-A v0.0.0

require github.com/sheldon/submodule-B v1.2.3

replace github.com/sheldon/submodule-A => ../submodule-A

replace github.com/sheldon/submodule-B => ../submodule-B
