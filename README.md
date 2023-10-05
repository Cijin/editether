# editether
A real time document collaboration tool built with Go &amp; Htmx

Todo:

X. **Project Setup**:
   - Initialize a new Go project.
   - Set up your Go environment and organize your folder structure.
   - Choose a Go web framework (e.g., Gorilla Mux or Gin).

X. **Basic Server**:
   - Build a basic HTTP server in Go that serves an HTML page.
   - Set up routes for the main page and any additional endpoints.

WIP
   - Choose a database (e.g., PostgreSQL, SQLite).
   - Create tables for storing document data (content, edits, users, etc.).

4. **WebSocket Setup**:
   - Integrate the Gorilla WebSocket package.
   - Create a route specifically for WebSocket connections.
   - Implement WebSocket handlers to manage connection, disconnection, broadcasting messages, etc.

5. **Frontend with HTMX**:
   - Serve the main editing page using HTMX. This allows for partial page updates without full page reloads.
   - Implement the content-editable attribute for the document's content area.
   - Set up HTMX requests to sync content changes to the server.

6. **Collaborative Editing**:
   - On content change, send the changes via WebSockets to the Go server.
   - The server broadcasts these changes to all connected clients.
   - On the client side, apply changes to the document content in real-time.

7. **User Management (optional)**:
   - Implement user registration and authentication.
   - Manage permissions (e.g., who can edit or view the document).

8. **Document Versioning (optional)**:
   - Save document snapshots at intervals.
   - Allow users to view and revert to previous versions.

9. **Optimizations**:
   - Implement Operational Transformation (OT) or Conflict-Free Replicated Data Types (CRDT) for conflict-free collaborative editing.
   - Enhance performance by optimizing WebSocket messages (e.g., sending diffs instead of full content).

10. **Testing**:
   - Write unit tests for your Go functions.
   - Test WebSocket functionality.
   - Test collaboration with multiple users simultaneously.

11. **Deployment**:
   - Choose a hosting provider or platform.
   - Deploy your Go server and database.
   - Set up necessary environment variables and configurations.

12. **Scaling (advanced)**:
   - If your application gains traction, consider load balancing, database sharding, and other scaling strategies.
