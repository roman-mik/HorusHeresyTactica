
# Horus Heresy Tactica

**Horus Heresy Tactica** is a dynamic application for tracking legions and simulating tactical battles in the Warhammer 40k universe, focusing on the Horus Heresy era. This project is built using a monorepo structure with a **Golang** backend and a **React/Next.js** frontend. The application is designed to combine immersive battle simulations with legion management features.

## Features

- **Legion Management**: Keep track of your favorite legions from the Horus Heresy era.
- **Tactical Simulations**: Set up and run tactical simulations with different legions.
- **Battle History Tracking**: Monitor the results of battles and track legion victories over time.
- **Dynamic UI**: A grim-dark themed UI reflecting the atmosphere of the Horus Heresy, but with a dynamic twist for the hopeful future envisioned during the Great Crusade.

## Project Structure

This project is organized as a monorepo containing both the frontend and backend:

```
horus-heresy-tactica/
├── backend/                # Go backend API for handling legion data and battle simulations
├── frontend/               # React/Next.js frontend for user interaction and visualizations
├── docker-compose.yml      # Docker Compose setup for running the project
└── README.md               # Project overview and setup instructions
```

### Backend

The backend is written in **Golang** and serves as the API for managing data and running simulations. It exposes endpoints for legion management, battle creation, and simulation.

### Frontend

The frontend is built using **React/Next.js** and serves the user interface for interacting with the application, visualizing legions, and running simulations.

## Getting Started

### Prerequisites

Make sure you have Docker installed to run the application locally.

### Running the Project

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/horus-heresy-tactica.git
   cd horus-heresy-tactica
   ```

2. Start the project using Docker Compose:

   ```bash
   docker-compose up --build
   ```

3. The frontend will be available at `http://localhost:3000` and the backend API at `http://localhost:8080`.

## Roadmap

- [ ] Add more legions from the Horus Heresy lore.
- [ ] Expand tactical simulation features.
- [ ] Create an account system for saving and tracking battles.
- [ ] Integrate PostgreSQL for persistent data storage.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
