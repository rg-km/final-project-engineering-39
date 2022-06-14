import React from "react";
import img from "../../../assets/img/regisImg.png";
import "./Navbar_Regis.css";

export default function Navbar() {
    return(

        <div className="container">

            <div class="row">
                <div className="col">
                    <span className="text-ladang1">Ladang </span>
                    <span className="text-materi1">Materi</span>
                </div>
                <div class="col-6">
                    <div className="img-con">
                        <img src = {img} alt="regis photo" />
                    </div>

                    <form>
                        <div class="form-group row mb-3">
                            <label for="inputUsername3" class="col-sm-2 col-form-label">Username</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" id="inputUsername3" placeholder="Username"></input>
                            </div>
                        </div>

                        <div class="form-group row mb-3">
                            <label for="inputEmail3" class="col-sm-2 col-form-label">Email</label>
                            <div class="col-sm-10">
                                <input type="email" class="form-control" id="inputEmail3" placeholder="Email"></input>
                            </div>
                        </div>
                    
                        <div class="form-group row mb-3">
                            <label for="inputPassword3" class="col-sm-2 col-form-label">Password</label>
                            <div class="col-sm-10">
                                <input type="password" class="form-control" id="inputPassword3" placeholder="Password"></input>
                            </div>
                        </div>

                        <div className="buttonRegister">
                            <button type="button" class="btn btn-success">Register</button>
                        </div>
                    </form>
                </div>
                <div className="col">
                    <h4>Register Page</h4>
                </div>
            </div>

        </div>
        
    );
}