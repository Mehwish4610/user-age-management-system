function UserTable({ users, onEdit, onDelete }) {
  return (
    <div className="card">
      <h2>Users List</h2>

      {users.length === 0 ? (
        <p className="empty-text">No users found.</p>
      ) : (
        <div className="table-wrapper">
          <table>
            <thead>
              <tr>
                <th>ID</th>
                <th>Name</th>
                <th>DOB</th>
                <th>Age</th>
                <th>Actions</th>
              </tr>
            </thead>

            <tbody>
              {users.map((user) => (
                <tr key={user.id}>
                  <td>{user.id}</td>
                  <td>{user.name}</td>
                  <td>{user.dob}</td>
                  <td>{user.age}</td>
                  <td>
                    <button className="btn edit" onClick={() => onEdit(user)}>
                      Update
                    </button>
                    <button
                      className="btn delete"
                      onClick={() => onDelete(user.id)}
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}

export default UserTable;