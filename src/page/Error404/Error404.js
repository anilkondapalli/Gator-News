import React from "react";
import { Link } from "react-router-dom";
import Error404Image from "../../assets/png/404-Error.png";
import Logo from "../../assets/png/gator-logo.png";
import "./Error404.scss";

export default function Error404() {

    return(
        <div className="error404">
        <img src={Logo} alt="GatorLogo" />
        <img src={Error404Image} alt="Error404" />
        <Link to="/">Back to top</Link>
      </div>
    )
}