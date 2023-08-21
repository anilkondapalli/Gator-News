import React, { useState} from "react";
import { Container, Row, Col, Button } from "react-bootstrap";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome"
import {faSearch,faUsers,faComment} from "@fortawesome/free-solid-svg-icons";
import "./SignInSingUp.scss";
import BasicModal from "../../components/Modal/BasicModal";
import GatorLogo from "../../assets/png/logo-white.png";
import SignUpForm from "../../components/SignUpForm"
import SignInForm from "../../components/SignInForm";

export default function SignInSingUp(props) {
  const {setRefreshCheckLogin}=props;
  const [showModal, setShowModal] = useState(false);
  const [contentModal, setContentModal] = useState(null);

  const openModal = content => {
    setShowModal(true);
    setContentModal(content);
  };

  return (
    <div className="signInSignUp">
      <Container className="signin-signup" fluid>
        <Row>
          <LeftComponent />
          <RightComponent
            openModal={openModal}
            setShowModal={setShowModal}
            setRefreshCheckLogin={setRefreshCheckLogin}
          />
        </Row>
      </Container>
      <BasicModal show={showModal} setShow={setShowModal}>
        {contentModal}
      </BasicModal>
      </div>

  );
}
function LeftComponent() {
  return (
    <Col className="signin-signup__left" xs={6}>
      <div>
        
      <h2>
          <FontAwesomeIcon icon={faSearch} />
          Connect with Gators.
        </h2>
        <h2>
          <FontAwesomeIcon icon={faUsers} />
          Community to share your thoughts.
        </h2>
        <h2>
          <FontAwesomeIcon icon={faComment} />
          Get to know your Co-Gators.
        </h2>
      </div>
    </Col>
  );
}

function RightComponent(props) {
  const { openModal, setShowModal, setRefreshCheckLogin} = props;
  return (
    <Col className="signin-signup__right" xs="6">
      <div className="rightPane">
      <div className="LogoImage">
      <img src={GatorLogo} alt="GatorLogo"/>
      </div>
      <div className="content">
      
      <h2>Welcome to Gator News!!</h2>
      <Button
          variant="primary"
          onClick={() => openModal(<SignUpForm setShowModal={setShowModal}></SignUpForm>)}
        >Register</Button>
        <Button
          variant="primary"
          onClick={() => openModal(<SignInForm setRefreshCheckLogin={setRefreshCheckLogin}/>)}
        >Login</Button>
      </div>
      </div>
    </Col>
  );
}
