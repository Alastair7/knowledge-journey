# ACTOR MODEL PATTERN

**[31/08/2024]**

When sending a message to another actor what it really does `actor.SendMessage()` is:

1. Obtain the actor from the actor system that will receive the message.
2. Send Message to the actor.

It confused me a little because it's like the actor is sending a message to itself once it's found.
