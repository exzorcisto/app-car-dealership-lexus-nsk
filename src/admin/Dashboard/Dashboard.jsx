import React from 'react';
import { Link } from 'react-router-dom';
import './Dashboard.css';

function Dashboard() {
  return (
    <div className="dashboard">
      <h1>Admin Dashboard</h1>
      <div className="dashboard-cards">
        <Link to="/admin/cars" className="dashboard-card">
          <h2>Cars Management</h2>
          <p>Manage your car inventory</p>
        </Link>
        <Link to="/admin/employees" className="dashboard-card">
          <h2>Employees</h2>
          <p>Manage staff accounts</p>
        </Link>
      </div>
    </div>
  );
}

export default Dashboard;