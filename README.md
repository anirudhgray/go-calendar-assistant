### **Google Calendar CLI Assistant**

EXC Course 2024: A command-line tool built with Go to interact with your Google Calendar. Easily list, add, and manage events from the terminal.

---

### **Setup Instructions**
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd calendar-assistant
   ```

2. Initialize the project:
   ```bash
   go mod tidy
   ```

3. Add your Google OAuth credentials:
   - Enable the Google Calendar API in the [Google Cloud Console](https://console.cloud.google.com/).
   - Download the `credentials.json` file and place it in the `credentials/` folder.

4. Run the application:
   ```bash
   go run main.go
   ```

5. **List Events**:
   ```bash
   go run main.go list
   ```

6. **Add Events**:
   ```bash
   go run main.go add
   ```
