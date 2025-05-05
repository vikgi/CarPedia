import React from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { FaSignOutAlt, FaCar } from 'react-icons/fa';

const Navbar = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("userId");
    localStorage.removeItem("username");
    navigate("/login");
  };

  const handleGoToCars = () => {
    navigate("/cars");
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light fixed-top shadow-sm py-2">
      <div className="container-fluid">
        <span className="navbar-brand fw-bold">CarPedia</span>

        <div className="ms-auto d-flex gap-2">
          <button
            onClick={handleGoToCars}
            className={`btn btn-outline-danger d-flex align-items-center gap-1 ${
              location.pathname === "/cars" ? "active fw-bold" : ""
            }`}
          >
            <FaCar /> Cars
          </button>

          <button
            className="btn btn-outline-danger d-flex align-items-center gap-1"
            onClick={handleLogout}
          >
            <FaSignOutAlt /> Logout
          </button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
