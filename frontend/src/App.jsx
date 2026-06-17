import { useEffect, useState } from "react";
import "./App.css";

import UserForm from "./components/UserForm";
import UserTable from "./components/UserTable";
import Pagination from "./components/Pagination";

import {
  getUsers,
  createUser,
  updateUser,
  deleteUser,
} from "./services/api";

function App() {
  const [users, setUsers] = useState([]);
  const [selectedUser, setSelectedUser] = useState(null);
  const [page, setPage] = useState(1);
  const [limit, setLimit] = useState(5);
  const [message, setMessage] = useState("");
  const [loading, setLoading] = useState(false);
  const [darkMode, setDarkMode] = useState(false);
  const [searchTerm, setSearchTerm] = useState("");

  const fetchUsers = async () => {
    try {
      setLoading(true);
      const response = await getUsers(page, limit);
      setUsers(response.data.data || []);
    } catch (error) {
      setMessage("Failed to fetch users");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, [page, limit]);

  const handleSubmit = async (formData) => {
    try {
      if (selectedUser) {
        await updateUser(selectedUser.id, formData);
        setMessage("User updated successfully");
        setSelectedUser(null);
      } else {
        await createUser(formData);
        setMessage("User created successfully");
      }

      fetchUsers();
    } catch (error) {
      setMessage(error.response?.data?.error || "Something went wrong");
    }
  };

  const handleEdit = (user) => {
    setSelectedUser(user);
    window.scrollTo({ top: 0, behavior: "smooth" });
  };

  const handleDelete = async (id) => {
    const confirmDelete = window.confirm(
      "Are you sure you want to delete this user?"
    );

    if (!confirmDelete) return;

    try {
      await deleteUser(id);
      setMessage("User deleted successfully");
      fetchUsers();
    } catch (error) {
      setMessage(error.response?.data?.error || "Failed to delete user");
    }
  };

  const handleCancel = () => {
    setSelectedUser(null);
  };

  const filteredUsers = users.filter((user) =>
    user.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const totalUsers = users.length;

  const averageAge =
    users.length > 0
      ? Math.round(users.reduce((sum, user) => sum + user.age, 0) / users.length)
      : 0;

  const youngestUser =
    users.length > 0
      ? users.reduce((youngest, user) =>
          user.age < youngest.age ? user : youngest
        )
      : null;

  const oldestUser =
    users.length > 0
      ? users.reduce((oldest, user) => (user.age > oldest.age ? user : oldest))
      : null;

  return (
    <div className={darkMode ? "app dark" : "app"}>
      <aside className="sidebar">
        <div className="logo" onClick={() => window.scrollTo({ top: 360, behavior: "smooth" })} >👥 User Age Management</div>


        <div className="status-box">
          <strong>● API Status</strong>
          <p>Connected</p>
          <small>Backend: localhost:8081</small>
        </div>
      </aside>

      <main className="main">
        <section className="hero">
          <div>
            <h1><strong>WELCOME 👋</strong></h1>
            <p>Manage users, store DOB, and view dynamically calculated ages.</p>
          </div>

          <button
            className="theme-toggle"
            onClick={() => setDarkMode(!darkMode)}
          >
            {darkMode ? "☀️ Light Mode" : "🌙 Dark Mode"}
          </button>
        </section>

        {message && (
          <div className="message">
            {message}
            <button onClick={() => setMessage("")}>×</button>
          </div>
        )}

        <section className="stats-grid">
          <div className="stat-card">
            <span>👥</span>
            <h3>{totalUsers}</h3>
            <p>Total Users</p>
          </div>

          <div className="stat-card">
            <span>🎂</span>
            <h3>{averageAge}</h3>
            <p>Average Age</p>
          </div>

          <div className="stat-card">
            <span>👶🏻</span>
            <h3>{youngestUser ? youngestUser.name : "-"}</h3>
            <p>Youngest User</p>
          </div>

          <div className="stat-card">
            <span>🧓🏻</span>
            <h3>{oldestUser ? oldestUser.name : "-"}</h3>
            <p>Oldest User</p>
          </div>
        </section>

        <UserForm
          onSubmit={handleSubmit}
          selectedUser={selectedUser}
          onCancel={handleCancel}
        />

        {loading ? (
          <div className="card loading">Loading users...</div>
        ) : (
          <>
            <div className="toolbar">
              <input
                type="text"
                placeholder="Search users by name..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
              />
            </div>

            <UserTable
              users={filteredUsers}
              onEdit={handleEdit}
              onDelete={handleDelete}
            />

            <Pagination
              page={page}
              limit={limit}
              onPrevious={() => setPage((prev) => Math.max(prev - 1, 1))}
              onNext={() => setPage((prev) => prev + 1)}
              onLimitChange={(newLimit) => {
                setLimit(newLimit);
                setPage(1);
              }}
            />
          </>
        )}
      </main>
    </div>
  );
}

export default App;