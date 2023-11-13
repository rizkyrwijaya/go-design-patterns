# Creational Pattern

Object Creations.

## [Singleton](singleton)

With <b> singleton </b>  we are ensuring that only one instance of an object is created. In the example created, I've shown how we can have one instance of the package.

### Use Cases

- When we need to make sure that there is only object instance

### Pros

- Ensures there is only a single instance meaning that other packages will hit the same instance
- Can easily control the initialization i.e. put it in ```init()``` and make sure its hit once

### Cons

- Global Variable