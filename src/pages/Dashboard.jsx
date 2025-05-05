// src/pages/Dashboard.jsx
import React, { useState, useEffect } from 'react';
import Addcarform from '../components/Addcarform';
import Navbar from '../components/Navbar';

function Dashboard() {
  const [newcar, setNewcar] = useState();
  const [username, setUsername] = useState('');
  const [greeting, setGreeting] = useState('');

  const addCar = (car) => {
    setNewcar([...newcar, { ...car, _id: Date.now() }]);
  };

  useEffect(() => {
    const storedUsername = localStorage.getItem('username');
    if (storedUsername) {
      setUsername(storedUsername);
    }

    const hour = new Date().getHours();
    if (hour < 12) setGreeting('Good Morning');
    else if (hour < 18) setGreeting('Good Afternoon');
    else setGreeting('Good Evening');
  }, []);

  return (
    <>
    <Navbar />
    <div className="container mt-5 pt-3" style={{ maxWidth: '800px' }}>
      <div className="text-center mb-3 pt-5">
      <h2 className="fw-semibold">{greeting}, {username || 'Guest'}!</h2>
      <h6 className="text-muted">Welcome back — ready to add a new ride?</h6>
    </div>
  
  <div className="overflow-auto" style={{ maxHeight: '800px' }}>
    <Addcarform/>
  </div>
</div>

    </>
  );
}

export default Dashboard;
