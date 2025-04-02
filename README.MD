## ✅ README.md — go - react learning project


---

## 📌 Table of Contents
- [📖 About the Project](#📖-about-the-project)
- [⚙️ Installation & Configuration](#⚙️-installation--configuration)
- [🚀 Running the Server](#🚀-running-the-server)
- [🔑 Authentication & Security](#🔑-authentication--security)
- [📡 API Endpoints](#📡-api-endpoints)
- [📞 Support & Contributions](#📞-support--contributions)

---

## 📖 About the Project
The project is a learning for go backend api using postgres for database ,
and react for frontend . 

### Features:
- Schema-per-user isolation (multi-tenant ready)
- Automatic DB/user/schema setup on boot
- Configurable via `.env` file
- Ready for REST APIs, and admin portals



## ⚙️ Installation & Configuration

### 1️⃣ Clone the Repository

<!-- ```md
git clone https://github.com/Picasso-K400/server.git
cd K400-server
``` -->

### 2️⃣ Create a `.env` File

```md
DB_HOST=localhost
DB_PORT=5432
DB_ADMIN_USER=postgres
DB_ADMIN_PASSWORD=<your_postgres_password>
DB_ADMIN_DATABASE=postgres

APP_DB_NAME=<k400db>
APP_DB_USER=<k400user>
APP_DB_PASSWORD=<securepass123>


```

> Make sure `.env` is in your `.gitignore`

### 3️⃣ Install Go Dependencies

```md
cd server 
go mod tidy
```

---

## 🚀 Running the Server

```md
go run cmd/main.go
```

By default, the server runs on:  
`http://localhost:8080`


## 🔑 Authentication & Security

<!-- - The system is built to support **JWT-based auth** (coming soon).
- You can configure the initial admin via `.env`. -->

---

## 📡 API Endpoints (WIP)

| Method | Endpoint        | Description              | Auth     |
|--------|------------------|--------------------------|----------|
| `POST` | `/users/login`   | Login and receive token  | Public   |
| `POST` | `/users/create`  | Create new user          | Admin    |
| `GET`  | `/users/`        | List all users           | Admin    |

> More endpoints coming soon as `handlers/` and `routes/` are implemented.


# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type-aware lint rules:

```js
export default tseslint.config({
  extends: [
    // Remove ...tseslint.configs.recommended and replace with this
    ...tseslint.configs.recommendedTypeChecked,
    // Alternatively, use this for stricter rules
    ...tseslint.configs.strictTypeChecked,
    // Optionally, add this for stylistic rules
    ...tseslint.configs.stylisticTypeChecked,
  ],
  languageOptions: {
    // other options...
    parserOptions: {
      project: ['./tsconfig.node.json', './tsconfig.app.json'],
      tsconfigRootDir: import.meta.dirname,
    },
  },
})
```

You can also install [eslint-plugin-react-x](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-x) and [eslint-plugin-react-dom](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-dom) for React-specific lint rules:

```js
// eslint.config.js
import reactX from 'eslint-plugin-react-x'
import reactDom from 'eslint-plugin-react-dom'

export default tseslint.config({
  plugins: {
    // Add the react-x and react-dom plugins
    'react-x': reactX,
    'react-dom': reactDom,
  },
  rules: {
    // other rules...
    // Enable its recommended typescript rules
    ...reactX.configs['recommended-typescript'].rules,
    ...reactDom.configs.recommended.rules,
  },
})
```




## 📞 Support & Contributions

Need help or want to contribute?

- Open an **Issue**
- Submit a **Pull Request**
- Fork and improve the project!

---

🔥  — developer-friendly, production-ready, and built to scale.