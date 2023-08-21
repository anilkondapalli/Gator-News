import React from "react";
import { Modal } from "react-bootstrap";
import Gator from "../../../assets/png/logo-white.png";

import "./BasicModal.scss";

export default function BasicModal(props) {
  const { show, setShow, children } = props;

  return (
    <div data-testid="modal-header"> 
    <Modal 
      className="basic-modal"
      show={show}
      onHide={() => setShow(false)}
      centered
      size="lg"
    >
      <Modal.Header >
        <Modal.Title >
          <img title="image" data-testid="gator_logo" src={Gator} alt="Twittor" />
        </Modal.Title>
      </Modal.Header>
      <Modal.Body >{children}</Modal.Body>
    </Modal>
    </div>
  );
}