# 🎨 Multiplayer Whiteboard

A simple real-time multiplayer whiteboard built for learning WebSockets, Vue, and Go.

Users can:
- Register and login
- Create and join drawing rooms
- Draw together in real-time
- See drawings instantly from other users

This project follows the KISS principle:
> Keep It Simple, Stupid

The goal of this project is learning:
- Real-time communication
- WebSockets
- Canvas drawing
- Authentication
- Frontend/Backend integration

---

# 🛠 Tech Stack

## Frontend
- Vue 3
- Pinia
- UnoCSS
- Canvas API
- Vue Router

## Backend
- Go
- Gorilla WebSocket
- SQLite
- JWT
- bcrypt

---

# 📂 Project Structure

```text
whiteboard/
├── backend/
│   ├── main.go
│   ├── auth.go
│   ├── db.go
│   ├── ws.go
│   └── whiteboard.db
│
└── frontend/
    ├── src/
    │   ├── App.vue
    │   ├── Login.vue
    │   ├── Board.vue
    │   ├── store.js
    │   └── main.js
```
🚀 Features

✅ User registration
✅ User login
✅ JWT authentication
✅ Real-time drawing
✅ WebSocket communication
✅ Room-based collaboration
✅ Drawing history saving
✅ Multiple users in same room
⚡ How It Works
Authentication
Users register and login using:

Username
Password
Passwords are hashed using bcrypt.

After login:

Backend generates a JWT token
Frontend stores it in localStorage
Real-Time Drawing
When a user draws:

Canvas captures mouse movement
Drawing coordinates are converted to JSON
JSON is sent through WebSocket
Backend broadcasts data to all users in same room
Other users instantly see the drawing
🧠 Learning Goals
This project helps understand:

HTTP APIs
Authentication flow
JWT tokens
Password hashing
WebSockets
Real-time systems
Canvas API
Vue reactivity
State management with Pinia
Frontend/Backend communication
▶️ Running The Project
Backend
Bash

cd backend

go mod tidy

go run .
Server runs on:

text

http://localhost:8080
Frontend
Bash

cd frontend

npm install

npm run dev
Frontend runs on:

text

http://localhost:5173
🔌 API Endpoints
Auth
Register
http

POST /api/register
Login
http

POST /api/login
Rooms
Create Room
http

POST /api/rooms
Room History
http

GET /api/rooms/:id/history
WebSocket
text

ws://localhost:8080/ws?room=ROOM_ID
📸 Example Workflow
User registers
User logs in
User enters room
WebSocket connection opens
User starts drawing
Backend broadcasts strokes
Other users receive drawing instantly
📌 Notes
This project is intentionally simple.

It does NOT include:

Advanced permissions
Scaling architecture
Redis
Docker
Production optimizations
The focus is:

Understanding
Building manually
Learning fundamentals
✅ Future Improvements
Possible upgrades:

Eraser tool
Shapes
Undo / Redo
Private rooms
User cursors
Room list UI
Mobile support
Better toolbar
Drawing optimization
📄 License
MIT
