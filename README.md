# Peer coding exercise

A simple exercise for coding in golang
* This is not an exam
* Nor is a memory test, feel free to use the internet

# Prerequisites
* VS code
* Live share extension
* Golang extension, from the Google team 

Refer https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code

# Execise

## Implement simple cache
 > don't worry about effencies or optimisation or that no would do this in a real production system.

 The cache has two actions a get and a put.

 The get will return the cached string or an error if there is no value in the cache.

 The put will add the value to the cache. 
 * The cache will have a limited length. 
 * If the cache is full when an item is added the item with the oldest access time will be dropped.