import React, { useEffect, useState } from "react";
import { NavLink, useNavigate } from "react-router-dom";
import "./Navbar.css";
import img from "../../assets/img/img-login.png";
import GetCookie from "../../hooks/GetCookie";
import SetCookie from "../../hooks/SetCookie";
import jwt_decode from "jwt-decode";
import axios from "axios";

const App = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  // const [token, setToken] = useState("");
  let navigate = useNavigate();

  let sectionStyle = {
    width: "100%",
    height: "100vh",

    // backgroundPosition: 'center',
    backgroundSize: "cover",
    backgroundRepeat: "no-repseat",
    backgroundImage: `url(${img})`,
  };

  const submit = async (e) => {
    e.preventDefault();

    try {
      const res = await axios.post(
        "http://localhost:8008/Login",
        {
          username: username,
          password: password,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
      console.log(res.data.code);
      SetCookie("token", res.data.data.token);

      // decode token
      let decodeRole = GetCookie("token");
      let newRole = jwt_decode(decodeRole);
      console.log(newRole.Role);
      if (newRole.Role === "user") {
        navigate("/home");
      } else if (newRole.Role === "admin") {
        navigate("/admin");
      }
    } catch (error) {
      alert("Wrong username or password");
    }
  };

  return (
    <div className="main-login" style={sectionStyle}>
      <div className="login-contain">
        <div className="flex-row align-items-end">
          <div className="container">
            <div className="row justify-content-end">
              <div className="col-md-6 login-col-custom ms-5 me-5">
                <div className="card p-3 cek">
                  <div className="card-body">
                    {/* redirect ke home */}
                    <form onSubmit={submit}>
                      {/* {token && <Navigate to="/home" replace={true} />} */}
                      <div className="form-group row mb-2">
                        <NavLink
                          className="nav-link active"
                          aria-current="page"
                          to="/"
                        ></NavLink>
                        <h2 className="text-center">Login To Your Account </h2>
                        <label
                          htmlFor="inputUsername3"
                          className="col-sm-3 col-form-label"
                        >
                          Username
                        </label>
                        <div className="col-sm-10"> </div>
                        <input
                          onChange={(e) => setUsername(e.target.value)}
                          value={username}
                          name="username"
                          type="username"
                          id="inputUsername3"
                          placeholder="Enter username"
                          className="form-control"
                        />
                      </div>

                      <div className="form-group row mb-3">
                        <label
                          htmlFor="inputPassword3"
                          className="col-sm-3 col-form-label"
                        >
                          Password
                        </label>
                        <div className="col-sm-10"></div>
                        <input
                          onChange={(e) => setPassword(e.target.value)}
                          value={password}
                          type="password"
                          name="password"
                          placeholder="Enter password"
                          className="form-control"
                        />
                      </div>

                      <div className="buttonLogin row mb-3">
                        <button
                          className="btn btn-success-custom"
                          type="submit"
                        >
                          Login
                        </button>
                      </div>
                      <p className="text-black">
                        Don't have an account? <a href="register"> Register </a>
                      </p>
                    </form>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
export default App;
