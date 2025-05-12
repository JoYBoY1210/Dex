# 🐧 DEX - Bookmark Manager Chrome Extension

DEX is a minimalist, category-wise bookmark manager Chrome extension built using **React** and **Golang**. With DEX, you can save links and organize them into categories — all from a sleek, user-friendly interface.

---

## 🚀 Features

- 📌 Save links with titles and categories  
- 📁 Organize bookmarks by category  
- ⚡️ Fast and lightweight  
- 🌐 Built with Go (backend) and React (frontend)

---

## 🛠️ Tech Stack

- **Frontend**: React + Vite + Chrome Extension APIs  
- **Backend**: Go (Golang), net/http, SQLite

---

## 📦 Installation (For Development)

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/dex-extension.git
cd dex-extension
```

### 2. Start the Backend (Go)

```bash
cd dex-backend
go run main.go
```

This starts the backend server (default port: `http://localhost:8080`).

> ⚠️ Make sure your backend is running before using the extension.

### 3. Build the Frontend

```bash
cd dex-frontend
npm install
npm run build
```

This creates the production-ready files in the `dist/` folder.

---

### 4. Load the Extension in Chrome

1. Open Chrome and navigate to `chrome://extensions`  
2. Enable **Developer Mode** (top right)  
3. Click **Load unpacked**  
4. Select the `dex-frontend/dist/` folder  
5. You're done! 🎉

---

## 📁 Folder Structure

```
dex-extension/
├── dex-backend/        # Go backend
│   └── main.go
├── dex-frontend/       # React extension
│   ├── src/
│   └── dist/           # Build output loaded by Chrome
├── README.md
```

---

## 🔐 Security Note

Currently, bookmarks are stored **without user authentication**, which means:
- All users share the same bookmark storage.
- Do **NOT** use this in production without implementing user-based separation.

We recommend adding authentication if you plan to publish the extension to the Chrome Web Store.

---

## 📬 Contact

Built with ❤️ by [Your Name]  
GitHub: [@yourusername](https://github.com/yourusername)

---

## 📜 License

This project is licensed under the MIT License.
