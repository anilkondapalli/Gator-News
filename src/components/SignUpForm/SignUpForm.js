import React, { useState } from "react";
import { Row, Col, Form, Button, Spinner } from "react-bootstrap";
import { values, size } from "lodash";
import { toast } from "react-toastify";
import { isEmailValid } from "../../utils/validations";
import { signUpApi } from "../../api/auth";

import "./SignUpForm.scss";

export default function SignUpForm(props) {
  const { setShowModal } = props;
  const [formData, setFormData] = useState(initialFormValue());
  const [signUpLoading, setSignUpLoading] = useState(false);

  const onSubmit = e => {
    e.preventDefault();

    let validCount = 0;
    values(formData).some(value => {
      value && validCount++;
      return null;
    });

    if (validCount !== size(formData)) {
      toast.warning("Complete All form details");
    } else {
      if (!isEmailValid(formData.email)) {
        toast.warning("Email invalid");
      } else if (formData.password !== formData.confirmPassword) {
        toast.warning("Password did not match");
      } else if (size(formData.password) < 6) {
        toast.warning("Password characters less than 6");
      } else {
        setSignUpLoading(true);
        signUpApi(formData)
          .then(response => {
            if (response.code) {
              toast.warning(response.message);
            } else {
              toast.success("Successful Created");
              setShowModal(false);
              setFormData(initialFormValue());
            }
          })
          .catch(() => {
            toast.error("Server error, please try again later!");
          })
          .finally(() => {
            setSignUpLoading(false);
          });
      }
    }
  };

  const onChange = e => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  return (
    <div className="sign-up-form">
      <Form onSubmit={onSubmit} onChange={onChange}>
        <Form.Group>
          <Row>
            <Col>
              <Form.Control
                type="text"
                placeholder="Name"
                name="name"
                defaultValue={formData.name}
              />
            </Col>
            <Col>
              <Form.Control
                type="text"
                placeholder="Last Name"
                name="lastname"
                defaultValue={formData.lastname}
              />
            </Col>
          </Row>
        </Form.Group>
        <Form.Group>
          <Form.Control
            type="email"
            placeholder="Email Address"
            name="email"
            defaultValue={formData.email}
          />
        </Form.Group>
        <Form.Group>
          <Row>
            <Col>
              <Form.Control
                type="password"
                placeholder="Password"
                name="password"
                defaultValue={formData.password}
              />
            </Col>
            <Col>
              <Form.Control
                type="password"
                placeholder="Confirm Password"
                name="confirmPassword"
                defaultValue={formData.confirmPassword}
              />
            </Col>
          </Row>
        </Form.Group>

        <Button data-testid="modalButton" variant="primary" type="submit" role="submit" >
          {!signUpLoading ? "Register" : <Spinner animation="border" />}
        </Button>
      </Form>
    </div>
  );
}

function initialFormValue() {
  return {
    name: "",
    lastname: "",
    email: "",
    password: "",
    confirmPassword: ""
  };
}