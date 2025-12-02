# Go: Real Time Chat Application (WIP)

Real-time chat application built on GO that allows users to send and receive messages in real time 
 

 ## Services descriptions
*   **back-end-service**: contains the Websocket that the front-end uses in the chat-room
*   **front-end-service**: Serve the front-end templates
*   **broker-service**: crud chat_message db (wip)
 

### Stack
<ul>
<li>Go</li>
<li>Docker</li>
</ul>

### How to run the project 
*   **Start containers**: cd to project -> then run the following commands 
  ```bash
$ make up_build
```
The command will build the binary files for the `backend` and `url-front-end` first, and  then   it will start the containers

### How to test the chat-room
*   **Browser #1**: once project is up & running, access to the following URl  
  ```bash
 http://localhost:3000/
```
*   **Username field**: Type your username 
*   **Message field**: Type the message
*   **Send button**: click the `send` button to send the message 


 *  **Browser #2**: on a different browser, access to the following URL
  ```bash
 http://localhost:3000/
```
*   **Username field**: Type your username 
*   **Message field**: Type the message
*   **Send button**: click the `send` button to send the message 


As soon as the user types their username, the username will appear under the `Who's online` panel in real time, and as soon as the user clicks the `send` button, everybody in the chatroom will see the message 
 
<img width="1139" height="610" alt="Screenshot 2025-11-30 at 6 26 37 PM" src="https://github.com/user-attachments/assets/50c4f672-5cf6-4a04-b7b9-c9ac3c8cf088" />
