import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './Sidebar.css';

function Sidebar() {
  const navigate = useNavigate();
  
  const handleLogout = () => {
    localStorage.removeItem('adminToken');
    navigate('/admin/login');
  };

  return (
    <aside className="admin-sidebar">
      <h2>Admin Panel</h2>
      <nav>
        <ul>
          <li><Link to="/admin/dashboard">Dashboard</Link></li>
          <li><Link to="/admin/cars">Cars Management</Link></li>
          <li><Link to="/admin/employees">Employees</Link></li>
        </ul>
      </nav>
      <button onClick={handleLogout} className="logout-btn">Logout</button>
    </aside>
  );
}

export default Sidebar;