import { Route, Routes } from 'react-router-dom';
import { Layout } from './components/Layout';
import Homepage from "./pages/Homepage/Homepage";
import ModelRange from "./pages/ModelRange/ModelRange";
import AvailableCars from "./pages/AvailableCars/AvailableCars";
import LexusWorld from "./pages/LexusWorld/LexusWorld";
import Contacts from "./pages/Contacts/Contacts";
import CarDetails from "./pages/CarDetails/CarDetails";
import { useState, useEffect } from 'react';
import axios from 'axios';
import AdminLayout from './admin/AdminLayout/AdminLayout';
import ProtectedRoute from './admin/components/ProtectedRoute';
import Login from './admin/auth/Login/Login';
import Dashboard from './admin/Dashboard/Dashboard';
import CarList from './admin/cars/CarList/CarList';
import AddCar from './admin/cars/AddCar/AddCar';
import EmployeeList from './admin/employees/EmployeeList/EmployeeList';
import './index.css';

function App() {
  const [cars, setCars] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchCars = async () => {
      try {
        const response = await axios.get('http://localhost:8000/cars');
        setCars(response.data);
      } catch (error) {
        console.error('Error fetching cars:', error);
      } finally {
        setLoading(false);
      }
    };
    fetchCars();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <Routes>
        {/* Public Routes */}
        <Route path="/" element={<Layout />}>
          <Route index element={<Homepage />} />
          <Route path="modelrange" element={<ModelRange />} />
          <Route path="availablecars" element={<AvailableCars cars={cars} />} />
          <Route path="availablecars/:id" element={<CarDetails cars={cars} />} />
          <Route path="lexusworld" element={<LexusWorld />} />
          <Route path="contacts" element={<Contacts />} />
        </Route>

        {/* Admin Routes */}
        <Route path="/admin/login" element={<Login />} />
        <Route path="/admin" element={<ProtectedRoute><AdminLayout /></ProtectedRoute>}>
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="cars" element={<CarList />} />
          <Route path="cars/add" element={<AddCar />} />
          <Route path="employees" element={<EmployeeList />} />
        </Route>
      </Routes>
    </>
  );
}

export default App;