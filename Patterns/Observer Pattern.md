# Observer Pattern

Observer is a behavioral pattern that allows some objects to notify other objects about changes in their state. 

The Observer pattern suggests that you add a subscription mechanism to the publisher class so individual objects can subscribe to or unsubscribe from a stream of events coming from that publisher. Fear not! Everything isn’t as complicated as it sounds. In reality, this mechanism consists of 1) an array field for storing a list of references to subscriber objects and 2) several public methods which allow adding subscribers to and removing them from that list.

<img width="495" alt="Screenshot 2024-02-01 at 3 11 09 PM" src="https://github.com/ahenrie/Go/assets/103060170/d74e9788-afbe-4674-b945-ac1d8e0db72a">

