# Simple Chatroom using RPC

This is a simple chatroom implemented in **Go** using **Remote Procedure Call (RPC)**.

## 🧠 Overview
- The server listens for incoming RPC connections.
- Each client connects to the server, enters their name, and can send messages.
- The server stores all chat messages in memory.
- The full chat history is returned to every client each time someone sends a new message.
- The chat continues until the client types `exit`.

---

## ⚙️ How to Run

1. **Start the server**:
   ```
 go run server.go
Output:
Server started on port 1234...
```
2. **Run the client(s) in other terminals**:
   ```
 go run client.go
```
3. **Enter your name and start chatting!**

4. **Type exit to leave the chat.**

Example Output :
Server started on port 1234...
[Abdelfattah]: hello fares how are you ?
[Fares]: i'm fine abdo how are you
[Fares]: what are you doing

**Client (Abdelfattah)**
Enter your name: Abdelfattah
> hello fares how are you ?
--- Chat History ---
[Abdelfattah]: hello fares how are you ?
--------------------

**Client (Fares)**
Enter your name: Fares
> i'm fine abdo how are you
--- Chat History ---
[Abdelfattah]: hello fares how are you ?
[Fares]: i'm fine abdo how are you
--------------------

**Demo Video**

