# Self Register Factory

In this main idea is to create a self registering factory pattern with minimum code changes needed.

In OOP languages it is possible to do this with polymorphism as referenced from here[https://accu.org/journals/overload/6/27/bellingham_597/]. 

Here is my attempt to create something close to this.

I tried to leverage the fact that init is run when a package is imported. With this approach I've found this.

- In order to minimize the development, we can create a new package without having to change the factory.
- We can decrease the need of changes in the factory.go it self (maybe adding additional lines to initialize the method implementation)
- Still need to trigger the init() by importing it.

Based on this, it may be feasible but I dont think it will be that useful.