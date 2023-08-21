import React, {useState} from 'react';
import "./LeftMenu.scss"
import LogoWhite from "../../assets/png/logo-white.png"
import {Button} from "react-bootstrap";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import{
    faHome,
    faUser,
    faUsers,
    faPowerOff
}from "@fortawesome/free-solid-svg-icons";
import TweetModal from '../Modal/TweetModal/TweetModal';
import { logoutApi } from "../../api/auth";

import useAuth from "../../hooks/useAuth";

import { Link } from 'react-router-dom';
export default function LeftMenu(props) {
  const { setRefreshCheckLogin } = props;
  const [showModal, setShowModal] = useState(false);
  const user = useAuth();

  console.log(user);

  const logout = () => {
    logoutApi();
    window.location.reload();
    //setRefreshCheckLogin(true);
  };


  return (
    <div className="left-menu">
      <img  className='logo' src={LogoWhite} alt="Twittor" />
      <Link to="/">
        <FontAwesomeIcon icon={faHome} /> Home
      </Link>
      <Link to="/users">
        <FontAwesomeIcon icon={faUsers} /> Users
      </Link>
      <Link to={`/${user?._id}`}>
        <FontAwesomeIcon icon={faUser} /> Profile
      </Link>
      <Link to="" onClick={logout}>
        <FontAwesomeIcon icon={faPowerOff} /> Logout
      </Link>
      <Button onClick={()=>setShowModal(true)}>Buzz</Button>
      <TweetModal show={showModal} setShow={setShowModal}>

      </TweetModal>
    </div>
  );
}