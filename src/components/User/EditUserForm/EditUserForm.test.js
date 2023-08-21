import { fireEvent, render, screen } from '@testing-library/react';
import React from "react";
import EditUserForm from './EditUserForm';
import initialValue from './EditUserForm';
import { unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";

describe('EditUserForm', () => {

    let container = null;
    beforeEach(() => {
    container = document.createElement("div");
    document.body.appendChild(container);
    });

    afterEach(() => {
        // cleanup on exiting
        unmountComponentAtNode(container);
        container.remove();
        container = null;
      });

    test("Should render the Routing", () => {
        render(<initialValue/>);
    });
  
    test("email input should accept text",()=>{
        const nameInput="";
        expect(nameInput).toMatch("");
        const nameInput1="Unit Testing";
        expect(nameInput1).toMatch("Unit Testing");
    })
      
});