import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import "./Login.css"; // Reusing styles from Login.css

const Signup = () => {

  const [formData, setFormData] = useState({
    name: '',
    email: '',
    country: '',
    password: ''
  });
  
  const [errorMsg, setErrorMsg] = useState("");
  const navigate = useNavigate();

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSignup = async (e) => {
    e.preventDefault();
    console.log(formData);

    try {
      const res = await axios.post('http://localhost:5000/api/signup', formData);
      console.log("Signup success:", res.data);
      localStorage.setItem("token", res.data.token);
      localStorage.setItem("username", res.data.username);
      localStorage.setItem("userId", res.data.userId);
      navigate("/dashboard");
    } catch (err) {
      console.error("Signup error:", err.response?.data || err.message);
    }
  };

  return (
    <div className="login-container">
      <h2 className="login-title">SignUp</h2>
      <form onSubmit={handleSignup} className="login-form">

        <input
          type="text"
          placeholder="Name"
          name="name"
          value={formData.name}
          className="login-input"
          onChange={handleChange}
          required
        />
        <input
          type="text"
          placeholder="Email"
          name="email"
          value={formData.email}
          className="login-input"
          onChange={handleChange}
          required
        />
        <input
          type="text"
          placeholder="Country"
          name="country"
          value={formData.country}
          className="login-input"
          onChange={handleChange}
          required
        />

        <input
          type="password"
          placeholder="Password"
          name="password"
          value={formData.password}
          className="login-input"
          onChange={handleChange}
          required
        />

        <button type="submit" className="login-button">SignUp</button>

        {errorMsg && <p className="login-error">{errorMsg}</p>}
      </form>
    </div>
  );
};

export default Signup;
