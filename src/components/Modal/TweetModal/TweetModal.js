import React, { useState } from "react";
import { Modal, Form, Button } from "react-bootstrap";
import "./TweetModal.scss"
import classNames from "classnames";
import { Camera } from "../../../utils/Icons";
import { toast } from "react-toastify";

import {addTweetApi} from "../../../api/tweet";

export default function TweetModal(props) {
    const { show, setShow } = props;
    const [message, setMessage] = useState("");
    const maxLength=280;


    const onSubmit = (e) => {
      e.preventDefault();
  
      if (message.length > 0 && message.length <= maxLength) {
        console.log("OK");
        addTweetApi(message)
          .then((response) => {
            console.log(response);
            if (response?.code >= 200 && response?.code < 300) {
              toast.success(response.message);
              setShow(false);
              //window.location.reload();
            }
          })
          .catch(() => {
            toast.warning("Error sending gator tweet, please try again later");
          });
      }
    };
  return (
    <Modal
      className="tweet-modal"
      show={show}
      onHide={() => setShow(false)}
      centered
      size="lg"
    >
        <Modal.Header>
        <Modal.Title>
          <Camera onClick={() => setShow(false)} />
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
      <Form onSubmit={onSubmit}>
          <Form.Control
            as="textarea"
            rows="6"
            placeholder="Whats Happening Gator?"
            onChange={(e) =>{
                e.preventDefault(); 
                setMessage(e.target.value);
                console.log(message);
            }
        }
          />
          <span
            className={classNames("count", {
              error: message.length > maxLength,
            })}
          >
            {message.length}
          </span>
          <Button role="submit"
            type="submit"
            disabled={message.length > maxLength || message.length < 1}
          >
              Buzz
              </Button>
          </Form>
          </Modal.Body>
    </Modal>
  );
}
