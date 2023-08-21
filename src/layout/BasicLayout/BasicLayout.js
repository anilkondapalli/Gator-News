import React from "react";

import { Container, Row, Col } from "react-bootstrap";
import LeftMenu from "../../components/LeftMenu";
import "./BasicLayout.scss";

export default function BasicLayout(props) {
    const { className, setRefreshCheckLogin, children } = props;

    return(
            <Row>
                <Col lg={2} className="basic-layout__menu">
                    <LeftMenu setRefreshCheckLogin={setRefreshCheckLogin}></LeftMenu>
                </Col>
                <Col lg={10} className="basic-layout__content">
                    {children}
                </Col>
                <col>
                </col>
            </Row>
    );
}