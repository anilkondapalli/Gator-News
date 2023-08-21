import { fireEvent, render, screen } from '@testing-library/react';
import React from "react";
import SignInForm from './SignInForm';
import {isEmailValid} from "../../utils/validations";

describe ("login",()=>{
  test ("validate function should pass on correct input",()=>{
      const test='alekhyagollam68@gmail.com'
      expect(isEmailValid(test)).toBe(true);
  });

  test ("validate function should fall on correct input",()=>{
    const test='ABCDEFGH'
    expect(isEmailValid(test)).not.toBe(true);
});
  test("Email field should have Placeholder",()=>{
    render(<SignInForm/>);
    screen.queryByPlaceholderText(/email/i)
  });
  test("email input should accept text",()=>{
      render(<SignInForm/>);
      const emailInput=screen.getByLabelText("Email");
      expect(emailInput.value).toMatch("");
      fireEvent.change(emailInput,{target:{value:"Unit Testing"}});
      expect(emailInput.value).toMatch("Unit Testing");
      fireEvent.change(emailInput,{target:{value:"unit Testing@"}});
  })
  test("user should be able to submit the form",()=>{
      const mockFunction=jest.fn()
      render(<SignInForm handleSubmit={mockFunction}/>)
      const buttonNode=screen.getByRole("submitForm");
      fireEvent.submit(buttonNode);

  })
})
