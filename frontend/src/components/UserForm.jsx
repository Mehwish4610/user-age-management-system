import { useState, useEffect } from "react";

function UserForm({ onSubmit, selectedUser, onCancel }) {
  const [formData, setFormData] = useState({
    name: "",
    dob: "",
  });

  useEffect(() => {
    if (selectedUser) {
      setFormData({
        name: selectedUser.name,
        dob: selectedUser.dob,
      });
    }
  }, [selectedUser]);

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    if (!formData.name || !formData.dob) {
      alert("Please fill all fields");
      return;
    }

    onSubmit(formData);

    if (!selectedUser) {
      setFormData({
        name: "",
        dob: "",
      });
    }
  };

  return (
    <div className="card">
      <h2>{selectedUser ? "Update User" : "Add New User"}</h2>

      <form onSubmit={handleSubmit} className="user-form">
        <div className="form-group">
          <label>Name</label>
          <input
            type="text"
            name="name"
            placeholder="Enter user name"
            value={formData.name}
            onChange={handleChange}
          />
        </div>

        <div className="form-group">
          <label>Date of Birth</label>
          <input
            type="date"
            name="dob"
            value={formData.dob}
            onChange={handleChange}
          />
        </div>

        <div className="form-actions">
          <button type="submit" className="btn primary">
            {selectedUser ? "Update User" : "Add User"}
          </button>

          {selectedUser && (
            <button type="button" className="btn secondary" onClick={onCancel}>
              Cancel
            </button>
          )}
        </div>
      </form>
    </div>
  );
}

export default UserForm;