Your task, should you choose to accept it, is to implement an autonomous agent.

Autonomous agents have a number of defining characteristics:
- Communicate with the environment via asynchronous messages
- Display reactiveness (handling messages) and proactiveness (generating new messages based on internal state or local time)
- Can be thought of as representing a human, organisation, or thing in a specific domain and tasks

Your autonomous agent should support these operations and characteristics:
- Continuously consume messages (of different types) from an InBox
- Emit messages to an OutBox
- Allow for registration of message handlers to handle a given message type with its specific handler (reactive: if this message then that is done)
- Allow for registration of behaviours (proactive: if this internal state or local time is reached then this message is created)

Once the generic autonomous agent exists, create a concrete instance which:
- Has one handler that filters messages for the keyword “hello” and prints the whole message to stdout
- Has one behaviour that generates random 2-word messages from an alphabet of 10 words (“hello”, “sun”, “world”, “space”, “moon”, “crypto”, “sky”, “ocean”, “universe”, “human”) every 2 seconds

Run two instances of your concrete agents where the InBox of agent 1 is the OutBox of agent 2 and vice versa.
