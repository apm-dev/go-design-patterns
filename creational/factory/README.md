## Factory Method

### **Problem**
Imagine that you’re creating a payment management application.
 The first version of your app can only handle
 payments by cash, so you wrote a `type Cach struct{...}` and a method like `func (c *Cash) payByCash() (string
 , error)`.

After a while, your app becomes pretty popular. 
Each day you receive dozens of requests from users to incorporate credit card into the app.

Great news, right? But how about the code? At present, 
most of your code is coupled to the `Cash` struct and `payByCash` method. 
Adding credit card into the app would require making changes to 
the entire codebase. Moreover, if later you decide to add
another type of payment method to the app, 
you will probably need to make all of these changes again.

As a result, you will end up with pretty nasty code,
riddled with conditionals that switch the app’s behavior 
depending on some `payBy` methods.

### **Solution**
The Factory Method pattern suggests that you replace 
direct struct creation and method calls (for each payment method) 
with calls to a special factory method.
Don’t worry: the structs are still created, but it’s being called from within the factory method
. references returned by a factory method are often referred to as products.

At first glance, this change may look pointless: we just moved the structs and methods call from one part of the
 program to another
. However, consider this: now you can override the factory method and change the struct of products being created by
 the method.

There’s a slight limitation though: factory method may return different types of products only if these products
 implemented a common interface. Also, the factory method in the base class should have its return type declared as
  this interface.
